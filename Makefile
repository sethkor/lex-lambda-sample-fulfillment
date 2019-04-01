include env.mk

all: zip deploy

zip: linux
	zip -qu9 $(EXE-NAME).zip ./$(EXE-NAME)

linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.version=$(VERSION)-$(BRANCH)' -X 'main.commit=$(COMMIT)' -X 'main.date=$(DATE)'" -o $(EXE-NAME)

package:
	sam package --profile $(BUILD-PROFILE) --region $(REGION) --template-file template.yaml --s3-bucket $(SAM-BUCKET) --output-template-file packaged.yaml

deploy: package
	sam deploy --profile $(DEPLOY-PROFILE) --region $(REGION) --template-file packaged.yaml --stack-name $(STACK-NAME) --capabilities CAPABILITY_IAM

remove:
	aws --profile $(DEPLOY-PROFILE) --region $(REGION) cloudformation delete-stack  --stack-name $(STACK-NAME)
	aws --profile $(DEPLOY-PROFILE) --region $(REGION) cloudformation wait stack-delete-complete --stack-name $(STACK-NAME)

clean:
	rm connect-sla $(EXE-NAME).zip

notify:
	aws --profile $(DEPLOY-PROFILE) --region $(REGION) lambda add-permission --function-name function:$(LAMBDA-NAME)  --statement-id 100 --action "lambda:InvokeFunction" --principal lex.amazonaws.com --source-account 293499315857 --source-arn "$(CONNECT-ARN)"

unnotify:
	aws --profile $(DEPLOY-PROFILE) lambda remove-permission --function-name $(LAMBDA-NAME)  --statement-id 100
