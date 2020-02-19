package nmcli

import (
	"github.com/sam7r/nordnm/utils"
	"os/exec"
)

func checkHasNmcli() (err error) {
	cmd := exec.Command("nmcli", "-v")
	if err = cmd.Run(); err != nil {
		utils.Logger.Info("nmcli command not found")
	}
	return err
}
