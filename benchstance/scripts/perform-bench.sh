#!/bin/bash
set -xe
## Deploy
task deploy && sleep 600 
## Get all data 
task get-data