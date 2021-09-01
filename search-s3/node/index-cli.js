const listAllKeys = require('./count')

var AWS = require( "aws-sdk");
AWS.config.update({
    region: "eu-west-1"
});
var s3 = new AWS.S3();
const bucket="letsbuild-cli";
var params = {
    Bucket: bucket
};

count = 0;

async function  doit(){
    await listAllKeys(s3, params);
    console.log(count)
}

doit()
