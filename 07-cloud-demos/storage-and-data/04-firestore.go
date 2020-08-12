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
	"google.golang.org/api/option"
)

// [END fs_initialize]
func createClient(ctx context.Context) *firestore.Client {
	jsonPath := "creds.json"

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

type City struct {
	Name       string   `firestore:"name,omitempty"`
	State      string   `firestore:"state,omitempty"`
	Country    string   `firestore:"country,omitempty"`
	Capital    bool     `firestore:"capital,omitempty"`
	Population int64    `firestore:"population,omitempty"`
	Regions    []string `firestore:"regions,omitempty"`
}

func main() {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	cities := []struct {
		id string
		c  City
	}{
		{
			id: "SF",
			c: City{Name: "San Francisco", State: "CA", Country: "USA",
				Capital: false, Population: 860000,
				Regions: []string{"west_coast", "norcal"}},
		},
		{
			id: "LA",
			c: City{Name: "Los Angeles", State: "CA", Country: "USA",
				Capital: false, Population: 3900000,
				Regions: []string{"west_coast", "socal"}},
		},
		{
			id: "DC",
			c: City{Name: "Washington D.C.", Country: "USA",
				Capital: false, Population: 680000,
				Regions: []string{"east_coast"}},
		},
		{
			id: "TOK",
			c: City{Name: "Tokyo", Country: "Japan",
				Capital: true, Population: 9000000,
				Regions: []string{"kanto", "honshu"}},
		},
		{
			id: "BJ",
			c: City{Name: "Beijing", Country: "China",
				Capital: true, Population: 21500000,
				Regions: []string{"jingjinji", "hebei"}},
		},
	}
	for _, c := range cities {
		if _, err := client.Collection("cities").Doc(c.id).Set(ctx, c.c); err != nil {
			fmt.Println("Error ", err)
		}
	}

}
