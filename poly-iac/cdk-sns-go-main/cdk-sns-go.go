package sns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkSnsGoStackProps struct {
	awscdk.StackProps
}

func NewCdkSnsGoStack(scope constructs.Construct, id string, props *CdkSnsGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// as an example, here's how you would define an AWS SNS topic:
	awssns.NewTopic(stack, jsii.String("MyTopic"), &awssns.TopicProps{
		DisplayName: jsii.String("MyCoolTopic"),
	})

	return stack
}
