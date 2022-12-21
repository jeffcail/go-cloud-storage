package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jeffcail/cloud-storage/server/models"

	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

func TestXorm(t *testing.T) {
	dbDsn := "root:123456@tcp(127.0.0.1:3306)/cloud_storage?charset=utf8mb4&parseTime=True&loc=Local"
	engine, er := xorm.NewEngine("mysql", dbDsn)
	if er != nil {
		t.Fatal(er)
	}
	err := engine.Ping()
	if er != nil {
		t.Fatal(err)
	}
	data := make([]*models.StorageUser, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
