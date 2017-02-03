package main

import (
	"bungie"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"gdestiny/auth"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const destinyURL = "https://www.bungie.net/platform/Destiny"

var logger *log.Logger
var destinyDB *sql.DB
var destinyClient *http.Client

func makeCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  value,
		Secure: true,
	}
}

func callBungieService(loc string, options map[string]string) []byte {
	req, err := http.NewRequest("GET", loc, nil)
	if err != nil {
		logger.Fatal(err)
	}
	q := req.URL.Query()
	for opt, val := range options {
		q.Set(opt, val)
	}

	// TODO(wcn): clean this up.
	if destinyClient == nil {
		// The HTTP client needs to be long-lived to take advantage of keepalives
		jar, err := cookiejar.New(nil)
		if err != nil {
			logger.Fatal(err)
		}
		bungled := makeCookie("bungled", auth.Bungled)
		bungleatk := makeCookie("bungleatk", auth.Bungleatk)
		cookieURL, _ := url.Parse(destinyURL)
		jar.SetCookies(cookieURL, []*http.Cookie{bungled, bungleatk})
		destinyClient = &http.Client{
			Jar: jar,
		}
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("X-API-KEY", auth.APIKey)
	resp, err := destinyClient.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		logger.Fatal(err)
	}

	var header bungie.Header
	err = json.Unmarshal(content, &header)

	if header.ErrorCode != 1 {
		// TODO(wcn): the Bungie API never seems to return a non-zero throttle value, so
		// I just delay randomly.
		delay := int(header.ThrottleSeconds)
		if delay < 5 {
			delay = 5
		}
		time.Sleep(time.Duration(1+rand.Intn(delay)) * time.Second)
	}

	return content
}

func prettyPrintJSON(data []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(out.String())
}

func getBungieIDForGamerTag(tag string) string {
	data := callBungieService(fmt.Sprintf("%s/SearchDestinyPlayer/1/%s/", destinyURL, tag), nil)
	var msg bungie.LookupResponse
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logger.Printf("Data: %s", string(data))
		logger.Fatal(err)
	}

	if len(msg.Response) != 1 {
		logger.Fatal("Too many Bungie memberships: %s", string(data))
	}
	return msg.Response[0].MembershipID
}

func dump(url string, options map[string]string) {
	prettyPrintJSON(callBungieService(url, options))
}

func getPGCR(id int64) []byte {
	filename := fmt.Sprintf("reports/%d", id)

	if b, err := ioutil.ReadFile(filename); err == nil {
		return b
	}
	data := callBungieService(fmt.Sprintf("%s/Stats/PostGameCarnageReport/%d/", destinyURL, id), nil)
	ioutil.WriteFile(filename, data, 0644)
	return data
}

func forClass(class string) func([]bungie.CharacterData) string {
	return func(chars []bungie.CharacterData) string {
		for _, char := range chars {
			charClass := getClassForID(char.CharacterBase.ClassHash)
			logger.Printf("Class = %+v\n", charClass)
			if class == charClass {
				return char.CharacterBase.CharacterID
			}
		}
		logger.Fatal("Couldn't find specified class for member")
		return ""
	}
}
func getLastUsed(chars []bungie.CharacterData) string {
	var latest time.Time
	var latestChar bungie.CharacterData

	for _, char := range chars {
		logger.Printf("character = %+v\n", char.CharacterBase.CharacterID)
		logger.Printf("Class = %+v\n", getClassForID(char.CharacterBase.ClassHash))
		last := char.CharacterBase.DateLastPlayed
		logger.Printf("Last used = %+v\n", last)
		ts, err := time.Parse(time.RFC3339, last)
		if err != nil {
			logger.Fatal(err)
		}
		if ts.After(latest) {
			latest = ts
			latestChar = char
		}
	}
	logger.Printf("character = %+v\n", latestChar.CharacterBase.CharacterID)
	logger.Printf("Class = %+v\n", getClassForID(latestChar.CharacterBase.ClassHash))
	logger.Printf("latest = %+v\n", latest)
	return latestChar.CharacterBase.CharacterID
}

func getCharacter(member string, matcher func([]bungie.CharacterData) string) string {
	options := map[string]string{}
	data := callBungieService(fmt.Sprintf("%s/1/Account/%s/Summary/", destinyURL, member), options)
	var msg bungie.SummaryResponse
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logger.Fatal(err)
	}

	return matcher(msg.Response.Data.Characters)
}

