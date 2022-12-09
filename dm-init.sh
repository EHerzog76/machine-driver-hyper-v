alias dm=docker-machine

dmPREFIX="hyper-v_"
read "ENDPOINT?ENDPOINT: "
export ${dmPREFIX}ENDPOINT
read "CLUSTER?CLUSTER: "
export ${dmPREFIX}CLUSTER
export ${dmPREFIX}INSECURE=true
read "VM_IMAGE?VM_IMAGE: "
export ${dmPREFIX}VM_IMAGE
export ${dmPREFIX}VM_IMAGE_SIZE=20
read "VM_NETWORK?VM_NETWORK: "
export ${dmPREFIX}VM_NETWORK
read "USERNAME?USERNAME: "
export ${dmPREFIX}USERNAME
read -s "PASSWORD?PASSWORD: "
export ${dmPREFIX}PASSWORD
