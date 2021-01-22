# Keypress

Simple go program to type the reccuring error handling code for Go:
```go
if err != nil {
    log.Println(err.Error())
}
```

You must build the binary with `go build simulate.go` the put in somehere in your path (I use `$home/bin`)

## Add keybinding (i3)
add this to your i3 config file :
```tar
bindsym $mod+b exec $home/bin/keypress
```

## Setup uinput permissions
The file /dev/uinput must belong to the `input` group
```bash
sudo chown root:input /dev/uinput
```

### Load the uinput module
```bash
sudo modprobe uinput
```

### Add udev rules
```bash
sudo groupadd input
sudo usermod -a -G input $(whoami)
sudo udevadm control --reload-rules
echo "SUBSYSTEM==\"misc\", KERNEL==\"uinput\", GROUP=\"input\", MODE=\"0660\"" | sudo tee /etc/udev/rules.d/uinput.rules
echo input | sudo tee /etc/modules-load.d/uinput.conf
```
