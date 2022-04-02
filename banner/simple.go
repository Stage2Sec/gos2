package banner

import (
	"fmt"
	"time"
)

func Banner(engagement, vuln, details string) {
	dt := time.Now()
	fmt.Print("\033[H\033[2J")
	fmt.Printf("\n\n")
	fmt.Printf("      ▄▄▄▄▄▄▄▄   ▄▄▄▄▄▄▄▄      ||   Stage 2 Security\n")
	fmt.Printf("    ██▀▀▀▀▀▀██▌▐█▀▀▀▀▀▀▀██     ||\n")
	fmt.Printf("   ██ ██▀▀▀▀▀▀  ▀▀▀▀▀▀▀██ █▌   ||   Engagement:    %s\n", engagement)
	fmt.Printf("   ██ ██▄▄▄▄         ▄▄██ █▌   ||   Date:          %s\n", dt.Format(time.UnixDate))
	fmt.Printf("    ██▄▄▄▀▀▀██  ▄██▀▀▀▄▄██     ||   Vulnerability: %s\n", vuln)
	fmt.Printf("      ▀▀▀▀██ █▌▐█▌ ██▀▀        ||   Details:       %s\n", details)
	fmt.Printf("   ▄▄▄▄▄▄▄██ █▌▐█ ██▄▄▄▄▄▄▄    ||\n")
	fmt.Printf("   ██▀▀▀▀▀▀▄▄█ ▐█▌▀▀▀▀▀▀▀██    ||\n")
	fmt.Printf("   ▀▀▀▀▀▀▀▀▀▀   ▀▀▀▀▀▀▀▀▀▀▀    ||\n")
	fmt.Printf("\n\n")
}
