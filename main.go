// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// main
package main

import (
	"fmt"
	"io"
	"os"
)

//nolint:gochecknoglobals
var out io.Writer = os.Stdout

func main() {
	_, _ = fmt.Fprintln(out, "Hello, World!")
}
