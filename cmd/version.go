package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current gg version and build information",
	Long: `Display the current gg version and build information like
	os , golang version, current kernel, etc.`,

	Run: func(cmd *cobra.Command, args []string) {
		printBuildContext()
		return
	},
}

func printBuildContext() {
	var (
		version    = "gg0.10"
		gitVersion []byte
		goVersion  []byte
		kernel     []byte
		os         []byte
	)

	goVersion, _ = exec.Command("go", []string{"version"}...).Output()
	kernel, _ = exec.Command("uname", []string{"-r"}...).Output()
	gitVersion, _ = exec.Command("git", []string{"version"}...).Output()
	os, _ = exec.Command("uname", []string{"-s"}...).Output()

	fmt.Println("Build Information:")
	fmt.Println("------------------")
	fmt.Printf("Version: %v\n", version)
	fmt.Printf("Go Version: %v", string(goVersion))
	fmt.Printf("Git Version: %v", string(gitVersion))
	fmt.Printf("Current Kernel Version: %v", string(kernel))
	fmt.Printf("Current OS Type: %v", string(os))
}
