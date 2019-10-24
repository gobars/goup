package goup

import (
	"bytes"
	"fmt"
	"github.com/gobars/cmd"
	"goup/util"
	"io/ioutil"
	"text/template"
	"time"
)

var mysqlInventory = `
{{range $index, $element := .inventory_hosts}}
[mysql{{$element.ID}}]
{{$element.IP}}

[mysql{{$element.ID}}:vars]
ansible_ssh_user={{$element.Username}}
ansible_ssh_pass={{$element.Password}}
{{end}}
`

type MySQLInventory struct {
	Master1Addr string
	Master1User string
	Master1Pass string
	Master2Addr string
}

func InstallMysqlHa(settings util.Settings) (err error) {
	inventory := "mysql_ha_inventory"
	playbook := "mysql_ha.yml"

	fmt.Println("> genegene mysql_ha_inventory file")
	if err = tplToFile([]byte(mysqlInventory), inventory, mysqlInventoryParams(settings)); err != nil {
		return
	}

	fmt.Println("> genegene mysql_ha.yml file")
	if err = tplToFile(util.StatiqFS.Files["/mysql_ha.tpl.yml"].Data, playbook, settings.MySQLHa); err != nil {
		return
	}

	fmt.Println("> begin exec ansible-playbook")
	exec(inventory, playbook)
	fmt.Println("> end exec ansible-playbook")
	return
}

func mysqlInventoryParams(settings util.Settings) map[string]interface{} {
	hosts := make([]interface{}, 2)
	for _, s := range settings.Servers {
		if s.IP == settings.MySQLHa.Master1Addr {
			hosts[0] = InstallHosts{
				Servers: s,
				ID:      1,
			}
		} else if s.IP == settings.MySQLHa.Master2Addr {
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

func tplToFile(tplStr []byte, filename string, data interface{}) (err error) {
	var tpl *template.Template
	if tpl, err = template.New(filename).Parse(string(tplStr[:])); err != nil {
		return
	}
	var tplResult bytes.Buffer
	if err = tpl.Execute(&tplResult, data); err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, tplResult.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func exec(inventory, playbook string) {
	shell := fmt.Sprintf(`ansible-playbook -i %s %s`, inventory, playbook)
	fmt.Printf("$ %s\n", shell)
	p := cmd.NewCmdOptions(cmd.Options{Buffered: true, Streaming: true}, "/bin/bash", "-c", shell)
	statusChan := p.Start()
	timeout := time.After(10 * time.Minute)
	for {
		select {
		case curLine := <-p.Stderr:
			fmt.Println(curLine)
		case curLine := <-p.Stdout:
			fmt.Println(curLine)
		case status := <-statusChan:
			fmt.Printf("exec Complete:%+v,exit:%+v \n", status.Complete, status.Exit)
			return
		case <-timeout:
			fmt.Println("timeout reading streaming output")
			return
		}
	}
}
