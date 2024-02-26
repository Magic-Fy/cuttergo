package cuttergo

import "testing"

func TestCutter_Init(t *testing.T) {
	c := &Cutter{}
	err := c.Init("./nonexist.txt")
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCutter_Cut(t *testing.T) {
	content := "北京清华大学研究生院"
	c := &Cutter{}
	res, err := c.Cut(content)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(res) != 6 && res[0] != "北京" && res[1] != "清华" && res[5] != "院" {
		t.Errorf("Error: res length = %v, not 5", len(res))
	}
}
