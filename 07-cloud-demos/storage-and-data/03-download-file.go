package main

// [START storage_list_files]
import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	jsonPath := "creds.json"

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bucket := "bucketName"
	fileName := "customers.json"
	rc, err := client.Bucket(bucket).Object(fileName).NewReader(ctx)
	if err != nil {
		fmt.Errorf("Object(%q).NewReader: %v", fileName, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	fmt.Println("Blob %v downloaded.\n", fileName)

	err = ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		fmt.Println("Error writing to a local file")
	}
}
