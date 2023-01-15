package main

import "TestApp-LoginServer-Go/demo/util"

func main() {
	//wsStart.WebSocket()
	//start.LoginController()
	data := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	temp := make([]byte, 4)
	util.Arraycopy(data, 0, temp, 0, 4)

	for _, b := range temp {
		println(b)
	}
}
