import { aws_s3 as aws_s3, CfnOutput, RemovalPolicy, Stack, StackProps } from 'aws-cdk-lib';               // stable module
import { aws_s3_deployment as aws_s3_deployment } from 'aws-cdk-lib';               // stable module
import { aws_ec2 as aws_ec2 } from 'aws-cdk-lib';               // stable module
import { aws_iam as aws_iam } from 'aws-cdk-lib';               // stable module
import { Construct } from 'constructs';


// ** local imports
import { BenchstanceVPCStack } from '../lib/benchstance-vpc-stack';

// *** Non cdk imports
import * as path from 'path'
import { readFileSync } from 'fs';


export class BenchstanceStack extends Stack {
  constructor(scope: Construct, id: string, vpcstack: BenchstanceVPCStack, props?: StackProps) {
    super(scope, id, props);
    const userDataFile = 'bootstrap/userdata.sh'

    
    const testVpc = vpcstack.vpc;
    const sg = vpcstack.sg;

    // The bucket for output data
    const benchData = new aws_s3.Bucket(this, "BenchData", {
      removalPolicy: RemovalPolicy.RETAIN
    });
    new CfnOutput(this, "dataBucket",
      {
        value: benchData.bucketName,
        exportName: "benchdata"
      })
    const benchdeploy = new aws_s3.Bucket(this, "benchdeploy", {
      removalPolicy: RemovalPolicy.RETAIN
    });
    new CfnOutput(this, "benchDeployment-out",
      {
        value: benchdeploy.bucketName,
        exportName: "benchdeploy"
      })

    
    // Copy assets
    const deployment = new aws_s3_deployment.BucketDeployment(this, 'DeployMe', {
      sources: [aws_s3_deployment.Source.asset(path.join(__dirname, '../assets'))],
      destinationBucket: benchdeploy,
      retainOnDelete: false, // default is true, which will block the integration test cleanup
    });

    // Update userdata dynamic attributes
    var userdata = readFileSync(userDataFile, 'utf8').toString();
    userdata = userdata.replace(/BUCKET/g, benchData.bucketName);
    userdata = userdata.replace(/DEPLOYMENT/g, benchdeploy.bucketName);
    userdata = userdata.replace(/REGION/g, this.region);

    var userdataArm = userdata;

    userdata = userdata.replace(/IMDS/g, "ec2-imds-linux");
    userdataArm = userdataArm.replace(/IMDS/g, "ec2-imds-linux-arm");


    // ****************************************************************
    // *    Instances
    // *    Update ami and instance type
    // *    Update userdata for other os
    var amiLinuxX86_64 = new aws_ec2.AmazonLinuxImage({
      cpuType: aws_ec2.AmazonLinuxCpuType.X86_64,
      edition: aws_ec2.AmazonLinuxEdition.STANDARD,
      generation: aws_ec2.AmazonLinuxGeneration.AMAZON_LINUX_2
    })
    // *
    var instances: aws_ec2.Instance[] = new Array<aws_ec2.Instance>();

    const instanceAMD001 = new aws_ec2.Instance(this, 'testedInstanceC4XL', {
      instanceType: aws_ec2.InstanceType.of(aws_ec2.InstanceClass.C4, aws_ec2.InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: aws_ec2.UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets: { subnetType: aws_ec2.SubnetType.PUBLIC },
      securityGroup: sg
    })
    instances.push(instanceAMD001);

    const testedInstance002 = new aws_ec2.Instance(this, 'testedInstanceC5XL', {
      instanceType: aws_ec2.InstanceType.of(aws_ec2.InstanceClass.C5, aws_ec2.InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: aws_ec2.UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets: { subnetType: aws_ec2.SubnetType.PUBLIC },
      securityGroup: sg
    })
    instances.push(testedInstance002);

    const testedInstanceC6i = new aws_ec2.Instance(this, 'testedInstanceC6iXL', {
      instanceType: aws_ec2.InstanceType.of(aws_ec2.InstanceClass.C6I, aws_ec2.InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: aws_ec2.UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets: { subnetType: aws_ec2.SubnetType.PUBLIC },
      securityGroup: sg
    })
    instances.push(testedInstanceC6i);



    var amiLinuxG = new aws_ec2.AmazonLinuxImage({
      cpuType: aws_ec2.AmazonLinuxCpuType.ARM_64,
      edition: aws_ec2.AmazonLinuxEdition.STANDARD,
      generation: aws_ec2.AmazonLinuxGeneration.AMAZON_LINUX_2
    })

    const testedInstanceG01 = new aws_ec2.Instance(this, 'testedInstanceC6G', {
      instanceType: aws_ec2.InstanceType.of(aws_ec2.InstanceClass.C6G, aws_ec2.InstanceSize.LARGE),
      machineImage: amiLinuxG,
      userData: aws_ec2.UserData.custom(userdataArm),
      vpc: testVpc,
      vpcSubnets: { subnetType: aws_ec2.SubnetType.PUBLIC },
      securityGroup: sg
    })
    instances.push(testedInstanceG01);

    const testedInstanceG3_01 = new aws_ec2.Instance(this, 'C7g', {
      instanceType: aws_ec2.InstanceType.of(aws_ec2.InstanceClass.C7G, aws_ec2.InstanceSize.LARGE),
      machineImage: amiLinuxG,
      userData: aws_ec2.UserData.custom(userdataArm),
      vpc: testVpc,
      vpcSubnets: { subnetType: aws_ec2.SubnetType.PUBLIC },
      securityGroup: sg
    })
    instances.push(testedInstanceG3_01);

    for (var index = 0; index < instances.length; index++) {
      let currentInstance = instances[index];
      // * Access rights for instance
      benchData.grantReadWrite(currentInstance);
      benchdeploy.grantRead(currentInstance);
      currentInstance.addToRolePolicy(new aws_iam.PolicyStatement(
        {
          effect: aws_iam.Effect.ALLOW,
          resources: ["arn:aws:ec2:" + this.region + ":" + this.account + ":instance/*"],
          actions: ["ec2:StopInstances"]
        }
      ))
    }
    // *
    // ****************************************************************



  }
}
