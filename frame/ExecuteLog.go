/*
 *@author  chengkenli
 *@project pjstack
 *@package main
 *@file    HandleLog
 *@date    2024/10/10 13:52
 */

package frame

import (
	"fmt"
	"github.com/chengkenli/dos/shell"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"pjstack/conn"
	"pjstack/tools"
	"pjstack/util"
	"strings"
	"sync"
)

// ExecuteLog 检查关键字日志返回
func ExecuteLog() {
	if !tools.InitCommon() {
		return
	}
	fmt.Println("日志搜索模式。")
	c := color.New()
	var host []string
	if util.P.IP == "" {
		db, err := conn.StarRocks(util.P.App)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch util.P.Te {
		case "fe":
			host = tools.FrontendsIP(db)
		case "be":
			host = tools.BackendsIP(db)
		case "broker":
			host = tools.BackendsIP(db)
		default:
			c.Add(color.FgHiRed).Sprint("service name is err.")
			fmt.Println(c.Add(color.FgHiRed).Sprint("service name is err."))
			return
		}
	} else {
		host = strings.Split(util.P.IP, ",")
	}

	var path string
	// 根据集群标记找到fe和be的日志目录，以/结尾
	for _, m := range util.ConnectLink {
		if m["app"].(string) == util.P.App {
			switch util.P.Te {
			case "fe":
				path = m["fe_log_path"].(string)
			case "be":
				path = m["be_log_path"].(string)
			}
		}
	}
	if len(path) == 0 {
		fmt.Println("集群名称或日志路径异常！")
		return
	}
	// end
	fmt.Println(fmt.Sprintf("获取到%s节点：%s", util.P.Te, strings.Join(host, ",")))

	for i, ip := range host {
		c := color.New()
		fmt.Println(i, c.Add(color.FgHiYellow).Sprint(ip))

		t1, ok := juet(fmt.Sprintf("find %s -type f %s", path, util.P.Mtime), ip)
		if !ok {
			fmt.Println(c.Add(color.FgHiRed).Sprint("file is nil."))
			continue
		}
		fmt.Println(c.Add(color.FgHiBlue).Sprint(t1))

		files := strings.Split(t1, "\n")
		if util.P.Rtar {
			var f []string
			for _, file := range files {
				if strings.Contains(file, "WARNING") || strings.Contains(file, "warn") || strings.Contains(file, "jni.INFO") || strings.Contains(file, ".tbz") {
					continue
				}
				f = append(f, file)
			}

			sname := fmt.Sprintf("/tmp/%s.%s.tbz", util.P.Te, ip)
			command := fmt.Sprintf("tar -zcf %s %s", sname, strings.Join(f, " "))
			fmt.Println("tar:\n", c.Add(color.FgHiYellow).Sprint(strings.Join(f, "\n")), "\n", c.Add(color.FgHiGreen).Sprint(command))
			juet(command, ip)

			var localPath string
			if len(util.Config.GetString("common.LogPath")) != 0 {
				localPath = util.Config.GetString("common.LogPath")
			} else {
				localPath = "./"
			}
			command2 := fmt.Sprintf("scp starrocks@%s:%s %s", ip, sname, localPath)
			err := shell.RunShell(command2, fmt.Sprintf("%s.log", filepath.Base(os.Args[0])), 600)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			command3 := fmt.Sprintf("rm -f %s", sname)
			juet(command3, ip)
		}
		if util.P.Key == "" {
			continue
		}

		var wg sync.WaitGroup
		ch := make(chan struct{}, 2)
		for _, file := range files {
			wg.Add(1)
			go func(file string) {
				defer func() {
					<-ch
					wg.Done()
				}()

				ch <- struct{}{}
				if len(file) == 0 {
					return
				}
				if strings.Contains(file, "WARNING") {
					return
				}

				comm := fmt.Sprintf("grep -arn '%s' %s", util.P.Key, file)
				fmt.Println("R:", len(file))
				t2, ok := juet(comm, ip)
				if !ok {
					return
				}
				c := color.New()
				fmt.Println("file: ", c.Add(color.FgHiGreen).Sprint(file))
				fmt.Println(c.Add(color.FgHiYellow).Sprint(t2))

			}(file)
		}
		wg.Wait()
	}
}

func juet(command, host string) (string, bool) {
	ot := conn.ConnectSSH(
		&util.ConnSSH{
			User:           util.Config.GetString("common.User"),
			Host:           host,
			Port:           util.Config.GetInt("common.Port"),
			PrivateKeyFile: util.Config.GetString("common.Private"),
			Command:        command,
		})
	if len(ot.(string)) != 0 {
		return ot.(string), true
	}
	return ot.(string), false
}
