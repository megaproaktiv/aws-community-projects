package sftpkey

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	certificatemanager "github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	route53 "github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	ssm "github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	transfer "github.com/aws/aws-cdk-go/awscdk/v2/awstransfer"
	"github.com/aws/constructs-go/constructs/v10"

	// SDK
	"github.com/aws/aws-sdk-go-v2/aws"
)

type SftpKeyStackProps struct {
	awscdk.StackProps
}

func NewSftpKeyStack(scope constructs.Construct, id string, props *SftpKeyStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Bucket
	sftpBucket := s3.NewBucket(stack, aws.String("sftpbucket"),&s3.BucketProps{})

	// Paramater
	// region := ssm.StringParameter_ValueForStringParameter(stack, aws.String("region"))
	domainName := ssm.StringParameter_FromStringParameterName(stack, aws.String("domainParameter"), aws.String("/sftp/serverdns")).StringValue()
	hostedZoneId := ssm.StringParameter_FromStringParameterName(stack,aws.String("zoneidParamater"), aws.String("zoneId")).StringValue()

	// Route 53 ============
	hostedZone1 := route53.HostedZone_FromHostedZoneAttributes(stack, aws.String("zone"), &route53.HostedZoneAttributes{
		HostedZoneId: hostedZoneId,
		ZoneName:     domainName,
	})

	// Certificate =============
	certificate := certificatemanager.NewDnsValidatedCertificate(stack, aws.String("portalCertificate"),
		&certificatemanager.DnsValidatedCertificateProps{
			DomainName:              domainName,
			Validation:              certificatemanager.CertificateValidation_FromDns(hostedZone1),
			HostedZone:              hostedZone1,
		})
		

	// Log Policy ******************
	LogAccessRole := iam.NewRole(stack, aws.String("sftp-server-logging-role"),
		&iam.RoleProps{
			AssumedBy: iam.NewServicePrincipal(aws.String("transfer.amazonaws.com"), &iam.ServicePrincipalOpts{}),
		},
	)
	LogAccessPolicy := iam.NewPolicy(stack, aws.String("sftp-server-logging-policy"),
		&iam.PolicyProps{
			PolicyName: aws.String("sftp-server-logging-policy"),
		})
	LogAccessPolicy.AddStatements(iam.NewPolicyStatement(&iam.PolicyStatementProps{
		Actions: &[]*string{
			aws.String("logs:CreateLogStream"),
			aws.String("logs:DescribeLogStreams"),
			aws.String("logs:CreateLogGroup"),
			aws.String("logs:PutLogEvents"),
		},
		Effect:    iam.Effect_ALLOW,
		Resources: &[]*string{aws.String("*")},
	}))

	//    endpointType: "PUBLIC",
	//    identityProviderType: "SERVICE_MANAGED",
	//    loggingRole: LogAccessRole.roleArn,
	//    protocols: [
	// 	   "SFTP"
	//    ],
	//    domain: "S3",
	//    securityPolicyName: "TransferSecurityPolicy-2020-06",
	//    certificate: certificate.certificateArn,
	// Server
	transfer.NewCfnServer(stack, aws.String("sftpserver"),
		&transfer.CfnServerProps{
			EndpointType:            aws.String("PUBLIC"),
			IdentityProviderType:    aws.String("SERVICE_MANAGED"),
			LoggingRole:             LogAccessRole.PhysicalName(),			
			Protocols:               &[]*string{
				aws.String("SFTP"),
			},
			Domain:                  aws.String("S3"),
			Certificate:             certificate.CertificateArn(),
			SecurityPolicyName:      aws.String("TransferSecurityPolicy-2020-06"),
		})

	// Access Policy ******************
	AccessRole := iam.NewRole(stack, aws.String("AdminAccessRole"),
		&iam.RoleProps{
			AssumedBy: iam.NewServicePrincipal(aws.String("transfer.amazonaws.com"), &iam.ServicePrincipalOpts{}),
		})

	AccessPolicy := iam.NewPolicy(stack, aws.String("sftp-admin-access-policy"),
		&iam.PolicyProps{
			PolicyName: aws.String("sftp-admin-access-policy"),
		})

	AccessPolicy.AddStatements(iam.NewPolicyStatement(&iam.PolicyStatementProps{
		Actions: &[]*string{
			aws.String("s3:ListBucket"),
		},
		Effect:    iam.Effect_ALLOW,
		Resources: &[]*string{sftpBucket.BucketArn()},
	}))

	sftpActions := []*string{
        aws.String("s3:PutObject"),
        aws.String("s3:GetObject"),
        aws.String("s3:DeleteObject"),              
        aws.String("s3:DeleteObjectVersion"),
        aws.String("s3:GetObjectVersion"),
        aws.String("s3:GetObjectACL"),
        aws.String("s3:PutObjectACL"),
	}

	allObjects := *sftpBucket.BucketArn() + "/*"
	AccessPolicy.AddStatements(iam.NewPolicyStatement(&iam.PolicyStatementProps{
		Actions: &sftpActions,
		Effect:  iam.Effect_ALLOW,
		Resources: &[]*string{
			sftpBucket.BucketArn(),
			&allObjects,
		},
	}))

	AccessRole.AttachInlinePolicy(AccessPolicy)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewSftpKeyStack(app, "SftpKeyStack", &SftpKeyStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
