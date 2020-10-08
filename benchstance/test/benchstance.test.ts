import { expect as expectCDK, matchTemplate, MatchStyle, haveResource, ResourcePart } from '@aws-cdk/assert';
import * as cdk from '@aws-cdk/core';
import * as Benchstance from '../lib/benchstance-stack';
import * as Vpc from '../lib/benchstance-vpc-stack';

test('Empty Stack', () => {
    const app = new cdk.App();
    // WHEN
    const vpc = new Vpc.BenchstanceVPCStack(app ,'Testvpc');
    const stack = new Benchstance.BenchstanceStack(app, 'MyTestStack', vpc);
    // THEN
    // Bucket is there
   expectCDK(stack).to(haveResource('AWS::S3::Bucket'));
   // bucket has policy< for retentions
   expectCDK(stack).to(haveResource('AWS::S3::Bucket',{
    DeletionPolicy: 'Retain'
   },ResourcePart.CompleteDefinition))

   // Instance
   expectCDK(stack).to(haveResource('AWS::EC2::Instance'))
});
