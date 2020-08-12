package main

// [START storage_list_files]
import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// listFiles lists objects within specified bucket.
// func listFiles(w io.Writer, bucket string) error {
// 	// bucket := "bucket-name"
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		return fmt.Errorf("storage.NewClient: %v", err)
// 	}

// }

// [END storage_list_files]

func main() {
	jsonPath := "creds.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucket := "bucketId"

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}
		fmt.Println(attrs.Name)
	}
}
