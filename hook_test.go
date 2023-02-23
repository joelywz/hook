package hook_test

import (
	"github.com/joelywz/hook"
	"sync"
	"testing"
	"time"
)

func TestHook_Trigger(t *testing.T) {
	h := hook.New[int]()

	var res = 0
	resMutex := sync.Mutex{}

	h1 := h.AddHandler(func(data int) {
		resMutex.Lock()
		res = data
		resMutex.Unlock()
	})

	h.Trigger(1)
	time.Sleep(1 * time.Millisecond)

	resMutex.Lock()
	if res != 1 {
		t.Fatalf("hook did not trigger, expected 1, got %d", res)
	}
	resMutex.Unlock()

	h.RemoveHandler(h1)

	h.Trigger(2)
	time.Sleep(1 * time.Millisecond)

	if res != 1 {
		t.Fatalf("handler was not removed, expected 1, got %d", res)
	}
}

func TestHook_MultiHandler(t *testing.T) {
	h := hook.New[int]()

	var res = 0
	var res2 = 0

	h1 := h.AddHandler(func(data int) {
		res = data
	})

	h2 := h.AddHandler(func(data int) {
		res2 = data
	})

	h.Trigger(1)
	time.Sleep(1 * time.Millisecond)

	if res != 1 {
		t.Fatalf("hook did not trigger handler 1, expected 1, got %d", res)
	}

	if res2 != 1 {
		t.Fatalf("hook did not trigger handler 2, expected 1, got %d", res2)
	}

	h.RemoveHandler(h1)

	h.Trigger(2)
	time.Sleep(1 * time.Millisecond)

	if res != 1 {
		t.Fatalf("handler 1 was not removed, expected 1, got %d", res)
	}

	if res2 != 2 {
		t.Fatalf("handler 2 was removed, expected 2, got %d", res2)
	}

	h.RemoveHandler(h2)
}
