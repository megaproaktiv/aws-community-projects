import { aws_ec2 as aws_ec2 } from 'aws-cdk-lib';               // stable module
import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';

// *** Non cdk imports
import { GetLocalIp } from './getip'




export class BenchstanceVPCStack extends Stack {
  public vpc: aws_ec2.Vpc;
  public sg: aws_ec2.SecurityGroup;

  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);
   
    
    const testVpc = new aws_ec2.Vpc(this, "TestVPC");
    this.vpc = testVpc;

    const sg = new aws_ec2.SecurityGroup(this, "DynamicSSHSG", {
      vpc: testVpc,
      securityGroupName: "SSH incoming",
      description: "SSH Incoming on current public ip",
      allowAllOutbound: true,
    });
    
    const clientIp = GetLocalIp();
  
    clientIp.then((ip) => {
      
      // Tags.of(this).add("Name","dynamicIncomingSSHClientTagBenchstance")
      
      sg.addIngressRule(aws_ec2.Peer.ipv4(ip), aws_ec2.Port.tcp(22), "Ssh Client incoming")
      
      
    });

    this.sg = sg;
  
  }
}
