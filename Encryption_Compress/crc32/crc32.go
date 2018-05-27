package main

import "fmt"
import "hash/crc32"
import "strconv"

func getUids(uids ...int) {
	fmt.Println(uids)
}

type Info struct {
	DbHash    int
	TableHash int
}

func main() {
	for i := 0; i < 10; i++ {
		value := Crc32Int(i)
		fmt.Println(value % 100)

	}

	fmt.Println(Crc32(6293011573812430849) % 2)
	fmt.Println(GetDbIndex(102990875, 1))
	fmt.Println(GetTableIndex(100050838, 64))

	fmt.Println(GenLikesInfo(102990875))
}

func Crc32(input int64) int64 {
	return int64(crc32.ChecksumIEEE([]byte(strconv.FormatInt(input, 10))))
}
func Crc32Int(input int) int {
	return int(crc32.ChecksumIEEE([]byte(strconv.Itoa(input))))
}

func GetDbIndex(input int64, count int) int {
	input = Crc32(input)

	return (int(input) >> 5) % count
}

func GetTableIndex(input int64, count int) int {
	return int(input) % count
}

func GenLikesInfo(input int64) Info {
	input = Crc32(input)
	dbIndex := GetDbIndex(input, 1)
	tableIndex := GetTableIndex(input, 64)
	hashInfo := Info{DbHash: dbIndex, TableHash: tableIndex}
	return hashInfo
}
