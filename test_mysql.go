package main

import (
	"context"
	"fmt"
	"time"

	"github.com/horm-database/common/proto"
	"github.com/horm-database/go-horm/horm"
)

// TestMySQL 测试 mysql
func TestMySQL() {
	insertMysql()
}

func insertMysql() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	var ret proto.ModResult

	data := horm.Map{
		"_id":        1,
		"article_id": 1,
		"title":      "推进美丽中国建设",
	}

	_, err := horm.NewQuery("student").
		Replace(data).
		Exec(ctx, &ret)

	fmt.Println(err)
}
