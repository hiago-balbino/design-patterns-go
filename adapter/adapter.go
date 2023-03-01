package main

import "fmt"

// Computer is an interface to handle the lightning connector of any computer type.
type Computer interface {
	// ConnectLightningPort is a function that plugs the lightning connector into the computer port.
	ConnectLightningPort()
}

// DesktopMachine is a struct that represents a desktop machine and implements the Computer interface.
type DesktopMachine struct{}

// ConnectLightningPort is a function that plugs the lightning connector into the desktop computer port.
func (m *DesktopMachine) ConnectLightningPort() {
	fmt.Println("Lightning connector is plugged into desktop machine.")
}

// LaptopMachine is a struct that represents a laptop machine and implements the Computer interface.
type LaptopMachine struct{}

func (w *LaptopMachine) connectUSBPort() {
	fmt.Println("USB connector is plugged into laptop machine.")
}

// LaptopMachineAdapter is a struct that represents an adapter of a laptop machine and implements the Computer interface.
type LaptopMachineAdapter struct {
	laptopMachine *LaptopMachine
}

// ConnectLightningPort is a function that plugs the lightning connector into the laptop computer port.
func (w *LaptopMachineAdapter) ConnectLightningPort() {
	fmt.Println("Adapter converts lightning signal to USB.")
	w.laptopMachine.connectUSBPort()
}

// Client is a struct to handle any type of computer to connect lightning connectors.
type Client struct{}

// ConnectLightningConnectorIntoComputer is an intermediate function that connects a lightning connector to any type of computer received on a parameter.
func (c *Client) ConnectLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lightning connector into computer.")
	com.ConnectLightningPort()
}

func main() {
	client := &Client{}

	desktopMachine := &DesktopMachine{}
	client.ConnectLightningConnectorIntoComputer(desktopMachine)

	laptopMachine := &LaptopMachine{}
	laptopMachineAdapter := &LaptopMachineAdapter{laptopMachine: laptopMachine}
	client.ConnectLightningConnectorIntoComputer(laptopMachineAdapter)
}
