package repository

// repository constants
const (
	TableColumnDeviceId            = "deviceId"
	TableColumnCreatedAt           = "createdAt"
	MsgErrorNotFound               = "not found"
	MsgErrorMoreThanOneFound       = "more than one found"
	MsgErrorNotAbleMarshal         = "not able to marshal the result"
	MsgErrorNotCreatedUpdateItem   = "not created Update Item Input"
	MsgErrorDynamoClientUpdateItem = "not able to execute the dynamo UpdateItem"
)
