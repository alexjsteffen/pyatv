// Command atvremote is a CLI tool for interacting with Apple TV devices.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/alexjsteffen/goatv/pkg/pyatv"
)

var (
	host       = flag.String("s", "", "IP address or hostname of the device")
	name       = flag.String("n", "", "Name of the device")
	timeout    = flag.Duration("t", 5*time.Second, "Timeout for scanning")
	identifier = flag.String("id", "", "Device identifier")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	cmd := args[0]
	ctx := context.Background()

	switch cmd {
	case "scan":
		runScan(ctx)
	case "playing":
		runPlaying(ctx)
	case "pause":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Pause(ctx)
		})
	case "play":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Play(ctx)
		})
	case "next":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Next(ctx)
		})
	case "previous":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Previous(ctx)
		})
	case "up":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Up(ctx, pyatv.InputActionSingleTap)
		})
	case "down":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Down(ctx, pyatv.InputActionSingleTap)
		})
	case "left":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Left(ctx, pyatv.InputActionSingleTap)
		})
	case "right":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Right(ctx, pyatv.InputActionSingleTap)
		})
	case "select":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Select(ctx, pyatv.InputActionSingleTap)
		})
	case "menu":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Menu(ctx, pyatv.InputActionSingleTap)
		})
	case "home":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().Home(ctx, pyatv.InputActionSingleTap)
		})
	case "top_menu":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.RemoteControl().TopMenu(ctx)
		})
	case "volume_up":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.Audio().VolumeUp(ctx)
		})
	case "volume_down":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.Audio().VolumeDown(ctx)
		})
	case "turn_on":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.Power().TurnOn(ctx, false)
		})
	case "turn_off":
		runCommand(ctx, func(atv pyatv.AppleTV) error {
			return atv.Power().TurnOff(ctx, false)
		})
	case "app_list":
		runAppList(ctx)
	case "power_state":
		runPowerState(ctx)
	case "version":
		fmt.Printf("goatv version %s\n", pyatv.Version)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: atvremote [options] <command>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -s <address>    IP address or hostname of the device")
	fmt.Println("  -n <name>       Name of the device")
	fmt.Println("  -id <id>        Device identifier")
	fmt.Println("  -t <duration>   Timeout for scanning (default: 5s)")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  scan            Scan for Apple TV devices on the network")
	fmt.Println("  playing         Show what is currently playing")
	fmt.Println("  pause           Pause playback")
	fmt.Println("  play            Start playback")
	fmt.Println("  next            Skip to next track")
	fmt.Println("  previous        Go to previous track")
	fmt.Println("  up              Press up button")
	fmt.Println("  down            Press down button")
	fmt.Println("  left            Press left button")
	fmt.Println("  right           Press right button")
	fmt.Println("  select          Press select button")
	fmt.Println("  menu            Press menu button")
	fmt.Println("  home            Press home button")
	fmt.Println("  top_menu        Go to top menu")
	fmt.Println("  volume_up       Increase volume")
	fmt.Println("  volume_down     Decrease volume")
	fmt.Println("  turn_on         Turn on device")
	fmt.Println("  turn_off        Turn off device")
	fmt.Println("  app_list        List installed apps")
	fmt.Println("  power_state     Show power state")
	fmt.Println("  version         Show version")
	fmt.Println("  help            Show this help")
}

func runScan(ctx context.Context) {
	opts := pyatv.ScanOptions{
		Timeout: *timeout,
	}

	if *identifier != "" {
		opts.Identifier = *identifier
	}

	if *host != "" {
		opts.Hosts = []string{*host}
	}

	devices, err := pyatv.Scan(ctx, opts)
	if err != nil {
		fmt.Printf("Error scanning: %v\n", err)
		os.Exit(1)
	}

	if len(devices) == 0 {
		fmt.Println("No devices found")
		return
	}

	fmt.Printf("Found %d device(s):\n\n", len(devices))
	for i, device := range devices {
		fmt.Printf("%d. %s\n", i+1, device.Name)
		fmt.Printf("   Address: %s\n", device.Address)
		if device.DeviceInfo != nil {
			fmt.Printf("   Model: %s\n", device.DeviceInfo.Model)
			if device.DeviceInfo.Version != "" {
				fmt.Printf("   Version: %s\n", device.DeviceInfo.Version)
			}
		}
		fmt.Printf("   Identifier: %s\n", device.Identifier)
		fmt.Printf("   Services:\n")
		for _, svc := range device.Services {
			fmt.Printf("     - %s (port %d)\n", svc.Protocol, svc.Port)
		}
		fmt.Println()
	}
}

func findDevice(ctx context.Context) (*pyatv.Config, error) {
	opts := pyatv.ScanOptions{
		Timeout: *timeout,
	}

	if *identifier != "" {
		opts.Identifier = *identifier
	}

	if *host != "" {
		opts.Hosts = []string{*host}
	}

	devices, err := pyatv.Scan(ctx, opts)
	if err != nil {
		return nil, err
	}

	if len(devices) == 0 {
		return nil, fmt.Errorf("no devices found")
	}

	// If name is specified, find by name
	if *name != "" {
		for _, d := range devices {
			if d.Name == *name {
				return d, nil
			}
		}
		return nil, fmt.Errorf("device '%s' not found", *name)
	}

	// Return first device
	return devices[0], nil
}

func connectDevice(ctx context.Context) (pyatv.AppleTV, error) {
	device, err := findDevice(ctx)
	if err != nil {
		return nil, err
	}

	atv, err := pyatv.Connect(ctx, device, pyatv.ConnectOptions{})
	if err != nil {
		return nil, err
	}

	return atv, nil
}

func runCommand(ctx context.Context, fn func(pyatv.AppleTV) error) {
	atv, err := connectDevice(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer atv.Close()

	if err := fn(atv); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("OK")
}

func runPlaying(ctx context.Context) {
	atv, err := connectDevice(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer atv.Close()

	playing, err := atv.Metadata().Playing(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("  Media type: %s\n", playing.MediaType)
	fmt.Printf("Device state: %s\n", playing.DeviceState)

	if playing.Title != "" {
		fmt.Printf("       Title: %s\n", playing.Title)
	}
	if playing.Artist != "" {
		fmt.Printf("      Artist: %s\n", playing.Artist)
	}
	if playing.Album != "" {
		fmt.Printf("       Album: %s\n", playing.Album)
	}
	if playing.Genre != "" {
		fmt.Printf("       Genre: %s\n", playing.Genre)
	}

	if playing.Position != nil && playing.TotalTime != nil {
		percent := float64(*playing.Position) / float64(*playing.TotalTime) * 100
		fmt.Printf("    Position: %d/%ds (%.1f%%)\n", *playing.Position, *playing.TotalTime, percent)
	}

	if playing.Repeat != nil {
		fmt.Printf("      Repeat: %s\n", playing.Repeat)
	}
	if playing.Shuffle != nil {
		fmt.Printf("     Shuffle: %s\n", playing.Shuffle)
	}
}

func runAppList(ctx context.Context) {
	atv, err := connectDevice(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer atv.Close()

	apps, err := atv.Apps().AppList(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(apps) == 0 {
		fmt.Println("No apps found")
		return
	}

	fmt.Printf("Installed apps:\n")
	for _, app := range apps {
		fmt.Printf("  - %s (%s)\n", app.Name, app.Identifier)
	}
}

func runPowerState(ctx context.Context) {
	atv, err := connectDevice(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer atv.Close()

	state := atv.Power().PowerState()
	fmt.Printf("Power state: %s\n", state)
}
