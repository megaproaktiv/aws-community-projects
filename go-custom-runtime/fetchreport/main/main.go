package main

import (
	"context"
	"fetchreport"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func main() {
	var lambdaName string

	flag.StringVar(&lambdaName, "lambda", "", "Name of the Lambda function")
	flag.Parse()

	// Check if the required parameters are provided
	if lambdaName == "" {
		fmt.Println("Error: Lambda function name is required.")
		os.Exit(1)
	}

	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS SDK configuration:", err)
		return
	}

	// Create a CloudWatch Logs client
	client := cloudwatchlogs.NewFromConfig(cfg)

	// Set the query string
	queryString := `fields @timestamp, @message, @logStream, @log
		| sort @timestamp desc
		| limit 20
		| filter @message like /REPORT/`

	// Start the query
	name := lambdaName
	queryInput := &cloudwatchlogs.StartQueryInput{
		LogGroupName: aws.String("/aws/lambda/" + name), // Replace with your CloudWatch Log Group name
		StartTime:    aws.Int64(0),                      // Set the start time for the query, 0 means all data available
		EndTime:      aws.Int64(time.Now().UnixMilli()), // Set the end time for the query to the current time
		QueryString:  aws.String(queryString),
	}

	startQueryOutput, err := client.StartQuery(context.TODO(), queryInput)
	if err != nil {
		fmt.Println("Error starting CloudWatch Logs query:", err)
		return
	}

	// Get the query results
	getQueryResultsInput := &cloudwatchlogs.GetQueryResultsInput{
		QueryId: startQueryOutput.QueryId,
	}

	maxCount := 10
	count := 1
	for {
		count++
		if count > maxCount {
			break
		}

		getQueryResultsOutput, err := client.GetQueryResults(context.TODO(), getQueryResultsInput)
		if err != nil {
			fmt.Println("Error getting CloudWatch Logs query results:", err)
			return
		}

		for _, result := range getQueryResultsOutput.Results {
			for _, field := range result {
				// Print the message
				if *field.Field == "@message" {
					line := fetchreport.Split(*field.Value)
					line.Name = name
					line.Print()
				}
			}

		}

		if getQueryResultsOutput.Status != types.QueryStatusComplete {
			// Query is still running, sleep for a while and then get the results again
			time.Sleep(2 * time.Second)
		} else {
			// Query has completed
			break
		}
	}
}
