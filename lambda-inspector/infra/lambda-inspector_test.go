package lambdainspector_test

import (
	"lambdainspector"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/jsii-runtime-go"
)

func TestLambdaInspectorStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := lambdainspector.NewLambdaInspectorStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.ResourceCountIs(jsii.String("AWS::Lambda::Function"), aws.Float64(3))
}
