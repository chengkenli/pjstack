/*
 *@author  chengkenli
 *@project StarRocksQueris
 *@package util
 *@file    conf
 *@date    2024/8/7 14:42
 */

package util

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Printf("\nUsage: %s [-s adhoc] [-h]\n\nOptions:\n(log,pstack,jstack,iostat)日志打包工具.\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	fmt.Println()
}

func init() {
	execDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	defaultConf := fmt.Sprintf("%s/.%s.yaml", execDir, filepath.Base(os.Args[0]))

	flag.StringVar(&P.ConfPath, "c", defaultConf, "conf file")
	c := color.New()
	flag.BoolVar(&P.Help, "h", false, "show help information")
	flag.StringVar(&P.IP, "ip", "", c.Add(color.FgHiGreen).Sprint("ip host?"))
	flag.StringVar(&P.Mtime, "t", "-mmin -5", "查找当前目录及其子目录下，<time>前到现在修改过的文件，"+c.Add(color.FgHiYellow).Sprint("10分钟之内：<-mmin -10>，2天之内：<-mtime -2>，12个小时之内：<-mmin -720>"))
	flag.StringVar(&P.Key, "k", "", "搜索日志中的"+c.Add(color.FgHiYellow).Sprint("<关键字>"))
	flag.BoolVar(&P.Rtar, "r", false, c.Add(color.FgHiRed).Sprint("压缩+回传本地"))
	flag.StringVar(&P.App, "s", "sr-adhoc", "集群名称")
	flag.StringVar(&P.Te, "e", "fe", "<fe/be/broker>")
	flag.StringVar(&P.Action, "a", "log", fmt.Sprintf("<%s,%s,%s,%s>",
		c.Add(color.FgHiGreen).Sprint("log"),
		c.Add(color.FgHiYellow).Sprint("pstack"),
		c.Add(color.FgHiYellow).Sprint("jstack"),
		c.Add(color.FgHiYellow).Sprint("iostat"),
	))

	flag.Parse()
	flag.Usage = usage

	if P.Help {
		flag.Usage()
		os.Exit(-1)
	}

	paths, name := filepath.Split(P.ConfPath)
	Config = viper.New()
	Config.SetConfigFile(fmt.Sprintf("%s%s", paths, name))
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
	Logrus()
}

func Init() {

}
