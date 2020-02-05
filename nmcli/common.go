package nmcli

import (
	"bufio"
	"io"
	"nordnm/logger"
	"os/exec"
)

func checkHasNmcli() (err error) {
	cmd := exec.Command("nmcli", "-v")
	if err = cmd.Run(); err != nil {
		logger.Stdout.Info("nmcli command not found")
	}
	return err
}

func getStdoutText(r io.Reader) (stdout []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		stdout = append(stdout, scanner.Text())
	}
	return stdout
}
