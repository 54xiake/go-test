package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "LmXc2i3b0BlrIriie9NyC_UnfsjzAvnyLMnbb8hDQje6DMU1A9tPh_Bg2xi55v9V2nCH-lG0LPigW10qvfa4Mg=="
	const bucket = "test"
	const org = "test"

	client := influxdb2.NewClient("http://localhost:8086", token)
	// always close client at the end
	defer client.Close()

	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 22.5, 55.5))
	// Flush writes
	writeAPI.Flush()

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 66.6},
		time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 77.7).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()

	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"stat\")", bucket)
	// Get query client
	queryAPI := client.QueryAPI(org)
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}
