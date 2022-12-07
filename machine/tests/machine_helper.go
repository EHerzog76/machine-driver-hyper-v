package tests

import (
	"fmt"
	"os/exec"
)

// const (
// 	driver = "hyper-v"
// )

// MachineConfig structure
type MachineConfig struct {
	VMMem       int
	VMCPUs      int
	VMCores     int
	MachineName string
	VNICs       []string
	VDisks      []string
	SSHUser     string
	SSHPass     string
	User string
	Pass string
	Endpoint    string
}

// CreateMachine test new machine creation
func (c MachineConfig) CreateMachine() ([]byte, error) {
	vmMem := c.VMMem
	if vmMem == 0 {
		vmMem = 1024
	}
	vmCores := c.VMCores
	if vmCores == 0 {
		vmCores = 1
	}
	vmCPUs := c.VMCPUs
	if vmCPUs == 0 {
		vmCPUs = 1
	}
	args := []string{
		"create",
		"-d",
		"hyper-v",
		"--username",
		c.User,
		"--password",
		c.Pass,
		"--endpoint",
		c.Endpoint,
		"--vm-mem",
		fmt.Sprintf("%d", vmMem),
		"--vm-cpus",
		fmt.Sprintf("%d", vmCPUs),
		"--vm-cores",
		fmt.Sprintf("%d", vmCores),
		"--vm-ssh-username",
		c.SSHUser,
		"--vm-ssh-password",
		c.SSHPass,
	}

	for _, nic := range c.VNICs {
		args = append(args, "--vm-network-uuid", nic)
	}

	for _, disk := range c.VDisks {
		args = append(args, "--vm-disk-uuid", disk)
	}

	args = append(args, c.MachineName)

	cmd := exec.Command("docker-machine", args...)

	fmt.Println(cmd.Args)

	return cmd.CombinedOutput()
}

// DeleteMachine test delete machine creation
func (c MachineConfig) DeleteMachine() ([]byte, error) {
	args := []string{
		"rm",
		c.MachineName,
	}
	cmd := exec.Command("docker-machine", args...)
	return cmd.CombinedOutput()
}

// StartMachine test start machine
func (c MachineConfig) StartMachine() ([]byte, error) {
	args := []string{
		"start",
		c.MachineName,
	}
	cmd := exec.Command("docker-machine", args...)
	return cmd.CombinedOutput()
}

// StopMachine test stop machine
func (c MachineConfig) StopMachine() ([]byte, error) {
	args := []string{
		"stop",
		c.MachineName,
	}
	cmd := exec.Command("docker-machine", args...)
	return cmd.CombinedOutput()
}
