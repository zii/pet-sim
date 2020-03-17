// 角色: 人物 or 宠物
package biz

import (
	"math/rand"
	"sync"
)

const (
	CHAR_MAXUPLEVEL = 140
)

// 角色类型
const (
	CHAR_TYPENONE = 0
	CHAR_TYPEPLAYER
	CHAR_TYPEENEMY
	CHAR_TYPEPET
	CHAR_TYPEDOOR
	CHAR_TYPEBOX
	CHAR_TYPEMSG
	CHAR_TYPEWARP
	CHAR_TYPESHOP
	CHAR_TYPEHEALER
	CHAR_TYPEOLDMAN
	CHAR_TYPEROOMADMIN
	CHAR_TYPETOWNPEOPLE
	CHAR_TYPEDENGON
	CHAR_TYPEADM
	CHAR_TYPETEMPLE
	CHAR_TYPESTORYTELLER
	CHAR_TYPERANKING
	CHAR_TYPEOTHERNPC
	CHAR_TYPEPRINTPASSMAN
	CHAR_TYPENPCENEMY
	CHAR_TYPEACTION
	CHAR_TYPEWINDOWMAN
	CHAR_TYPESAVEPOINT
	CHAR_TYPEWINDOWHEALER
	CHAR_TYPEITEMSHOP
	CHAR_TYPESTONESHOP
	CHAR_TYPEDUELRANKING
	CHAR_TYPEWARPMAN
	CHAR_TYPEEVENT
	CHAR_TYPEMIC
	CHAR_TYPELUCKYMAN
	CHAR_TYPEBUS
	CHAR_TYPECHARM
	CHAR_TYPECHECKMAN
	CHAR_TYPEJANKEN
	CHAR_TYPETRANSMIGRATION
	CHAR_TYPEFMWARPMAN        // 家族ＰＫ场管理员
	CHAR_TYPEFMSCHEDULEMAN    // 家族ＰＫ场登记员
	CHAR_TYPEMANORSCHEDULEMAN // 庄园ＰＫ场预约人
	CHAR_GAMBLEBANK
	CHAR_NEWNPCMAN

	CHAR_GAMBLEROULETTE
	CHAR_GAMBLEMASTER
	CHAR_TRANSERMANS

	CHAR_ITEMCHANGENPC
	CHAR_FREESKILLSHOP
	CHAR_PETRACEMASTER // 宠物竞速
	CHAR_PETRACEPET

	CHAR_TYPEALLDOMAN

	CHAR_TYPEPETMAKER // petmaker

	CHAR_TYPENUM
)

type Char struct {
	Id         int
	Name       string
	ImgNo      int
	WhichType  int // 角色类型 CHAR_TYPE
	DuelPoint  int
	Vital      int // 体力
	Str        int // 力量
	Tough      int // 坚韧
	Dex        int // 敏捷
	EarthAT    int
	WaterAT    int
	FireAT     int // 火属性
	WindAT     int
	ModAI      int
	VariableAI int
	Slot       int
	Poison     int
	Paralysis  int
	Sleep      int
	Stone      int
	Drunk      int
	Confusion  int
	Rare       int
	PetId      int // EnemyBase.No, TEMPNO
	Critical   int
	Counter    int
	Luck       int
	PetSkill1  int
	PetSkill2  int
	PetSkill3  int
	PetSkill4  int
	PetSkill5  int
	PetSkill6  int
	PetSkill7  int
	PetRank    int
	Exp        int
	AllocPoint [4]int

	WorkTactics            int
	WorkTacticsOption      string
	WorkBattleActContition string
	WorkPetFlag            int
	WorkModCaptureDefault  int // EnemyBase.Get

	Hp int // 当前剩余血量
	Mp int // 当前剩余法力

	BornLv    int    // 初始级别
	Lv        int    // 当前级别
	BornPoint [4]int // 初始四维
	// 当前四维
	WorkMaxHp    int // 当前最大血量
	WorkMaxMp    int // 当前最大法力
	WorkFixVital int // 当前的体力档
	WorkFixStr   int // 当前攻
	WorkFixTough int // 当前防
	WorkFixDex   int // 当前敏
	// 成长率
	GrowthHp    float32 // 血成长率
	GrowthStr   float32 // 攻成长率
	GrowthTough float32 // 防成长率
	GrowthDex   float32 // 敏成长率
	Growth      float32 // 总成长率
}

var charIDCounter int
var charSet *sync.Map

func InitChar() {
	charSet = new(sync.Map)
}

func NewCharID() int {
	charIDCounter += 1
	return charIDCounter
}

func InitNewChar(char *Char) {
	char.Id = NewCharID()
	charSet.Store(char.Id, char)
}

func GetChar(id int) *Char {
	v, ok := charSet.Load(id)
	if !ok {
		return nil
	}
	c, _ := v.(*Char)
	return c
}

// 只是对四维形状的一个修整, 不影响大局
func Char_initCharWork(char *Char) {
	char.WorkFixDex = char.Dex / 100
	char.WorkFixVital = char.Vital / 100
	char.WorkFixStr = int(float32(char.Str)*1.0+float32(char.Tough)*0.1+float32(char.Vital)*0.1+float32(char.Dex)*0.05) / 100
	char.WorkFixTough = int(float32(char.Tough)*1.0+float32(char.Str)*0.1+float32(char.Vital)*0.1+float32(char.Dex)*0.05) / 100
	char.WorkMaxHp = (char.Vital*4 + char.Str*1 + char.Tough*1 + char.Dex*1) / 100
}

// 调整角色参数
func Char_complianceParameter(char *Char) {
	Char_initCharWork(char)
	char.Hp = char.WorkMaxHp
	char.Mp = char.WorkMaxMp
}

func PetLevelUp(char *Char) {
	if char.Lv >= CHAR_MAXUPLEVEL {
		return
	}
	var ranktab = []struct {
		min int
		max int
	}{
		{450, 500},
		{470, 520},
		{490, 540},
		{510, 560},
		{530, 580},
		{550, 600},
	}

	var param = [4]float32{}
	for i := 0; i < 10; i++ {
		j := rand.Intn(4)
		param[j] = 1.0
	}
	petrank := char.PetRank
	// 这个frand和param 只是让每一次升级过程看起来玄妙一些 但100多级平均下来 是很稳定的一个期望值 不影响总成长
	frand := float32(ranktab[petrank].min+rand.Intn(ranktab[petrank].max-ranktab[petrank].min+1)) * 0.01
	vital := float32(char.AllocPoint[0])
	str := float32(char.AllocPoint[1])
	tgh := float32(char.AllocPoint[2])
	dex := float32(char.AllocPoint[3])
	vital = (vital + param[0]) * frand
	str = (str + param[1]) * frand
	tgh = (tgh + param[2]) * frand
	dex = (dex + param[3]) * frand
	char.Vital += int(vital)
	char.Str += int(str)
	char.Tough += int(tgh)
	char.Dex += int(dex)

	char.Lv += 1
	Char_complianceParameter(char)

	// 算成长率
	lvd := char.Lv - char.BornLv
	if lvd > 0 {
		char.GrowthHp = float32(char.WorkMaxHp-char.BornPoint[0]) / float32(lvd)
		char.GrowthStr = float32(char.WorkFixStr-char.BornPoint[1]) / float32(lvd)
		char.GrowthTough = float32(char.WorkFixTough-char.BornPoint[2]) / float32(lvd)
		char.GrowthDex = float32(char.WorkFixDex-char.BornPoint[3]) / float32(lvd)
		char.Growth = char.GrowthStr + char.GrowthTough + char.GrowthDex
	}
}
