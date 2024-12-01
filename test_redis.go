package main

import (
	"context"
	"fmt"
	"time"

	"github.com/horm-database/go-horm/horm"
)

// TestRedis 测试 redis
func TestRedis() {
	queryMultiReturn()
}

func queryMultiReturn() {
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	birthday, _ := time.Parse("2006-01-02", "1987-08-27")
	data := Student{
		Identify: 430602198702221111,
		Gender:   1,
		Age:      19,
		Name:     "caohao",
		Score:    89.7,
		Image:    []byte("IMAGE.PCG"),
		Article:  "Artificial Intelligence",
		ExamTime: "15:30:00",
		Birthday: birthday,
	}

	//horm 会对结构体参数自动编解码
	isNil, err := horm.NewQuery("student_score_range").
		ZAdd("student_score", data, data.Score).Exec(ctx)

	results := make([]*Student, 0)
	scores := make([]float64, 0)
	isNil, err = horm.NewQuery("student_score_range").
		ZRangeByScore("student_score", 70, 100, true).Exec(ctx, &results, &scores)
	fmt.Println(isNil, err)
}
