// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
)

var counter = 1

const bucket = "letsbuild-cli"

func main() {
	bucketName := bucket
	bucketp := &bucketName
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "eu-west-1"
	})
	input := &s3.ListObjectsV2Input{
		Bucket: bucketp,
	}
	counter := listAllKeys(client, input)
	fmt.Println(counter)
}
func listAllKeys(client *s3.Client, input *s3.ListObjectsV2Input) int {
	resp, err := client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		return 0
	}
	lastname := ""
	for _, item := range resp.Contents {
		name := *item.Key
		endsWith := strings.HasSuffix(name, "LICENSE") // true
		if endsWith {
			// fmt.Println("Name:", *item.Key)
			counter = counter + 1
		}
		lastname = name
	}
	if resp.IsTruncated {
		bucketName := bucket
		bucketp := &bucketName
		input := &s3.ListObjectsV2Input{
			Bucket:            bucketp,
			ContinuationToken: resp.ContinuationToken,
			StartAfter:        aws.String(lastname),
		}
		listAllKeys(client, input)
	}
	return counter
}