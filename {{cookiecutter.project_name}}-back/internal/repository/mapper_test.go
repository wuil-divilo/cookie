package repository

import (
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"reflect"
	"testing"
)

func Test_fromModelDevice(t *testing.T) {
	type args struct {
		device *model.Device
	}
	tests := []struct {
		name string
		args args
		want *deviceTable
	}{
		{
			name: "convert to device model",
			args: args{
				device: &model.Device{
					DeviceId:      "deviceId",
					Model:         "model",
					HwVersion:     "hwversion",
					OSVersion:     "osversion",
					AppVersion:    "appversion",
					ApiLevel:      "apilevel",
					SecurityPatch: "securitypatch",
					NFCAvailable:  true,
					NFCEnabled:    false,
					CreatedAt:     1,
					UpdatedAt:     2,
				},
			},
			want: &deviceTable{
				DeviceId: "deviceId",
				Meta: metaDeviceTable{
					Model:         "model",
					HwVersion:     "hwversion",
					OSVersion:     "osversion",
					AppVersion:    "appversion",
					ApiLevel:      "apilevel",
					SecurityPatch: "securitypatch",
					NFCAvailable:  true,
					NFCEnabled:    false,
				},
				CreatedAt: 1,
				UpdatedAt: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fromModelDevice(tt.args.device)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromModelDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toModelDevice(t *testing.T) {
	type args struct {
		deviceTable *deviceTable
	}
	tests := []struct {
		name string
		args args
		want *model.Device
	}{
		{
			name: "success",
			args: args{
				deviceTable: &deviceTable{
					DeviceId: "deviceId",
					Meta: metaDeviceTable{
						Model:         "model",
						HwVersion:     "hwversion",
						OSVersion:     "osversion",
						AppVersion:    "appversion",
						ApiLevel:      "apilevel",
						SecurityPatch: "securitypatch",
						NFCAvailable:  true,
						NFCEnabled:    true,
					},
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			want: &model.Device{
				DeviceId:      "deviceId",
				Model:         "model",
				HwVersion:     "hwversion",
				OSVersion:     "osversion",
				AppVersion:    "appversion",
				ApiLevel:      "apilevel",
				SecurityPatch: "securitypatch",
				NFCAvailable:  true,
				NFCEnabled:    true,
				CreatedAt:     1,
				UpdatedAt:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toModelDevice(tt.args.deviceTable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toModelDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}
