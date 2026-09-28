//go:build !windows
// +build !windows

package daemon

import (
	"fmt"
	"gossh/internal/log"
	"os"
	"syscall"
)

func Daemonize(logger log.Logger, args []string, path string) {
	// If this is already a child process, then return
	if os.Getenv("DAEMON_ENV") == "1" {
		return
	}
	procname, err := os.Executable()
	if err != nil {
		logger.Error("Error finding the executable path")
		return
	}
	// Add DAEMON_ENV=1 for identifying the child process
	env := append(
		os.Environ(),
		"DAEMON_ENV=1",
	)

	attr := &os.ProcAttr{
		Dir: "/",
		Env: env,
		Sys: &syscall.SysProcAttr{
			Setsid: true, // Setsid doesn't exist within the SysProcAttr struct on Windows.
		},
	}

	proc, err := os.StartProcess(procname, args, attr)
	if err != nil {
		logger.Error("Unable to start child process: %s", err.Error())
		return
	}

	// Create a pid file to log the process id (to be used later to identify and kill child process)
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		proc.Kill()
		proc.Wait()
		return
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%d", proc.Pid)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		proc.Kill()
		proc.Wait()
		return
	}
}
