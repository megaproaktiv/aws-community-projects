// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
package main

import (
	"fmt"
	"strings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var counter = 1

const bucket = "letsbuild-cli"

func main() {
	bucketName := bucket
	bucketp := &bucketName
	
	sess, err := session.NewSession(&aws.Config{
        Region: aws.String("eu-west-1")},
    )
	if err == nil { // resp is now filled
		fmt.Println(err)
	}
	input := &s3.ListObjectsV2Input{
		Bucket: bucketp,
	}
	counter := listAllKeys(sess, input)
	fmt.Println(counter)
}
func listAllKeys(sess *session.Session, params *s3.ListObjectsV2Input) int {

	client := s3.New(sess)
	resp, err := client.ListObjectsV2( params)
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
	if *resp.IsTruncated {
		bucketName := bucket
		bucketp := &bucketName
		input := &s3.ListObjectsV2Input{
			Bucket:            bucketp,
			ContinuationToken: resp.ContinuationToken,
			StartAfter:        aws.String(lastname),
		}
		listAllKeys(sess, input)
	}
	return counter
}
