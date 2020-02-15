package nmcli

import (
	"bufio"
	"github.com/sam7r/nordnm/logger"
	"io"
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
		if lineOut := scanner.Text(); lineOut != "" {
			stdout = append(stdout, lineOut)
		}
	}
	return stdout
}
