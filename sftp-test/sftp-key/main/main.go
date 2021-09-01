package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"sftpkey"
)


func main() {
	app := awscdk.NewApp(nil)

	sftpkey.NewSftpKeyStack(app, "SftpKeyStack", &sftpkey.SftpKeyStackProps{})

	app.Synth(nil)
}
