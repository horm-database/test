package main

import (
	"time"
)

type Student struct {
	Id        int       `orm:"id,int,omitempty" json:"id"`
	Identify  int64     `orm:"identify,bigint" json:"identify"`                     //身份证件ID
	Gender    int8      `orm:"gender,tinyint,omitinsertempty" json:"gender"`        //1-male 2-female
	Age       uint      `orm:"age,int unsigned,omitreplaceempty" json:"age"`        //年龄
	Name      string    `orm:"name,varchar,omitupdateempty" json:"name"`            //名称
	Score     float64   `orm:"score,double,omitempty" json:"score"`                 //分数
	Image     []byte    `orm:"image,blob,omitempty" json:"image"`                   //image
	Article   string    `orm:"article,text,omitempty" json:"article"`               //publish article
	ExamTime  string    `orm:"exam_time,time,omitempty" json:"exam_time"`           //考试时间
	Birthday  time.Time `orm:"birthday,date" json:"birthday"`                       //出生日期
	CreatedAt time.Time `orm:"created_at,timestamp,oncreatetime" json:"created_at"` //创建时间
	UpdatedAt time.Time `orm:"updated_at,datetime,onupdatetime" json:"updated_at"`  //修改时间
}
