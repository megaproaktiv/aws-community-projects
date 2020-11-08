package main

import(
	"letsbuild13"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func main(){
	cfg, err := config.LoadDefaultConfig(config.WithRegion("eu-central-1"))
    if err != nil {
        panic("unable to load SDK config, " + err.Error())
	}
	
	client := cloudformation.NewFromConfig(cfg);

	count := letsbuild13.Count(client);

	fmt.Println("Counting CloudFormation Stacks: ",count)
}