package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/cshengqun/myblog/env"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", Env.Conf.DbUser+ ":" + Env.Conf.DbPasswd + "@tcp("+ Env.Conf.DbIp + ":" + Env.Conf.DbPort + ")/" + Env.Conf.DbName + "?charset=utf8");
	if err != nil {
		panic(err)
	}

	err = SqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
