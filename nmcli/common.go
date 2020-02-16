package nmcli

import (
	"os/exec"

	"github.com/sam7r/nordnm/logger"
)

func checkHasNmcli() (err error) {
	cmd := exec.Command("nmcli", "-v")
	if err = cmd.Run(); err != nil {
		logger.Stdout.Info("nmcli command not found")
	}
	return err
}
