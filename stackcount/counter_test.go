package letsbuild13_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"letsbuild13"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/stretchr/testify/assert"
)

func TestCountStacks(t *testing.T) {
	expectedValues := 2;

	// bytes := []byte(str_emp)
	// var res Response
	// json.Unmarshal(bytes, &res)

	mockedCounterInterface := &letsbuild13.CounterInterfaceMock{
            DescribeStacksFunc: func(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error) {
				var cloudformationOutput cloudformation.DescribeStacksOutput
				// Read json file
				data, err := ioutil.ReadFile("test/cloudformation.json")
				if err != nil {
					fmt.Println("File reading error", err)
				}
				json.Unmarshal(data, &cloudformationOutput);
				return &cloudformationOutput,nil;

            },
	}
	
	computedValue := letsbuild13.Count(mockedCounterInterface)

	assert.Equal(t,expectedValues, computedValue)

}