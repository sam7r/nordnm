package nmcli

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sam7r/nordnm/utils"
)

// ImportOvpnConnection imports connection from openvpn file
func ImportOvpnConnection(filepath string) ([]string, error) {
	if err := checkHasNmcli(); err != nil {
		return nil, err
	}

	execCmd := []string{"connection", "import", "type", "openvpn", "file", filepath}
	utils.Logger.Infof("running nmcli %v", execCmd)
	cmd := exec.Command("nmcli", execCmd...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		utils.Logger.Info("cmd out pipe produced err")
		panic(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		utils.Logger.Info("cmd out err produced")
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		utils.Logger.Info("cmd failed to start")
		panic(err)
	}

	data := utils.GetStdoutText(stdout)
	if errOut := utils.GetStdoutText(stderr); len(errOut) > 0 {
		utils.Logger.Infof("cmd resp -> data %v", data)
		utils.Logger.Infof("cmd resp -> error %v", errOut)
		return nil, fmt.Errorf(strings.Join(errOut, " "))
	}

	cmd.Wait()
	return data, err

}
