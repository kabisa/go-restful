package restful

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"strings"
)

type acceptHeader struct {
	value string
}

func (a acceptHeader) acceptsAny() bool {
	return len(a.value) == 0 || strings.Contains(a.value, "*/*")
}

func (a acceptHeader) acceptsMIME(mime string) bool {
	return strings.Contains(a.value, mime)
}
