package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// データベース接続機能のテスト
//
func TestConnect(t *testing.T) {
	connector := MySqlConnector()
	db, _ := connector.Connect()
	defer db.Close()
	assert.NotNil(t, db)
}
