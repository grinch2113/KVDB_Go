package main

import (
	handler "GO_KVDB/KVDBHandler"
	"fmt"
)

func help() {
	fmt.Println("welcome to kv database system")
	fmt.Println("1. set a k-v group into current database")
	fmt.Println("2. get a value by key from current database")
	fmt.Println("3. quit")
	fmt.Println(`type "help" to get options`)
}
func main() {
	help()
	for {
		var option string
		fmt.Scanln(&option)
		switch option {
		case "1":
			var key, value string
			fmt.Scanln(&key, &value)
			handler.Set(key, value)
		case "2":
			var key string
			fmt.Scanln(&key)
			value := handler.Get(key)
			if len(value) == 0 {
				fmt.Println("未找到对应的value")
			} else {
				fmt.Println(value[:len(value)-1])
			}

		case "3":
			return
		case "help":
			help()
		default:
			fmt.Println("invalid option, type 5 to get options")
		}
	}

}
