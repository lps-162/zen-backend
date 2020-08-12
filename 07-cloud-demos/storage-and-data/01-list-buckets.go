package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// explicit reads credentials from the specified path.
func explicit(jsonPath, projectID string) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Buckets:")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(battrs.Name)
	}
}

func main() {
	filePath := "creds.json"
	projectID := "xxxxxxxxx"
	explicit(filePath, projectID)
}
