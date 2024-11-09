/*
 *@author  chengkenli
 *@project pjstack
 *@package main
 *@file    HandlePjstack
 *@date    2024/9/9 15:45
 */

package frame

import (
	"fmt"
	"pjstack/conn"
	"pjstack/tools"
	"pjstack/util"
	"strings"
	"sync"
	"time"
)

// ExecuteTrace 堆栈根据主体
func ExecuteTrace() {
	db, err := conn.StarRocks(util.P.App)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	FeHosts := tools.FrontendsIP(db)
	BeHosts := tools.BackendsIP(db)

	stime := time.Now()
	ch := make(chan struct{}, 2)
	var wg sync.WaitGroup
	for _, host := range strings.Split(util.P.IP, ",") {
		wg.Add(1)
		go func(host string) {
			defer func() {
				<-ch
				wg.Done()
			}()
			ch <- struct{}{}

			var sign string
			Ok := tools.StrInSlice(host, FeHosts)
			if Ok {
				sign = "fe"
			}
			Year := tools.StrInSlice(host, BeHosts)
			if Year {
				sign = "be"
			}
			if sign != "fe" && sign != "be" {
				fmt.Println(fmt.Sprintf("%s is an unknown host...", host))
				return
			}
			switch sign {
			case "fe":
				execJstack(host)
			case "be":
				execPstack(host)
			}
		}(host)
	}
	wg.Wait()

	fmt.Println(fmt.Sprintf("all done. total edtime:[%s]", time.Now().Sub(stime).String()))
}
