/*
   fileName: bstore
   author: diogoxiang@qq.com
   date: 2019/7/9
*/
package bstore

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func testBadgerStore(t testing.TB) (*BadgerStore, string) {
	path, err := ioutil.TempDir("", "raftbadger")
	if err != nil {
		t.Fatalf("err. %s", err)
	}
	os.RemoveAll(path)

	// Successfully creates and returns a store
	store, err := NewBadgerStore(path)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return store, path
}

func TestBadgerStore_Set_Get(t *testing.T) {
	store, path := testBadgerStore(t)
	defer store.Close()
	defer os.Remove(path)

	// Returns error on non-existent key
	if _, err := store.Get([]byte("bad")); err != ErrKeyNotFound {
		t.Fatalf("expected not found error, got: %q", err)
	}

	k, v := []byte("hello"), []byte("world")

	// Try to set a k/v pair
	if err := store.Set(k, v); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Try to read it back
	val, err := store.Get(k)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if !bytes.Equal(val, v) {
		t.Fatalf("bad: %v", val)
	}
}

func TestBadgerStore_SetUint64_GetUint64(t *testing.T) {
	store, path := testBadgerStore(t)
	defer store.Close()
	defer os.Remove(path)

	// Returns error on non-existent key
	if _, err := store.GetUint64([]byte("bad")); err != ErrKeyNotFound {
		t.Fatalf("expected not found error, got: %q", err)
	}

	k, v := []byte("abc"), uint64(123)

	// Attempt to set the k/v pair
	if err := store.SetUint64(k, v); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Read back the value
	val, err := store.GetUint64(k)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if val != v {
		t.Fatalf("bad: %v", val)
	}
}

func BenchmarkBadgerStore_Set(b *testing.B) {
	store, path := testBadgerStore(b)
	defer store.Close()
	defer os.Remove(path)

	// raftbench.Set(b, store)

	set(b,*store)


}

func BenchmarkBadgerStore_Get(b *testing.B) {
	store, path := testBadgerStore(b)
	defer store.Close()
	defer os.Remove(path)

	// raftbench.Get(b, store)
	get(b,*store)
}

func set(b *testing.B, store BadgerStore)  {

	for n := 0; n < b.N; n++ {
		if err := store.Set([]byte{byte(n)}, []byte("val")); err != nil {
			b.Fatalf("err: %s", err)
		}
	}

}

func get(b *testing.B, store BadgerStore)  {
	for i := 1; i < 10; i++ {
		if err := store.Set([]byte{byte(i)}, []byte("val")); err != nil {
			b.Fatalf("err: %s", err)
		}
	}
	b.ResetTimer()

	// Run Get a number of times
	for n := 0; n < b.N; n++ {
		if _, err := store.Get([]byte{0x05}); err != nil {
			b.Fatalf("err: %s", err)
		}
	}
}