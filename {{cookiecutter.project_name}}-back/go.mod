module github.com/divilo/{{cookiecutter.project_name}}-back

go 1.16

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/aws/aws-sdk-go-v2 v1.15.0
	github.com/aws/aws-sdk-go-v2/config v1.15.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.8.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.4.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.0
	github.com/aws/smithy-go v1.11.1
	github.com/divilo/aws-go v0.4.0
	github.com/divilo/utils-go v0.2.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/golang/mock v1.6.0
	github.com/oxyno-zeta/gomock-extra-matcher v1.1.0
	github.com/sarulabs/di/v2 v2.4.2
	go.uber.org/zap v1.21.0
)
