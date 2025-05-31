package partyrobot

import "fmt"

func switchOperatingSystem(operatingSystem string) string {
	switch operatingSystem {
	case "windows":
		return fmt.Sprintf("The operating system is %s", operatingSystem)
	case "linux":
		return fmt.Sprintf("The operating system is %s", operatingSystem)
	case "macos":
		return fmt.Sprintf("The operating system is %s", operatingSystem)
	default:
		return fmt.Sprintf("The operating system is %s", operatingSystem)
	}
}
