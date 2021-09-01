package main

import (
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	jsii "github.com/aws/jsii-runtime-go"
	"sns"
	"os"
)

func main() {
	app := awscdk.NewApp(nil)

	sns.NewCdkSnsGoStack(app, "sns", &sns.CdkSnsGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {

	return &awscdk.Environment{
	 Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
