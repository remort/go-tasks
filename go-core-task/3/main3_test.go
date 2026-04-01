package main

import (
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	if m == nil {
		t.Error("NewStringIntMap returned nil")
	}
	if m.data == nil {
		t.Error("NewStringIntMap created nil map")
	}
	if len(m.data) != 0 {
		t.Error("NewStringIntMap should create empty map")
	}
}

func TestAdd(t *testing.T) {
	m := NewStringIntMap()

	m.Add("key1", 10)
	if val, ok := m.data["key1"]; !ok {
		t.Error("Add failed to add key1")
	} else if val != 10 {
		t.Errorf("Add added wrong value: got %d, want 10", val)
	}

	m.Add("key2", 20)
	if val, ok := m.data["key2"]; !ok {
		t.Error("Add failed to add key2")
	} else if val != 20 {
		t.Errorf("Add added wrong value: got %d, want 20", val)
	}

	// Test overwriting existing key
	m.Add("key1", 100)
	if val, ok := m.data["key1"]; !ok {
		t.Error("Add failed to overwrite key1")
	} else if val != 100 {
		t.Errorf("Add overwrote with wrong value: got %d, want 100", val)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()

	// Remove non-existent key should not panic
	m.Remove("nonexistent")

	m.Add("key1", 10)
	m.Add("key2", 20)

	m.Remove("key1")
	if _, ok := m.data["key1"]; ok {
		t.Error("Remove failed to delete key1")
	}
	if _, ok := m.data["key2"]; !ok {
		t.Error("Remove deleted wrong key")
	}

	m.Remove("key2")
	if len(m.data) != 0 {
		t.Error("Remove should empty the map")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()

	// Test copy of empty map
	copy1 := m.Copy()
	if copy1 == nil {
		t.Error("Copy returned nil for empty map")
	}
	if len(copy1.data) != 0 {
		t.Error("Copy should return empty map")
	}

	m.Add("key1", 10)
	m.Add("key2", 20)
	m.Add("key3", 30)

	copy2 := m.Copy()

	// Verify copy has all elements
	if len(copy2.data) != 3 {
		t.Errorf("Copy size wrong: got %d, want 3", len(copy2.data))
	}

	for k, v := range m.data {
		if copyVal, ok := copy2.data[k]; !ok {
			t.Errorf("Copy missing key: %s", k)
		} else if copyVal != v {
			t.Errorf("Copy wrong value for %s: got %d, want %d", k, copyVal, v)
		}
	}

	// Verify copy is independent
	m.Add("key4", 40)
	if _, ok := copy2.data["key4"]; ok {
		t.Error("Copy should not reflect changes to original")
	}

	copy2.Add("key5", 50)
	if _, ok := m.data["key5"]; ok {
		t.Error("Original should not reflect changes to copy")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()

	if m.Exists("key1") {
		t.Error("Exists returned true for non-existent key")
	}

	m.Add("key1", 10)

	if !m.Exists("key1") {
		t.Error("Exists returned false for existing key")
	}

	m.Remove("key1")

	if m.Exists("key1") {
		t.Error("Exists returned true after removal")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()

	val, ok := m.Get("key1")
	if ok {
		t.Error("Get returned ok=true for non-existent key")
	}
	if val != 0 {
		t.Errorf("Get returned non-zero value for non-existent key: got %d, want 0", val)
	}

	m.Add("key1", 42)

	val, ok = m.Get("key1")
	if !ok {
		t.Error("Get returned ok=false for existing key")
	}
	if val != 42 {
		t.Errorf("Get returned wrong value: got %d, want 42", val)
	}

	m.Add("key2", 100)

	val, ok = m.Get("key2")
	if !ok {
		t.Error("Get returned ok=false for existing key")
	}
	if val != 100 {
		t.Errorf("Get returned wrong value: got %d, want 100", val)
	}

	m.Remove("key1")

	val, ok = m.Get("key1")
	if ok {
		t.Error("Get returned ok=true after removal")
	}
	if val != 0 {
		t.Errorf("Get returned non-zero value after removal: got %d, want 0", val)
	}
}

func TestIntegration(t *testing.T) {
	m := NewStringIntMap()

	// Add multiple values
	for i := 0; i < 100; i++ {
		key := string(rune('a'+i%26)) + string(rune('0'+i/26))
		m.Add(key, i)
	}

	// Verify all exist
	for i := 0; i < 100; i++ {
		key := string(rune('a'+i%26)) + string(rune('0'+i/26))
		if !m.Exists(key) {
			t.Errorf("Key %s should exist", key)
		}
	}

	// Get all values
	for i := 0; i < 100; i++ {
		key := string(rune('a'+i%26)) + string(rune('0'+i/26))
		val, ok := m.Get(key)
		if !ok {
			t.Errorf("Get failed for key %s", key)
		}
		if val != i {
			t.Errorf("Get wrong value for %s: got %d, want %d", key, val, i)
		}
	}

	// Copy and verify
	copied := m.Copy()
	if len(copied.data) != 100 {
		t.Errorf("Copy size wrong: got %d, want 100", len(copied.data))
	}

	// Remove half
	for i := 0; i < 50; i++ {
		key := string(rune('a'+i%26)) + string(rune('0'+i/26))
		m.Remove(key)
	}

	// Verify removal
	for i := 0; i < 100; i++ {
		key := string(rune('a'+i%26)) + string(rune('0'+i/26))
		exists := m.Exists(key)
		if i < 50 && exists {
			t.Errorf("Key %s should have been removed", key)
		}
		if i >= 50 && !exists {
			t.Errorf("Key %s should still exist", key)
		}
	}

	// Verify copy unaffected
	if len(copied.data) != 100 {
		t.Errorf("Copy changed after original modifications: got %d, want 100", len(copied.data))
	}
}
