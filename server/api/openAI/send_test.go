package openAI

import (
	"fmt"
	"github.com/xgpc/util"
	"testing"
)

func TestDemo(t *testing.T) {
	body := `{"message":[{"role":"user","content":"1"},{"role":"user","content":"12"},{"role":"user","content":"12"},{"role":"user","content":"12"},{"role":"user","content":"1"},{"role":"user","content":"1"}]}`

	var param OpenAISendReq

	err := util.JsonDecode([]byte(body), &param)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(param)
}
