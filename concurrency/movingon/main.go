package main

import "database/sql"

// Db db connection
type DB sql.DB

var db *DB

func init() {
	// init db connection
}

// example of moving on
func main() {
	Query(db, "sql ...")
}

// oving on 示例；
// go 的闭包中，查询语句直接在 select 语句的 case 中；按道理 ch 是无 buffer 的通道如果没有接收会阻塞
// 但是在这里，如果 ch <- db.QueryRow(query) 无法立即返回要阻塞的话，会转而执行 default 分支
func Query(db *DB, query string) Result {
	ch := make(chan Result)
	for i := 0; i <= 10; i++ {
		go func(db *DB) {
			select {
			case ch <- db.QueryRow(query):
			default:
			}
		}(db)
	}
	return <-ch
}

type Result struct{}

func (d *DB) QueryRow(query string) Result {
	d.QueryRow(query)
	return Result{}
}
