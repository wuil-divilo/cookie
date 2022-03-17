package internal

import "github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"

func toModelDevice(dvcReq *model.CreateDeviceRequest) *model.Device {
	return &model.Device{
		DeviceId:      dvcReq.DeviceId,
		Model:         dvcReq.Model,
		HwVersion:     dvcReq.HardwareVersion,
		OSVersion:     dvcReq.OperatingSystemVersion,
		AppVersion:    dvcReq.AppVersion,
		ApiLevel:      dvcReq.ApiLevel,
		SecurityPatch: dvcReq.SecurityPatch,
		NFCAvailable:  dvcReq.IsNfcAvailable,
		NFCEnabled:    dvcReq.IsNfcEnabled,
	}
}
