package repository

// {{cookiecutter.model_name.capitalize()}} table model
type {{cookiecutter.model_name}}Table struct {
	{{cookiecutter.model_name.capitalize()}}Id  string          `dynamodbav:"{{cookiecutter.model_name}}Id"`
	Meta      meta{{cookiecutter.model_name.capitalize()}}Table `dynamodbav:"meta"`
	CreatedAt int64           `dynamodbav:"createdAt"`
	UpdatedAt int64           `dynamodbav:"updatedAt"`
}

// meta{{cookiecutter.model_name.capitalize()}}Table
type meta{{cookiecutter.model_name.capitalize()}}Table struct {
	Model         string `dynamodbav:"model"`
	HwVersion     string `dynamodbav:"hwVersion"`
	OSVersion     string `dynamodbav:"osVersion"`
	AppVersion    string `dynamodbav:"appVersion"`
	ApiLevel      string `dynamodbav:"apiLevel"`
	SecurityPatch string `dynamodbav:"securityPatch"`
	NFCAvailable  bool   `dynamodbav:"nfcAvailable"`
	NFCEnabled    bool   `dynamodbav:"nfcEnabled"`
}
