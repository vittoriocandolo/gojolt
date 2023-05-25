package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/godbus/dbus/v5"
)

const (
	dbusDest      = "org.gnome.SessionManager"
	dbusPath      = "/org/gnome/SessionManager"
	dbusInterface = "org.gnome.SessionManager"
	inhibitReason = "Preventing screen blank"
	inhibitFlags  = 0x8 // inhibit flag for idle
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <minutes>\n", os.Args[0])
		os.Exit(1)
	}

	minutes, err := strconv.Atoi(os.Args[1])
	if err != nil || minutes < 1 {
		fmt.Println("Invalid number of minutes. Please enter a positive integer.")
		os.Exit(1)
	}

	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to session bus: %s\n", err)
		os.Exit(1)
	}

	obj := conn.Object(dbusDest, dbus.ObjectPath(dbusPath))
	var inhibitCookie uint32
	err = obj.Call("org.gnome.SessionManager.Inhibit", 0, os.Args[0], uint32(0), inhibitReason, uint32(inhibitFlags)).Store(&inhibitCookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call Inhibit: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Inhibiting screen blank for %d minutes...\n", minutes)
	time.Sleep(time.Duration(minutes) * time.Minute)

	err = obj.Call("org.gnome.SessionManager.Uninhibit", 0, inhibitCookie).Store()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call Uninhibit: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Screen blank inhibition ended.")
}
