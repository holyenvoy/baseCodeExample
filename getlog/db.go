package main

import (
	"database/sql"
	"fmt"
	"strings"

	seelog "github.com/cihub/seelog"
)

type MySQL struct {
	mysql *sql.DB
	log   seelog.LoggerInterface
}

func (m *MySQL) Init(c *Cfg, log seelog.LoggerInterface) bool {
	m.log = log
	connStr := c.MDB + ":" + c.MPwd + "@tcp(" + c.MHost + ":" + c.MPort + ")/" + c.MDB + "?timeout=" + fmt.Sprintf("%d", c.MTimeout) + "s&charset=utf8"
	var err error
	if m.mysql, err = sql.Open("mysql", connStr); err != nil {
		m.log.Errorf("mysql db open err:", err.Error())
		return false
	}
	m.mysql.SetMaxOpenConns(c.MMaxConn)
	m.mysql.SetMaxIdleConns(c.MMaxIdle)
	m.mysql.Ping()
	log.Info("init mysql ok")
	return true
}

func (m *MySQL) QueryRow(s string) *sql.Row {
	var row *sql.Row
	if m.mysql != nil {
		if row = m.mysql.QueryRow(s); row == nil {
			m.log.Errorf("row == nil:%s", s)
		}
	} else {
		m.log.Error("mysql == nil")
	}
	return row
}

func (m *MySQL) Query(s string) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	if m.mysql != nil {
		if rows, err = m.mysql.Query(s); err != nil && strings.Contains(err.Error(), "closed") {
			//log.Error(err.Error())
			rows, err = m.mysql.Query(s)
		}
	} else {
		m.log.Error("mysql == nil")
	}
	return rows, err
}

func (m *MySQL) Exec(s string) (sql.Result, error) {
	var res sql.Result
	var err error
	if m.mysql != nil {
		res, err = m.mysql.Exec(s)
		if err != nil && strings.Contains(err.Error(), "closed") {
			m.log.Error(err.Error())
			res, err = m.mysql.Exec(s)
		}
	} else {
		m.log.Error("mysql == nil")
	}
	return res, err
}

type DB struct {
	MySQL MySQL
}

func (d *DB) Init(c *Cfg, log seelog.LoggerInterface) {
	d.MySQL.Init(c, log)
}
