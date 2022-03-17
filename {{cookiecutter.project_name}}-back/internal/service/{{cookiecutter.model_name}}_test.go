package service

import (
	"context"
	"errors"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/repository"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/testing/repository/mock"
	"github.com/divilo/utils-go/service/logger"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestNew{{cookiecutter.model_name}}ervice(t *testing.T) {
	ctrl := gomock.NewController(t)
	dvcRepoMock := mock.NewMockDeviceRepository(ctrl)
	lgr := logger.NewEmpty()
	type args struct {
		logger     *zap.SugaredLogger
		deviceRepo repository.DeviceRepository
	}
	tests := []struct {
		name string
		args args
		want *{{cookiecutter.model_name}}ervice
	}{
		{
			name: "instantiates",
			args: args{
				logger:     lgr,
				deviceRepo: dvcRepoMock,
			},
			want: &{{cookiecutter.model_name}}ervice{
				lgr:             lgr,
				deviceRepo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := New{{cookiecutter.model_name}}ervice(tt.args.logger, tt.args.deviceRepo).(*{{cookiecutter.model_name}}ervice)

			//if ok is false is ERROR
			if !ok || got.deviceRepo != tt.want.deviceRepo || got.lgr != tt.want.lgr || got.createValidator == nil {
				t.Errorf("New{{cookiecutter.model_name}}ervice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_{{cookiecutter.model_name}}ervice_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	dvcRepoMock := mock.NewMockDeviceRepository(ctrl)
	lgr := logger.NewEmpty()
	type fields struct {
		lgr             *zap.SugaredLogger
		deviceRepo      repository.DeviceRepository
		createValidator *validator.Validate
	}
	type args struct {
		ctx    context.Context
		device model.Device
	}
	tests := []struct {
		name           string
		fields         fields
		configureMocks func()
		args           args
		want           model.Device
		wantErr        bool
	}{
		{
			name: "successful created",
			fields: fields{
				lgr:             lgr,
				deviceRepo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
			configureMocks: func() {
				dvcRepoMock.EXPECT().Upsert(context.TODO(), model.Device{
					DeviceId:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
					Model:         "Pixel 123",
					HwVersion:     "5",
					OSVersion:     "98",
					AppVersion:    "1",
					ApiLevel:      "125",
					SecurityPatch: "2050-05-06",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     0,
					UpdatedAt:     0,
				}).Return(model.Device{
					DeviceId:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
					Model:         "Pixel 123",
					HwVersion:     "5",
					OSVersion:     "98",
					AppVersion:    "1",
					ApiLevel:      "125",
					SecurityPatch: "2050-05-06",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     123456789,
					UpdatedAt:     123456789,
				}, nil)
			},
			args: args{
				ctx: context.TODO(),
				device: model.Device{
					DeviceId:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
					Model:         "Pixel 123",
					HwVersion:     "5",
					OSVersion:     "98",
					AppVersion:    "1",
					ApiLevel:      "125",
					SecurityPatch: "2050-05-06",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     0,
					UpdatedAt:     0,
				},
			},
			want: model.Device{
				DeviceId:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
				Model:         "Pixel 123",
				HwVersion:     "5",
				OSVersion:     "98",
				AppVersion:    "1",
				ApiLevel:      "125",
				SecurityPatch: "2050-05-06",
				NFCAvailable:  true,
				NFCEnabled:    false,
				CreatedAt:     123456789,
				UpdatedAt:     123456789,
			},
			wantErr: false,
		},
		{
			name: "validation error",
			fields: fields{
				lgr:             lgr,
				deviceRepo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
			configureMocks: func() {
				return
			},
			args: args{
				ctx: context.TODO(),
				device: model.Device{
					DeviceId:      "e2fa2039-",
					Model:         "",
					HwVersion:     "5",
					OSVersion:     "98",
					AppVersion:    "1",
					ApiLevel:      "125",
					SecurityPatch: "2050-05-06",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     123456789,
					UpdatedAt:     0,
				},
			},
			want:    model.Device{},
			wantErr: true,
		},
		{
			name: "save error",
			fields: fields{
				lgr:             lgr,
				deviceRepo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
			configureMocks: func() {
				dvcRepoMock.EXPECT().Upsert(context.TODO(), model.Device{
					DeviceId:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
					Model:         "Pixel 123",
					HwVersion:     "5",
					OSVersion:     "98",
					AppVersion:    "1",
					ApiLevel:      "125",
					SecurityPatch: "2050-05-06",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     0,
					UpdatedAt:     0,
				}).Return(model.Device{}, errors.New("unexpected error"))
			},
			args: args{
				ctx: context.TODO(),
				device: model.Device{
					DeviceId:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
					Model:         "Pixel 123",
					HwVersion:     "5",
					OSVersion:     "98",
					AppVersion:    "1",
					ApiLevel:      "125",
					SecurityPatch: "2050-05-06",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     0,
					UpdatedAt:     0,
				},
			},
			want:    model.Device{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &{{cookiecutter.model_name}}ervice{
				lgr:             tt.fields.lgr,
				deviceRepo:      tt.fields.deviceRepo,
				createValidator: tt.fields.createValidator,
			}
			tt.configureMocks()
			got, err := ds.Create(tt.args.ctx, tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
