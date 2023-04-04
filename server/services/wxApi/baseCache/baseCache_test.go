// Package baseCache
// @Author:        asus
// @Description:   $
// @File:          baseCache_test.go
// @Data:          2021/12/3117:06
//
package baseCache

import (
	"testing"
)

func TestCache_SetTag(t *testing.T) {
	var cache Cache

	cache.SetTag("tag")

	if cache.Tag() != "tag:" {
		t.Fatal(cache.tag)
	}

}
