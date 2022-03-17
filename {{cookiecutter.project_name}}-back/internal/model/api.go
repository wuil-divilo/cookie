// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package model

const (
	JwtBearerScopes = "jwtBearer.Scopes"
)

// CreateDeviceRequest defines model for CreateDeviceRequest.
type CreateDeviceRequest struct {
	// API revision offered by a version of the Android platform
	ApiLevel string `json:"apiLevel" mod:"trim" validate:"required,max=100"`

	// Version de Divilo APP
	AppVersion string `json:"appVersion" mod:"trim" validate:"required,max=100"`

	// Device identification
	DeviceId string `json:"deviceId" mod:"trim" validate:"required,uuid4"`

	// Mobile hardware version
	HardwareVersion string `json:"hardwareVersion" mod:"trim" validate:"required,max=100"`

	// Indicates if NFC is available
	IsNfcAvailable bool `json:"isNfcAvailable" validate:"required"`

	// Indicates if NFC is enabled
	IsNfcEnabled bool `json:"isNfcEnabled" validate:"required"`

	// Mobile model name
	Model string `json:"model" mod:"trim" validate:"required,max=100"`

	// Mobile operating system version
	OperatingSystemVersion string `json:"operatingSystemVersion" mod:"trim" validate:"required,max=100"`

	// Android software security patch update date
	SecurityPatch string `json:"securityPatch" mod:"trim" validate:"required,max=100"`
}

// CreatedeviceJSONBody defines parameters for Createdevice.
type CreatedeviceJSONBody CreateDeviceRequest

// CreatedeviceJSONRequestBody defines body for Createdevice for application/json ContentType.
type CreatedeviceJSONRequestBody CreatedeviceJSONBody
