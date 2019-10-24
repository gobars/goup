package goup

import (
	"fmt"
	"goup/util"
	"strings"
)

var kafkaInventoryTpl = `
[kafka]{{range $index, $element := .inventory_hosts}}
{{$element.IP}} ansible_ssh_user={{$element.Username}} ansible_ssh_pass={{$element.Password}} kafka_id={{$element.ID}} kafka_ip="{{$element.IP}}"{{end}}
`

func InstallKafkaHa(settings util.Settings) (err error) {
	inventory := "kafka_ha_inventory"
	playbook := "kafka_ha.yml"

	fmt.Printf("> genegene %s file\n", inventory)
	if err = tplToFile([]byte(kafkaInventoryTpl), inventory, kafkaInventoryParams(settings)); err != nil {
		return
	}

	fmt.Printf("> genegene %s file", playbook)
	zk := append([]string(nil), settings.KafkaHa.Brokers...) // Notice the ... splat
	for i := range zk {
		zk[i] = strings.TrimSpace(zk[i]) + ":2181"
	}

	var config = map[string]interface{}{
		"kafka_zk_quorum": strings.Join(zk, ",") + "/kafka",
		"zookeeper_hosts": settings.KafkaHa.Brokers,
	}
	if err = tplToFile(util.StatiqFS.Files["/kafka_ha.tpl.yml"].Data, playbook, config); err != nil {
		return
	}

	fmt.Println("> begin exec ansible-playbook")
	exec(inventory, playbook)
	fmt.Println("> end exec ansible-playbook")
	return
}

func kafkaInventoryParams(settings util.Settings) interface{} {
	hosts := make([]InstallHosts, 0)
	for _, s := range settings.Servers {
		if i := util.Contains(settings.KafkaHa.Brokers, s.IP); i != -1 {
			hosts = append(hosts, InstallHosts{
				Servers: s,
				ID:      i + 1,
			})
		}
	}
	return map[string]interface{}{
		"inventory_hosts": hosts,
	}
	return nil
}
