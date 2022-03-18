package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/divilo/aws-go/service/dynamodb"
	"github.com/divilo/aws-go/service/ssm"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/repository"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/service"
	"github.com/divilo/{{cookiecutter.project_name}}-back/lambda/create{{cookiecutter.model_name}}/internal"
	"github.com/divilo/utils-go/interfaces"
	"github.com/divilo/utils-go/middleware/cors"
	"github.com/divilo/utils-go/service/config"
	"github.com/divilo/utils-go/service/eventmapper"
	"github.com/divilo/utils-go/service/logger"
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
)

const (
	awsConfig          = "awsConfig"
	globalConfig       = "globalConfig"
	loggerService      = "loggerService"
	corsMiddleware     = "corsMiddleware"
	eventMapperService = "eventMapperService"
	dynamodbClient     = "dynamodbClient"
	{{cookiecutter.model_name}}Repository   = "{{cookiecutter.model_name}}Repository"
	{{cookiecutter.model_name}}Service      = "{{cookiecutter.model_name}}Service"
	handler            = "handler"
)

func bootstrap() *di.Builder {
	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: awsConfig,
		Build: func(ctn di.Container) (interface{}, error) {
			return awsconfig.LoadDefaultConfig(context.TODO())
		},
	})
	builder.Add(di.Def{
		Name: globalConfig,
		Build: func(ctn di.Container) (interface{}, error) {
			awsCfg := ctn.Get(awsConfig).(aws.Config)
			ssmSrv := ssm.NewFromConfig(awsCfg)
			cfg := &internal.Config{}
			err := config.New().WithSSMClient(ssmSrv).LoadDefaultConfig(cfg)
			if err != nil {
				return nil, err
			}
			return cfg, nil
		},
	})
	builder.Add(di.Def{
		Name: loggerService,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(globalConfig).(*internal.Config)
			loggerCfg := logger.Config{
				Level:      cfg.LogLevel,
				IsCloud:    true,
				DomainName: cfg.DomainName,
			}
			return logger.New(loggerCfg), nil
		},
	})
	builder.Add(di.Def{
		Name: corsMiddleware,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(globalConfig).(*internal.Config)
			return cors.New(cfg.CorsOrigins), nil
		},
	})
	builder.Add(di.Def{
		Name: eventMapperService,
		Build: func(ctn di.Container) (interface{}, error) {
			return eventmapper.New(), nil
		},
	})
	builder.Add(di.Def{
		Name: dynamodbClient,
		Build: func(ctn di.Container) (interface{}, error) {
			awsCfg := ctn.Get(awsConfig).(aws.Config)
			return dynamodb.NewFromConfig(awsCfg), nil
		},
	})
	builder.Add(di.Def{
		Name: {{cookiecutter.model_name}}Repository,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.New{{cookiecutter.model_name.capitalize()}}Repository(
				ctn.Get(loggerService).(*zap.SugaredLogger),
				ctn.Get(dynamodbClient).(dynamodb.ServiceDynamo),
				ctn.Get(globalConfig).(*internal.Config).{{cookiecutter.model_name.capitalize()}}sTableName,
			), nil
		},
	})
	builder.Add(di.Def{
		Name: {{cookiecutter.model_name}}Service,
		Build: func(ctn di.Container) (interface{}, error) {
			return service.New{{cookiecutter.model_name.capitalize()}}Service(
				ctn.Get(loggerService).(*zap.SugaredLogger),
				ctn.Get({{cookiecutter.model_name}}Repository).(repository.{{cookiecutter.model_name.capitalize()}}Repository),
			), nil
		},
	})
	builder.Add(di.Def{
		Name: handler,
		Build: func(ctn di.Container) (interface{}, error) {
			return internal.New(
				ctn.Get(loggerService).(*zap.SugaredLogger),
				ctn.Get(corsMiddleware).(interfaces.APIGatewayProxyMiddleware),
				ctn.Get(eventMapperService).(eventmapper.ServiceEventMapper),
				ctn.Get({{cookiecutter.model_name}}Service).(service.{{cookiecutter.model_name.capitalize()}}Service),
			), nil
		},
	})
	return builder
}

func newHandler(builder *di.Builder) interfaces.APIGatewayProxyHandlerFunc {
	ctn := builder.Build()
	return ctn.Get(handler).(interfaces.APIGatewayProxyHandler).HandleProxy()
}
