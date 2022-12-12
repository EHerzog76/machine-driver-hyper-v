module docker-machine-driver-hyper-v

go 1.19

require github.com/docker/machine v0.16.2

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/docker/docker v20.10.21+incompatible // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace github.com/EHerzog76/machine-driver-hyper-v/machine/driver => ../machine-driver-hyper-v/machine/driver
