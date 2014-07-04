package web

import "testing"

func TestMuxPath(t *testing.T) {
	// test default
	if muxPath("/a") != "/a" || muxPath("/") != "/" {
		t.Error("test default error")
	}

	// test add /
	if muxPath("/a/") != "/a" {
		t.Error("test add / error")
	}

	// test toggle case
	if muxPath("/A/") != "/a" {
		t.Error("test toggle case error")
	}
}
