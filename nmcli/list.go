package nmcli

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sam7r/nordnm/logger"
	"github.com/sam7r/nordnm/utils"
)

// NetworkConnection connection information provided by nmcli
type NetworkConnection struct {
	Name string
	UUID string
	Type string
}

// NetworkConnections array of NetworkConnection
type NetworkConnections []NetworkConnection

// FilterByType excludes given type string
func (ncs *NetworkConnections) FilterByType(typeName string) {
	filtered := NetworkConnections{}
	for _, conn := range *ncs {
		if conn.Type == typeName {
			filtered = append(filtered, conn)
		}
	}

	logger.Stdout.Info(filtered)
	*ncs = filtered
}

// ListConnections shows connections via nmcli command
func ListConnections(onlyActive bool) (out NetworkConnections, err error) {
	if err := checkHasNmcli(); err != nil {
		return nil, err
	}

	execCmd := []string{"-c", "no", "conn", "show"}
	if onlyActive == true {
		execCmd = append(execCmd, "--active")
	}

	logger.Stdout.Infof("running nmcli %v", execCmd)
	cmd := exec.Command("nmcli", execCmd...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	data := utils.GetStdoutText(stdout)
	if errOut := utils.GetStdoutText(stderr); errOut != nil {
		err = fmt.Errorf(strings.Join(errOut, " "))
	}

	cmd.Wait()

	for _, conn := range data {
		row := strings.Fields(conn)
		logger.Stdout.Infof("%v", row)
		out = append(out, NetworkConnection{row[0], row[1], row[2]})
	}
	return out, err
}
