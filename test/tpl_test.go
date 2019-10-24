package main

import (
	"goup"
	"log"
	"os"
	"testing"
	"text/template"
)

func Test_KafkaHaFile(t *testing.T) {
	tpl, err := template.ParseFiles("goup/res/kafka_ha.tpl.yml")
	if err != nil {
		log.Print(err)
		return
	}
	var config = map[string]interface{}{
		"kafka_zk_quorum": "192.168.33.110:2181,192.168.33.111:2181,192.168.33.112:2181/kafka",
		"zookeeper_hosts": []string{"192.168.33.110", "192.168.33.111", "192.168.33.112"},
	}
	tpl.Execute(os.Stdout, config)
}

func Test_RedisHaFile(t *testing.T) {
	tpl, err := template.ParseFiles("goup/res/redis_ha.tpl.yml")
	if err != nil {
		log.Print(err)
		return
	}
	config := goup.RedisHa{
		MasterAddr: "192.168.33.110",
		SlaveAddr:  "192.168.33.111",
		Port:       6379,
		Password:   "123456",
		VipAddr:    "192.168.33.116",
		IfaceName:  "enp0s8",
	}
	tpl.Execute(os.Stdout, config)
}

func Test_MySQLHaFile(t *testing.T) {
	tpl, err := template.ParseFiles("goup/res/mysql_ha.tpl.yml")
	if err != nil {
		log.Print(err)
		return
	}
	config := goup.MySQLHa{
		Master1Addr:  "192.168.33.110",
		Master2Addr:  "192.168.33.111",
		Port:         3306,
		Password:     "1qazzaq1",
		VipAddr:      "192.168.33.115",
		IfaceName:    "enp0s8",
		ReplUsr:      "rep_3306",
		ReplPassword: "123456",
	}
	tpl.Execute(os.Stdout, config)
}
