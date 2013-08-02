package mlib

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	m0 := &MusicEntry {
		"1", "My heart will go on", "Celion dion", "pop",
		"http://localhost", "MP3"}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManger.Add() failed.")
	}
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name
		|| m.Source != m0.Source ||m.Type != m0.Type {
		t.Error("MusicManager.FInd() failed, found item mismatch")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.remove() failed", err)
	}
}
