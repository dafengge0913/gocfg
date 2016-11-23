package gocfg

import (
	"fmt"
	"github.com/dafengge0913/gocfg"
	"runtime"
	"testing"
)

func TestIni(t *testing.T) {
	fmt.Println("TestIni")
	cfg, err := gocfg.ParseIni("demo.ini")
	if err != nil {
		t.Error("error: ", err)
		return
	}

	for k, v := range cfg.GetAllData() {
		t.Log(k, "->", v)
	}

	equal(t, cfg.GetString("username"), "root")

	n, err := cfg.GetInt("password")
	if err != nil {
		t.Error("error: ", err)
		return
	}
	equal(t, 123456, n)

	for _, s := range cfg.GetStringList("key") {
		t.Log(s)
	}

	equal(t, true, cfg.GetBool("admin"))
}

func TestJson(t *testing.T) {
	fmt.Println("TestIni")
	cfg, err := gocfg.ParseJson("demo.json")
	if err != nil {
		t.Error("error: ", err)
		return
	}

	for k, v := range cfg.GetAllData() {
		t.Log(k, "->", v)
	}

	equal(t, cfg.GetString("username"), "root")

	n, err := cfg.GetInt("password")
	if err != nil {
		t.Error("error: ", err)
		return
	}
	equal(t, 123456, n)

	t.Log(cfg.GetStringList("key"))

	for _, s := range cfg.GetStringList("key") {
		t.Log(s)
	}
}

func equal(t *testing.T, a, b interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if a != b {
		t.Error(file, ":", line, ":", a, "!=", b)
	}
}
