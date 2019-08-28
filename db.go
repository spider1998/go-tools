package tools

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strconv"
)

//根据数据集查询目标字段集合（包含重复记录，一一对应）
func QueryRelated(dsn string, tableName, key, source string, array interface{}) (result []string, err error) {
	T := reflect.TypeOf(array)
	if T.Kind() != reflect.Slice {
		panic("only slice allowed")
	}
	s := reflect.ValueOf(array)
	to := make([]interface{}, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		to = append(to, s.Index(i).Interface())
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	defer db.Close()
	var values string
	var quote string
	types := reflect.TypeOf(to[0])
	switch types.String() {
	case "string":
		for i, ar := range to {
			if i == len(to)-1 {
				quote = ""
			} else {
				quote = ","
			}
			values += "'" + ar.(string) + "'" + quote
		}
	case "int":
		for i, ar := range to {
			if i == len(to)-1 {
				quote = ""
			} else {
				quote = ","
			}
			values += strconv.Itoa(ar.(int)) + quote
		}
	}
	sq := "select `" + key + "`,`" + source + "` from " + tableName + " where `" + key + "` in (" + values + ")"

	rows, err := db.Query(sq)
	if err != nil {
		return
	}
	mresult := make(map[interface{}]interface{})
	var (
		k string
		v string
	)
	for rows.Next() {
		err = rows.Scan(&k, &v)
		if err != nil {
			return
		}
		mresult[k] = v
	}
	_ = rows.Close()
	for _, t := range to {
		result = append(result, mresult[t].(string))
	}
	return
}
