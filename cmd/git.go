package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func status() string {
	command := []string{"status"}
	stdout, err := runCmd(command)
	if err != nil {
		errCmd(command, err)
	}
	return fmt.Sprintf("\n%v\n", string(stdout))
}

func addAll() {
	command := []string{"add", "."}
	_, err := runCmd(command)
	if err != nil {
		errCmd(command, err)
	}
}

func addSelected(files []string) {
	for _, file := range files {
		command := []string{"add", file}
		_, err := runCmd(command)
		if err != nil {
			errCmd(command, err)
		}
	}
}

func commit(message string) {
	command := []string{"commit", "-m", message}
	_, err := runCmd(command)
	if err != nil {
		errCmd(command, err)
	}
}

func push() {
	command := []string{"push"}
	out, err := runCmd(command)
	if err != nil {
		errCmd(command, err)
	}
	fmt.Println(out)
	fmt.Println("Pushing: Ok")
}

func isDiverged() bool {
	stdout, _ := runCmd([]string{"remote", "update"})
	fmt.Printf("Updating Remotes %v \n", string(stdout))

	local, _ := runCmd([]string{"rev-parse", "@"})
	remote, _ := runCmd([]string{"rev-parse", "@{u}"})

	if bytes.Equal(local, remote) {
		return false
	}

	fmt.Printf("LOCAL *HEAD REF: %v \n ", string(local))
	fmt.Printf("REMOTE *HEAD REF: %v \n ", string(remote))
	return true
}

func isRepo() bool {
	command := []string{"rev-parse", "--is-inside-work-tree"}
	stdout, _ := runCmd(command)
	if len(stdout) == 0 {
		return false
	}
	return true
}

func getOptions() (payload []string) {

	unstaged := []string{"diff", "--name-only"}
	stdout, err := runCmd(unstaged)
	if err != nil {
		errCmd(unstaged, err)
	}

	untracked := []string{"ls-files", "--others", "--exclude-standard"}
	untrackedOut, err := runCmd(untracked)
	if err != nil {
		errCmd(untracked, err)
	}

	for _, u := range untrackedOut {
		stdout = append(stdout, u)
	}

	subfields := bytes.Fields(stdout)
	for _, field := range subfields {
		fmt.Println(string(field))
		payload = append(payload, string(field))
	}
	return payload
}

func runCmd(commands []string) (stdout []byte, err error) {
	var cmd = "git"
	stdout, err = exec.Command(cmd, commands...).Output()
	if err != nil {
		return stdout, err
	}
	return stdout, err
}

func errCmd(command []string, err error) error {
	fmtErr := fmt.Errorf("could not execute command %v %v\n", command, err)
	return errors.Unwrap(fmtErr)
}
