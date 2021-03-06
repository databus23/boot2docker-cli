package driver

import (
	"fmt"
	"net"

	flag "github.com/ogier/pflag"
)

// Machine config.
type MachineConfig struct {
	// Gereral flags.
	Init    bool
	Verbose bool
	Driver  string

	// basic config
	Clobber              bool
	ForceUpgradeDownload bool
	SSH                  string // SSH client executable
	SSHGen               string // SSH keygen executable
	SSHKey               string // SSH key to send to the vm
	VM                   string // virtual machine name
	Dir                  string // boot2docker directory
	ISOURL               string // Source URL to retrieve the ISO from
	ISO                  string // boot2docker ISO image path
	DiskSize             uint   // VM disk image size (MB)
	Memory               uint   // VM memory size (MB)

	// NAT network: port forwarding
	SSHPort    uint16 // host SSH port (forward to port 22 in VM)
	DockerPort uint16 // host Docker port (forward to port 2376 in VM)

	// host-only network
	HostIP      net.IP
	DHCPIP      net.IP
	NetMask     net.IPMask
	LowerIP     net.IP
	UpperIP     net.IP
	DHCPEnabled bool

	// Serial console pipe/socket
	Serial     bool
	SerialFile string

	// boot2docker init retry settings
	Waittime int
	Retries  int

	DriverCfg map[string]interface{}
}

type ConfigFunc func(B2D *MachineConfig, flags *flag.FlagSet) error

var configs map[string]ConfigFunc // optional map of driver ConfigFunc

func init() {
	configs = make(map[string]ConfigFunc)
}

// optional - allows a driver to add its own commandline parameters
func RegisterConfig(driver string, configFunc ConfigFunc) error {
	if _, exists := configs[driver]; exists {
		return fmt.Errorf("Driver already registered %s", driver)
	}
	configs[driver] = configFunc

	return nil
}

func ConfigFlags(B2D *MachineConfig, flags *flag.FlagSet) error {
	for _, configFunc := range configs {
		if err := configFunc(B2D, flags); err != nil {
			return err
		}
	}
	return nil
}
