module docker-machine-driver-hyper-v

go 1.19

require github.com/docker/machine v0.16.2

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/EHerzog76/machine-driver-hyper-v/machine/driver v0.0.0-20221212084813-dfeaaa6e9214 // indirect
	github.com/EHerzog76/machine-driver-hyper-v/machine/driver/client v0.0.0-20221212090225-ac7c9140d064 // indirect
	github.com/PaesslerAG/gval v1.0.0 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1 // indirect
	github.com/docker/docker v20.10.21+incompatible // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/zapr v1.2.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/EHerzog76/machine-driver-hyper-v/machine/driver => ../machine-driver-hyper-v/machine/driver
