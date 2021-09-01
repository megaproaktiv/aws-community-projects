
var allKeys = [];

async function listAllKeys(client, params ) {
    client.listObjectsV2(params, function (err, data) {
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

module.exports = listAllKeys;
