package internal

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/service"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/testing/service/mock"
	"github.com/divilo/utils-go/interfaces"
	"github.com/divilo/utils-go/service/eventmapper"
	"github.com/divilo/utils-go/service/logger"
	utilsmock "github.com/divilo/utils-go/testing/mock"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	lgr := logger.NewEmpty()
	corsMidMock := utilsmock.NewMockAPIGatewayProxyMiddleware(ctrl)
	evtMapMock := utilsmock.NewMockServiceEventMapper(ctrl)
	dvcSrvMock := mock.NewMock{{cookiecutter.model_name.capitalize()}}Service(ctrl)
	type args struct {
		lgr            *zap.SugaredLogger
		corsMiddleware interfaces.APIGatewayProxyMiddleware
		eventMapper    eventmapper.ServiceEventMapper
		{{cookiecutter.model_name}}Service  service.{{cookiecutter.model_name.capitalize()}}Service
	}
	tests := []struct {
		name string
		args args
		want interfaces.APIGatewayProxyHandler
	}{
		{
			name: "creates instance",
			args: args{
				lgr:            lgr,
				corsMiddleware: corsMidMock,
				eventMapper:    evtMapMock,
				{{cookiecutter.model_name}}Service:  dvcSrvMock,
			},
			want: &handler{
				lgr:            lgr,
				corsMiddleware: corsMidMock,
				eventMapper:    evtMapMock,
				{{cookiecutter.model_name}}Service:  dvcSrvMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.lgr, tt.args.corsMiddleware, tt.args.eventMapper, tt.args.{{cookiecutter.model_name}}Service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_HandleProxy(t *testing.T) {
	ctrl := gomock.NewController(t)
	lgr := logger.NewEmpty()
	corsMidMock := utilsmock.NewMockAPIGatewayProxyMiddleware(ctrl)
	evtMapMock := utilsmock.NewMockServiceEventMapper(ctrl)
	dvcSrvMock := mock.NewMock{{cookiecutter.model_name.capitalize()}}Service(ctrl)
	type fields struct {
		lgr            *zap.SugaredLogger
		corsMiddleware interfaces.APIGatewayProxyMiddleware
		eventMapper    eventmapper.ServiceEventMapper
		{{cookiecutter.model_name}}Service  service.{{cookiecutter.model_name.capitalize()}}Service
	}
	type args struct {
		ctx   context.Context
		event *events.APIGatewayProxyRequest
	}
	tests := []struct {
		name           string
		fields         fields
		configureMocks func()
		args           args
		want           *events.APIGatewayProxyResponse
	}{
		{
			name: "handle successfully upsert {{cookiecutter.model_name}}",
			fields: fields{
				lgr:            lgr,
				corsMiddleware: corsMidMock,
				eventMapper:    evtMapMock,
				{{cookiecutter.model_name}}Service:  dvcSrvMock,
			},
			configureMocks: func() {
				corsMidMock.EXPECT().ProxyMiddleware(gomock.Any()).DoAndReturn(
					func(next interfaces.APIGatewayProxyHandlerFunc) interfaces.APIGatewayProxyHandlerFunc {
						return next
					},
				)
				evtMapMock.EXPECT().FromProxyRequest(&events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "user-id",
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
				}, &handlerRequest{}).DoAndReturn(func(event *events.APIGatewayProxyRequest, output *handlerRequest) error {
					output.UserID = "user-id"
					output.{{cookiecutter.model_name.capitalize()}}Id = "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27"
					output.Model = "Google Pixel 3"
					output.HardwareVersion = "Qualcomm Snapdragon 845"
					output.AppVersion = "1.2.2"
					output.OperatingSystemVersion = "Android 9.0"
					output.ApiLevel = "28"
					output.SecurityPatch = "2021-04-20"
					output.IsNfcAvailable = true
					output.IsNfcEnabled = true
					return nil
				})
				evtMapMock.EXPECT().ToProxyResponse(http.StatusNoContent, "").Return(&events.APIGatewayProxyResponse{
					StatusCode: 204,
					Headers: map[string]string{
						"Content-type": "application/json",
					},
					Body: `""`,
				}, nil)
				dvcSrvMock.EXPECT().Create(context.TODO(), model.{{cookiecutter.model_name.capitalize()}}{
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
				}).Return(model.{{cookiecutter.model_name.capitalize()}}{
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
				context.TODO(),
				&events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "user-id",
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
					"Content-type": "application/json",
				},
				MultiValueHeaders: nil,
				Body:              `""`,
				IsBase64Encoded:   false,
			},
		},
		{
			name: "handle bad request error",
			fields: fields{
				lgr:            lgr,
				corsMiddleware: corsMidMock,
				eventMapper:    evtMapMock,
				{{cookiecutter.model_name}}Service:  dvcSrvMock,
			},
			configureMocks: func() {
				corsMidMock.EXPECT().ProxyMiddleware(gomock.Any()).DoAndReturn(
					func(next interfaces.APIGatewayProxyHandlerFunc) interfaces.APIGatewayProxyHandlerFunc {
						return next
					},
				)
				evtMapMock.EXPECT().FromProxyRequest(&events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "user-id",
						},
					},
					Body: `{"{{cookiecutter.model_name}}Id":"", "model":""}`,
				}, &handlerRequest{}).Return(errors.New("error"))
				evtMapMock.EXPECT().ToProxyResponse(http.StatusBadRequest, "").Return(&events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers: map[string]string{
						"Content-type": "application/json",
					},
					Body: "",
				}, nil)
			},
			args: args{
				context.TODO(),
				&events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "user-id",
						},
					},
					Body: `{"{{cookiecutter.model_name}}Id":"", "model":""}`,
				},
			},
			want: &events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers: map[string]string{
					"Content-type": "application/json",
				},
				Body: "",
			},
		},
		{
			name: "handle service error",
			fields: fields{
				lgr:            lgr,
				corsMiddleware: corsMidMock,
				eventMapper:    evtMapMock,
				{{cookiecutter.model_name}}Service:  dvcSrvMock,
			},
			configureMocks: func() {
				corsMidMock.EXPECT().ProxyMiddleware(gomock.Any()).DoAndReturn(
					func(next interfaces.APIGatewayProxyHandlerFunc) interfaces.APIGatewayProxyHandlerFunc {
						return next
					},
				)
				evtMapMock.EXPECT().FromProxyRequest(&events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "user-id",
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
				}, &handlerRequest{}).DoAndReturn(func(event *events.APIGatewayProxyRequest, output *handlerRequest) error {
					output.UserID = "user-id"
					output.{{cookiecutter.model_name.capitalize()}}Id = "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27"
					output.Model = "Google Pixel 3"
					output.HardwareVersion = "Qualcomm Snapdragon 845"
					output.AppVersion = "1.2.2"
					output.OperatingSystemVersion = "Android 9.0"
					output.ApiLevel = "28"
					output.SecurityPatch = "2021-04-20"
					output.IsNfcAvailable = true
					output.IsNfcEnabled = true
					return nil
				})
				evtMapMock.EXPECT().ToProxyResponse(http.StatusInternalServerError, "").Return(&events.APIGatewayProxyResponse{
					StatusCode: 500,
					Headers: map[string]string{
						"Content-type": "application/json",
					},
					Body: "",
				}, nil)
				dvcSrvMock.EXPECT().Create(context.TODO(), model.{{cookiecutter.model_name.capitalize()}}{
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
				}).Return(model.{{cookiecutter.model_name.capitalize()}}{}, errors.New("service error"))
			},
			args: args{
				context.TODO(),
				&events.APIGatewayProxyRequest{
					RequestContext: events.APIGatewayProxyRequestContext{
						Authorizer: map[string]interface{}{
							"username": "user-id",
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
				StatusCode: 500,
				Headers: map[string]string{
					"Content-type": "application/json",
				},
				Body: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				lgr:            tt.fields.lgr,
				corsMiddleware: tt.fields.corsMiddleware,
				eventMapper:    tt.fields.eventMapper,
				{{cookiecutter.model_name}}Service:  tt.fields.{{cookiecutter.model_name}}Service,
			}
			tt.configureMocks()
			if got, _ := h.HandleProxy()(tt.args.ctx, tt.args.event); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleProxy()() = %v, want %v", got, tt.want)
			}
		})
	}
}
