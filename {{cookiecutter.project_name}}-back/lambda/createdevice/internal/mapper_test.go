package internal

import (
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"reflect"
	"testing"
)

func Test_toModel{{cookiecutter.model_name.capitalize()}}(t *testing.T) {
	type args struct {
		dvcReq *model.Create{{cookiecutter.model_name.capitalize()}}Request
	}
	tests := []struct {
		name string
		args args
		want *model.{{cookiecutter.model_name.capitalize()}}
	}{
		{
			name: "convert to {{cookiecutter.model_name}} model",
			args: args{
				dvcReq: &model.Create{{cookiecutter.model_name.capitalize()}}Request{
					ApiLevel:               "apilevel",
					AppVersion:             "appversion",
					{{cookiecutter.model_name.capitalize()}}Id:               "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
					HardwareVersion:        "hardware version",
					IsNfcAvailable:         true,
					IsNfcEnabled:           false,
					Model:                  "model1",
					OperatingSystemVersion: "android os x.x",
					SecurityPatch:          "2021-10-25",
				},
			},
			want: &model.{{cookiecutter.model_name.capitalize()}}{
				{{cookiecutter.model_name.capitalize()}}Id:      "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
				Model:         "model1",
				HwVersion:     "hardware version",
				OSVersion:     "android os x.x",
				AppVersion:    "appversion",
				ApiLevel:      "apilevel",
				SecurityPatch: "2021-10-25",
				NFCAvailable:  true,
				NFCEnabled:    false,
				CreatedAt:     0,
				UpdatedAt:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toModel{{cookiecutter.model_name.capitalize()}}(tt.args.dvcReq)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toModel{{cookiecutter.model_name.capitalize()}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
