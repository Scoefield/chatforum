package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	driverName = "mysql"
	userName = "root"
	passWord = "123456"
	port = "3306"
	host = "192.168.7.3"
	dbName = "chitchat"

)

var Db *sql.DB

func init() {
	var err error
	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?harset=utf8&parseTime=true", userName, passWord, host, port, dbName)
	//Db, err = sql.Open("mysql", "root:123456@tcp(192.168.7.137:3306)/chitchat?charset=utf8&parseTime=true")
	Db, err = sql.Open(driverName, dbPath)
	if err != nil {
		log.Fatal("sql.Open err ", err)
	}
	return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Connot generate uuid", err)
	}
	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])

	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
