package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)


func notify() error {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return err
	}
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call(
		"org.freedesktop.Notifications.Notify",
		0,
		"systemd-mon",
		uint32(0), // replaces id 
		"", // app icon
		"hello",
		"there, world",
		[]string{}, // actions
		map[string]dbus.Variant{}, // hints
		int32(5000),
	)
	if call.Err != nil {
		return fmt.Errorf("error sending notification: %w", call.Err)
	}
	var ret uint32
	err = call.Store(&ret)
	if err != nil {
		return fmt.Errorf("error getting uint32 ret value: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("Hello")
	notify()
	fmt.Println("Sent")
}
