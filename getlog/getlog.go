package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
)

var db DB

var cfg Cfg

//var log logger.LoggerInterface

type Result struct {
	XMLName      xml.Name          `xml:"rtmp"`
	Applications []ApplicationType `xml:"server>application"`
}

type ApplicationType struct {
	Name   string       `xml:"name"`
	Stream []StreamType `xml:"live>stream"`
}

type StreamType struct {
	Name      string  `xml:"name"`
	Bytes_in  float32 `xml:"bytes_in"`               //Audio video bit rate
	Video_fps float32 `xml:"meta>video>frame_rate"`  //video frame rate
	Audio_fps float32 `xml:"meta>audio>sample_rate"` //audio frame rate
	Data      float32
}

func InitAll() bool {
	logger, err := log.LoggerFromConfigAsFile("./conf/logConfig.xml")
	if err != nil {
		log.Errorf("load logConfig.xml fail: %s:", err)
		panic(err)
	}
	log.ReplaceLogger(logger)
	fmt.Println("wwwwwww")
	defer log.Flush()
	log.Info("InitAll success.")
	cfg.Init()
	db.Init(&cfg, logger)
	//	log.ReplaceLogger(logger)

	return true
}

func FilterParameters(res *Result) {

	if res == nil {
		log.Errorf("[FilterParameters]res is nil")
		return
	}

	ApplicationNum := len(res.Applications)
	//log.Info(num)
	for i := 0; i < ApplicationNum; i++ {
		//	log.Info(i)
		ApplicationName := res.Applications[i].Name
		if ApplicationName == "mt_test" {

			StreamNum := len(res.Applications[i].Stream)

			for j := 0; j < StreamNum; j++ {
				NameId := res.Applications[i].Stream[j].Name
				Bytes := res.Applications[i].Stream[j].Bytes_in
				Afps := res.Applications[i].Stream[j].Audio_fps
				Vfps := res.Applications[i].Stream[j].Video_fps
				Data := res.Applications[i].Stream[j].Data
				s := fmt.Sprintf("update stream set speed = '%f',audio = '%f',video = '%f',data='%f' where streamid = '%s' limit 1", Bytes, Afps, Vfps, Data, NameId)
				log.Info("sql：", s)
				if _, err := db.MySQL.Exec(s); err != nil {
					log.Errorf(err.Error())
				}
			}

		} else {
			continue
		}
	}
	return

}

func main() {
	defer log.Flush()

	if !InitAll() {
		log.Error("[main] InitAll fail.")
	}
	//	fmt.Println("type", reflect.TypeOf(time.Second*10))

	log.Infof("[main] InitAll success.")
	var Source string = cfg.Mxml
	var num = cfg.MTime
	times := time.Duration(num)

	for {
		resp, err := http.Get(Source)
		if err != nil {
			log.Error(err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		//body, err := ioutil.ReadFile("rtmp.xml")
		if err != nil {
			log.Error(err.Error())

		}
		var result Result
		err = xml.Unmarshal(body, &result)
		if err != nil {
			log.Error(err.Error())
		}

		FilterParameters(&result)
		time.Sleep(time.Second * times)

	}
	//	log.Info(len(result.Applications))
	//	log.Info(result.Applications[1].Stream[0].Data)
	//	log.Info(result)
}
