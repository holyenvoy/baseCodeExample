package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"time"

	"crypto/sha256"
)

type Info struct {
	DbHash    int
	TableHash int
}

func (i Info) Equals(info Info) bool {
	return (i.DbHash == info.DbHash) && (i.TableHash == info.TableHash)
}

func Crc32(input int64) int64 {
	return int64(crc32.ChecksumIEEE([]byte(strconv.FormatInt(input, 10))))
}

func GetDbIndex(input int64, count int) int {
	return (int(input) >> 5) % count
}

func GetTableIndex(input int64, count int) int {
	return int(input) % count
}

func GenInfo(input int64) Info {
	input = Crc32(input)
	dbIndex := GetDbIndex(input, 1)
	tableIndex := GetTableIndex(input, 2)
	hashInfo := Info{DbHash: dbIndex, TableHash: tableIndex}
	return hashInfo
}

func GenGroupInfo(input int64) Info {
	input = Crc32(input)
	dbIndex := GetDbIndex(input, 32)
	tableIndex := GetTableIndex(input, 32)
	hashInfo := Info{DbHash: dbIndex, TableHash: tableIndex}
	return hashInfo
}

func makeSha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func GetFake(timestamp int64) int64 {
	return int64(timestamp << 22)
}

func main() {

	var fromId, toId int64
	var agroaId1 uint32
	var agroaId2 int32

	fromId = 4294967295
	agroaId1 = uint32(fromId)
	agroaId2 = int32(fromId)
	fmt.Println(agroaId1)
	fmt.Println(agroaId2)

	toId = int64(-1)
	fmt.Println(toId)

	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(uint32(now))

	Info := GenGroupInfo(6319115602384783361)
	fmt.Printf("6319115602384783361:%v\n", Info)

	Info = GenGroupInfo(6319116191588026369)
	fmt.Printf("6319116191588026369:%v\n", Info)

	//hash = hex.EncodeToString(makeSha256(s.Body))
	//6286583896473600000
	fmt.Printf("uuid:%v\n", GetFake(1507710136000))

}
