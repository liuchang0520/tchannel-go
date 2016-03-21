// Copyright (c) 2015 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt32(t *testing.T) {
	atom := &Int32{42}

	require.Equal(t, int32(42), atom.Load(), "Load didn't work.")
	require.Equal(t, int32(44), atom.Add(2), "Add didn't work.")
	require.Equal(t, int32(45), atom.Inc(), "Inc didn't work.")
	require.Equal(t, int32(44), atom.Dec(), "Dec didn't work.")

	require.True(t, atom.CAS(44, 0), "CAS didn't report a swap.")
	require.Equal(t, int32(0), atom.Load(), "CAS didn't set the correct value.")

	require.Equal(t, int32(0), atom.Swap(1), "Swap didn't return the old value.")
	require.Equal(t, int32(1), atom.Load(), "Swap didn't set the correct value.")

	atom.Store(42)
	require.Equal(t, int32(42), atom.Load(), "Store didn't set the correct value.")
}

func TestInt64(t *testing.T) {
	atom := &Int64{42}

	require.Equal(t, int64(42), atom.Load(), "Load didn't work.")
	require.Equal(t, int64(44), atom.Add(2), "Add didn't work.")
	require.Equal(t, int64(45), atom.Inc(), "Inc didn't work.")
	require.Equal(t, int64(44), atom.Dec(), "Dec didn't work.")

	require.True(t, atom.CAS(44, 0), "CAS didn't report a swap.")
	require.Equal(t, int64(0), atom.Load(), "CAS didn't set the correct value.")

	require.Equal(t, int64(0), atom.Swap(1), "Swap didn't return the old value.")
	require.Equal(t, int64(1), atom.Load(), "Swap didn't set the correct value.")

	atom.Store(42)
	require.Equal(t, int64(42), atom.Load(), "Store didn't set the correct value.")
}

func TestUint32(t *testing.T) {
	atom := &Uint32{42}

	require.Equal(t, uint32(42), atom.Load(), "Load didn't work.")
	require.Equal(t, uint32(44), atom.Add(2), "Add didn't work.")
	require.Equal(t, uint32(45), atom.Inc(), "Inc didn't work.")
	require.Equal(t, uint32(44), atom.Dec(), "Dec didn't work.")

	require.True(t, atom.CAS(44, 0), "CAS didn't report a swap.")
	require.Equal(t, uint32(0), atom.Load(), "CAS didn't set the correct value.")

	require.Equal(t, uint32(0), atom.Swap(1), "Swap didn't return the old value.")
	require.Equal(t, uint32(1), atom.Load(), "Swap didn't set the correct value.")

	atom.Store(42)
	require.Equal(t, uint32(42), atom.Load(), "Store didn't set the correct value.")
}

func TestUint64(t *testing.T) {
	atom := &Uint64{42}

	require.Equal(t, uint64(42), atom.Load(), "Load didn't work.")
	require.Equal(t, uint64(44), atom.Add(2), "Add didn't work.")
	require.Equal(t, uint64(45), atom.Inc(), "Inc didn't work.")
	require.Equal(t, uint64(44), atom.Dec(), "Dec didn't work.")

	require.True(t, atom.CAS(44, 0), "CAS didn't report a swap.")
	require.Equal(t, uint64(0), atom.Load(), "CAS didn't set the correct value.")

	require.Equal(t, uint64(0), atom.Swap(1), "Swap didn't return the old value.")
	require.Equal(t, uint64(1), atom.Load(), "Swap didn't set the correct value.")

	atom.Store(42)
	require.Equal(t, uint64(42), atom.Load(), "Store didn't set the correct value.")
}
