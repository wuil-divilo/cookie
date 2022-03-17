package repository

// Device table model
type deviceTable struct {
	DeviceId  string          `dynamodbav:"deviceId"`
	Meta      metaDeviceTable `dynamodbav:"meta"`
	CreatedAt int64           `dynamodbav:"createdAt"`
	UpdatedAt int64           `dynamodbav:"updatedAt"`
}

// metaDeviceTable
type metaDeviceTable struct {
	Model         string `dynamodbav:"model"`
	HwVersion     string `dynamodbav:"hwVersion"`
	OSVersion     string `dynamodbav:"osVersion"`
	AppVersion    string `dynamodbav:"appVersion"`
	ApiLevel      string `dynamodbav:"apiLevel"`
	SecurityPatch string `dynamodbav:"securityPatch"`
	NFCAvailable  bool   `dynamodbav:"nfcAvailable"`
	NFCEnabled    bool   `dynamodbav:"nfcEnabled"`
}
