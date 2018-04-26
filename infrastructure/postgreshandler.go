package infrastructure


import (
	"database/sql"
	//"fmt"
	_ "github.com/lib/pq"
	//"time"
	"fmt"
	"log"
	"github.com/ieltof/interfaces"
)

type Type string

type Info struct {
	Type Type
	Postgres PostgresInfo
}

type PostgresInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

func DSN(ci PostgresInfo) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		ci.Username, ci.Password, ci.Name)
}

type PostgresHandler struct {
	Conn *sql.DB
}

func (handler *PostgresHandler) Execute(statement string) {
	var err error
	_, err = handler.Conn.Exec(statement)

	if err != nil {
		log.Fatal(err)
	}
}

func (handler *PostgresHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(PostgersRow)
	}
	row := new(PostgersRow)
	row.Rows = rows
	return row
}

type PostgersRow struct {
	Rows *sql.Rows
}

func (r PostgersRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r PostgersRow) Next() bool {
	return r.Rows.Next()
}

func NewPostgresHandler(d Info) *PostgresHandler {
	var err error
	conn, _ := sql.Open("postgres", DSN(d.Postgres))
	if conn, err = sql.Open("postgres", DSN(d.Postgres)); err != nil {
		log.Println("SQL Driver Error", err)
	}
	postgresHandler := new(PostgresHandler)
	postgresHandler.Conn = conn
	return postgresHandler
}