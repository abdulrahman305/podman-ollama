package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func main() {
	// Create client
	url := "https://us-east-1-1.aws.cloud2.influxdata.com"
	token := os.Getenv("INFLUXDB_TOKEN")

	// Create a new client using an InfluxDB server base URL and an authentication token
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:  url,
		Token: token,
	})

	if err != nil {
		panic(err)
	}
	// Close client at the end and escalate error if present
	defer func(client *influxdb3.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	database := "buket"
	data := map[string]map[string]interface{}{
		"point1": {
			"location": "Klamath",
			"species":  "bees",
			"count":    23,
		},
		"point2": {
			"location": "Portland",
			"species":  "ants",
			"count":    30,
		},
		"point3": {
			"location": "Klamath",
			"species":  "bees",
			"count":    28,
		},
		"point4": {
			"location": "Portland",
			"species":  "ants",
			"count":    32,
		},
		"point5": {
			"location": "Klamath",
			"species":  "bees",
			"count":    29,
		},
		"point6": {
			"location": "Portland",
			"species":  "ants",
			"count":    40,
		},
	}

	// Write data
	// options := influxdb3.WriteOptions{ // Removed options
	// 	Database: database,
	// }
	for key := range data {
		// Create a new point with tags and fields
		point, err := influxdb3.NewPoint(
			"census", // Measurement name
			map[string]string{ // Tags
				"location": data[key]["location"].(string),
			},
			map[string]interface{}{ // Fields
				data[key]["species"].(string): data[key]["count"],
			},
			time.Now(), // Timestamp
		)
		if err != nil {
			panic(err)
		}

		// Write the point to InfluxDB
		writeAPI := client.WriteAPIOptions(database) // Get the WriteAPI
		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second) // separate points by 1 second
	}

	fmt.Println("Data written successfully!") // Print a success message
}
// Execute query
query := `SELECT *
          FROM 'census'
          WHERE time >= now() - interval '1 hour'
            AND ('bees' IS NOT NULL OR 'ants' IS NOT NULL)`

queryOptions := influxdb3.QueryOptions{
  Database: database,
}
iterator, err := client.QueryWithOptions(context.Background(), &queryOptions, query)

if err != nil {
  panic(err)
}

for iterator.Next() {
  value := iterator.Value()

  location := value["location"]
  ants := value["ants"]
  bees := value["bees"]
  fmt.Printf("in %s are %d ants and %d bees\n", location, ants, bees)
}