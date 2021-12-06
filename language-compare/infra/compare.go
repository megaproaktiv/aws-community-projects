package compare

import (
	// Go base
	"log"
	"os"

	//more
	"path/filepath"

	//sdk
	"github.com/aws/aws-sdk-go-v2/aws"

	//cdk
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	logs "github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	ssm "github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	python "github.com/aws/aws-cdk-go/awscdklambdapythonalpha/v2"

)

type CompareStackProps struct {
	awscdk.StackProps
}

func NewCompareStack(scope constructs.Construct, id string, props *CompareStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// ************ GO ************
	lambdaGOPath := filepath.Join(path, "../lambda/go/dist/main.zip")
	fnGO := lambda.NewFunction(stack, aws.String("simplelambda"),
	&lambda.FunctionProps{
		Description:                  aws.String("go lambda fibonacci"),
		FunctionName:                 aws.String("compare-go"),
		LogRetention:                 logs.RetentionDays_THREE_MONTHS,
		MemorySize:                   aws.Float64(1024),
		Timeout:                      awscdk.Duration_Seconds(aws.Float64(10)),
		Code: lambda.Code_FromAsset(&lambdaGOPath, &awss3assets.AssetOptions{}),
		Handler: aws.String("main"),
		Runtime:                      lambda.Runtime_GO_1_X(),

	})
	
	ssm.NewStringParameter(stack, aws.String("gofn"), &ssm.StringParameterProps{
		Description:    aws.String("go compare fn"),
		ParameterName:  aws.String("/compare/go"),
		StringValue:    fnGO.FunctionName(),
	})
	// ************ GO ************
	
	// *********** Java ***********
	lambdaJDKPath := filepath.Join(path, "../lambda/java/")

	buildCommands := &[]*string{
		aws.String("/bin/sh"),
		aws.String("-c"),
		aws.String("cd ./FunctionOne && mvn clean install  && cp /asset-input/FunctionOne/target/functionone.jar /asset-output/"),
		
	}
	fnJava := lambda.NewFunction(stack, aws.String("comparejdk"),
	&lambda.FunctionProps{
		Runtime:                      lambda.Runtime_JAVA_11(),
		Code: lambda.Code_FromAsset(&lambdaJDKPath, &awss3assets.AssetOptions{
			Bundling:       &awscdk.BundlingOptions{
				Image: lambda.Runtime_JAVA_11().BundlingImage(),
				Command:          buildCommands,
				User: aws.String("root"),
				OutputType: awscdk.BundlingOutput_ARCHIVED,
			},
		}),
		Handler: aws.String("compare.App"),
		Description:                  aws.String("jdk lambda fibonacci"),
		FunctionName:                 aws.String("compare-jdk"),
		LogRetention:                 logs.RetentionDays_THREE_MONTHS,
		MemorySize:                   aws.Float64(1024),
		Timeout:                      awscdk.Duration_Seconds(aws.Float64(10)),
	
	})
	
	ssm.NewStringParameter(stack, aws.String("jdkfn"), &ssm.StringParameterProps{
		Description:    aws.String("jdk compare fn"),
		ParameterName:  aws.String("/compare/jdk"),
		StringValue:    fnJava.FunctionName(),
	})
	// *********** Java ***********
	
	// *********** Python ***********
	lambdaPyPath := filepath.Join(path, "../lambda/python3")
	
	fnPy := python.NewPythonFunction(stack, aws.String("python"),
		&python.PythonFunctionProps{
			Entry: &lambdaPyPath,
			Description:                  aws.String("Python fibonaci"),
			FunctionName:                 aws.String("compare-py"),
			LogRetention:                  logs.RetentionDays_THREE_MONTHS,
			MemorySize:                   aws.Float64(1024),
			Timeout:                     awscdk.Duration_Seconds(aws.Float64(10)),
			Runtime:                      lambda.Runtime_PYTHON_3_7(),
		})
	ssm.NewStringParameter(stack, aws.String("pyfn"), &ssm.StringParameterProps{
		Description:    aws.String("py compare fn"),
		ParameterName:  aws.String("/compare/py"),
		StringValue:    fnPy.FunctionName(),
	})
	// *********** Python ***********

	
	// *********** Node.JS ***********
	lambdaNodePath := filepath.Join(path, "../lambda/nodejs")

	fnNode := lambda.NewFunction(stack, aws.String("nodefn"), 
		&lambda.FunctionProps{
			Runtime: lambda.Runtime_NODEJS_14_X(),
			FunctionName: aws.String("compare-node"),
			Description: aws.String("Compare - Node"),
			Code: lambda.AssetCode_FromAsset(&lambdaNodePath,nil),
			Handler: aws.String("index.handler"),
			Timeout:                     awscdk.Duration_Seconds(aws.Float64(10)),

	})

	ssm.NewStringParameter(stack, aws.String("nodefnout"), &ssm.StringParameterProps{
		Description:    aws.String("node compare fn"),
		ParameterName:  aws.String("/compare/node"),
		StringValue:    fnNode.FunctionName(),
	})
	// *********** Node.JS ***********





	return stack
}

