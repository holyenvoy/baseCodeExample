// 当前程序的包名
package main

// 导入其它的包
import (
	"encoding/json"
	"fmt"
)

func main() {
	map2json2map()
}

func map2json2map() {

	map1 := make(map[string]interface{})
	map1["10000"] = 6265725593732712449
	map1["0"] = 6265802261981759489

	//return []byte
	str, err := json.Marshal(map1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json\n", string(str))

	//json([]byte) to map
	map2 := make(map[string]interface{})
	err = json.Unmarshal(str, &map2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json to map :%v \n", map2)

	fmt.Printf("The value of key1 is:%v", map2["0"])
}
