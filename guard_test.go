package xguard

import "testing"

func TestNewGuard(t *testing.T) {
	g := NewGuard()
	if g == nil {
		t.Error("NewGuard() should not return nil")
	}
}

func TestGuard_Start(t *testing.T) {
	g := NewGuard()
	g.Start()
	status := g.GetStatus()
	if status != "running" {
		t.Errorf("GetStatus() should return 'running', but got %s", status)
	} else {
		t.Log("GetStatus() return 'running'")
	}
}

func TestGuard_Stop(t *testing.T) {
	g := NewGuard()
	g.Stop()
	status := g.GetStatus()
	if status != "stopped" {
		t.Errorf("GetStatus() should return 'stopped', but got %s", status)
	} else {
		t.Log("GetStatus() return 'stopped'")
	}
}
