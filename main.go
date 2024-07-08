package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HiWay-Media/youtube-go-example/env"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

func main() {

	config := env.GetEnvConfig()
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     config.ClientID,     // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		ClientSecret: config.ClientSecret, // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/o/oauth2/auth",
			TokenURL: "https://provider.com/o/oauth2/token",
		},
		Scopes: []string{youtube.YoutubeUploadScope},
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

	service, err := youtube.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test Title",
			Description: "Test Description", // can not use non-alpha-numeric characters
			CategoryId:  "22",
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	upload.Snippet.Tags = []string{"test", "upload", "api"}
	log.Println("Uploading video...", service)
	//call := service.Videos.Insert([]string{"snippet", "status"}, upload)
}
