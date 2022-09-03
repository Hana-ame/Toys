package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DAO struct {
	db  *sql.DB
	err error
}

var Dao DAO

func (dao *DAO) onError() {
	if dao.err != nil {
		fmt.Println(dao.err)
		dao.err = nil
	}
}
func (dao *DAO) Open(driver, dataSource string) {
	dao.db, dao.err = sql.Open(driver, dataSource)
	dao.onError()
}
func (dao *DAO) Close() {
	dao.db.Close()
}
func (dao *DAO) Ping() {
	dao.err = dao.db.Ping()
	dao.onError()
}
func (dao *DAO) AlwaysPing() {
	go func() {
		for {
			time.Sleep(time.Hour)
			dao.Ping()
		}
	}()
}
func (dao *DAO) Prepare(s string) (stmt *sql.Stmt) {
	stmt, dao.err = dao.db.Prepare(s)
	dao.onError()
	return
}
func (dao *DAO) Query(s string, paras ...interface{}) (rows *sql.Rows) {
	// fmt.Println(paras...)
	rows, dao.err = dao.db.Query(s, paras...)
	dao.onError()
	return
}

// func (dao *DAO) QueryMetaData(start, end int) result []MetaData {
// 	return "Select * From metadata Limit ?,? ;"
// }
// func (dao *DAO) StrQueryBulletinData(bulletinID int) string {
// 	return "Select * From Bulletine" + strconv.Itoa(bulletinID) + " ;"
// }
// func (dao *DAO) StrPrepareCreatBulletin(name, describe string) string {
// 	return "Insert Into metadata (bulletin_name,bulletin_describe) values(?,?) ;"
// }
// func (dao *DAO) Query(query string) (res *sql.Rows) {
// 	res, dao.err = dao.db.Query(query)
// 	return res
// }

func DaoInit() {

	Dao = DAO{}
	Dao.Open("mysql", "kaban:serval@tcp(localhost:3306)/bubble")
	defer Dao.Close()
	fmt.Println("ready")
	for {
		time.Sleep(time.Hour)
		Dao.Ping()
	}
	fmt.Println("ready for die")
}

func InsertBubble(bubble Bubble) {
	stmt := Dao.Prepare("INSERT INTO bubbles(  pos_x,   pos_y,  size_x,  size_y,  scale,  background,  content,  meta_imgsrc,  meta_title,  meta_author  ) VALUES (?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	_, err := stmt.Exec(
		int(bubble.Pos.X),
		int(bubble.Pos.Y),
		int(bubble.Size.X),
		int(bubble.Size.Y),
		bubble.Scale,
		bubble.Background,
		bubble.Content,
		bubble.Meta.Src,
		bubble.Meta.Title,
		bubble.Meta.Author,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func QueryBubbles(page int, bn string) (l []Bubble) {
	i := 0
	l = make([]Bubble, 100)
	res := Dao.Query("SELECT  id,pos_x,  pos_y, size_x, size_y, scale, background, content, meta_imgsrc, meta_title, meta_author, create_time  FROM bubbles where board_name = ? LIMIT ?, 100;", bn, page*100)
	for res.Next() {
		err := res.Scan(
			&l[i].Id,
			&l[i].Pos.X,
			&l[i].Pos.Y,
			&l[i].Size.X,
			&l[i].Size.Y,
			&l[i].Scale,
			&l[i].Background,
			&l[i].Content,
			&l[i].Meta.Src,
			&l[i].Meta.Title,
			&l[i].Meta.Author,
			&l[i].Meta.Timestamp,
		)
		if err != nil {
			log.Fatal(err)
		}
		i += 1
	}
	l = l[:i]
	return
}
