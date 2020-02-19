package nmcli

import (
	"fmt"
	"github.com/sam7r/nordnm/utils"
	"os/exec"
	"strings"
)

// OvpnConnectionDefaults settings to apply to an existing NetworkManager connection
type OvpnConnectionDefaults struct {
	DNS          string
	AuthSettings Auth
	IgnoreIPV6   bool
}

// Auth holds authentication details for setting up a connection
type Auth struct {
	Mode string
	User string
	Pass string
}

// ModifyConnection modifies the specified connection using settings provided
func ModifyConnection(connectionID string, settings OvpnConnectionDefaults) ([]string, error) {
	if err := checkHasNmcli(); err != nil {
		return nil, err
	}

	execCmd := []string{"connection", "modify", connectionID,
		"ipv4.dns-priority -1", // prevents DNS leak in /etc/resolv.conf
		fmt.Sprintf("ipv4.dns %s", settings.DNS),
		"ipv4.ignore-auto-dns true",
	}
	if settings.AuthSettings.Mode == "non_encrypted" {
		passwordCmd := []string{
			"+vpn.data password-flags=0",
			fmt.Sprintf("+vpn.data username=%s", settings.AuthSettings.User),
			fmt.Sprintf("vpn.secrets password=%s", settings.AuthSettings.Pass),
		}
		execCmd = append(execCmd, passwordCmd...)
	}
	if settings.IgnoreIPV6 {
		execCmd = append(execCmd, "ipv6.method ignore")
	}

	var cmdFields []string
	for _, cmd := range execCmd {
		cmdFields = append(cmdFields, strings.Fields(cmd)...)
	}
	utils.Logger.Infof("running nmcli %v", cmdFields)
	cmd := exec.Command("nmcli", cmdFields...)

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
	if errOut := utils.GetStdoutText(stderr); errOut != nil {
		return nil, fmt.Errorf(strings.Join(errOut, " "))
	}

	cmd.Wait()
	return data, err

}
