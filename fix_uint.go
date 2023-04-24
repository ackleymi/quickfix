// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"errors"
	"strconv"
)

const (
	// ASCII - char.
	asciiMinus = 45

	// ASCII numbers 0-9.
	ascii0 = 48
	ascii9 = 57
)


// atoi is similar to the function in strconv, but is tuned for ints appearing in FIX field types.
func atoi(d []byte) (int, error) {
	if d[0] == asciiMinus {
		n, err := parseUInt(d[1:])
		return (-1) * int(n), err
	}
	n, err := parseUInt(d)
	return int(n), err
}

// atoui is similar to the function in strconv, but is tuned for uints appearing in FIX field types.
func atoui(d []byte) (uint, error) {
	if d[0] == asciiMinus {
		return 0, errors.New("invalid format")

	}

	return parseUInt(d)
}

// parseUInt is similar to the function in strconv, but is tuned for uints appearing in FIX field types.
func parseUInt(d []byte) (n uint, err error) {
	if len(d) == 0 {
		err = errors.New("empty bytes")
		return
	}

	for _, dec := range d {
		if dec < ascii0 || dec > ascii9 {
			err = errors.New("invalid format")
			return
		}

		n = n*10 + (uint(dec) - ascii0)
	}

	return
}

// FIXUInt is a FIX INT/SEQNUM Value, implements FieldValue.
type FIXUInt uint

// UInt converts the FIXUInt value to uint.
func (f FIXUInt) UInt() uint { return uint(f) }

func (f *FIXUInt) Read(bytes []byte) error {
	i, err := atoui(bytes)
	if err != nil {
		return err
	}
	*f = FIXUInt(i)
	return nil
}

func (f FIXUInt) Write() []byte {
	return strconv.AppendUint(nil, uint64(f), 10)
}
