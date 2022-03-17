package internal

import (
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"reflect"
	"testing"
)

func Test_toModelDevice(t *testing.T) {
	type args struct {
		dvcReq *model.CreateDeviceRequest
	}
	tests := []struct {
		name string
		args args
		want *model.Device
	}{
		{
			name: "convert to device model",
			args: args{
				dvcReq: &model.CreateDeviceRequest{
					ApiLevel:               "apilevel",
					AppVersion:             "appversion",
					DeviceId:               "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
					HardwareVersion:        "hardware version",
					IsNfcAvailable:         true,
					IsNfcEnabled:           false,
					Model:                  "model1",
					OperatingSystemVersion: "android os x.x",
					SecurityPatch:          "2021-10-25",
				},
			},
			want: &model.Device{
				DeviceId:      "1f066b37-c8f5-40cb-bf7b-5b7eda60dd27",
				Model:         "model1",
				HwVersion:     "hardware version",
				OSVersion:     "android os x.x",
				AppVersion:    "appversion",
				ApiLevel:      "apilevel",
				SecurityPatch: "2021-10-25",
				NFCAvailable:  true,
				NFCEnabled:    false,
				CreatedAt:     0,
				UpdatedAt:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toModelDevice(tt.args.dvcReq)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toModelDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}
