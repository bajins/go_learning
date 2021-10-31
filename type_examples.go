package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"
)

func main() {
	bytes_, err := os.ReadFile("E:\test.txt")
	if err != nil {
		return
	}

	bytesBuffer := bytes.NewBuffer(bytes_)
	var x int32
	err = binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(int(x))
	i := binary.LittleEndian.Uint32(bytes_)
	fmt.Println(i)
	i = binary.BigEndian.Uint32(bytes_)
	fmt.Println(i)
}

// Int2Byte 把int的每个字节取出来放入byte数组中，存储采用Littledian
func Int2Byte(data int) (ret []byte) {
	var len_ uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len_)
	var tmp int = 0xff
	var index uint = 0
	for index = 0; index < uint(len_); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

// Byte2Int 把byte　Slice 中的每个字节取出来，　按Littledian端拼成一个int
func Byte2Int(data []byte) int {
	var ret int = 0
	var _len int = len(data)
	var i uint = 0
	for i = 0; i < uint(_len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}

// BytesToIntU 字节数(大端)组转成int(无符号的)
func BytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// BytesToIntS 字节数(大端)组转成int(有符号)
func BytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// IntToBytes 整形转换成字节
func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		if err != nil {
			return nil, err
		}
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		if err != nil {
			return nil, err
		}
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		if err != nil {
			return nil, err
		}
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}
