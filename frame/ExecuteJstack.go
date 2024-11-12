/*
 *@author  chengkenli
 *@project pjstack
 *@package main
 *@file    ExecuteJstack
 *@date    2024/11/9 18:35
 */
package frame

import (
	"fmt"
	"pjstack/conn"
	"pjstack/tools"
	"pjstack/util"
	"strings"
	"time"
)

// 打印jstack返回
func execJstack(host string) {
	fmt.Println("JStack模式。")
	stime := time.Now()
	command := "ps aux | grep 'fe-' | grep -v grep | awk '{print$2}'"
	ot := conn.ConnectSSH(
		&util.ConnSSH{
			User:           util.Config.GetString("common.User"),
			Host:           host,
			Port:           util.Config.GetInt("common.Port"),
			PrivateKeyFile: util.Config.GetString("common.Private"),
			Command:        command,
		})
	pid := strings.NewReplacer("\n", "").Replace(ot.(string))

	command = util.Config.GetString("common.Jstack") + " " + pid
	ot = conn.ConnectSSH(
		&util.ConnSSH{
			User:           util.Config.GetString("common.User"),
			Host:           host,
			Port:           util.Config.GetInt("common.Port"),
			PrivateKeyFile: util.Config.GetString("common.Private"),
			Command:        command,
		})
	logfile := fmt.Sprintf("%s/fe.%s.%s.jstack", util.Config.GetString("common.LogPath"), host, time.Now().Format("20060102150405"))
	tools.WriteFile(logfile, ot.(string))
	fmt.Println(fmt.Sprintf("%-13s jstack %s edtime:%s", host, logfile, time.Now().Sub(stime).String()))
}
