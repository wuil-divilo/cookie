package internal

// Config lambda configuration
type Config struct {
	DomainName       string `config:"DOMAIN_NAME"`
	LogLevel         string `config:"ssm.LOG_LEVEL"`
	CorsOrigins      string `config:"ssm.CORS_ALLOWED_ORIGINS"`
	{{cookiecutter.model_name}}TableName string `config:"{{cookiecutter.model_name}}_TABLE_NAME"`
}
