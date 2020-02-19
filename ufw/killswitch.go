package ufw

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sam7r/nordnm/utils"
)

func getUfwRules(isDelete bool) map[string][]string {
	rules := make(map[string][]string)

	// allow tun0 interface explicitly when default is being denied
	rules["enable_tun0_io"] = []string{
		"allow out on tun0 from any to any",
		"allow in on tun0 from any to any",
	}

	// allow ports needed to establish new VPN connections
	rules["enable_vpn_io"] = []string{
		"allow out 443/tcp",
		"allow out 1194/udp",
	}

	if isDelete {
		// loop through and prefix with delete
		for k, ruleGroup := range rules {
			removeRules := []string{}
			for _, rule := range ruleGroup {
				removeRules = append(removeRules, fmt.Sprintf("delete %s", rule))
			}
			rules[k] = removeRules
		}

		// remove outgoing lock
		rules["allow_default_io"] = []string{
			"default deny incoming",
			"default allow outgoing",
		}
	} else {
		// deny all io on default
		rules["deny_default_io"] = []string{
			"default deny incoming",
			"default deny outgoing",
		}

	}

	return rules
}

// EnableKillswitch enables UFW killswitch
func EnableKillswitch(dryRun bool) (out []string, err error) {
	if err := checkHasUFW(); err != nil {
		return out, err
	}

	rules := getUfwRules(false)

	groupOut := make(map[string][]string)
	for ruleKey, ruleGroup := range rules {
		utils.Logger.Infof("running ufw %s", ruleKey)
		for _, rule := range ruleGroup {
			if dryRun {
				rule = "--dry-run " + rule
			}
			utils.Logger.Infof("ufw %s", strings.Fields(rule))
			cmd := exec.Command("ufw", strings.Fields(rule)...)

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

			groupOut[ruleKey] = append(groupOut[ruleKey], strings.Join(utils.GetStdoutText(stdout), "\n"))

			if errOut := utils.GetStdoutText(stderr); errOut != nil {
				return nil, fmt.Errorf(strings.Join(errOut, " "))
			}

			cmd.Wait()
		}
	}

	for ruleKey, ruleOut := range groupOut {
		out = append(out, ruleKey, strings.Join(ruleOut, "\n"))
	}

	return out, err
}

// DisableKillswitch disables UFW killswitch
func DisableKillswitch(dryRun bool) (out []string, err error) {
	if err := checkHasUFW(); err != nil {
		return out, err
	}

	rules := getUfwRules(true)
	groupOut := make(map[string][]string)
	for ruleKey, ruleGroup := range rules {
		utils.Logger.Infof("running ufw %s", ruleKey)
		for _, rule := range ruleGroup {

			if dryRun {
				rule = "--dry-run " + rule
			}
			utils.Logger.Infof("ufw %s", strings.Fields(rule))
			cmd := exec.Command("ufw", strings.Fields(rule)...)

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

			groupOut[ruleKey] = append(groupOut[ruleKey], strings.Join(utils.GetStdoutText(stdout), "\n"))

			if errOut := utils.GetStdoutText(stderr); errOut != nil {
				return nil, fmt.Errorf(strings.Join(errOut, " "))
			}

			cmd.Wait()
		}
	}

	for ruleKey, ruleOut := range groupOut {
		out = append(out, ruleKey, strings.Join(ruleOut, "\n"))
	}

	return out, err
}

func checkHasUFW() (err error) {
	cmd := exec.Command("ufw", "version")
	if err = cmd.Run(); err != nil {
		utils.Logger.Info("ufw command not found")
	}
	return err
}
