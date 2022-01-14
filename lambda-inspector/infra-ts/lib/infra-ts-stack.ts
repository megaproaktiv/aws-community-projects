import { Duration, Stack, StackProps } from 'aws-cdk-lib';
import { aws_lambda } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import path = require('path');


export class InfraTsStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    var dockerfile = path.join(__dirname, '../../app/node')
    const dockerlambda = new aws_lambda.DockerImageFunction(this, "lambdainspector-node",
    {
      architecture: aws_lambda.Architecture.X86_64,
      functionName: "lambdainspector-node",
      timeout: Duration.seconds(1024),
      code: aws_lambda.DockerImageCode.fromImageAsset(dockerfile)
    })
    
  }
}
