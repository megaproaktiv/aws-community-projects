import { Construct } from "constructs";
import { App, TerraformStack } from "cdktf";
import {
  AwsProvider,
  SnsTopic 
  } from './.gen/providers/aws'

class MyStack extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    new AwsProvider(this, "aws", {
      region: "eu-central-1",
    });
    // define resources here
    new  SnsTopic(this, "tf-sns-topic", {
      displayName: "tf-sns-topic"
    } )
  }
}

const app = new App();
new MyStack(app, "cdktf-sns-ts");
app.synth();
