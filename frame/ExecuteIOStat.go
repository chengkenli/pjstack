/*
 *@author  chengkenli
 *@project pjstack
 *@package main
 *@file    ConnectIOStat
 *@date    2024/9/9 15:41
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

// ExecuteIoStat 打印iostat 返回
func ExecuteIoStat() {
	if !tools.InitCommon() {
		return
	}
	fmt.Println("IOStat模式。")
	stime := time.Now()
	for _, host := range strings.Split(util.P.IP, ",") {
		command := "iostat -mx 1 1"
		ot := conn.ConnectSSH(
			&util.ConnSSH{
				User:           util.Config.GetString("common.User"),
				Host:           host,
				Port:           util.Config.GetInt("common.Port"),
				PrivateKeyFile: util.Config.GetString("common.Private"),
				Command:        command,
			})

		fmt.Println()
		fmt.Println("iostat>")
		fmt.Println(host + ":")
		fmt.Println(ot.(string))

		logfile := fmt.Sprintf("%s/iostat.%s.%s.jstack", util.Config.GetString("common.LogPath"), host, time.Now().Format("20060102150405"))
		tools.WriteFile(logfile, ot.(string))
		fmt.Println(fmt.Sprintf("%-13s iostat %s edtime:%s", host, logfile, time.Now().Sub(stime).String()))
	}
}
