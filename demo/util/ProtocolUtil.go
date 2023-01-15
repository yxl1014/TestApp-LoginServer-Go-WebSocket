package util

/*
自定义协议解析类
*/
import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
)

// DecodeProtocol 自定义协议解码
func DecodeProtocol(data []byte) []byte {
	temp := make([]byte, 4)

	Arraycopy(data, 0, temp, 0, 4)
	l := binary.LittleEndian.Uint32(temp)

	Arraycopy(data, 4, temp, 0, 4)
	if strings.Compare(string(temp), MagicNumber) != 0 {
		return nil
	}

	end := make([]byte, 5)
	Arraycopy(data, 4+4+4+int(l), end, 0, 5)
	if strings.Compare(string(end), END) != 0 {
		return nil
	}

	result := make([]byte, int(l))
	Arraycopy(data, 4+4+4, result, 0, int(l))
	return result
}

// GetType 获取此段数据的类型
func GetType(data []byte) int {
	temp := make([]byte, 4)
	Arraycopy(data, 4+4, temp, 0, 4)
	return int(binary.LittleEndian.Uint32(temp))
}

// EncodeProtocol 自定义协议编码
func EncodeProtocol(data []byte, len int /*, TestProto.Types type*/) []byte {
	temp := make([]byte, len+17)
	buf := bytes.NewBuffer(temp)
	binary.Write(buf, binary.LittleEndian, len)
	magic, _ := strconv.Atoi(MagicNumber)
	binary.Write(buf, binary.LittleEndian, magic)
	//TODO 这里需要填入proto对象的类型
	binary.Write(buf, binary.LittleEndian, 12)
	buf.Write(data)
	buf.WriteString(END)
	return buf.Bytes()
}

// Arraycopy 数组拷贝
func Arraycopy(src []byte, srcPos int, dest []byte, destPos int, length int) {
	for i := 0; i < length; i++ {
		dest[destPos+i] = src[srcPos+i]
	}
}
