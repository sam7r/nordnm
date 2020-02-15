package nordvpn

import (
	"fmt"
	"github.com/sam7r/nordnm/logger"
	"io/ioutil"
	"net/http"
)

// TODO: Get from config, env $TMPDIR or /tmp
var tempDir = "/tmp"

// GetNordVpnConfigFile fetches ovnpn config file from nordvpn
func GetNordVpnConfigFile(hostID string, tech string) (ovpnConfig []byte, err error) {
	resourceURI := fmt.Sprintf("https://downloads.nordcdn.com/configs/files/ovpn_%s/servers/%s.nordvpn.com.%s.ovpn", tech, hostID, tech)
	logger.Stdout.Infof("Fetching OVPN file from %s", resourceURI)
	resp, err := http.Get(resourceURI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fetching ovpn file returned a status of %s", resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}
