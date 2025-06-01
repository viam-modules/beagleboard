package beagley

import "go.viam.com/rdk/components/board/genericlinux"

const bbYAi = "bb_YAI"

var boardInfoMappings = map[string]genericlinux.BoardInformation{
	bbYAi: {
		PinDefinitions: []genericlinux.PinDefinition{
			{Name: "3", DeviceName: "gpiochip0", LineNumber: 18, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "5", DeviceName: "gpiochip0", LineNumber: 17, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "7", DeviceName: "gpiochip1", LineNumber: 38, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "8", DeviceName: "gpiochip2", LineNumber: 14, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "10", DeviceName: "gpiochip2", LineNumber: 13, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "11", DeviceName: "gpiochip2", LineNumber: 8, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "12", DeviceName: "gpiochip2", LineNumber: 11, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "13", DeviceName: "gpiochip1", LineNumber: 33, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "15", DeviceName: "gpiochip1", LineNumber: 41, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "16", DeviceName: "gpiochip0", LineNumber: 7, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "18", DeviceName: "gpiochip0", LineNumber: 10, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "19", DeviceName: "gpiochip0", LineNumber: 3, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "21", DeviceName: "gpiochip0", LineNumber: 4, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "22", DeviceName: "gpiochip1", LineNumber: 42, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "23", DeviceName: "gpiochip0", LineNumber: 2, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "24", DeviceName: "gpiochip0", LineNumber: 0, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "26", DeviceName: "gpiochip0", LineNumber: 9, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "29", DeviceName: "gpiochip2", LineNumber: 15, PwmChipSysfsDir: "23000000.pwm", PwmID: 0},
			{Name: "31", DeviceName: "gpiochip2", LineNumber: 17, PwmChipSysfsDir: "23010000.pwm", PwmID: 0},
			{Name: "32", DeviceName: "gpiochip2", LineNumber: 16, PwmChipSysfsDir: "23000000.pwm", PwmID: 1},
			{Name: "33", DeviceName: "gpiochip2", LineNumber: 18, PwmChipSysfsDir: "23010000.pwm", PwmID: 1},
			{Name: "35", DeviceName: "gpiochip2", LineNumber: 12, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "36", DeviceName: "gpiochip2", LineNumber: 7, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "37", DeviceName: "gpiochip1", LineNumber: 36, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "38", DeviceName: "gpiochip2", LineNumber: 10, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "40", DeviceName: "gpiochip2", LineNumber: 9, PwmChipSysfsDir: "", PwmID: -1},
		},
		Compats: []string{"beagle,am67a-beagley-ai", "ti,j722s"},
	},
}
