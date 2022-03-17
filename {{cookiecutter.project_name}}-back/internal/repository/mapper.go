package repository

import (
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
)

func toModelDevice(deviceTable *deviceTable) *model.Device {
	return &model.Device{
		DeviceId:      deviceTable.DeviceId,
		Model:         deviceTable.Meta.Model,
		HwVersion:     deviceTable.Meta.HwVersion,
		OSVersion:     deviceTable.Meta.OSVersion,
		AppVersion:    deviceTable.Meta.AppVersion,
		ApiLevel:      deviceTable.Meta.ApiLevel,
		SecurityPatch: deviceTable.Meta.SecurityPatch,
		NFCAvailable:  deviceTable.Meta.NFCAvailable,
		NFCEnabled:    deviceTable.Meta.NFCEnabled,
		CreatedAt:     deviceTable.CreatedAt,
		UpdatedAt:     deviceTable.UpdatedAt,
	}
}

func fromModelDevice(device *model.Device) *deviceTable {
	return &deviceTable{
		DeviceId: device.DeviceId,
		Meta: metaDeviceTable{
			Model:         device.Model,
			HwVersion:     device.HwVersion,
			OSVersion:     device.OSVersion,
			AppVersion:    device.AppVersion,
			ApiLevel:      device.ApiLevel,
			SecurityPatch: device.SecurityPatch,
			NFCAvailable:  device.NFCAvailable,
			NFCEnabled:    device.NFCEnabled,
		},
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}
}
