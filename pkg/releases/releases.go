package releases

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	guuid "github.com/google/uuid"
)

type ReleaseType uint8
const (
	Single 		ReleaseType = iota
	Draft 		ReleaseType = iota
	Remix  		ReleaseType = iota
	ReRelease ReleaseType = iota
)

type DropboxCredentials struct {
	APP_KEY    string 
	APP_SECRET string
}

type Release struct {
	ID 						int32  `json:"id"`
	UID 					string `json:"uid"`
	CID 					string `json:"cid" validate:"omitempty,min=10"`
	Name 					string `json:"name,omitempty" validate:"omitempty,min=2"`
	Artist 				string `json:"name,omitempty" validate:"omitempty,min=2"`
	Credits 			string `json:"name,omitempty" validate:"omitempty,min=2"`
	Formats				string `json:"formats,omitempty"`
	Label					string `json:"label,omitempty"`
	Thumbnail			string `json:"label,omitempty"`
	ReleaseDate		string `json:"release_date,omitempty"`
	RemixOf				string `json:"remix_of,omitempty"`
	Genres				string `json:"genres,omitempty"`
	Barcode				string `json:"barcode,omitempty"`
	CatNum				string `json:"catalog_number,omitempty"`
	Country				string `json:"country,omitempty"`
	Downloadable  bool `json:"downloadable,omitempty"`
	IsAiGenerated bool `json:"is_ai_generated,omitempty"`
}

func releases(credentials DropboxCredentials) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     credentials.APP_KEY,
		ClientSecret: credentials.APP_SECRET,
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/o/oauth2/auth",
			TokenURL: "https://provider.com/o/oauth2/token",
		},
	}

	// use PKCE to protect against CSRF attacks
	// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
	verifier := oauth2.GenerateVerifier()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")
}