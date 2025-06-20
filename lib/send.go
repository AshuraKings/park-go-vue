package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
)

func StmtClose(stmt *sql.Stmt, w http.ResponseWriter) {
	DefaultCatch(w)
	if stmt != nil {
		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}
}

func RowsClose(rows *sql.Rows, w http.ResponseWriter) {
	DefaultCatch(w)
	if rows != nil {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}
}

func DbClose(db *sql.DB, w http.ResponseWriter) {
	DefaultCatch(w)
	if db != nil {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}
}

func TxClose(tx *sql.Tx, w http.ResponseWriter) {
	if tx != nil {
		if r := recover(); r != nil {
			msg := fmt.Sprint("", r)
			log.Println(msg)
			log.Println(string(debug.Stack()))
			if strings.HasPrefix(msg, "method") {
				w.WriteHeader(405)
			} else if strings.HasPrefix(msg, "Token") {
				w.WriteHeader(403)
			} else if strings.HasPrefix(msg, "bad:") {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			if err := tx.Rollback(); err != nil {
				panic(err)
			}
			SendJson(map[string]string{"msg": msg}, w)
		} else {
			if err := tx.Commit(); err != nil {
				panic(err)
			}
		}
	} else {
		DefaultCatch(w)
	}
}

func QueryToMap(query url.Values) map[string]string {
	result := map[string]string{}
	for k, v := range query {
		result[k] = strings.Join(v, ", ")
	}
	return result
}

func DefaultCatch(w http.ResponseWriter) {
	if r := recover(); r != nil {
		msg := fmt.Sprint("", r)
		log.Println(msg)
		log.Println(string(debug.Stack()))
		if strings.HasPrefix(msg, "method") {
			w.WriteHeader(405)
		} else if strings.HasPrefix(msg, "Token") {
			w.WriteHeader(403)
		} else if strings.HasPrefix(msg, "bad:") {
			w.WriteHeader(400)
		} else if msg == "Not Found" {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(500)
		}
		SendJson(map[string]string{"msg": msg}, w)
	}
}

func SendJson(resp any, w http.ResponseWriter) {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
