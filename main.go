/*
 *@author  chengkenli
 *@project pjstack
 *@package pjstack
 *@file    main
 *@date    2024/11/9 17:05
 */
package main

import (
	"pjstack/frame"
	"pjstack/util"
)

func main() {
	printStarRocks()
	util.Init()
	ConfigDB()
	/*GO*/
	switch util.P.Action {
	case "log":
		frame.ExecuteLog()
	case "pstack", "jstack":
		frame.ExecuteTrace()
	case "iostat":
		frame.ExecuteIoStat()
	default:

		return
	}
}
