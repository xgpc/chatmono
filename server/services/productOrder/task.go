// Package productOrder
// @Author:        asus
// @Description:   $
// @File:          task
// @Data:          2022/2/2116:02
//
package productOrder

import (
	"time"
)

func task() {
	go func() {
		for {
			select {
			case <-time.After(6 * time.Minute):
				taskCheckUP()
			}
		}
	}()

}
