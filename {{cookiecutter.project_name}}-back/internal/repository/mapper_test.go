package repository

import (
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"reflect"
	"testing"
)

func Test_fromModel{{cookiecutter.model_name.capitalize()}}(t *testing.T) {
	type args struct {
		{{cookiecutter.model_name}} *model.{{cookiecutter.model_name.capitalize()}}
	}
	tests := []struct {
		name string
		args args
		want *{{cookiecutter.model_name}}Table
	}{
		{
			name: "convert to {{cookiecutter.model_name}} model",
			args: args{
				{{cookiecutter.model_name}}: &model.{{cookiecutter.model_name.capitalize()}}{
					{{cookiecutter.model_name.capitalize()}}Id:      "{{cookiecutter.model_name}}Id",
					Model:         "model",
					HwVersion:     "hwversion",
					OSVersion:     "osversion",
					AppVersion:    "appversion",
					ApiLevel:      "apilevel",
					SecurityPatch: "securitypatch",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     1,
					UpdatedAt:     2,
				},
			},
			want: &{{cookiecutter.model_name}}Table{
				{{cookiecutter.model_name.capitalize()}}Id: "{{cookiecutter.model_name}}Id",
				Meta: meta{{cookiecutter.model_name.capitalize()}}Table{
					Model:         "model",
					HwVersion:     "hwversion",
					OSVersion:     "osversion",
					AppVersion:    "appversion",
					ApiLevel:      "apilevel",
					SecurityPatch: "securitypatch",
					NFCAvailable:  true,
					NFCEnabled:    false,
				},
				CreatedAt: 1,
				UpdatedAt: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fromModel{{cookiecutter.model_name.capitalize()}}(tt.args.{{cookiecutter.model_name}})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromModel{{cookiecutter.model_name.capitalize()}}() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toModel{{cookiecutter.model_name.capitalize()}}(t *testing.T) {
	type args struct {
		{{cookiecutter.model_name}}Table *{{cookiecutter.model_name}}Table
	}
	tests := []struct {
		name string
		args args
		want *model.{{cookiecutter.model_name.capitalize()}}
	}{
		{
			name: "success",
			args: args{
				{{cookiecutter.model_name}}Table: &{{cookiecutter.model_name}}Table{
					{{cookiecutter.model_name.capitalize()}}Id: "{{cookiecutter.model_name}}Id",
					Meta: meta{{cookiecutter.model_name.capitalize()}}Table{
						Model:         "model",
						HwVersion:     "hwversion",
						OSVersion:     "osversion",
						AppVersion:    "appversion",
						ApiLevel:      "apilevel",
						SecurityPatch: "securitypatch",
						NFCAvailable:  true,
						NFCEnabled:    true,
					},
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			want: &model.{{cookiecutter.model_name.capitalize()}}{
				{{cookiecutter.model_name.capitalize()}}Id:      "{{cookiecutter.model_name}}Id",
				Model:         "model",
				HwVersion:     "hwversion",
				OSVersion:     "osversion",
				AppVersion:    "appversion",
				ApiLevel:      "apilevel",
				SecurityPatch: "securitypatch",
				NFCAvailable:  true,
				NFCEnabled:    true,
				CreatedAt:     1,
				UpdatedAt:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toModel{{cookiecutter.model_name.capitalize()}}(tt.args.{{cookiecutter.model_name}}Table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toModel{{cookiecutter.model_name.capitalize()}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
