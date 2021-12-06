package main

import (
	"analyse"
	mylambda "analyse/lambda"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

var client *lambda.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = lambda.NewFromConfig(cfg)
}

func main() {
	names, err := mylambda.GetFunctioNames()
	if err != nil {
		fmt.Print(err)
		panic("count to ten, then ")
	}
	fmt.Printf("%v,%v\n", "Lang", "Duration")
	for i := 0; i < 10; i++ {
		for _, name := range names {

			report, err := callFunction(name)
			report = analyse.ExtractSpeedTest(report).DurationMilliSeconds
			if err != nil {
				log.Fatal("Error: ", err)
			}
			fmt.Printf("%v,%v\n", *name, *report)
		}
		r := rand.Intn(100)
		time.Sleep(time.Duration(r) * time.Microsecond)
	}
}

func callFunction(name *string) (*string, error) {
	parms := &lambda.InvokeInput{
		FunctionName: name,
		LogType:      types.LogTypeTail,
		Payload:      []byte{},
	}
	resp, err := client.Invoke(context.TODO(), parms)
	if err != nil {
		fmt.Println("Error call l: ", err)
		return nil, err
	}
	sDec, err := base64.StdEncoding.DecodeString(*resp.LogResult)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return nil, err
	}
	// Report is last line
	lines := strings.Split(string(sDec), "\n")
	var report string
	for _, l := range lines {
		if strings.HasPrefix(l, "REPORT") {
			report = l
		}
	}
	return &report, nil
}
