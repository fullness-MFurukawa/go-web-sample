package mysql

import (
	"database/sql"
	"go-web-sample/infrastructure"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlConnector struct {
}

// MySQLデータベースに接続する
func (con *mysqlConnector) Connect() (*sql.DB, error) {
	conn, err := sql.Open("mysql", "root:password@/web_sample?parseTime=true")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// コンストラクタ
func MySqlConnector() infrastructure.Connector {
	connector := mysqlConnector{}
	return &connector
}
