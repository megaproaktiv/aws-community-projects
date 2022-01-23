package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-sdk-go-v2/aws"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	// "github.com/aws/jsii-runtime-go"
	"github.com/aws/jsii-runtime-go"
	"github.com/awslabs/goformation/v5/cloudformation/s3"
)

type S3AnalyticsStackProps struct {
	awscdk.StackProps
}

func NewS3AnalyticsStack(scope constructs.Construct, id string, props *S3AnalyticsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// **************** Use goformation type

	// *************

	helper := awss3.NewBucket(stack, aws.String("helper"), nil)
	
	myBucket := awss3.NewBucket(stack, aws.String("BucketsWithAnalyticsConfig"), &awss3.BucketProps{
		BlockPublicAccess:      awss3.BlockPublicAccess_BLOCK_ALL(),
	})
	var cfnBucketGoformation awss3.CfnBucket

	jsii.Get(myBucket.Node(), "defaultChild", &cfnBucketGoformation)

	var analyticsConfigurationFromFile  []s3.Bucket_AnalyticsConfiguration
	
	data, err := os.ReadFile("analyticsconfig.json")
	if err != nil {
			fmt.Println("Cant read json data: ", err)
	}
	json.Unmarshal(data, &analyticsConfigurationFromFile)
	if err != nil {
			fmt.Println("JSON unmarshall error data: ", err)
	}

	analyticsConfigurationFromFile[0].StorageClassAnalysis.DataExport.Destination.BucketArn = *helper.BucketArn()
	cfnBucketGoformation.AddPropertyOverride(aws.String("AnalyticsConfigurations"),analyticsConfigurationFromFile)

	// *************


	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewS3AnalyticsStack(app, "S3AnalyticsStack", &S3AnalyticsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
