package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/testing/repository/mock"
	"github.com/divilo/{{cookiecutter.project_name}}-back/lambda/create{{cookiecutter.model_name}}/internal"
	"github.com/divilo/utils-go/service/logger"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	ctrl := gomock.NewController(t)
	dvcRepoMock := mock.NewMock{{cookiecutter.model_name.capitalize()}}Repository(ctrl)
	builder := bootstrap()
	builder.Set(globalConfig, &internal.Config{
		DomainName:       "testing",
		LogLevel:         "DEBUG",
		CorsOrigins:      "testing",
		{{cookiecutter.model_name.capitalize()}}sTableName: "{{cookiecutter.model_name}}s-table",
	})
	builder.Set(loggerService, logger.New(logger.Config{
		Level:      "DEBUG",
		IsCloud:    false,
		DomainName: "testing",
	}))
	builder.Set({{cookiecutter.model_name}}Repository, dvcRepoMock)
	handler := newHandler(builder)
	type args struct {
		ctx   context.Context
		event *events.APIGatewayProxyRequest
	}
	tests := []struct {
		name           string
		configureMocks func()
		args           args
		want           *events.APIGatewayProxyResponse
	}{
		{
			name: "new {{cookiecutter.model_name}} success",
			configureMocks: func() {
				dvcRepoMock.EXPECT().Upsert(gomock.Any(), gomock.Any()).Return(model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
					Model:         "Google Pixel 3",
					HwVersion:     "Qualcomm Snapdragon 845",
					OSVersion:     "Android 9.0",
					AppVersion:    "1.2.2",
					ApiLevel:      "28",
					SecurityPatch: "2021-04-20",
					NFCAvailable:  true,
					NFCEnabled:    true,
					CreatedAt:     0,
					UpdatedAt:     0,
				}, nil)

			},
			args: args{
				ctx: context.TODO(),
				event: &events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "12345678Z",
						},
					},
					Body: `{
					  "{{cookiecutter.model_name}}Id": "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
					  "model": "Google Pixel 3",
					  "hardwareVersion": "Qualcomm Snapdragon 845",
					  "operatingSystemVersion": "Android 9.0",
					  "appVersion": "1.2.2",
					  "apiLevel": "28",
					  "securityPatch": "2021-04-20",
					  "isNfcAvailable": true,
					  "isNfcEnabled": true
					}`,
				},
			},
			want: &events.APIGatewayProxyResponse{
				StatusCode: 204,
				Headers: map[string]string{
					"Content-type":                "application/json",
					"Access-Control-Allow-Origin": "testing",
					"Vary":                        "Origin",
				},
				Body: `""`,
			},
		},
		{
			name: "{{cookiecutter.model_name}} bad request in body",
			configureMocks: func() {
				return
			},
			args: args{
				ctx: context.TODO(),
				event: &events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{},
					Body: `{
					  "{{cookiecutter.model_name}}Id": "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
					  "model": "Google Pixel 3"
					}`,
				},
			},
			want: &events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers: map[string]string{
					"Content-type":                "application/json",
					"Access-Control-Allow-Origin": "testing",
					"Vary":                        "Origin",
				},
				Body: `""`,
			},
		},
		{
			name: "{{cookiecutter.model_name}} maxlength request in body",
			configureMocks: func() {
				return
			},
			args: args{
				ctx: context.TODO(),
				event: &events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "12345678Z",
						},
					},
					Body: `{
					  "{{cookiecutter.model_name}}Id": "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
					  "model": "Google Pixel 3",
					  "hardwareVersion": "Qualcomm Snapdragon 845Qualcomm Snapdragon 845Qualcomm Snapdragon 845Qualcomm Snapdragon 845Qualcomm Snapdragon 845Qualcomm Snapdragon 845",
					  "operatingSystemVersion": "Android 9.0",
					  "appVersion": "1.2.2",
					  "apiLevel": "28",
					  "securityPatch": "2021-04-20",
					  "isNfcAvailable": true,
					  "isNfcEnabled": true
					}`,
				},
			},
			want: &events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers: map[string]string{
					"Content-type":                "application/json",
					"Access-Control-Allow-Origin": "testing",
					"Vary":                        "Origin",
				},
				Body: `""`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.configureMocks()
			if got, _ := handler(tt.args.ctx, tt.args.event); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleProxy()() = %v, want %v", got, tt.want)
			}
		})
	}
}
