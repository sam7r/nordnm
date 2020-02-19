# nordnm - Nord VPN, Network Manager CLI with UFW Killswitch


Util written in Go for linux users who like to manage their Nord VPN connections via the standard NetworkManager package, this is not a background daemon and will only operate whilst executing the given command.

The main object here is to relieve the manual process required to find the best Nord VPN server, download the ovpn file and import the configuration into NetworkManger.

For this to work it requires youhaving the following packages installed:
1) [Network Manager](https://www.archlinux.org/packages/extra/x86_64/networkmanager/)
2) [Network Manager OpenVPN](https://www.archlinux.org/packages/extra/x86_64/networkmanager-openvpn/)
3) [UFW (Uncomplicated Firewall)](https://www.archlinux.org/packages/community/any/ufw/) -- optional if you use a different firewall or have your own kill switch toggle

## Features
The `nordnm` comes with a set of commands to help with:
- Identify the best server within a specified country filtering by vpn group (standard, double vpn etc...) and connection type (udp, tcp)
- Downloads the chosen server ovpn file and creates a new connection in NetworkManager
- Toggle a UFW kill switch

# Installation

This package requires **Go** >= **1.13**<br />
As mentioned above you will need to have the following packages already installed and configured for this to work
```sh
pacman -S networkmanager networkmanager-openvpn go
```
If this was your first time attempting to install **Go** you will most likely need to check your `go env` settings are correct before continuing

Install the package
```sh
go get -u github.com/sam7r/nordnm
```
This will download and install in  `$GOPATH/src`, now you can run the binary and follow on with the examples below

# Examples
## Nord VPN server finder
```sh
nordnm vpn list [--OPTIONS]
```

### Example using all available flags
```sh
# list top 3 standard UDP servers in the UK
nordnm vpn list \
    --country  227 \
    --group legacy_standard \
    --technology openvpn_udp \
    --limit 3
```
Short hands for the above can be used, run `nordnm vpn list -h` for more details

Instead of entering these in manually, you can setup a config file in `$HOME/.nordnmrc`<br/>
These will become the default, however any given flags at runtime will override these settings

```json
{
    "preferences": {
        "countryCode": 227,
        "groupIdentifier": "legacy_double_vpn",
        "technologyIdentifier": "openvpn_tcp"
    }
}

```

To find any country code, group identifier or technology identifier available within Nord VPN, you can use the `nordnm vpn show` command

```sh
# get the list of NordVPN technologies
nordnm vpn show tech

# get the list of NordVPN country codes
nordnm vpn show countries

# get the list of NordVPN groups
nordnm vpn show groups
```


## NetworkManager connections
```sh
nordnm conn [--OPTIONS]
```
### List vpn connections
```sh
# lists only VPN NetworkManager connections
nordnm conn list

# lists all active NetworkManager connections
nordnm conn list --all --active
```

### Create a new connection
```sh
nordnm conn create [--OPTIONS]
```

Example using all available flags
```sh
nordnm conn create /
    --host uk1515 # required
    --tech UDP # required
    --dns "1.1.1.1,1.0.0.1"
    --username user@email.com
    --password mypass123
    --ignoreIPV6 true
```

Instead of entering these credentials manually, you can add them to the `$HOME/.nordnmrc` settings file 
```json
{
    "connection": {
        "dns": "103.86.99.100,103.86.96.100",
        "password": "NORD_PASS",
        "username": "NORD_USER",
        "ignoreIPV6": true
    }
}
```
Credentials are not required as part of the connection creation stage as these can be added in manually within the NetworkManager GUI

**When adding connections there are two settings that are included by default to help stop DNS leaks**
```sh
# prevents other DNS settings leaking in to /etc/resolv.conf
ipv4.dns-priority -1 
ipv4.ignore-auto-dns true
```

## Killswitch
The following command will add the necessary UFW rules for your vpn connection to be the only outgoing source

### Example enabling killswitch rules, must be ran in `sudo` as UFW requires it
```sh
sudo nordnm killswitch enable [--OPTIONS]
```

If you want to see what rules would be affected without making any actual changes you can run with the `--dry-run` flag

### Example removing killswitch rules
```sh
sudo nordnm killswitch disable [--OPTIONS]
```

Both of the above commands are essentially shortcut for applying or removing the following rules
```sh
# deny all traffic
default deny incoming
default deny outgoing

# allow traffic on tun interface
allow out on tun0 from any to any
allow in on tun0 from any to any

# allow out on ports needed to establish connection to VPN
allow out 443/tcp
allow out 1194/udp
```
There is no firewall refresh so running enable/disable will not affect any of your existing rules

## Help
If at any point when using this you are unsure of the sub commands or flags available use the `--help` or `-h` flag
```sh
nordnm vpn --help
```

# Contributing
Contributing to this project is welcome, so far I have only built this to serve my own particular use cases

To debug or track events in more detail you can run in verbose output mode
```sh
nordnm vpn list --verbose
```
