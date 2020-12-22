package letsbuild13

import (
	"context"
	// "errors"
	// "fmt"
	// "github.com/awslabs/smithy-go"
	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

//go:generate moq -out counter_moq_test.go . CounterInterface

// CounterInterface Interface for counting CloudFormation
type CounterInterface interface {
	DescribeStacks(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error)
}

// Count counts the number of Stacks in the current account
func Count(client CounterInterface) (int){
	input := &cloudformation.DescribeStacksInput{}

	resp, _ := client.DescribeStacks(context.TODO(), input)
	count := len(resp.Stacks)
	return count
}


