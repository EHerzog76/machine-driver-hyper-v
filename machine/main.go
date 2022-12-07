package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/EHerzog76/docker-machine/machine/driver"
)

func main() {
	plugin.RegisterDriver(driver.NewDriver("", ""))
}
