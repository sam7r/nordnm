# [WIP] nordnm
## Nord VPN, Network Manager CLI with UFW Killswitch

**Features below are a work in progress and will change**
**In it's current state there is very limited functionality so please stay tuned for the first release**

This utility package was created for linux users who use NetworkManager and like to manage NordVPN connections manually via OVPN files.
The main purpose of this package is to speed up the process of listing nord vpn servers and managing new/existing ovpn connections.

In its current state this package has some strict dependencies but this will change over time.
1) Network Manager (nmcli)
2) Network Manager OpenVPN
3) UFW (Uncomplicated Firewall) -- optional: if you use a different firewall or have your own killswitch toggle

The package is written in Go and CLI created with Cobra.

## Features
The `nordnm` comes with a set of commands to help with:
- Pulling the files needed from NordVPN required to establish a connection
- Identifying the best server to choose with any set requirements
- Loading the chosen nordvpn connection into Network Manager
- Toggle a UFW killswitch

## Examples
Change the configuration, this can also be managed by creating a .nordnm config file
nordnm config set <option> <value>

### VPN querying
```
nordnm vpn [--OPTIONS]
```
 - list available servers, can be filtered with optional args<br>
 - show groups, technologies and country codes available


### Connection management
```sh
nordnm conn [--OPTIONS]
```
 -list (Default): list all created connection profiles <br>
 -create: create a new connection <br>
 -remove: remove an existing connection <br>


### Killswitch, this will create/remove a set of rules within UFW firewall.
```
nordnm killswitch (enable | disable)
```


## Config
### Recomended settings in Network Manager
### Keychain vs Password
