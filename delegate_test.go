package delegate

import (
	"testing"
)

func TestDelegate(t *testing.T) {
	var delegate Delegate
	counter := 0
	delegate.Add(func(i interface{}) {
		counter += 1
	})
	hdl := delegate.Add(func(i interface{}) {
		counter += 2
	})
	delegate.Apply(nil)
	if counter != 3 {
		t.FailNow()
	}
	hdl.Cancel()
	delegate.Apply(nil)
	if counter != 4 {
		t.FailNow()
	}
}
