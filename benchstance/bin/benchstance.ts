#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';

import {BenchstanceVPCStack} from '../lib/benchstance-vpc-stack';
import { BenchstanceStack } from '../lib/benchstance-stack';

const app = new cdk.App();
const vpc = new BenchstanceVPCStack(app,"BenchstanceVPCStack", {
    env: {  region: 'us-east-1' },
});

new BenchstanceStack(app, 'BenchstanceStack',vpc,{
    description: "Benchmarking EC2 with 7zip ",
    env: { region: 'us-east-1' },
});
