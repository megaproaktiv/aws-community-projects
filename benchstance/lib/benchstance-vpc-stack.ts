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

// *** Non cdk imports
import { GetLocalIp } from './getip'
import * as path  from 'path'
import { readFileSync } from 'fs';
import { Effect, PolicyStatement } from '@aws-cdk/aws-iam';


export class BenchstanceVPCStack extends cdk.Stack {
  public vpc: Vpc;
  public sg: SecurityGroup;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);
   
    const selfDestruct = new SelfDestruct(this, "selfDestructor", {
      timeToLive: Duration.minutes(60)
    });
    
    const testVpc = new Vpc(this, "TestVPC");
    testVpc.node.addDependency(selfDestruct);
    this.vpc = testVpc;

    const sg = new SecurityGroup(this, "DynamicSSHSG", {
      vpc: testVpc,
      securityGroupName: "SSH incoming",
      description: "SSH Incoming on current public ip",
      allowAllOutbound: true,
    });
    
    const clientIp = GetLocalIp();
  
    clientIp.then((ip) => {
      
      Tags.of(this).add("Name","dynamicIncomingSSHClientTagBenchstance")
      
      sg.addIngressRule(Peer.ipv4(ip), Port.tcp(22), "Ssh Client incoming")
      
      
    });

    this.sg = sg;
  
  }
}
