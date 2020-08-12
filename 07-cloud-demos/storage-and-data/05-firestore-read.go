// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample firestore_quickstart demonstrates how to connect to Firestore, and add and list documents.
package main

// [START fs_initialize]
import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// [END fs_initialize]
func createClient(ctx context.Context) *firestore.Client {
	jsonPath := "creds.json"

	// Sets your Google Cloud Platform project ID.
	//projectID := "207820166423"
	projectID := "projectId"
	client, err := firestore.NewClient(ctx, projectID,
		option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func main() {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	fmt.Println("All cities:")
	iter := client.Collection("cities").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(doc.Data())
	}

}
