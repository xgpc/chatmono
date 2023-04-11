package openAI

import (
	"fmt"
	"testing"
)

func TestSession(t *testing.T) {
	var info SessionData
	info.Request.Model = "models"
	info.Request.Temperature = 1
	info.SessionType = 1
	info.TemplateID = 2

	data := info.String()

	fmt.Println(data)
	info2 := SessionDecode(data)

	fmt.Println(info2)

}
