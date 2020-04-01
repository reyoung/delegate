package delegate

import (
	"sync"
)

type callbackFunc func(interface{})
type Handler struct {
	parent     *Delegate
	callbackID uint32
	fn         callbackFunc
}

func (h Handler) Cancel() {
	h.parent.remove(h.callbackID)
}

type Delegate struct {
	callbacks map[uint32]callbackFunc
	locker    sync.RWMutex
	idCounter uint32
}

func (d *Delegate) Add(callback callbackFunc) Handler {
	result := Handler{
		parent: d,
		fn:     callback,
	}
	d.locker.Lock()
	defer d.locker.Unlock()
	if d.callbacks == nil {
		d.callbacks = make(map[uint32]callbackFunc)
	}
	result.callbackID = d.idCounter
	d.callbacks[d.idCounter] = callback
	d.idCounter += 1
	return result
}

func (d *Delegate) Apply(i interface{}) {
	d.locker.RLock()
	defer d.locker.RUnlock()

	for _, fn := range d.callbacks {
		fn(i)
	}
}

func (d *Delegate) remove(id uint32) {
	d.locker.Lock()
	defer d.locker.Unlock()
	delete(d.callbacks, id)
}
