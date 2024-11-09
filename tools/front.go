/*
 *@author  chengkenli
 *@project pjstack
 *@package tools
 *@file    front
 *@date    2024/11/9 18:22
 */
package tools

import (
	"fmt"
	"gorm.io/gorm"
)

func FrontendsIP(db *gorm.DB) []string {
	var f []map[string]interface{}
	r := db.Raw("show frontends").Scan(&f)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
		return nil
	}

	var frontIP []string
	for _, s := range f {
		frontIP = append(frontIP, s["IP"].(string))
	}
	return frontIP
}

func BackendsIP(db *gorm.DB) []string {
	var f []map[string]interface{}
	r := db.Raw("show backends").Scan(&f)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
		return nil
	}

	var backendIP []string
	for _, s := range f {
		backendIP = append(backendIP, s["IP"].(string))
	}
	return backendIP
}
