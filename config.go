/*
 *@author  chengkenli
 *@project StarRocksQueris
 *@package main
 *@file    config
 *@date    2024/10/21 23:13
 */

package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"pjstack/conn"
	"pjstack/util"
)

// ConfigDB 初始化连接数据库是否成功
func ConfigDB() {
	c := color.New()
	var err error
	util.Connect, err = conn.ConnectMySQL()
	if err != nil {
		util.Loggrs.Error(err)
		return
	}
	fmt.Println(c.Add(color.FgGreen).Sprint("配置数据库连接成功!"))
	ok, err := initDB()
	if !ok {
		fmt.Println(err.Error())
		return
	}
	ok, err = initCommon()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// 初始化StarRocks登录配置
func initDB() (bool, error) {
	condb := util.Config.GetString("configdb.Schema.Connect")
	if len(condb) == 0 {
		return false, errors.New("Schema.Connect is nil")
	}
	r := util.Connect.Raw(fmt.Sprintf("select * from %s where status >= 1", condb)).Scan(&util.ConnectLink)
	if r.Error != nil {
		fmt.Println(r.Error)
		return false, errors.New(r.Error.Error())
	}
	return true, nil
}

func initCommon() (bool, error) {
	pstack := util.Config.GetString("common.Pstack")
	if len(pstack) == 0 {
		return false, errors.New("common.Pstack is nil")
	}
	jstack := util.Config.GetString("common.Jstack")
	if len(jstack) == 0 {
		return false, errors.New("common.Jstack is nil")
	}
	user := util.Config.GetString("common.User")
	if len(user) == 0 {
		return false, errors.New("common.User is nil")
	}
	private := util.Config.GetString("common.Private")
	if len(private) == 0 {
		return false, errors.New("common.Private is nil")
	}
	logpath := util.Config.GetString("common.LogPath")
	if len(logpath) == 0 {
		return false, errors.New("common.LogPath is nil")
	}
	return true, nil
}
