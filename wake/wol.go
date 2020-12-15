package wake

import (
	"errors"
	"github.com/mdlayher/wol"
	"github.com/tcdw/telegram-home-master/config"
	"net"
)

func ByMachine(addr string, tar string, password []byte) error {
	target, err := net.ParseMAC(tar)
	c, err := wol.NewClient()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.WakePassword(addr, target, password)
}

func ByName(conf config.Config, target string) error {
	var item *config.Computer
	for _, e := range conf.Computers {
		if e.Name == target {
			item = &e
			break
		}
	}
	if item == nil {
		return errors.New("no such machine")
	}
	var ip = "255.255.255.255:9"
	if item.IP != nil {
		ip = *item.IP
	}
	var password = ""
	if item.Password != nil {
		password = *item.Password
	}
	err := ByMachine(ip, item.Mac, []byte(password))
	return err
}
