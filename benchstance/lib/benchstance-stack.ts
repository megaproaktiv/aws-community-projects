import { BucketDeployment } from '@aws-cdk/aws-s3-deployment/lib/bucket-deployment';
import { CfnOutput, Duration, RemovalPolicy, Tags } from '@aws-cdk/core';
import * as cdk from '@aws-cdk/core';
import { Action, Instance, InstanceClass, InstanceSize, InstanceType, Peer, Port, SecurityGroup,SubnetType,UserData, Vpc } from '@aws-cdk/aws-ec2';
import { AmazonLinuxCpuType , AmazonLinuxEdition, AmazonLinuxGeneration, AmazonLinuxImage } from '@aws-cdk/aws-ec2';
import { Bucket } from '@aws-cdk/aws-s3';
import { Source } from '@aws-cdk/aws-s3-deployment';

// ** additional CDK imports
import { DestroyableBucket } from 'destroyable-bucket'
import { SelfDestruct} from 'cdk-time-bomb';

// ** local imports
import {BenchstanceVPCStack} from '../lib/benchstance-vpc-stack';

// *** Non cdk imports
import { GetLocalIp } from './getip'
import * as path  from 'path'
import { readFileSync } from 'fs';
import { Effect, PolicyStatement } from '@aws-cdk/aws-iam';


export class BenchstanceStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string,vpcstack: BenchstanceVPCStack ,props?: cdk.StackProps) {
    super(scope, id, props);
    const userDataFile = 'bootstrap/userdata.sh'

    const selfDestruct = new SelfDestruct(this, "selfDestructor", {
      timeToLive: Duration.minutes(15)
    });
    
    const testVpc = vpcstack.vpc;
    const sg = vpcstack.sg;

    // The bucket for output data
    const benchData = new Bucket(this, "BenchData",{
      removalPolicy: RemovalPolicy.RETAIN
    });
    new CfnOutput(this, "dataBucket", 
    {
      value: benchData.bucketName,
      exportName: "benchdata"
    })

    // Deployment bucket for assets
    const benchDeployment = new DestroyableBucket(this, "BenchDeployment",{
      publicReadAccess: false,
      removalPolicy: RemovalPolicy.DESTROY
    });
    new CfnOutput(this, "deploymentBucket", 
    {
      value: benchData.bucketName,
      exportName: "benchdeploy"
    })

    // Copy assets
    const deployment=new BucketDeployment(this, 'DeployMe', {
      sources: [Source.asset(path.join(__dirname, '../assets'))],
      destinationBucket: benchDeployment,
      retainOnDelete: false, // default is true, which will block the integration test cleanup
    });
    deployment.node.addDependency(benchDeployment);
    
    // Update userdata dynamic attributes
    var userdata = readFileSync(userDataFile, 'utf8').toString();
    userdata = userdata.replace(/BUCKET/g, benchData.bucketName);
    userdata = userdata.replace(/DEPLOYMENT/g, benchDeployment.bucketName);
    userdata = userdata.replace(/REGION/g, this.region);
    
    var userdataArm = userdata;
    
    userdata = userdata.replace(/IMDS/g, "ec2-imds-linux");
    userdataArm = userdataArm.replace(/IMDS/g, "ec2-imds-linux-arm");


    // ****************************************************************
    // *    Instances
    // *    Update ami and instance type
    // *    Update userdata for other os
    var amiLinuxX86_64 = new AmazonLinuxImage({
      cpuType: AmazonLinuxCpuType.X86_64,
      edition: AmazonLinuxEdition.STANDARD,
      generation: AmazonLinuxGeneration.AMAZON_LINUX_2
    })
    // *
    var instances: Instance[] = new Array<Instance>();
    
    const testedInstanceC4 = new Instance(this, 'testedInstanceC4', {
      instanceType: InstanceType.of(InstanceClass.COMPUTE4, InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceC4);

    const testedInstanceC5 = new Instance(this, 'testedInstance', {
      instanceType: InstanceType.of(InstanceClass.COMPUTE5, InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceC5);

    const testedInstanceM5 = new Instance(this, 'testedInstancecm5', {
      instanceType: InstanceType.of(InstanceClass.M5, InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceM5);

    const testedInstanceM5a = new Instance(this, 'testedInstancecm5a', {
      instanceType: InstanceType.of(InstanceClass.M5A, InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceM5a);

    const testedInstanceT2 = new Instance(this, 'testedInstanceT2', {
      instanceType: InstanceType.of(InstanceClass.T2, InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceT2);

    const testedInstanceT3 = new Instance(this, 'testedInstanceT3', {
      instanceType: InstanceType.of(InstanceClass.T3, InstanceSize.LARGE),
      machineImage: amiLinuxX86_64,
      userData: UserData.custom(userdata),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceT3);

    var amiLinuxG = new AmazonLinuxImage({
      cpuType: AmazonLinuxCpuType.ARM_64,
      edition: AmazonLinuxEdition.STANDARD,
      generation: AmazonLinuxGeneration.AMAZON_LINUX_2
    })

    const testedInstanceM6G = new Instance(this, 'testedInstanceM6G', {
      instanceType: InstanceType.of(InstanceClass.M6G, InstanceSize.LARGE),
      machineImage: amiLinuxG,
      userData: UserData.custom(userdataArm),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceM6G);

    const testedInstanceT4G = new Instance(this, 'testedInstanceT4G', {
      instanceType: new InstanceType("t4g.large"),
      machineImage: amiLinuxG,
      userData: UserData.custom(userdataArm),
      vpc: testVpc,
      vpcSubnets:  {subnetType: SubnetType.PUBLIC},
      securityGroup: sg
    })
    instances.push(testedInstanceT4G);

    for ( var index = 0; index <instances.length ; index++ )   {
      let currentInstance = instances[index];
      // * Access rights for instance
      benchData.grantReadWrite( currentInstance);
      benchDeployment.grantRead( currentInstance);  
      currentInstance.addToRolePolicy(new PolicyStatement(
        {
          effect: Effect.ALLOW,
          resources: ["arn:aws:ec2:"+this.region+":"+this.account+":instance/*"],
          actions: ["ec2:StopInstances"]
        }
      ))
    }
    // *
    // ****************************************************************
    


  }
}
