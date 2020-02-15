# nordnm
## Nord VPN, Network Manager CLI with UFW Killswitch


Util written in Go for linux users who like to manage thier NordVPN network connections via the standard NetworkManager package.
It aims to take away the pain of this manual process away by providing easy interface to query NordVPN, find out the best server, download the configuration and import into NetworkManger.

1) [Network Manager](https://www.archlinux.org/packages/extra/x86_64/networkmanager/)
2) [Network Manager OpenVPN](https://www.archlinux.org/packages/extra/x86_64/networkmanager-openvpn/)
3) [UFW (Uncomplicated Firewall)](https://www.archlinux.org/packages/community/any/ufw/) -- optional: if you use a different firewall or have your own killswitch toggle

## Features
The `nordnm` comes with a set of commands to help with:
- Identify the best server within a specified country, vpn group and connection type
- Downloads the chosen ovpn file and creates a new connection in NetworkManager
- Toggle a UFW killswitch

# Installation

This package requires Golang >= 1.13<br />
You will need to have the following packages already installed and configured
```sh
pacman -S networkmanager networkmanager-openvpn ufw go
```
Install the package
```sh
go get -u github.com/sam7r/nordnm
```


# Examples
## NordVPN server finder
```sh
nordnm vpn list [--OPTIONS]
```

Example using all available flags
```sh
# list top servers standard UDP servers in the UK
nordnm vpn list /
    --country 227 /
    --group legacy_standard /
    --tech openvpn_udp /
    --limit 3
```

Instead of entering these in manually, you can setup a config file in `$HOME/.nordnmrc`<br/>
Any given flags will override configuration settings

```json
{
    "preferences": {
        "countryCode": 227,
        "groupIdentifier": "legacy_double_vpn",
        "technologyIdentifier": "openvpn_tcp"
    }
}

```

To find out which country code code, group identifier or technology identifier to enter you can use the `nordnm vpn show` command

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
List vpn connectionscrea
```sh
# lists only VPN NetworkManager connections
nordnm conn list

# lists all active NetworkManager connections
nordnm conn list --all --active
```

Create a new NetworkManager vpn connection
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

Instead of entering these credentials manually, you can create and add them to `$HOME/.nordnmrc`
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
Credentials are not required as part of the connection creation as these can be added manually in the NetworkManager GUI or potentially a setup where they are stored in your operating systems key manager.

**When adding connections there are two options that are on by default to help stop DNS leaks**
```sh
# prevents Wired/Wifi DNS leaking in to /etc/resolv.conf
ipv4.dns-priority -1 
ipv4.ignore-auto-dns true
```

## Help
If at any point when using this you are unsure of the sub commands or flags available use the `--help` or `-h` flag
```sh
nordnm vpn --help
```

# Contributing
Contributing to this project is welcome, so far I have only built this to serve my own particular use case but am happy for this to extend beyond NordVPN and the `nmcli` interface.

To debug or track events in more detail you can run in verbose output mode
```sh
nordnm vpn list --verbose
```
