# ipl

## ✒️ Description
> Simple tool for displaying information about network interfaces

`ipl` displays the interface name, IPv4 address, and mask in one line:

```
❯ ipl 
[lo] 127.0.0.1/255.0.0.0
[wlo1] 192.168.18.2/255.255.255.0
[virbr0] 192.168.122.1/255.255.255.0
```
## ⚒️ To build

```sh
go build -o ./ipl
```

Or you can use just:

```sh
just
```
