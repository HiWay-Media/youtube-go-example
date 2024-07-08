package main

import (
	"log"

	"google.golang.org/api/youtube/v3"
)

func main() {

	//appConfig := env.GetEnvConfig()

	client := getClient(youtube.YoutubeUploadScope)

	service, err := youtube.New(client)

	// Create a new YouTube service
	//service, err := youtube.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, tok)))
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
	call := service.Videos.Insert([]string{"snippet", "status"}, upload)
	log.Println("Uploading video...", call)
}
