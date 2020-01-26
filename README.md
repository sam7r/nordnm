# nordnmcli
## Nord VPN, Network Manager CLI with UFW Killswitch

This utility package was created for linux users who wish to continue using Network Mananger whilst using Nord VPN's command line interface.
Once setup it should make connection management and remain very stable.

In its current state this package has some strict dependencies but this will change over time.
1) Network Manager (nmcli)
2) Network Manager OpenVPN
3) UFW (Uncomplicated Firewall)

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
 -list (Default): lists previously downloaded vpn connection profiles <br>
 -empty-cache: remove all files downloaded not currently in use from previous 'conn create' connections <br>


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
