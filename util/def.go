/*
 *@author  chengkenli
 *@project StarRocksQueris
 *@package util
 *@file    def
 *@date    2024/8/7 14:44
 */

package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Config      *viper.Viper
	Loggrs      *logrus.Logger
	Connect     *gorm.DB
	P           ArvgParms
	ConnectLink []map[string]interface{}
)

type ArvgParms struct {
	Help     bool
	ConfPath string
	IP       string
	Mtime    string
	Key      string
	Rtar     bool
	App      string
	Te       string
	Action   string
	Filename string
}

type ConnSSH struct {
	User           string
	Host           string
	Port           int
	PrivateKeyFile string
	Command        string
}
