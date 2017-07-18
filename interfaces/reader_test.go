// Copyright 2014 Jens Rantil. All rights reserved.  Use of this source code is
// governed by a BSD-style license that can be found in the LICENSE file.

package interfaces

import (
	"bytes"
	oldcsv "encoding/csv"
	"testing"

	thiscsv "github.com/eltorocorp/go-csv"
)

func TestReaderInterface(t *testing.T) {
	t.Parallel()

	var iface Reader
	iface = thiscsv.NewReader(new(bytes.Buffer))
	iface = thiscsv.NewDialectReader(new(bytes.Buffer), thiscsv.Dialect{})
	iface = oldcsv.NewReader(new(bytes.Buffer))

	// To get rid of compile-time warning that this variable is not used.
	iface.Read()
}
