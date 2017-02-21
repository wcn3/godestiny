package bungieoauth

// bungieoauth implements the Bungie Oauth2-style authentication flow in a way
// that is drop-in compatible with the standard golang.org/x/oauth2 code and
// can use custom TokenSources designed for that library. Bungie's official documentation
// for this API is described here: https://www.bungie.net/en/Help/Article/45481
import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// Bungie API response that contains OAuth tokens.
type bungieOAuth2Response struct {
	ErrorCode   int64    `json:"ErrorCode"`
	ErrorStatus string   `json:"ErrorStatus"`
	Message     string   `json:"Message"`
	MessageData struct{} `json:"MessageData"`
	Response    struct {
		AccessToken struct {
			Expires int64  `json:"expires"`
			Readyin int64  `json:"readyin"`
			Value   string `json:"value"`
		} `json:"accessToken"`
		RefreshToken struct {
			Expires int64  `json:"expires"`
			Readyin int64  `json:"readyin"`
			Value   string `json:"value"`
		} `json:"refreshToken"`
		Scope int64 `json:"scope"`
	} `json:"Response"`
	ThrottleSeconds int64 `json:"ThrottleSeconds"`
}

// BungieConfig wraps an oauth2.Config to allow use of that library.
type BungieConfig struct {
	cfg oauth2.Config
}

// NewBungieConfig creates a configuration using the supplied URL and API key.
func NewBungieConfig(authURL, apiKey string) *BungieConfig {
	return &BungieConfig{
		cfg: oauth2.Config{
			ClientID:     apiKey,
			ClientSecret: "",
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: "https://www.bungie.net/Platform/App/GetAccessTokensFromCode/",
			},
		},
	}
}

// AuthCodeURL builds the URL for authorizing an API, using the supplied state value.
func (b *BungieConfig) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	// TODO(wcn): filter out options we don't care about.
	return b.cfg.AuthCodeURL(state, opts...)
}

func getTokenFromRequest(req *http.Request, apiKey string) (*oauth2.Token, error) {
	req.Header.Add("Content-Type", "application/json; charset=UTF-8;")
	req.Header.Set("X-API-Key", apiKey)
	hc := http.DefaultClient
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bungieoauth: HTTP status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("bungieoauth: cannot fetch token: %v", err)
	}

	var oauthResp bungieOAuth2Response

	err = json.Unmarshal(body, &oauthResp)
	if err != nil {
		return nil, fmt.Errorf("bungieoauth: cannot decode token: %v", err)
	}

	// Create the oauth2.Token from the Bungie API return values, filling in data
	// to match the spec. We can't really do anything with the refresh token expiration
	// in the library. When it expires, the API needs to be reauthorized, which is what
	// will happen, so dropping it is OK.
	token := &oauth2.Token{
		AccessToken:  oauthResp.Response.AccessToken.Value,
		RefreshToken: oauthResp.Response.RefreshToken.Value,
		Expiry:       time.Now().Add(time.Duration(oauthResp.Response.AccessToken.Expires) * time.Second),
		TokenType:    "Bearer",
	}

	return token, nil
}

// Exchange converts an authorization code into a token.
//
// It is used after a resource provider redirects the user back to the
// Redirect URI (the URL obtained from AuthCodeURL).
func (b *BungieConfig) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	req, err := http.NewRequest("POST", "https://www.bungie.net/Platform/App/GetAccessTokensFromCode/", strings.NewReader(fmt.Sprintf("{\"code\":\"%s\"}", code)))
	if err != nil {
		return nil, err
	}

	return getTokenFromRequest(req, b.cfg.ClientID)
}

type refreshTokenSource struct {
	t      *oauth2.Token
	apiKey string
	mu     sync.Mutex
}

func newTokenSource(t *oauth2.Token, apiKey string) refreshTokenSource {
	return refreshTokenSource{
		t:      t,
		apiKey: apiKey,
	}
}

func (r refreshTokenSource) Token() (*oauth2.Token, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.t.Valid() {
		return r.t, nil
	}

	// Use the refresh token to get a new access token.
	req, err := http.NewRequest("POST", "https://www.bungie.net/Platform/App/GetAccessTokensFromRefreshToken/", strings.NewReader(fmt.Sprintf("{\"refreshToken\":\"%s\"}", r.t.RefreshToken)))
	if err != nil {
		return nil, err
	}

	return getTokenFromRequest(req, r.apiKey)
}

// TokenSource returns a TokenSource that refreshes tokens using the Bungie API.
func (b *BungieConfig) TokenSource(t *oauth2.Token) oauth2.TokenSource {
	return newTokenSource(t, b.cfg.ClientID)
}
