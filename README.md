# lex-lambda-sample-go
A sample Lambda that can be invoked by lex.  It also makes a good template for a lambda in Go.  [AWS CLI](https://aws.amazon.com/cli/) and [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html) is used for Lambda deployment.  Make sure you've installed these and [Go](https://golang.org/).  [Brew](https://brew.sh/) is usefull for all fo this if your on a mac.

Install this package with:

```
$ go get github.com/versent/connect-sample-lambda-go

```
To build and deploy we use Makefiles, AWS SAM and the AWS CLI.
You will need to edit the env.mk file and replace a couple of values with your own:

Value               | Description
------------------- | ----------------
`SAM-BUCKET`        | The S3 bucket SAM can use for deployments.  Make sure the`BUILD-PROFILE` has access to this bucket
`STACK-NAME`        | The name to use for the CFN stack that deploys your Lambda
`BUILD-PROFILE`     | The aws profile to use for building the Lambda
`DEPLOY-PROFILE` | The deployment profile to use
`EXE-NAME`         | The name of the executable that is built
`CONNECT-ARN`       | ARN of your AWS Connect Instance
`LAMBDA-NAME`       | The name of your Lambda function 


Sample values:

```
SAM-BUCKET=versent-connect-sam
STACK-NAME=connect-sla-s3-lambda
BUILD-PROFILE=versent
DEPLOY-PROFILE=versent
EXE-NAME=connect-sample-lambda-go
CONNECT-ARN=arn:aws:connect:ap-southeast-2:293499315857:instance/cfec03c9-bf43-4a82-b660-2c98aecddc61
LAMBDA-NAME=ConnectSampleLambdaGo
```
Make targets:

Target   | Description
---------|-------------
all | `$ make zip package deploy` This is also the default target.
zip | To compile your go source into the executable needed by lambda
package | Prepare Lambda S3 package and CFN template from the template.yaml.  Load the prepared package to S3
deploy | Deploy your lambda
remove | delete your lambda
notify | set up notification permissions of Lambda function to allow triggers from AWS Connect
unnotify | remove notification permissions of Lambda function





