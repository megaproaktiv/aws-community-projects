# see https://alexwlchan.net/2019/07/listing-s3-keys/
import boto3

def get_matching_s3_objects(bucketname, prefix=""):
    """
    Generate objects in an S3 bucket.
    :param bucket: Name of the S3 bucket.
    :param prefix: Only fetch objects whose key starts with
        this prefix (optional).
    """
    s3 = boto3.client("s3",region_name="eu-west-1")
    paginator = s3.get_paginator("list_objects_v2")
    # We can pass the prefix directly to the S3 API.  If the user has passed
    # a tuple or list of prefixes, we go through them one by one.
    for page in paginator.paginate(Bucket=bucketname):
        try:
            contents = page["Contents"]
        except KeyError:
            print("Key Error")
            break
        for obj in contents:
            key = obj["Key"]
            yield obj

def get_matching_s3_keys_as_set(bucketname, suffix="") -> set:
    """
    Generate the keys in an S3 bucket as set.
    :param bucket: Name of the S3 bucket.
    :param prefix: Only fetch keys that start with this prefix (optional).
    """
    result = set()
    for obj in get_matching_s3_objects(bucketname):
        key = obj["Key"]
        if (key.endswith(suffix)) : result.add( key)
    return result

count = len(get_matching_s3_keys_as_set("letsbuild-cli","LICENSE"))
print("Count: ",count, " times " )
#get_matching_s3_keys_as_set("letsbuild-cli","LICENSE")
