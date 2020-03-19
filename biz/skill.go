package biz

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/zii/pet-sim/base"
)

type Skill struct {
	Id   int
	Name string
	Des  string
}

var skillSet = new(sync.Map)

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, ",")
}

func InitSkill(filename string) {
	f, err := os.Open(filename)
	base.Raise(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		l = strings.TrimSpace(l)
		rows := strings.Split(l, ",")
		if len(rows) < 12 {
			continue
		}
		sk := new(Skill)
		sk.Id = base.ToInt(rows[6])
		sk.Name = rows[0]
		sk.Des = rows[1]
		sk.Des = compressStr(sk.Des)
		skillSet.Store(sk.Id, sk)
	}
}

// 精灵
func InitMagic(filename string) {
	f, err := os.Open(filename)
	base.Raise(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		l = strings.TrimSpace(l)
		rows := strings.Split(l, ",")
		if len(rows) < 9 {
			log.Fatalln("magic字段缺失!", l)
		}
	}
}

func GetSkill(id int) *Skill {
	v, ok := skillSet.Load(id)
	if !ok {
		return nil
	}
	sk, _ := v.(*Skill)
	return sk
}
