package main

import (
	"github.com/EHerzog76/machine-driver-hyper-v/machine/driver"

	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(driver.NewDriver("", ""))
}
