import * as cdk from 'aws-cdk-lib';
import * as 04CdkSnsTs from '../lib/04-cdk-sns-ts-stack';

test('Empty Stack', () => {
    const app = new cdk.App();
    // WHEN
    const stack = new 04CdkSnsTs.04CdkSnsTsStack(app, 'MyTestStack');
    // THEN
    const actual = app.synth().getStackArtifact(stack.artifactId).template;
    expect(actual.Resources ?? {}).toEqual({});
});
