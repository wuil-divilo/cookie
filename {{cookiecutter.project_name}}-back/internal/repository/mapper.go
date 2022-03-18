package repository

import (
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
)

func toModel{{cookiecutter.model_name.capitalize()}}({{cookiecutter.model_name}}Table *{{cookiecutter.model_name}}Table) *model.{{cookiecutter.model_name.capitalize()}} {
	return &model.{{cookiecutter.model_name.capitalize()}}{
		{{cookiecutter.model_name.capitalize()}}Id:      {{cookiecutter.model_name}}Table.{{cookiecutter.model_name.capitalize()}}Id,
		Model:         {{cookiecutter.model_name}}Table.Meta.Model,
		HwVersion:     {{cookiecutter.model_name}}Table.Meta.HwVersion,
		OSVersion:     {{cookiecutter.model_name}}Table.Meta.OSVersion,
		AppVersion:    {{cookiecutter.model_name}}Table.Meta.AppVersion,
		ApiLevel:      {{cookiecutter.model_name}}Table.Meta.ApiLevel,
		SecurityPatch: {{cookiecutter.model_name}}Table.Meta.SecurityPatch,
		NFCAvailable:  {{cookiecutter.model_name}}Table.Meta.NFCAvailable,
		NFCEnabled:    {{cookiecutter.model_name}}Table.Meta.NFCEnabled,
		CreatedAt:     {{cookiecutter.model_name}}Table.CreatedAt,
		UpdatedAt:     {{cookiecutter.model_name}}Table.UpdatedAt,
	}
}

func fromModel{{cookiecutter.model_name.capitalize()}}({{cookiecutter.model_name}} *model.{{cookiecutter.model_name.capitalize()}}) *{{cookiecutter.model_name}}Table {
	return &{{cookiecutter.model_name}}Table{
		{{cookiecutter.model_name.capitalize()}}Id: {{cookiecutter.model_name}}.{{cookiecutter.model_name.capitalize()}}Id,
		Meta: meta{{cookiecutter.model_name.capitalize()}}Table{
			Model:         {{cookiecutter.model_name}}.Model,
			HwVersion:     {{cookiecutter.model_name}}.HwVersion,
			OSVersion:     {{cookiecutter.model_name}}.OSVersion,
			AppVersion:    {{cookiecutter.model_name}}.AppVersion,
			ApiLevel:      {{cookiecutter.model_name}}.ApiLevel,
			SecurityPatch: {{cookiecutter.model_name}}.SecurityPatch,
			NFCAvailable:  {{cookiecutter.model_name}}.NFCAvailable,
			NFCEnabled:    {{cookiecutter.model_name}}.NFCEnabled,
		},
		CreatedAt: {{cookiecutter.model_name}}.CreatedAt,
		UpdatedAt: {{cookiecutter.model_name}}.UpdatedAt,
	}
}
