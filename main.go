package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/youtube/v3"
)

var (
	filename = flag.String("filename", "BigBuckBunny_320x180.mp4", "Name of video file to upload")
)

func main() {

	//appConfig := env.GetEnvConfig()

	flag.Parse()

	if *filename == "" {
		log.Fatalf("You must provide a filename of a video file to upload")
	}

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

	file, err := os.Open(*filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", *filename, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}
