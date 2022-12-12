package main

import (
	"docker-machine-driver-hyper-v/machine/driver"

	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(driver.NewDriver("", ""))
}