func getClassForID(id int64) string {
	// TODO(wcn): use a map or something cleaner.
	if id == 2271682572 {
		return "Warlock"
	}
	if id == 3655393761 {
		return "Titan"
	}
	if id == 671679327 {
		return "Hunter"
	}
	return "UNKNOWN"
}

// My member ID: 4611686018433876772
// warlock characterId: 2305843009395074356
// floor for season 3 PGCR: 5497427370

func getHistoryIDs(membershipID, characterID string) []string {
	done := false
	var ids []string
	total := 0
	for page := 0; !done; page++ {
		options := map[string]string{
			"mode":        "AllPvE",
			"definitions": "true",
			"page":        fmt.Sprintf("%d", page),
			"count":       "250",
		}
		data := callBungieService(fmt.Sprintf("%s/Stats/ActivityHistory/1/%s/%s/", destinyURL, membershipID, characterID), options)
		var msg bungie.ActivityHistory
		err := json.Unmarshal(data, &msg)
		if err != nil {
			logger.Fatal(err)
		}
		if len(msg.Response.Data.Activities) == 0 {
			done = true
			continue
		}

		for _, a := range msg.Response.Data.Activities {
			ids = append(ids, a.ActivityDetails.InstanceID)
		}
		total += len(msg.Response.Data.Activities)
	}
	return ids
}

func getWeaponName(id int64) string {
	rows, err := destinyDB.Query("select json from DestinyInventoryItemDefinition where id + 4294967296 = ? or id = ?", id, id)
	if err != nil {
		// TODO(wcn): it returns non-actionable errors in some cases. I need to revisit this.
		//log.Fatal("Query error: " + err.Error())
	}

	for rows.Next() {
		var data []byte
		err := rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
		var dst map[string]interface{}
		err = json.Unmarshal(data, &dst)
		name := fmt.Sprintf("%s", dst["itemName"])
		// Getting empty names is a good indicator the DB is out of date. I noticed
		// this using a content snapshot that didn't have Dawning content in it.
		if name == "" {
			return fmt.Sprintf("WAT:%d", id)
		}
		return name
	}
	return fmt.Sprintf("UNKNOWN:%d", id)
}

func main() {
	// TODO(wcn): check the versioning, download new version if it exists.
	db, err := sql.Open("sqlite3", "./destiny.content")
	if err != nil {
		logger.Fatal(err)
	}
	destinyDB = db

	logger = log.New(os.Stderr, "logger: ", log.Lshortfile|log.Ltime|log.Lmicroseconds)

	logger.Printf("Arg length: %d", len(os.Args))
	logger.Printf("Fetching info for %s", os.Args[1])
	membershipID := getBungieIDForGamerTag(os.Args[1])
	var characterID string
	if len(os.Args) == 3 {
		characterID = getCharacter(membershipID, forClass(os.Args[2]))
	} else {
		characterID = getCharacter(membershipID, getLastUsed)
	}

	ids := getHistoryIDs(membershipID, characterID)
	idsChan := make(chan int64)
	go func() {
		for _, id := range ids {
			num, _ := strconv.Atoi(id)
			idsChan <- int64(num)
		}
		close(idsChan)
	}()

	out := getReports(20, idsChan)

	kills := make(map[int64]int64)
	precisionKills := make(map[int64]int64)
	for x := range out {
		var msg bungie.PGCR
		err = json.Unmarshal(x, &msg)
		if err != nil {
			logger.Fatal(err)
		}

		for _, c := range msg.Response.Data.Entries {
			if c.CharacterID != characterID {
				continue
			}
			for _, w := range c.Extended.Weapons {
				kills[w.ReferenceID] += int64(w.Values.UniqueWeaponKills.Basic.Value)
				precisionKills[w.ReferenceID] += int64(w.Values.UniqueWeaponPrecisionKills.Basic.Value)
			}
		}
	}

	for w, c := range kills {
		fmt.Printf("%s,%d,%d,%.2f%%\n", getWeaponName(w), c, precisionKills[w], float64(100*precisionKills[w])/float64(c))
	}
}

func getReports(workers int, ids <-chan int64) <-chan []byte {
	out := make(chan []byte)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for id := range ids {
				out <- getPGCR(id)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
