package lambdainspector

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	// "github.com/aws/jsii-runtime-go"
)

type LambdaInspectorStackProps struct {
	awscdk.StackProps
}

func NewLambdaInspectorStack(scope constructs.Construct, id string, props *LambdaInspectorStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	path, err := os.Getwd()
	if err != nil {
			log.Println(err)
	}

	lambdaArchitecture := awslambda.Architecture_X86_64()

	dockerfile := filepath.Join(path, "../app/go")
	awslambda.NewDockerImageFunction(stack,
		aws.String("lambdainspector-go"),
		&awslambda.DockerImageFunctionProps{
			Architecture:                 lambdaArchitecture,
			FunctionName:                 aws.String("lambdainspector-go"),
			MemorySize:                   aws.Float64(1024),
			Timeout:                      awscdk.Duration_Seconds(aws.Float64(300)),
			Code:                         awslambda.DockerImageCode_FromImageAsset(&dockerfile, &awslambda.AssetImageCodeProps{}),
		})

	dockerfile = filepath.Join(path, "../app/python")
	awslambda.NewDockerImageFunction(stack,
		aws.String("lambdainspector-py"),
		&awslambda.DockerImageFunctionProps{
			Architecture:                 lambdaArchitecture,
			FunctionName:                 aws.String("lambdainspector-py"),
			MemorySize:                   aws.Float64(1024),
			Timeout:                      awscdk.Duration_Seconds(aws.Float64(300)),
			Code:                         awslambda.DockerImageCode_FromImageAsset(&dockerfile, &awslambda.AssetImageCodeProps{}),
		})

	dockerfile = filepath.Join(path, "../app/node")
	awslambda.NewDockerImageFunction(stack,
		aws.String("lambdainspector-node"),
		&awslambda.DockerImageFunctionProps{
			Architecture:                lambdaArchitecture,
			FunctionName:                 aws.String("lambdainspector-node"),
			MemorySize:                   aws.Float64(1024),
			Timeout:                      awscdk.Duration_Seconds(aws.Float64(300)),
			Code:                         awslambda.DockerImageCode_FromImageAsset(&dockerfile, &awslambda.AssetImageCodeProps{}),
		})

	dockerfile = filepath.Join(path, "../app/java")
	awslambda.NewDockerImageFunction(stack,
		aws.String("lambdainspector-java"),
		&awslambda.DockerImageFunctionProps{
			Architecture:                lambdaArchitecture,
			FunctionName:                 aws.String("lambdainspector-java"),
			MemorySize:                   aws.Float64(2048),
			Timeout:                      awscdk.Duration_Seconds(aws.Float64(900)),
			Code:                         awslambda.DockerImageCode_FromImageAsset(&dockerfile, &awslambda.AssetImageCodeProps{}),
		})


	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewLambdaInspectorStack(app, "LambdaInspectorStack", &LambdaInspectorStackProps{
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
