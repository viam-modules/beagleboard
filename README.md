# [`beagleboard` module](https://github.com/viam-modules/beagleboard)

This [beagleboard module](https://app.viam.com/module/viam/beagleboard) implements a [BeagleBoard BeagleBone AI 64](https://www.beagleboard.org/boards/beaglebone-ai-64) using the [`rdk:component:board` API](https://docs.viam.com/appendix/apis/components/board/).

## Setup

Follow the [setup guide](https://docs.viam.com/installation/prepare/beaglebone-setup/) to prepare your BeagleBone for running `viam-server` before configuring a `beaglebone` board.

> [!NOTE]
> Before configuring your board, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add board / beagleboard:beaglebone to your machine](https://docs.viam.com/configure/#components).

## Configure your beaglebone board

### Configuration
```json
  {}
```

### Configuration with optional digital interrupts
```json
{
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
```
### Attributes
The following attributes are available for `viam:beagleboard:beaglebone` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name.|

### Digital interrupt configuration
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

| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |


## Configure your BeagleY-AI board

In the default OS, all PWM is software-PWM (except for the pins 29, 31, 32, 33 mapped to hardware-PWM). To use hardware-PWM, follow the [BeagleBoard docs](https://docs.beagleboard.org/boards/beagley/ai/demos/using-pwm.html) to set the correct PWM overlays expected by this module.

Add the following line to the board's `/boot/firmware/extlinux/extlinux.conf` to enable hardware PWM on pins 29, 31, 32, 33:

```
fdtoverlays /overlays/k3-am67a-beagley-ai-pwm-epwm0-gpio5-gpio12.dtbo /overlays/k3-am67a-beagley-ai-pwm-epwm1-gpio6-gpio13.dtbo
```

So the resulting section in the config will look something like this:

```
label microSD (default)
    kernel /Image
    append console=ttyS2,115200n8 root=/dev/mmcblk1p3 ro rootfstype=ext4 resume=/dev/mmcblk1p2 rootwait net.ifnames=0 quiet
    fdtdir /
    fdt /ti/k3-am67a-beagley-ai.dtb
    fdtoverlays /overlays/k3-am67a-beagley-ai-pwm-epwm0-gpio5-gpio12.dtbo /overlays/k3-am67a-beagley-ai-pwm-epwm1-gpio6-gpio13.dtbo
    #initrd /initrd.img
```

After rebooting your board, you should have hardware-PWM enabled for those pins.


### Configuration
```json
  {}
```

### Configuration with optional digital interrupts
```json
{
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
```
### Attributes
The following attributes are available for `viam:beagleboard:beagle-y-ai` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name.|

### Digital interrupt configuration
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

| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |


## Next Steps
- To test your board, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your board, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a board component, explore [these tutorials](https://docs.viam.com/tutorials/).
