package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/wcn3/godestiny/bungieoauth"

	"golang.org/x/oauth2"
)

var apiKey string
var authURL string
var cannedState = "an_unlikely_string_to_guess"
var tokenURL = "https://www.bungie.net/Platform/App/GetAccessTokensFromCode/"
var bcfg *bungieoauth.BungieConfig

type outhClientKeyType int

const oauthClientKey outhClientKeyType = 0

func newContextWithOauthClient(ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, error) {
	ts := fileTokenSerializer{filename: "/tmp/token.txt"}
	tok, err := ts.Read()
	if err != nil {
		return nil, err
	}
	oauthClient := oauth2.NewClient(context.Background(), serializedTokenSource{ts, bcfg.TokenSource(tok)})
	return context.WithValue(ctx, oauthClientKey, oauthClient), nil
}

func oauthClientFromContext(ctx context.Context) *http.Client {
	return ctx.Value(oauthClientKey).(*http.Client)
}

func oauth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx, err := newContextWithOauthClient(req.Context(), rw, req)
		if err != nil {
			http.Redirect(rw, req, bcfg.AuthCodeURL(cannedState+"&dest="+req.URL.Path), http.StatusFound)
			return
		}
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func pgcr() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		num := strings.Replace(req.URL.Path, "/pgcr/", "", 1)
		num = strings.Replace(num, "http:/guardian.gg/en/pgcr/", "", 1)
		url := "https://www.bungie.net/platform/Destiny/Stats/PostGameCarnageReport/" + num
		bungieHandler(rw, req, url)
	}
}

func vault() http.Handler {
	return oauth(forURL("https://www.bungie.net/platform/Destiny/1/MyAccount/Vault/Summary/"))
}

func forURL(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bungieHandler(w, r, url)
	}
}

func bungieHandler(rw http.ResponseWriter, req *http.Request, url string) {
	oauthClient := oauthClientFromContext(req.Context())
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("X-API-Key", apiKey)
	resp, err := oauthClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		log.Fatal(err)
	}

	rw.Write(body)
}

func extractCodeWithTokenSerializer(fts fileTokenSerializer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		code := req.URL.Query().Get("code")
		state := req.URL.Query().Get("state")
		dest := req.URL.Query().Get("dest")
		if state != cannedState {
			panic("Got MITMed")
		}

		tok, err := bcfg.Exchange(context.Background(), code)
		if err != nil {
			log.Fatal(err)
		}

		fts.Write(tok)
		http.Redirect(w, req, dest, http.StatusFound)
	}
}

func main() {
	data, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		log.Fatal(err)
	}
	apiKey = string(data)

	data, err = ioutil.ReadFile("auth_url.txt")
	if err != nil {
		log.Fatal(err)
	}
	authURL = string(data)

	bcfg = bungieoauth.NewBungieConfig(authURL, apiKey)
	ts := fileTokenSerializer{filename: "/tmp/token.txt"}

	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/bungieauth", extractCodeWithTokenSerializer(ts))
	http.HandleFunc("/ua", unauthed)
	http.Handle("/vault", vault())
	http.Handle("/user", oauth(forURL("https://www.bungie.net/platform/User/GetCurrentBungieAccount/")))
	http.Handle("/definitions", oauth(forURL("https://www.bungie.net/platform/Destiny/Stats/Definition/")))
	http.Handle("/stats", oauth(forURL("https://www.bungie.net/platform/Destiny/Stats/1/4611686018433876772/2305843009395074356/?modes=5&groups=0,1,2,3")))
	http.Handle("/genstats", oauth(forURL("https://www.bungie.net/platform/Destiny/Stats/Account/1/4611686018433876772/?modes=5&groups=0,1,2,3")))
	http.Handle("/pgcr/", oauth(pgcr()))
	http.Handle("/activities", oauth(forURL("https://www.bungie.net/platform/Destiny/Stats/ActivityHistory/1/4611686018433876772/2305843009395074356/?mode=None&page=0&count=250&definitions=true")))
	fmt.Println("Started!")

	// The Bungie API requires that the Redirect URL be https. This is a sensible precaution
	// but I wish there was an exception for localhost. Since there isn't, we self-sign
	// certs using the process described in
	// https://gist.github.com/denji/12b3a568f092ab951456 and run https locally.
	err = http.ListenAndServeTLS(":7070", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func unauthed(w http.ResponseWriter, req *http.Request) {
	req, err := http.NewRequest("GET", "https://www.bungie.net/platform/Destiny/1/MyAccount/Vault/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-API-Key", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(prettify(body)))
}

func prettify(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

// This is something I really wish was in the Go Oauth library.
// Composing TokenSources makes a lot of sense. The outer one they
// implement is a reusable TokenSource that uses a Token until it's
// no longer valid, at which point, it calls the inner TokenSource
// to get a token.
//
// The next TokenSource should be the one that gets it from whatever
// persistent storage the app needs. It then has the behavior implemented
// below. As long as a valid Token exists in the store, it should use that.
// However, if it has to call its TokenSource to get a new token, then
// the persistent store should be updated to record the new token.
// Subsequent calls would then consult the store and get the most recent
// token.
//
// I've captured this concept with the TokenSerializer interface and the
// serializedTokenSource uses this interface to implement the logic
// described above. This code uses a simple file to handle a single token.
// A more elaborate implementation could support multiple users by associating
// tokens with login credentials. The TokenSource doesn't know or care as
// long as the interface is satisfied.
//
// TODO(wcn): pitch this to the Go oauth2 maintainers.
type fileTokenSerializer struct {
	filename string
	mu       sync.Mutex
}

func (f fileTokenSerializer) Write(tok *oauth2.Token) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	file, err := os.Create(f.filename)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(file)
	return enc.Encode(tok)
}

func (f fileTokenSerializer) Read() (*oauth2.Token, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	file, err := os.Open(f.filename)
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	dec := json.NewDecoder(file)
	if err = dec.Decode(tok); err != nil {
		return nil, err
	}
	return tok, nil
}

// TokenSerializer defines methods for loading and storing an OAuth2 token.
type TokenSerializer interface {
	Write(*oauth2.Token) error
	Read() (*oauth2.Token, error)
}

type serializedTokenSource struct {
	store TokenSerializer
	new   oauth2.TokenSource
}

func (s serializedTokenSource) Token() (*oauth2.Token, error) {
	if tok, err := s.store.Read(); err == nil && tok.Valid() {
		return tok, err
	}

	tok, err := s.new.Token()
	if err != nil {
		return nil, err
	}

	err = s.store.Write(tok)
	if err != nil {
		return nil, err
	}
	return tok, err
}
