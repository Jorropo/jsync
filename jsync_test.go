// SPDX-License-Identifier: Apache-2.0 OR MIT
package jsync_test

import (
	"testing"

	. "github.com/Jorropo/jsync"
)

func TestFWaitGroup(t *testing.T) {
	var called bool
	f := func() {
		called = true
	}
	var fwg *FWaitGroup
	shouldNot := func() {
		t.Helper()
		fwg.Done()
		if called {
			t.Error("called is true even tho it should not")
		}
	}
	should := func() {
		t.Helper()
		fwg.Done()
		if !called {
			t.Error("called is false even tho it should not")
		}
	}

	fwg = NewFWaitGroup(f, 2)
	shouldNot()
	should()

	called = false
	fwg.Init(f, 2)
	shouldNot()
	fwg.Add()
	shouldNot()
	should()
}
