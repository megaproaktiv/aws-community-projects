# Benchmark Instances with 7zip

**Stop guessing capacity! - Start calculating.**

If you migrate older server to the AWS Cloud using EC2 instances, the prefered way is to start with a good guess and then rightsize with CloudWatch metric data.

But sometimes you got no clue, where to start. And: Did you thing an AWS vcpu is equal an AWS vpcu?  
No, not at all. The compute power of different instance types is - yes - different. 


See details on blogpost on [tecRacer AWS blog.de](www.aws-blog.de).

## Create

Each step is a task in `Taskfile.yml`.

So with installed `task.dev` and `cdk` you just start:

```bash
scripts/perform-bench.sh
```

This will create data <instancedid>.txt in `data`subfolder.


## Buckets and retain

The benchmark results should remain after the stack is deleted. So the bucket should remain, but should not hinder the red-deployment.

 A unit test checks, that the bucket really has the right policy.

```ts
   expectCDK(stack).to(haveResource('AWS::S3::Bucket',{
    DeletionPolicy: 'Retain'
   },ResourcePart.CompleteDefinition))
```

The bucket naming is unique and with a new deployment you get a new name.

## Metadata

You could just query the metadata with a curl like `curl http://169.254.169.254/latest/meta-data/instance-i`.
But this only works for instance metadata v2. So its better to use a script. As i want to support as many os as possible, python is not the preffered way, because there might be no python. So a static compiled go program will do the job.

With a little search i get: [ec2-imds](https://github.com/SWITCHCo/ec2-imds), which also works with imdv2.
