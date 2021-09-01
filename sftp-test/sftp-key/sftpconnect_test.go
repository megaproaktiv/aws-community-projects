package sftpkey_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"gotest.tools/assert"

	// Parameter
	paddle "github.com/PaddleHQ/go-aws-ssm"
)

var params *paddle.Parameters

type Config struct {
	addr string
}

var configuration = Config{}

func init() {
	pmstore, err := paddle.NewParameterStore()
	if err != nil {
		log.Fatal("Cant connect to Parameter Store")
	}
	//Requesting the base path
	params, err = pmstore.GetAllParametersByPath("/sftp/", true)
	if err != nil {
		log.Fatal("Cant get Parameter Store")
	}
	configuration.addr = params.GetValueByName("serverdns") + ":22"
}

// Update know_hosts with
// ssh-keyscan -H sftp.test.epm.hrdirekt.de >>known_hosts
const userWithReadRights = "gernot"
const privateKey = "/Users/silberkopf/.ssh/id_rsa"

func TestWrite(t *testing.T) {
	addr := configuration.addr

	t.Log("Test sftp server ", addr)
	key, err := ioutil.ReadFile(privateKey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: userWithReadRights,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		panic("Failed to create client: " + err.Error())
	}

	dstFile, err := client.OpenFile("./DUMMY.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	assert.NilError(t, err, "File write should work without error")
	if err != nil {
		log.Fatal(err)
	}
	dstFile.Close()

	err = client.Remove("./DUMMY.txt")
	assert.NilError(t, err, "File delete should work without error")

	// Close connection
	defer client.Close()
}
