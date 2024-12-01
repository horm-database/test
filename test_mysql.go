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
	//insertMysql()
	//findAllByGlobalClient()
	//findAllByClient()
	//findAllByClientWithParams()
	//queryModeSingle()
}

func insertMysql() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	var ret proto.ModResult

	birthday, _ := time.Parse("2006-01-02", "1987-08-27")
	data := Student{
		Identify:  430602198702221111,
		Gender:    1,
		Age:       19,
		Name:      "caohao",
		Score:     89.7,
		Image:     []byte("IMAGE.PCG"),
		Article:   "Artificial Intelligence",
		ExamTime:  "15:30:00",
		Birthday:  birthday,
		CreatedAt: time.Now(),
	}

	_, err := horm.NewQuery("student").Insert(data).Exec(ctx, &ret)

	fmt.Println(err)
}

func findAllByGlobalClient() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	var result = make([]*Student, 0)
	_, err := horm.NewQuery("student").FindAll().Exec(ctx, &result)

	fmt.Println(err)
}

func findAllByClient() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	cli := horm.NewClient("ws_test.app2.server2.service2")

	var result = make([]*Student, 0)
	_, err := horm.NewQuery("student").FindAll().WithClient(cli).Exec(ctx, &result)

	fmt.Println(err)
}

func findAllByClientWithParams() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	cli := horm.NewClient("ws_test.app2.server2.service2",
		horm.WithAppID(10099),
		horm.WithSecret("S499721834"))

	var result = make([]*Student, 0)
	_, err := horm.NewQuery("student").FindAll().WithClient(cli).Exec(ctx, &result)

	fmt.Println(err)
}

func queryModeSingle() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	var result Student
	var where = horm.Where{"id": 13}
	isNil, err := horm.NewQuery("student").Find(where).Exec(ctx, &result)

	fmt.Println(isNil, err)
}
