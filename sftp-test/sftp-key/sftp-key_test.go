package sftpkey_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"gotest.tools/assert"

	"github.com/tidwall/gjson"

	"sftpkey"
)

func TestSftpKeyStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := sftpkey.NewSftpKeyStack(app, "MyStack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	// Bucket
	deletionPolicy := template.Get("Resources.sftpbucket219BD667.DeletionPolicy").String()
	assert.Equal(t, "Retain", deletionPolicy)

	// Policy
	version := template.Get("Resources.portalCertificateCertificateRequestorFunctionServiceRoleDefaultPolicy692C9C94.Properties.Version").String()
	assert.Equal(t,"2012-10-17",version)

}
