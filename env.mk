VERSION=$(shell git describe --abbrev=0 --exact-match --tags)
BRANCH=$(shell git branch | grep \* | cut -d ' ' -f2)
DATE=$(shell date)
COMMIT=$(shell git rev-parse HEAD)
STACK-NAME=lex-lambda-sample-fulfilment
BUILD-PROFILE=versent
DEPLOY-PROFILE=versent
EXE-NAME=lex-lambda-sample-fulfilment
LAMBDA-NAME=LexLambdaSampleFulfilment
REGION=us-west-2
SAM-BUCKET=versent-sam-$(REGION)
