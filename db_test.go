package tools

import (
	"testing"
)

func TestQueryRelated(t *testing.T) {
	var value = []string{
		"XA011-0001-0027",
		"XA011-0001-0027",
		"XA011-0001-0026",
	}
	res, err := QueryRelated("root:kevin@tcp(192.168.35.190:3307)/tdms?charset=utf8mb4", "tool_device",
		"code_number", "id", value)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
