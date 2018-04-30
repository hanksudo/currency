package backup

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Start - backup
func Start() {
	fmt.Println("Start backup")
	if os.Getenv("DROPBOX_ACCESS_TOKEN") == "" {
		fmt.Println("Stop backup, DROPBOX_ACCESS_TOKEN missing.")
		os.Exit(2)
	}
	cmd := exec.Command("python", "scripts/backup_to_dropbox.py")
	outPipe, _ := cmd.StdoutPipe()
	errPipe, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(outPipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	scanner = bufio.NewScanner(errPipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	cmd.Wait()
}
