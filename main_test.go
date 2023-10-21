// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	buf := new(bytes.Buffer)
	out = buf

	main()

	got := buf.String()
	want := "Hello, World!\n"

	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}
