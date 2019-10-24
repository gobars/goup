package goup

import (
	"fmt"
	"goup/util"
)

var redisInventoryTpl = `
{{range $index, $element := .inventory_hosts}}
[redis{{$element.ID}}]
{{$element.IP}}

[redis{{$element.ID}}:vars]
ansible_ssh_user={{$element.Username}}
ansible_ssh_pass={{$element.Password}}
{{end}}
`

func InstallRedisHa(settings util.Settings) (err error) {
	inventory := "redis_ha_inventory"
	playbook := "redis_ha.yml"

	fmt.Printf("> genegene %s file\n", inventory)
	if err = tplToFile([]byte(redisInventoryTpl), inventory, redisInventoryParams(settings)); err != nil {
		return
	}

	fmt.Printf("> genegene %s file", playbook)
	if err = tplToFile(util.StatiqFS.Files["/redis_ha.tpl.yml"].Data, playbook, settings.RedisHa); err != nil {
		return
	}

	fmt.Println("> begin exec ansible-playbook")
	exec(inventory, playbook)
	fmt.Println("> end exec ansible-playbook")
	return
}

type InstallHosts struct {
	util.Servers
	ID int
}

func redisInventoryParams(settings util.Settings) interface{} {
	hosts := make([]InstallHosts, 2)
	for _, s := range settings.Servers {
		if s.IP == settings.RedisHa.MasterAddr {
			hosts[0] = InstallHosts{
				Servers: s,
				ID:      1,
			}
		} else if s.IP == settings.RedisHa.SlaveAddr {
			hosts[1] = InstallHosts{
				Servers: s,
				ID:      2,
			}
		}
	}
	return map[string]interface{}{
		"inventory_hosts": hosts,
	}
}
