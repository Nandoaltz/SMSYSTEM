package banco

import (
	"TCC/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB()(*sql.DB, error){
	db, erro := sql.Open("mysql", config.StringBanco)
	if erro != nil{
		return nil, erro
	}
	if erro := db.Ping(); erro != nil{
		return nil, erro
	}
	return db, nil
}