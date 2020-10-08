#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';

import {BenchstanceVPCStack} from '../lib/benchstance-vpc-stack';
import { BenchstanceStack } from '../lib/benchstance-stack';

const app = new cdk.App();
const vpc = new BenchstanceVPCStack(app,"BenchstanceVPCStack");

new BenchstanceStack(app, 'BenchstanceStack',vpc,{
    description: "Benchmarking EC2 with 7zip "
});
