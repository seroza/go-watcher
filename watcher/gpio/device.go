package gpio

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

var once sync.Once

type device struct {
}

func (d *device) init() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	preparePins()

	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

}

func preparePins() {
	for i := 2; i <= 27; i++ {
		pin := rpio.Pin(i)
		pin.Output()
		pin.High()
	}
}

func convertRigNumToPinNum(rig string) (int, error) {
	return strconv.Atoi(rig)
}

func delay(milliSec time.Duration) {
	time.Sleep(time.Millisecond * milliSec)
}

func implementCommand(implFunc func(rpio.Pin), rig string) {
	pinNum, err := convertRigNumToPinNum(rig)
	if err == nil {
		pin := rpio.Pin(pinNum)
		implFunc(pin)
	} else {
		log.Println("can not convert rig num to gpio pin")
	}
}
func (d *device) Reboot(rig string) {
	log.Println("reboot command processing")
	command := func(pin rpio.Pin) {
		pin.Low()
		delay(5000)
		pin.High()
		delay(1000)
		pin.Low()
		delay(1000)
		pin.High()
	}

	implementCommand(command, rig)
	log.Println("done")
}

func (d *device) TurnOnOff(rig string) {
	log.Println("turn on/off command processing")
	command := func(pin rpio.Pin) {
		pin.Low()
		delay(1000)
		pin.High()
	}

	implementCommand(command, rig)
	log.Println("done")
}

func (d *device) HardTurnOff(rig string) {
	log.Println("hard turn-off command processing")
	command := func(pin rpio.Pin) {
		pin.Low()
		delay(5000)
		pin.High()
	}

	implementCommand(command, rig)
	log.Println("done")
}

//DeviceInstance returns instance of device
func DeviceInstance() (instance *device) {
	once.Do(func() {
		instance = new(device)
		instance.init()
	})
	return instance
}
