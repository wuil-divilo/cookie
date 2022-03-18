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

func TestNew{{cookiecutter.model_name.capitalize()}}Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	dvcRepoMock := mock.NewMock{{cookiecutter.model_name.capitalize()}}Repository(ctrl)
	lgr := logger.NewEmpty()
	type args struct {
		logger     *zap.SugaredLogger
		{{cookiecutter.model_name}}Repo repository.{{cookiecutter.model_name.capitalize()}}Repository
	}
	tests := []struct {
		name string
		args args
		want *{{cookiecutter.model_name}}Service
	}{
		{
			name: "instantiates",
			args: args{
				logger:     lgr,
				{{cookiecutter.model_name}}Repo: dvcRepoMock,
			},
			want: &{{cookiecutter.model_name}}Service{
				lgr:             lgr,
				{{cookiecutter.model_name}}Repo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := New{{cookiecutter.model_name.capitalize()}}Service(tt.args.logger, tt.args.{{cookiecutter.model_name}}Repo).(*{{cookiecutter.model_name}}Service)

			//if ok is false is ERROR
			if !ok || got.{{cookiecutter.model_name}}Repo != tt.want.{{cookiecutter.model_name}}Repo || got.lgr != tt.want.lgr || got.createValidator == nil {
				t.Errorf("New{{cookiecutter.model_name.capitalize()}}Service() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_{{cookiecutter.model_name}}Service_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	dvcRepoMock := mock.NewMock{{cookiecutter.model_name.capitalize()}}Repository(ctrl)
	lgr := logger.NewEmpty()
	type fields struct {
		lgr             *zap.SugaredLogger
		{{cookiecutter.model_name}}Repo      repository.{{cookiecutter.model_name.capitalize()}}Repository
		createValidator *validator.Validate
	}
	type args struct {
		ctx    context.Context
		{{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}
	}
	tests := []struct {
		name           string
		fields         fields
		configureMocks func()
		args           args
		want           model.{{cookiecutter.model_name.capitalize()}}
		wantErr        bool
	}{
		{
			name: "successful created",
			fields: fields{
				lgr:             lgr,
				{{cookiecutter.model_name}}Repo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
			configureMocks: func() {
				dvcRepoMock.EXPECT().Upsert(context.TODO(), model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
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
				}).Return(model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
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
				{{cookiecutter.model_name}}: model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
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
			want: model.{{cookiecutter.model_name.capitalize()}}{
				{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
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
				{{cookiecutter.model_name}}Repo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
			configureMocks: func() {
				return
			},
			args: args{
				ctx: context.TODO(),
				{{cookiecutter.model_name}}: model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-",
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
			want:    model.{{cookiecutter.model_name.capitalize()}}{},
			wantErr: true,
		},
		{
			name: "save error",
			fields: fields{
				lgr:             lgr,
				{{cookiecutter.model_name}}Repo:      dvcRepoMock,
				createValidator: newCreateValidator(),
			},
			configureMocks: func() {
				dvcRepoMock.EXPECT().Upsert(context.TODO(), model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
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
				}).Return(model.{{cookiecutter.model_name.capitalize()}}{}, errors.New("unexpected error"))
			},
			args: args{
				ctx: context.TODO(),
				{{cookiecutter.model_name}}: model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "e2fa2039-f7c4-455f-a6ca-b05db4149ed2",
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
			want:    model.{{cookiecutter.model_name.capitalize()}}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &{{cookiecutter.model_name}}Service{
				lgr:             tt.fields.lgr,
				{{cookiecutter.model_name}}Repo:      tt.fields.{{cookiecutter.model_name}}Repo,
				createValidator: tt.fields.createValidator,
			}
			tt.configureMocks()
			got, err := ds.Create(tt.args.ctx, tt.args.{{cookiecutter.model_name}})
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
