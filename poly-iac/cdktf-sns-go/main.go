package main

import (
	"cdk.tf/go/stack/generated/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/aws/jsii-runtime-go"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	aws.NewVpc(stack, jsii.String("cdktfvpcgo"), &aws.VpcConfig{

	})
	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-sns-go")

	app.Synth()
}
