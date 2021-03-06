// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestReadWriteUvarint(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []int64{
		0x7fffffffffffffff,
		0x78787793485adcdf,
		0x0009043716000001,
		-256 * 256 * 256 * 256 * 256 * 256 * 256 * 128,
		-1,
		1,
		0xacf3514de,
		0x152dac09f,
		0,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteUvarint(w, uint64(v)))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadUvarint(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, int64(value), v)
	}
}

func TestReadWriteVarint(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []int64{
		0x7fffffffffffffff,
		0x78787793485adcdf,
		0x0009043716000001,
		-256 * 256 * 256 * 256 * 256 * 256 * 256 * 128,
		-1,
		1,
		0xacf3514de,
		0x152dac09f,
		0,
		126,
		127,
		128,
		256,
		-256 * 256,
		256 * 256 * 256,
		-256 * 256 * 256 * 256,
		256 * 256 * 256 * 256 * 256,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteVarint(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadVarint(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteUint64(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []uint64{
		0xffffff7387645173,
		0x78787793485adcdf,
		0x0009043716000001,
		0xacf3514de,
		0x152dac09f,
		0,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteUint64(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadUint64(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteUint32(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []uint32{
		0x07645173,
		0x485adcdf,
		0x90437160,
		0xcf3514de,
		0x52dac09f,
		0,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteUint32(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadUint32(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteUint16(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []uint16{
		0x5173,
		0xadcd,
		0x3716,
		0x14de,
		0xc09f,
		0xffff,
		0,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteUint16(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadUint16(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteUint8(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []uint8{
		0xff,
		0x7f,
		128,
		127,
		0,
		1,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteUint8(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadUint8(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteByte(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []uint8{
		0xff,
		0x7f,
		128,
		127,
		0,
		1,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteByte(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadByte(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteInt64(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []int64{
		-12909,
		0x93485adcdf,
		1290183847,
		-1,
		0,
		-256 * 256 * 256 * 256 * 256 * 256 * 256 * 128,
		-1193470394772,
		0x152dac09f,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteInt64(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadInt64(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteInt32(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []int32{
		0x76451730,
		-938493900,
		-15565,
		-1,
		0,
		-256 * 256 * 256 * 128,
		0x0ff3514d,
		0x752dac0f,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteInt32(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadInt32(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteInt16(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []int16{
		0x5173,
		0x7dcd,
		-256 * 128,
		256*127 + 255,
		-1,
		0x14de,
		0x609f,
	}
	for _, v := range tests {
		ensure.Nil(t, WriteInt16(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadInt16(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteInt8(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	for i := -128; i <= 127; i++ {
		ensure.Nil(t, WriteInt8(w, int8(i)))
	}
	w.Flush()

	var r = bytes.NewBuffer(buf.Bytes())
	for i := -128; i <= 127; i++ {
		value, err := ReadInt8(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, int8(i))
	}
}

func TestReadWriteBytes(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = [][]byte{
		[]byte("1234567890"),
		[]byte("abcdefghijklmn"),
		[]byte("_=ld,.a}[[,.;😀🤣"),
		[]byte("中文汉字"),
		[]byte(""),
		{},
	}
	for _, v := range tests {
		ensure.Nil(t, WriteBytes(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadBytesOfLength(r, uint32(len(v)))
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteVarBytes(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = [][]byte{
		[]byte("1234567890"),
		[]byte("abcdefghijklmn"),
		[]byte("_=ld,.a}[[,.;😀🤣"),
		[]byte("中文汉字"),
		[]byte(""),
		{},
	}
	for _, v := range tests {
		ensure.Nil(t, WriteVarBytes(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadVarBytes(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, v)
	}
}

func TestReadWriteHex(t *testing.T) {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)

	var tests = []string{
		"1234567890",
		"fc3a5db8e0",
		"ABCDEF1234567890",
		"abcdefABCDEF1234567890",
		"",
	}
	for _, v := range tests {
		ensure.Nil(t, WriteHex(w, v))
	}
	w.Flush()

	var r = bufio.NewReader(bytes.NewBuffer(buf.Bytes()))
	for _, v := range tests {
		value, err := ReadHex(r)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, value, strings.ToLower(v))
	}
}

func testReadWrite(t *testing.T) {
	var w = &bytes.Buffer{}

	var ts = "abdcedas;fjd🤠😘😙.-"
	var tui uint64 = 1267856
	var ti32 int32 = -1
	var tvi int32 = -256 * 256 * 256 * 128
	var tvi0 int32 = 1
	var tb byte = 1
	var thex = "abcdef12340987"
	var tfs = [8]byte{0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58}
	var tbytes = []byte{
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
		0x01, 0x12, 0x13, 0x14, 0x25, 0x36, 0x47, 0x58, 0x00, 0x89, 0xff, 0xe6, 0xf8, 0x7f, 0x09, 0x80,
	}

	ensure.Nil(t, WriteVarBytes(w, []byte(ts)))
	ensure.Nil(t, WriteUint64(w, tui))
	ensure.Nil(t, WriteInt32(w, ti32))
	for i := 0; i < 4096; i++ {
		ensure.Nil(t, WriteVarint(w, int64(tvi)))
	}
	ensure.Nil(t, WriteByte(w, tb))
	ensure.Nil(t, WriteHex(w, thex))
	ensure.Nil(t, WriteBytes(w, tfs[:]))
	ensure.Nil(t, WriteVarint(w, int64(tvi0)))
	ensure.Nil(t, WriteVarBytes(w, tbytes))

	var r = bytes.NewBuffer(w.Bytes())

	s, _ := ReadVarBytes(r)
	ensure.DeepEqual(t, string(s), ts)

	ui, _ := ReadUint64(r)
	ensure.DeepEqual(t, ui, tui)

	i32, _ := ReadInt32(r)
	ensure.DeepEqual(t, i32, ti32)

	for i := 0; i < 4096; i++ {
		vi, _ := ReadVarint(r)
		ensure.DeepEqual(t, vi, int64(tvi))
	}

	b, _ := ReadByte(r)
	ensure.DeepEqual(t, b, tb)

	hex, _ := ReadHex(r)
	ensure.DeepEqual(t, hex, thex)

	fs, _ := ReadBytesOfLength(r, 8)
	ensure.DeepEqual(t, fs, tfs[:])

	vi0, _ := ReadVarint(r)
	ensure.DeepEqual(t, vi0, int64(tvi0))

	bs, _ := ReadVarBytes(r)
	ensure.DeepEqual(t, bs, tbytes)
}

func TestReadWrite(t *testing.T) {
	testReadWrite(t)
}

func TestConcurrecy(t *testing.T) {
	for i := 0; i < 256; i++ {
		name := fmt.Sprintf("%04d", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testReadWrite(t)
		})
	}
}
