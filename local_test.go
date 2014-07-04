package web

import (
	"testing"
)

func TestMessage(t *testing.T) {
	local := newLocal()

	// test default return
	if local.value("en_US", "login") != "login" {
		t.Error("default return error")
	}

	// test default return
	local.setMessage("en_US", "login", "sign in")
	if local.value("en_US", "login") != "sign in" {
		t.Error("default return error")
	}
	if local.value("zh_CN", "login") != "sign in" {
		t.Error("default return error")
	}

	local.setMessage("zh_CN", "login", "登录")
	local.setDefault("zh_CN")
	if local.value("zh_CN", "login") != "登录" {
		t.Error("default return error")
	}
	local.setDefault("en_US")

	// test no sub default return
	if local.value("en_US", "login") != "sign in" {
		t.Error("no sub default return error")
	}
	if local.value("zh_CN", "login") != "登录" {
		t.Error("no sub default return error")
	}
	if local.value("zh_TW", "login") != "sign in" {
		t.Error("no sub default return error")
	}

	// test sub default return
	local.setMessage("zh", "default", "zh_CN")
	if local.value("zh_TW", "login") != "登录" {
		t.Error("sub default error")
	}

	// test fmt return
	local.setMessage("zh_CN", "welcome", "你好%s%d")
	if local.value("zh_CN", "welcome", "张三", 1) != "你好张三1" {
		t.Error("fmt return error")
	}
}
