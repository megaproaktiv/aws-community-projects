package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
	"coldcalls"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

var Client *lambda.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = lambda.NewFromConfig(cfg)
}

func main() {
	// Parse command-line arguments
	var lambdaName string
	var times int
	var memoryList string

	flag.StringVar(&lambdaName, "lambda", "", "Name of the Lambda function")
	flag.IntVar(&times, "times", 1, "Number of times to call the Lambda function")
	flag.StringVar(&memoryList, "memory", "1024M", "Comma-separated list of memory values")
	flag.Parse()

	// Check if the required parameters are provided
	if lambdaName == "" {
		fmt.Println("Error: Lambda function name is required.")
		os.Exit(1)
	}

	if memoryList == "" {
		fmt.Println("Error: Memory list is required.")
		os.Exit(1)
	}

	// Convert memoryList into an array of integers
	memories := make([]int, 0)
	for _, memStr := range strings.Split(memoryList, ",") {
		mem, err := strconv.Atoi(strings.TrimSuffix(memStr, "M"))
		if err != nil {
			fmt.Printf("Error converting memory value %s to integer: %s\n", memStr, err.Error())
		}
		memories = append(memories, mem)
	}

	// Clear All logs
	coldcalls.ClearLogs(lambdaName)

	// Loop over all memory parameters and invoke the Lambda function
	for _, memory := range memories {
		// Update the Lambda function's memory
		slog.Info("Update", "memory", memory)
		if err := updateLambdaMemory(Client, lambdaName, memory); err != nil {
			fmt.Printf("Error updating Lambda memory for %d MB: %s\n", memory, err.Error())
			os.Exit(1)
		}
		Wait(Client, lambdaName)

		// Call the Lambda function 'times' times with a standard event
		for i := 0; i < times; i++ {
			slog.Info("Update", "Environment", i)
			err := createEnvironmentVariable(Client, lambdaName, "count", strconv.Itoa(i))
			Wait(Client, lambdaName)
			if err != nil {
				fmt.Printf("Error creating environment variable for %d MB, invocation %d: %s\n", memory, i+1, err.Error())
			}
			slog.Info("Invoke")
			if err = invokeLambda(Client, lambdaName); err != nil {
				fmt.Printf("Error invoking Lambda for %d MB, invocation %d: %s\n", memory, i+1, err.Error())
				os.Exit(1)
			}
		}
	}
}

// updateLambdaMemory updates the memory size of the Lambda function.
func updateLambdaMemory(client *lambda.Client, lambdaName string, memory int) error {

	_, err := client.UpdateFunctionConfiguration(context.TODO(), &lambda.UpdateFunctionConfigurationInput{
		FunctionName: &lambdaName,
		MemorySize:   aws.Int32(int32(memory)),
	})
	if err != nil {
		return err
	}

	return err
}

// createEnvironmentVariable creates an environment variable in the Lambda function.
func createEnvironmentVariable(client *lambda.Client, lambdaName, key, value string) error {
	_, err := client.UpdateFunctionConfiguration(context.TODO(), &lambda.UpdateFunctionConfigurationInput{
		FunctionName: &lambdaName,
		Environment: &types.Environment{
			Variables: map[string]string{
				key: value,
			},
		},
	})
	if err != nil {
		return err
	}

	return err
}

// invokeLambda invokes the Lambda function with a standard event.
func invokeLambda(client *lambda.Client, lambdaName string) error {

	_, err := client.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: &lambdaName,
		Payload:      []byte(`{ "key": "value" }`), // Change the payload as needed
	})

	return err
}

// Wait for status of function to become active
func Wait(client *lambda.Client, lambdaName string) {
	slog.Info("Wait ")
	maxCount := 10
	// Get Lambda description
	parms := &lambda.GetFunctionInput{
		FunctionName: &lambdaName,
	}
	for i := 0; i < maxCount; i++ {
		resp, err := client.GetFunction(context.Background(), parms)
		if err != nil {
			slog.Error(err.Error(), "Unable to get function")
			return
		}
		if resp.Configuration.LastUpdateStatus == types.LastUpdateStatusSuccessful {
			slog.Info("Function is active")
			return
		}
		slog.Info("Function is not active", "Status", resp.Configuration.LastUpdateStatus)
		time.Sleep(1 * time.Second)
	}

	slog.Info("Continue ")
}
