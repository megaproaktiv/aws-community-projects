package lambda

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	
)

var client *ssm.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = ssm.NewFromConfig(cfg)
}

func GetFunctioNames() ([]*string,error){
	names := make([]*string,0)
	var parms *ssm.GetParameterInput

	const pre = "/compare/"

	for _, name := range []string{ "jdk", "py", "node", "go"} {
		parmName := pre+ name 
		parms = &ssm.GetParameterInput{
			Name:           aws.String(parmName),
		}
		resp, err := client.GetParameter(context.TODO(), parms)
		if err != nil {
			return nil, err
		}
		names = append(names, resp.Parameter.Value)
	}
	return names, nil
}
