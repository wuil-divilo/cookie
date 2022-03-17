package model

//Device domain model
type Device struct {
	DeviceId      string `validateCreate:"required,uuid"`
	Model         string `validateCreate:"required,max=100"`
	HwVersion     string `validateCreate:"required,max=100"`
	OSVersion     string `validateCreate:"required,max=100"`
	AppVersion    string `validateCreate:"required,max=100"`
	ApiLevel      string `validateCreate:"required,max=100"`
	SecurityPatch string `validateCreate:"required,max=100"`
	NFCAvailable  bool
	NFCEnabled    bool
	CreatedAt     int64 `validateCreate:"isdefault"`
	UpdatedAt     int64 `validateCreate:"isdefault"`
}
