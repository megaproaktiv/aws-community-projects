var AWS = require( "aws-sdk");
AWS.config.update({
    region: "eu-west-1"
});
var s3 = new AWS.S3();
const bucket="letsbuild-cli";
var params = {
    Bucket: bucket
};
var allKeys = [];
count = 0;
listAllKeys();

function listAllKeys() {
    s3.listObjectsV2(params, function (err, data) {
        if (err) {
            console.log(err, err.stack); // an error occurred
        } else {
            var contents = data.Contents;
            contents.forEach(function (content) {
                let name= content.Key
                if( name.endsWith("LICENSE")){
                    allKeys.push(content.Key);
                    count = count +1
                }
            });
            if (data.IsTruncated) {
                params.ContinuationToken = data.NextContinuationToken;
                listAllKeys();
            }else{
              console.log("Count:" , count)
            }
        }
    });
}
