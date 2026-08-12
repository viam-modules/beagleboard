package beaglebone

import "go.viam.com/rdk/components/board/genericlinux"

const bbAi = "bb_Ai64"

const (
	gpioChip1 = "gpiochip1"
	gpioChip2 = "gpiochip2"
)

var boardInfoMappings = map[string]genericlinux.BoardInformation{
	bbAi: {
		PinDefinitions: []genericlinux.PinDefinition{
			// PinNumberBoard {914} -> PinNameCVM3 "P9_14"

			// ******** DATA MAPPING ********************************
			// Hardware PWMs contain a number other than -1 in the last element of map
			// beaglebone pwm mapping sys/devices/platform/bus@100000/*.pwm
			// NOTE: each hardware PWM device can only drive 1 line at a time, even though 2 lines
			// are hooked up to each. For example, you can't have PWM signals running on lines 914
			// and 916 at the same time, even though both of them work on their own.
			// NOTE: pins with hardware PWM support don't work as GPIO input.

			// GPIO only pins
			// beaglebone gpio mapping uses directory sys/devices/platform/bus@100000/*.gpio
			// beaglebone 600000.gpio/ (128 lines) corresponds to gpiochip1 and /sys/class/gpio/gpiochip300/
			// beaglebone 601000.gpio/ (36 lines) corresponds to gpiochip2 and /sys/class/gpio/gpiochip264/

			{Name: "803", DeviceName: gpioChip1, LineNumber: 20, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "804", DeviceName: gpioChip1, LineNumber: 48, PwmChipSysfsDir: "", PwmID: -1}, // BOOTMODE2
			{Name: "805", DeviceName: gpioChip1, LineNumber: 33, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "806", DeviceName: gpioChip1, LineNumber: 34, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "807", DeviceName: gpioChip1, LineNumber: 15, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "808", DeviceName: gpioChip1, LineNumber: 14, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "809", DeviceName: gpioChip1, LineNumber: 17, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "810", DeviceName: gpioChip1, LineNumber: 16, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "811", DeviceName: gpioChip1, LineNumber: 60, PwmChipSysfsDir: "", PwmID: -1}, // BOOTMODE7
			{Name: "812", DeviceName: gpioChip1, LineNumber: 59, PwmChipSysfsDir: "", PwmID: -1},
			// Note that gpio input cannot be used on this pin.
			{Name: "813", DeviceName: gpioChip1, LineNumber: 89, PwmChipSysfsDir: "3000000.pwm", PwmID: 1}, // pwmchip0 V27 EHRPWM0_A
			{Name: "814", DeviceName: gpioChip1, LineNumber: 75, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "815", DeviceName: gpioChip1, LineNumber: 61, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "816", DeviceName: gpioChip1, LineNumber: 62, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "817", DeviceName: gpioChip1, LineNumber: 3, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "818", DeviceName: gpioChip1, LineNumber: 4, PwmChipSysfsDir: "", PwmID: -1},
			// Note that gpio input cannot be used on this pin.
			{Name: "819", DeviceName: gpioChip1, LineNumber: 88, PwmChipSysfsDir: "3000000.pwm", PwmID: 0}, // pwmchip0 V29 EHRPWM0_B
			{Name: "820", DeviceName: gpioChip1, LineNumber: 76, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "821", DeviceName: gpioChip1, LineNumber: 30, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "822", DeviceName: gpioChip1, LineNumber: 5, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "823", DeviceName: gpioChip1, LineNumber: 31, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "824", DeviceName: gpioChip1, LineNumber: 6, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "825", DeviceName: gpioChip1, LineNumber: 35, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "826", DeviceName: gpioChip1, LineNumber: 51, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "827", DeviceName: gpioChip1, LineNumber: 71, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "828", DeviceName: gpioChip1, LineNumber: 72, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "829", DeviceName: gpioChip1, LineNumber: 73, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "830", DeviceName: gpioChip1, LineNumber: 74, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "831", DeviceName: gpioChip1, LineNumber: 32, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "832", DeviceName: gpioChip1, LineNumber: 26, PwmChipSysfsDir: "", PwmID: -1}, // Timer-based PWM
			{Name: "833", DeviceName: gpioChip1, LineNumber: 25, PwmChipSysfsDir: "", PwmID: -1}, // Timer-based PWM
			{Name: "834", DeviceName: gpioChip1, LineNumber: 7, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "835", DeviceName: gpioChip1, LineNumber: 24, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "836", DeviceName: gpioChip1, LineNumber: 8, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "837", DeviceName: gpioChip1, LineNumber: 11, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "838", DeviceName: gpioChip1, LineNumber: 9, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "839", DeviceName: gpioChip1, LineNumber: 69, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "840", DeviceName: gpioChip1, LineNumber: 70, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "841", DeviceName: gpioChip1, LineNumber: 67, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "842", DeviceName: gpioChip1, LineNumber: 68, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "843", DeviceName: gpioChip1, LineNumber: 65, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "844", DeviceName: gpioChip1, LineNumber: 66, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "845", DeviceName: gpioChip1, LineNumber: 79, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "846", DeviceName: gpioChip1, LineNumber: 80, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "911", DeviceName: gpioChip1, LineNumber: 1, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "912", DeviceName: gpioChip1, LineNumber: 45, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "913", DeviceName: gpioChip1, LineNumber: 2, PwmChipSysfsDir: "", PwmID: -1},
			// Note that gpio input cannot be used on this pin.
			{Name: "914", DeviceName: gpioChip1, LineNumber: 93, PwmChipSysfsDir: "3020000.pwm", PwmID: 0},
			{Name: "915", DeviceName: gpioChip1, LineNumber: 47, PwmChipSysfsDir: "", PwmID: -1},
			// Note that gpio input cannot be used on this pin.
			{Name: "916", DeviceName: gpioChip1, LineNumber: 94, PwmChipSysfsDir: "3020000.pwm", PwmID: 1},
			{Name: "917", DeviceName: gpioChip1, LineNumber: 28, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "918", DeviceName: gpioChip1, LineNumber: 40, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "919", DeviceName: gpioChip1, LineNumber: 78, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "920", DeviceName: gpioChip1, LineNumber: 77, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "921", DeviceName: gpioChip1, LineNumber: 39, PwmChipSysfsDir: "3010000.pwm", PwmID: 0},
			{Name: "922", DeviceName: gpioChip1, LineNumber: 38, PwmChipSysfsDir: "3010000.pwm", PwmID: 1},
			{Name: "923", DeviceName: gpioChip1, LineNumber: 10, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "924", DeviceName: gpioChip1, LineNumber: 13, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "925", DeviceName: gpioChip1, LineNumber: 127, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "926", DeviceName: gpioChip1, LineNumber: 12, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "927", DeviceName: gpioChip1, LineNumber: 46, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "928", DeviceName: gpioChip1, LineNumber: 43, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "929", DeviceName: gpioChip2, LineNumber: 14, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "930", DeviceName: gpioChip2, LineNumber: 13, PwmChipSysfsDir: "", PwmID: -1}, // Timer-based PWM
			{Name: "931", DeviceName: gpioChip1, LineNumber: 52, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "933", DeviceName: gpioChip1, LineNumber: 50, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "935", DeviceName: gpioChip1, LineNumber: 55, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "936", DeviceName: gpioChip1, LineNumber: 56, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "937", DeviceName: gpioChip1, LineNumber: 57, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "938", DeviceName: gpioChip1, LineNumber: 58, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "939", DeviceName: gpioChip1, LineNumber: 54, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "940", DeviceName: gpioChip1, LineNumber: 81, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "941", DeviceName: gpioChip2, LineNumber: 0, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "942", DeviceName: gpioChip1, LineNumber: 123, PwmChipSysfsDir: "", PwmID: -1}, // Timer-based PWM
		},
		Compats: []string{"beagle,j721e-beagleboneai64", "ti,j721e"},
	},
}
