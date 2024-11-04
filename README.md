# [`beagleboard` module](https://github.com/viam-modules/beagleboard)

This [beagleboard module](https://app.viam.com/module/viam/beagleboard) implements a [BeagleBoard BeagleBone AI 64](https://www.beagleboard.org/boards/beaglebone-ai-64) using the [`rdk:component:board` API](https://docs.viam.com/appendix/apis/components/board/).

## Setup

The [BeagleBone AI-64](https://docs.beagleboard.org/latest/boards/beaglebone/ai-64/) from [BeagleBoard.org](https://beagleboard.org/) is an open-source single-board computer with a Debian GNU/Linux operating system based on the Texas Instruments TDA4VM processor.
Follow this guide to set up your BeagleBone AI-64 and prepare it for `viam-server` installation.

### Hardware requirements

You need the following hardware, tools, and software to install `viam-server` on a BeagleBone AI-64:

1. A [BeagleBone AI-64](https://www.beagleboard.org/boards/beaglebone-ai-64)
2. A 5V barrel jack (recommended) and/or USB-C power supply, to power the BeagleBone
3. Ethernet cable and/or WiFi dongle, to establish network connection on the BeagleBone
4. (Optional) A microSD card and a way to connect the microSD card to the computer (like a microSD slot or microSD reader)
   - This is required if you need to set up your BeagleBone for the first time or update your BeagleBone to the latest software image.

The following instructions mirror the instructions given in the [BeagleBoard Quick Start Guide](https://docs.beagleboard.org/latest/boards/beaglebone/ai-64/02-quick-start.html).

If you want additional help setting up your BeagleBone, you can follow the guides there and return to the Viam docs after SSH'ing into your BeagleBone.

### Power your BeagleBone

Power your board by plugging a 5VDC power source into the BeagleBone's barrel jack.
You can also power the BeagleBone with a USB-C cable, but a 5VDC power source is recommended for more reliable performance.

If the board has power, the LED on the board labeled _PWR_ or _ON_ is lit steadily.

### Enable a network connection

You need to enable a network connection on your BeagleBone to install `viam-server` on it.
You can do this in multiple ways:

- Connect an ethernet cable to your BeagleBone's ethernet port.
- If you are working on a macOS machine, use [internet sharing over USB](https://support.apple.com/guide/mac-help/share-internet-connection-mac-network-users-mchlp1540/mac) to share your connection.
  After enabling the option on your machine, SSH into your BeagleBone and run `sudo dhclient usb1`.
- If you are working on a Linux machine, read [these tips on enabling a network connection](https://elinux.org/Beagleboard:Terminal_Shells).
- If your personal computer supports mDNS (Multicast DNS), you can check to see if your BeagleBone board has established a network connection by visiting [beaglebone.local](https://beaglebone.local).

### SSH into your BeagleBone from your PC

You can SSH into your BeagleBone by running the following command in your terminal:

`ssh <your-username>@<your-hostname>.local`

By default, the hostname, username and password on a BeagleBone are:

- Hostname: `beaglebone`
- Username: `debian`
- Password: `temppwd`

Therefore, if you are using the default settings on your BeagleBone, the command is:

`ssh debian@beaglebone.local`

### Update your BeagleBone

After SSH'ing into your BeagleBone, verify all packages are up to date:

`sudo apt update && sudo apt dist-upgrade && sudo reboot`

### Troubleshooting

If you experience any issues getting Viam working on your BeagleBone, consult the [BeagleBone documentation](https://docs.beagleboard.org/latest/boards/beaglebone/ai-64/02-quick-start.html) for help updating your BeagleBone.

You can find additional assistance in the [Troubleshooting section](https://docs.viam.com/appendix/troubleshooting/).

> [!NOTE]
> Before configuring your board, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

## Configure your beaglebone board

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add board / beagleboard:beaglebone to your machine](https://docs.viam.com/configure/#components).

### Attributes

The following attributes are available for `viam:beagleboard:beaglebone` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name.|

For instructions on implementing digital interrupts, see [Digital interrupt configuration](#Digital-interrupt-configuration).

## Example configuration

### `viam:beagleboard:beaglebone`
```json
  {
     "name": "<your-beagleboard-beaglebone-board-name>",
      "model": "viam:beagleboard:beaglebone",
      "type": "board",
      "namespace": "rdk",
      "attributes": {
      },
      "depends_on": []
  }
```

### Next Steps
- To test your board, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your board, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a board component, explore [these tutorials](https://docs.viam.com/tutorials/).

## Digital interrupt configuration
[Interrupts](https://en.wikipedia.org/wiki/Interrupt) are a method of signaling precise state changes.
Configuring digital interrupts to monitor GPIO pins on your board is useful when your application needs to know precisely when there is a change in GPIO value between high and low.

- When an interrupt configured on your board processes a change in the state of the GPIO pin it is configured to monitor, it ticks to record the state change.
  You can stream these ticks with the board API's [`StreamTicks()`](https://docs.viam.com/appendix/apis/components/board/#streamticks), or get the current value of the digital interrupt with [`Value()`](https://docs.viam.com/appendix/apis/components/board/#value).
- Calling [`GetGPIO()`](https://docs.viam.com/appendix/apis/components/board/#getgpio) on a GPIO pin, which you can do without configuring interrupts, is useful when you want to know a pin's value at specific points in your program, but is less precise and convenient than using an interrupt.

Integrate `digital_interrupts` into your machine in the `attributes` of your board by adding the following to your board's `attributes` configuration:

```json {class="line-numbers linkable-line-numbers"}
{
  "digital_interrupts": [
    {
      "name": "<your-digital-interrupt-name>",
      "pin": "<your-digital-interrupt-pin-number>"
    }
  ]
}
```

### Attributes

The following attributes are available for `digital_interrupts`:

<!-- prettier-ignore -->
| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |

### Example configuration

```json {class="line-numbers linkable-line-numbers"}
{
  "components": [
    {
      "name": "<your-beagleboard-beaglebone-board-name>",
      "model": "viam:beagleboard:beaglebone",
      "type": "board",
      "namespace": "rdk",
      "attributes": {
        "digital_interrupts": [
          {
            "name": "your-interrupt-1",
            "pin": "15"
          },
          {
            "name": "your-interrupt-2",
            "pin": "16"
          }
        ]
      }
    }
  ]
}
```
