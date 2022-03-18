package repository

// repository constants
const (
	TableColumn{{cookiecutter.model_name.capitalize()}}Id            = "{{cookiecutter.model_name}}Id"
	TableColumnCreatedAt           = "createdAt"
	MsgErrorNotFound               = "not found"
	MsgErrorMoreThanOneFound       = "more than one found"
	MsgErrorNotAbleMarshal         = "not able to marshal the result"
	MsgErrorNotCreatedUpdateItem   = "not created Update Item Input"
	MsgErrorDynamoClientUpdateItem = "not able to execute the dynamo UpdateItem"
)
