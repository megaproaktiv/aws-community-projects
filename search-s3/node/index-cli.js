const listAllKeys = require('./count')

var AWS = require( "aws-sdk");
AWS.config.update({
    region: "eu-west-1"
});
const s3Client = new AWS.S3({apiVersion: '2006-03-01'});
const bucket="letsbuild-cli";
var params = {
    Bucket: bucket
};

count = 0;

async function  doit(){
    await listAllKeys(s3Client, params);
    console.log(count)
}

doit()
