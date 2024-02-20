package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	green = "\033[0;32m"
	reset = "\033[0m"
	red   = "\033[0;31m"
	blue  = "\033[0;34m"
	tick  = "\u2713"
	fail  = "\u2717"
)

func checkInternetConnection() bool {
	cmd := exec.Command("ping", "-c", "1", "8.8.8.8")
	err := cmd.Run()
	return err == nil
}

func checkDomainConnection(url string) bool {
	cmd := exec.Command("ping", "-c", "1", url)
	err := cmd.Run()
	return err == nil
}

func checkToolInstalled(tool string) bool {
	_, err := exec.LookPath(tool)
	return err == nil
}

func installGAU() {
	fmt.Printf("%s[+]%s GAU is not installed. Installing GAU ...\n", blue, reset)
	cmd := exec.Command("wget", "-q", "https://github.com/lc/gau/releases/download/v2.1.2/gau_2.1.2_linux_amd64.tar.gz")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s[+]%s Error downloading GAU: %s\n", red, reset, err)
		os.Exit(1)
	}
	cmd = exec.Command("tar", "xvf", "gau_2.1.2_linux_amd64.tar.gz")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s[+]%s Error extracting GAU: %s\n", red, reset, err)
		os.Exit(1)
	}
	cmd = exec.Command("sudo", "mv", "gau", "/usr/bin/gau")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s[+]%s Error moving GAU to /usr/bin/gau: %s\n", red, reset, err)
		os.Exit(1)
	}
	fmt.Printf("%s[+]%s GAU installed successfully!\n", blue, reset)
}

func gatherEndpoints(url, folder, domain string) {
	fmt.Printf("%s[+]%s Gathering The End-Points for %s%s%s\n", blue, reset, green, url, reset)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.Mkdir(folder, 0755); err != nil {
			fmt.Printf("%s[+]%s Error creating folder: %s\n", red, reset, err)
			os.Exit(1)
		}
	}

	for {
		gauOutput, err := exec.Command("gau", url).Output()
		if err != nil {
			fmt.Printf("%s[+]%s Error running gau: %s\n", red, reset, err)
			os.Exit(1)
		}

		if len(gauOutput) > 0 {
			endpoints := strings.Join(strings.Fields(string(gauOutput)), "\n")
			filePath := fmt.Sprintf("%s/%s.txt", folder, domain)
			if err := os.WriteFile(filePath, []byte(endpoints), 0644); err != nil {
				fmt.Printf("%s[+]%s Error writing endpoints to file: %s\n", red, reset, err)
				os.Exit(1)
			}
			fmt.Printf("%s[+]%s Endpoints updated at %s\n", green, reset, time.Now().Format(time.RFC3339))
		} else {
			fmt.Printf("%s[+]%s No EndPoints Found\n", red, reset)
			fmt.Printf("%s[+]%s Exiting\n", blue, reset)
			os.Exit(1)
		}

		// Sleep for 10 seconds
		time.Sleep(10 * time.Second)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [+]website[+]\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]
	folder := "trj_output"
	domain := strings.Split(url, ".")[0]

	// Checking the Internet Connection
	fmt.Printf("%s[+]%s Checking the Internet Connection ...\n", blue, reset)
	if !checkInternetConnection() {
		fmt.Printf("Internet Check Failed[%s%s%s]\n", red, fail, reset)
		os.Exit(1)
	}
	fmt.Printf("Internet Check Done [%s%s%s]\n", green, tick, reset)

	// Checking the Domain
	fmt.Printf("%s[+]%s Checking the Domain.\n", blue, reset)
	if !checkDomainConnection(url) {
		fmt.Printf("Domain Check Failed![%s%s%s]\n", red, fail, reset)
		os.Exit(1)
	}
	fmt.Printf("Domain Check Done[%s%s%s]\n", green, tick, reset)

	// Specify the name of the tools
	uro := "uro"
	gau := "gau"

	// Check if uro is installed
	fmt.Printf("%s[+]%s Checking if %s is installed...\n", red, reset, uro)
	if !checkToolInstalled(uro) {
		fmt.Printf("%s[+]%s %s is not Installed... Please Install it and rerun this Script...\n", red, reset, uro)
		os.Exit(1)
	}
	fmt.Printf("%s[+]%s %s is installed.\n", green, reset, uro)

	// Check if gau is installed
	fmt.Printf("%s[+]%s Checking if %s is installed...\n", red, reset, gau)
	if !checkToolInstalled(gau) {
		installGAU()
	}
	fmt.Printf("%s[+]%s %s is installed.\n", green, reset, gau)

	// Gather endpoints with a 10-second timer loop
	gatherEndpoints(url, folder, domain)
}
