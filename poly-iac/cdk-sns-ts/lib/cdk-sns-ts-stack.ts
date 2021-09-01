import { Stack, StackProps } from 'aws-cdk-lib';
import { aws_sns as sns } from 'aws-cdk-lib';
import { Construct } from 'constructs';

export class CdkSnsTsStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here
    new sns.Topic(scope, "cdksns", {
      
    })

  }
}
