package jedlik

import "fmt"

func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	noOfRounds := (trackDistance + c.speed - 1) / c.speed
	canCarFinish := c.battery >= (noOfRounds * c.batteryDrain)
	return canCarFinish
}
