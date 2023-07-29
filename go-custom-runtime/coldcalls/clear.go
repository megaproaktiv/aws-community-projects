package coldcalls

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// Clear all CLoudWatch Log entries from a lambda

func ClearLogs(functionName string) {
	
	// Load AWS session from shared config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS SDK configuration:", err)
		return
	}

	// Create a CloudWatch Logs client
	client := cloudwatchlogs.NewFromConfig(cfg)

	// Get the log group name for the Lambda function
	logGroupName := fmt.Sprintf("/aws/lambda/%s", functionName)

	// Describe the log streams in the log group
	describeInput := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	}

	// Retrieve the log streams associated with the log group
	describeResp, err := client.DescribeLogStreams(context.Background(), describeInput)
	if err != nil {
		fmt.Println("Error describing log streams:", err)
		return
	}

	// Clear the log events from each log stream
	for _, logStream := range describeResp.LogStreams {
		clearInput := &cloudwatchlogs.DeleteLogStreamInput{
			LogGroupName:  aws.String(logGroupName),
			LogStreamName: logStream.LogStreamName,
		}

		_, err = client.DeleteLogStream(context.Background(), clearInput)
		if err != nil {
			fmt.Println("Error clearing log events for log stream:", *logStream.LogStreamName, "-", err)
			return
		}

		fmt.Println("Cleared log events for log stream:", *logStream.LogStreamName)
	}

	fmt.Println("All CloudWatch log entries for the Lambda function have been cleared.")
}
