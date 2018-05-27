package main

import (
	"fmt"

	"github.com/spf13/viper"
)

//Cfg...
type Cfg struct {
	//listen
	Listen string

	//pili
	Piliver string
	//mysql
	MHost    string
	MPort    string
	MDB      string
	MPwd     string
	MMaxConn int
	MMaxIdle int
	MTimeout int
	MSplit   int
	Mxml     string
	MTime    int
	//from myql db
	Gid      string
	Domain   string
	Hosts    string
	AllHosts string

	//stream timeout
	StreamTimeOut int

	//cloud storage
	Fragments string
	Snapshots string

	//pfop path
	DealPath string

	//pfop auth check
	PfopAccessKey string
	PfopSecretKey string

	//static subdomain
	Static string

	//auth check
	Auth int
}

func (c *Cfg) Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("cfg")
	viper.SetConfigType("toml")

	viper.SetDefault("listen", ":80")
	viper.SetDefault("piliver", "v1")

	viper.SetDefault("mhost", "127.0.0.1")
	viper.SetDefault("mport", "3036")
	viper.SetDefault("mdb", "test")
	viper.SetDefault("mpwd", "test")
	viper.SetDefault("mmaxidle", 10)
	viper.SetDefault("mmaxopen", 100)
	viper.SetDefault("mtimeout", 10)
	viper.SetDefault("msplit", 256)

	viper.SetDefault("streamtimeout", 10)

	viper.SetDefault("auth", 0)

	err := viper.ReadInConfig()
	fmt.Printf("read ReadInConfig:%v\n", err)

	c.Piliver = viper.GetString("piliver")
	c.Listen = viper.GetString("listen")
	c.MHost = viper.GetString("mhost")
	c.MPort = viper.GetString("mport")
	c.MDB = viper.GetString("mdb")
	c.MPwd = viper.GetString("mpwd")
	c.MMaxConn = viper.GetInt("mmaxopen")
	c.MMaxIdle = viper.GetInt("mmaxidle")
	c.MTimeout = viper.GetInt("mtimeout")
	c.MSplit = viper.GetInt("msplit")
	c.StreamTimeOut = viper.GetInt("streamtimeout")
	c.Mxml = viper.GetString("xml")
	c.MTime = viper.GetInt("time")
	//auth check
	c.Auth = viper.GetInt("auth")

	fmt.Printf("host:%v\n", c.MHost)
}
