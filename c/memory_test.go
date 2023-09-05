package c_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
	"weizicoding.com/go-mem-manager/c"
)

func TestMemoryC(t *testing.T) {
	data := c.Malloc(4)
	fmt.Printf(" data %+v, %T\n", data, data)
	myData := (*uint32)(data)
	*myData = 4
	fmt.Printf(" data %+v, %T\n", *myData, *myData)

	c.Free(data)
}

func IsLittleEndian() bool {
	var n int32 = 0x01020304

	// 下面是为了将 int32 类型的指针转换成 byte 类型的指针
	u := unsafe.Pointer(&n)
	pb := (*byte)(u)

	// 取得 pb 位置对应的值
	b := *pb

	// 由于 b 是 byte 类型，最多保存 8 位，那么只能取得开始的 8 位
	// 小端: 04 (03 02 01)
	// 大端: 01 (02 03 04)
	return (b == 0x04)
}

func IntToBytes(n uint32) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	if IsLittleEndian() {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	binary.Write(bytesBuffer, order, x)

	return bytesBuffer.Bytes()
}

func TestMemoryC2(t *testing.T) {
	data := c.Malloc(4)
	fmt.Printf(" data %+v, %T\n", data, data)
	myData := (*uint32)(data)
	*myData = 5
	fmt.Printf(" data %+v, %T\n", *myData, *myData)

	var a uint32 = 100
	c.Memcpy(data, IntToBytes(a), 4)
	fmt.Printf(" data %+v, %T\n", *myData, *myData)

	c.Free(data)
}
