#!/bin/bash
set -xe
sudo yum update -y
# Amazon Linux 2
sudo amazon-linux-extras install epel -y
sudo yum install p7zip -y
## 169.254.169.254/latest/meta-data
export deploy="DEPLOYMENT"
aws s3 cp s3://${deploy}/IMDS .
chmod u+x IMDS
export instanceid=`./IMDS instance-id`
export benchfile=/tmp/${instanceid}.txt
export outout="BUCKET"
# Wait for other startup scripts to finish
sleep 300
7za b >${benchfile}
echo -n "instance-type: " >> ${benchfile}
./IMDS instance-type >> ${benchfile}
echo  >> ${benchfile}
echo -n "instance-id: " >> ${benchfile}
./IMDS instance-id >> ${benchfile}
aws s3 cp ${benchfile} s3://${outout}/${instanceid}.txt
aws ec2 stop-instances --instance-ids "${instanceid}" --region REGION
