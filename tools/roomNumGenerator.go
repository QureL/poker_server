package tools

import "sync"

var lock sync.Mutex
var room_num int = 0

func RoomNumGenerator() int {
	lock.Lock()
	defer lock.Unlock()
	room_num++
	var ret int = room_num
	return ret
}

func GetJsonString(json string) string {
	return json[2:]
}
