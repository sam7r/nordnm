package nmcli

import (
	"fmt"
	"nordnm/logger"
	"os/exec"
	"strings"
)

// ImportOvpnConnection imports connection from openvpn file
func ImportOvpnConnection(filepath string) ([]string, error) {
	if err := checkHasNmcli(); err != nil {
		return nil, err
	}

	execCmd := []string{"connection", "import", "type", "openvpn", "file", filepath}
	logger.Stdout.Infof("running nmcli %v", execCmd)
	cmd := exec.Command("nmcli", execCmd...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Stdout.Info("cmd out pipe produced err")
		panic(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.Stdout.Info("cmd out err produced")
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		logger.Stdout.Info("cmd failed to start")
		panic(err)
	}

	data := getStdoutText(stdout)
	if errOut := getStdoutText(stderr); len(errOut) > 0 {
		logger.Stdout.Infof("cmd resp -> data %v", data)
		logger.Stdout.Infof("cmd resp -> error %v", errOut)
		return nil, fmt.Errorf(strings.Join(errOut, " "))
	}

	cmd.Wait()
	return data, err

}
