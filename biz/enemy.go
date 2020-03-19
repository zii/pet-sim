package biz

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/zii/pet-sim/base"
)

// enemybase.txt [:7]前7个字段(都是字符串)
const (
	E_T_NAME         = 0 // 宠物品种名称
	E_T_ATOMFIXNAME1     // 素材1 ...
	E_T_ATOMFIXNAME2
	E_T_ATOMFIXNAME3
	E_T_ATOMFIXNAME4
	E_T_ATOMFIXNAME5
	E_T_DATACHARNUM
)

// enemybase.txt [7:]后面的字段(都是整型)
const (
	E_T_TEMPNO    = 0 // 种类编号
	E_T_INITNUM       // 初始成长值
	E_T_LVUPPOINT     // 野生成长率
	E_T_BASEVITAL     // 基础体力档
	E_T_BASESTR       // 基础力量档
	E_T_BASETGH       // 基础耐力档
	E_T_BASEDEX       // 基础敏捷档
	E_T_MODAI         // AI档
	E_T_GET           // 捕捉难度 越大越容易
	E_T_EARTHAT       // 地属性 10的整数倍
	E_T_WATERAT       // 水属性 10的整数倍
	E_T_FIREAT        // 火属性 10的整数倍
	E_T_WINDAT        // 风属性 10的整数倍
	E_T_POISON        // 毒抗性 0-1000
	E_T_PARALYSIS     // 麻痹抗性 0-1000
	E_T_SLEEP         // 睡眠抗性 0-1000
	E_T_STONE         // 石化抗性 0-1000
	E_T_DRUNK         // 酒醉抗性 0-1000
	E_T_CONFUSION     // 混乱抗性 0-1000
	E_T_PETSKILL1
	E_T_PETSKILL2
	E_T_PETSKILL3
	E_T_PETSKILL4
	E_T_PETSKILL5
	E_T_PETSKILL6
	E_T_PETSKILL7
	E_T_RARE         // 稀有度 越高越稀有 决定宠物价格和逃跑成功率
	E_T_CRITICAL     // 暴击率
	E_T_COUNTER      // 反击率
	E_T_SLOT         // 技能格数
	E_T_IMGNUMBER    // 宠物图片编号
	E_T_PETFLG       // 是否可捕获(没用)
	E_T_SIZE         // 大小 决定宠物在战场上所占有的位置 0:小 1:大 在刷怪时size=1时最多只能刷出5只
	E_T_ATOMBASEADD1 // 合成素材参数...
	E_T_ATOMFIXMIN1
	E_T_ATOMFIXMAX1
	E_T_ATOMBASEADD2
	E_T_ATOMFIXMIN2
	E_T_ATOMFIXMAX2
	E_T_ATOMBASEADD3
	E_T_ATOMFIXMIN3
	E_T_ATOMFIXMAX3
	E_T_ATOMBASEADD4
	E_T_ATOMFIXMIN4
	E_T_ATOMFIXMAX4
	E_T_ATOMBASEADD5
	E_T_ATOMFIXMIN5
	E_T_ATOMFIXMAX5
	E_T_LIMITLEVEL // 限制宠物等级
	// 最后一列是宠物系编号 为-1时不能参与融合 乌力系=0 布依系=1 加美系=6 ...
	E_T_DATAINTNUM
)

// enemy.txt [3:]字段int
const (
	ENEMY_ID           = 0 // 敌人ID
	ENEMY_TEMPNO           // 宠物ID
	ENEMY_LV_MIN           // 生成敌人时等级下限
	ENEMY_LV_MAX           // 生成敌人时等级上限
	ENEMY_CREATEMAXNUM     // 创造数量上限 (也就是刷怪的时候,1次最多能刷多少只)
	ENEMY_CREATEMINNUM     // 创造数量下限
	ENEMY_TACTICS          // 战术ID
	ENEMY_EXP              // 战斗后可得经验点数
	ENEMY_DUELPOINT        // DP点数
	ENEMY_STYLE            // 类型 (默认为0 1:斧 2:棍 3:枪 4:弓 5:回旋标 6:投掷石 7:投掷斧)
	ENEMY_PETFLG           // 宠物标记 (0:不可捕获，1:可以捕获)

	ENEMY_ITEM1 // 得到的道具和几率
	ENEMY_ITEM2
	ENEMY_ITEM3
	ENEMY_ITEM4
	ENEMY_ITEM5
	ENEMY_ITEM6
	ENEMY_ITEM7
	ENEMY_ITEM8
	ENEMY_ITEM9
	ENEMY_ITEM10
	ENEMY_ITEMPROB1
	ENEMY_ITEMPROB2
	ENEMY_ITEMPROB3
	ENEMY_ITEMPROB4
	ENEMY_ITEMPROB5
	ENEMY_ITEMPROB6
	ENEMY_ITEMPROB7
	ENEMY_ITEMPROB8
	ENEMY_ITEMPROB9
	ENEMY_ITEMPROB10

	ENEMY_DATAINTNUM
)

// enemy.txt 前3个字段string
const (
	ENEMY_NAME          = 0 // 敌人名 客户端用不到
	ENEMY_TACTICSOPTION     // 战斗策略
	ENEMY_ACT_CONDITION     // 特殊行为,只有个别BOSS用到 tn:2|wp:60307;15;18 萨登第二回合将玩家传送到卡罗的精神结界
	ENEMY_DATACHARNUM
)

// 宠物品种 enemybase.txt
// 举例: 乌力,石,木,皮,骨,线,1,10,4.50,20,12,15,25,150,11,80,20,0,0,0,0,0,0,0,0,1,,,,,,,0,1,1,7,100250,1,0,,700,700,,700,700,,700,700,,700,700,,700,700,,0
type EnemyBase struct {
	Name         string
	AtomFixName1 string
	AtomFixName2 string
	AtomFixName3 string
	AtomFixName4 string
	AtomFixName5 string
	No           int // TEMPNO, PETID
	InitNum      int
	LvUpPoint    int
	BaseVital    int
	BaseStr      int
	BaseTgh      int
	BaseDex      int
	ModAI        int
	Get          int
	EarthAT      int
	WaterAT      int
	FireAT       int
	WindAt       int
	Poison       int
	Paralysis    int
	Sleep        int
	Stone        int
	Drunk        int
	Confusion    int
	PetSkillIds  [9]int
	PetSkills    [9]*Skill
	Rare         int
	Critical     int
	Counter      int
	Slot         int
	ImgNo        int
	PetFlag      int
	Size         int
	LimitLevel   int
	Species      int
}

// enemy.txt
// 敌人, 主要用于战斗场景生成的作战单位, 比如可以配置两种不同战斗属性的[昆依]
// sai_e_113_1昆依,at:20;1;1|gu:0|es:1|wa:0;0;0;0;0;0;0;,,119,113,1,1,2,1,1,-1,-1,0,1,1234,,,,,,,,,,300,,,,,,,,,
// sai_e_113_2/3昆依,at:20;1;1|gu:0|es:1|wa:0;0;0;0;0;0;0;,,120,113,2,3,10,1,1,-1,-1,0,1,1234,,,,,,,,,,300,,,,,,,,,
type Enemy struct {
	Name          string // 这个名字不重要 起到备注作用
	TacticsOption string // 战术配置
	ActCondition  string // 只有个别BOSS用到 tn:2|wp:60307;15;18 萨登第二回合将玩家传送到卡罗的精神结界
	EnemyId       int
	TempNo        int // EnemyBase.No, PET_ID
	LvMin         int
	LvMax         int
	CreateMaxNum  int
	CreateMinNUm  int
	Tactics       int
	Exp           int
	DuelPoint     int
	Style         int
	PetFlag       int
}

// 宠物集合
var EnemyBaseSet *sync.Map
var EnemyNoList []int

func InitEnemyBase(filename string) {
	f, err := os.Open(filename)
	base.Raise(err)
	defer f.Close()

	EnemyBaseSet = new(sync.Map)
	reader := bufio.NewReader(f)
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		rows := strings.Split(l, ",")
		if len(rows) < 56 {
			log.Fatalln("字段缺失!", l)
		}
		eb := new(EnemyBase)
		eb.Name = rows[0]
		eb.AtomFixName1 = rows[1]
		eb.AtomFixName2 = rows[2]
		eb.AtomFixName3 = rows[3]
		eb.AtomFixName4 = rows[4]
		eb.AtomFixName5 = rows[5]
		eb.No = base.ToInt(rows[6])
		eb.InitNum = base.ToInt(rows[7])
		eb.LvUpPoint = int(base.ToFloat(rows[8]))
		eb.BaseVital = base.ToInt(rows[9])
		eb.BaseStr = base.ToInt(rows[10])
		eb.BaseTgh = base.ToInt(rows[11])
		eb.BaseDex = base.ToInt(rows[12])
		eb.ModAI = base.ToInt(rows[13])
		eb.Get = base.ToInt(rows[14])
		eb.EarthAT = base.ToInt(rows[15])
		eb.WaterAT = base.ToInt(rows[16])
		eb.FireAT = base.ToInt(rows[17])
		eb.WindAt = base.ToInt(rows[18])
		eb.Poison = base.ToInt(rows[19])
		eb.Paralysis = base.ToInt(rows[20])
		eb.Sleep = base.ToInt(rows[21])
		eb.Stone = base.ToInt(rows[22])
		eb.Drunk = base.ToInt(rows[23])
		eb.Confusion = base.ToInt(rows[24])
		eb.Slot = base.ToInt(rows[35])
		for j := 0; j < 7; j++ {
			s := rows[25+j]
			if s != "" {
				skid := base.ToInt(s)
				eb.PetSkillIds[j] = skid
				eb.PetSkills[j] = GetSkill(skid)
			}
		}
		eb.Rare = base.ToInt(rows[32])
		eb.Critical = base.ToInt(rows[33])
		eb.Counter = base.ToInt(rows[34])
		eb.ImgNo = base.ToInt(rows[36])
		eb.PetFlag = base.ToInt(rows[37])
		eb.Size = base.ToInt(rows[38])
		eb.LimitLevel = base.ToInt(rows[54])
		eb.Species = base.ToInt(rows[55])
		EnemyBaseSet.Store(eb.No, eb)
		EnemyNoList = append(EnemyNoList, eb.No)
	}
}

func GetEnemyBase(no int) *EnemyBase {
	v, ok := EnemyBaseSet.Load(no)
	if !ok {
		return nil
	}
	eb, _ := v.(*EnemyBase)
	return eb
}

func CreateEnemy(ebno int, baselevel int) *Char {
	eb := GetEnemyBase(ebno)
	if eb == nil {
		return nil
	}

	level := baselevel
	PARAM_CAL := func(v int) int {
		return ((level-1)*eb.LvUpPoint + eb.InitNum) * v
	}
	tp := *eb
	// 这里才是关键, 如果不加这一段, 同类的宠物几乎没有区别, 这初始的三项的变动幅度=总成长的浮动幅度
	// 也就是说总成长一开始就决定了
	tp.BaseVital += rand.Intn(5) - 2
	tp.BaseStr += rand.Intn(5) - 2
	tp.BaseTgh += rand.Intn(5) - 2
	tp.BaseDex += rand.Intn(5) - 2

	char := new(Char)
	char.Name = eb.Name
	char.WhichType = CHAR_TYPEENEMY
	char.Lv = baselevel
	char.AllocPoint = [4]int{tp.BaseVital, tp.BaseStr, tp.BaseTgh, tp.BaseDex}
	for i := 0; i < 10; i++ {
		w := rand.Intn(4)
		switch w {
		case 0:
			tp.BaseVital++
		case 1:
			tp.BaseStr++
		case 2:
			tp.BaseTgh++
		case 3:
			tp.BaseDex++
		}
	}
	char.Vital = PARAM_CAL(tp.BaseVital)
	char.Str = PARAM_CAL(tp.BaseStr)
	char.Tough = PARAM_CAL(tp.BaseTgh)
	char.Dex = PARAM_CAL(tp.BaseDex)
	char.EarthAT = tp.EarthAT
	char.WaterAT = tp.WaterAT
	char.FireAT = tp.FireAT
	char.WindAT = tp.WindAt
	char.ModAI = tp.ModAI
	char.VariableAI = 0
	char.ImgNo = tp.ImgNo
	char.Slot = tp.Slot
	char.Poison = tp.Poison
	char.Paralysis = tp.Paralysis
	char.Sleep = tp.Sleep
	char.Stone = tp.Stone
	char.Drunk = tp.Drunk
	char.Confusion = tp.Confusion
	char.Rare = tp.Rare
	char.PetId = tp.No
	char.Critical = tp.Critical
	char.Counter = tp.Counter
	char.PetSkills = tp.PetSkills
	for _, id := range tp.PetSkillIds {
		char.PetSkillIds = append(char.PetSkillIds, id)
	}
	char.PetRank = Enemy_getRank(ebno)
	InitNewChar(char)
	Char_complianceParameter(char)
	char.BornLv = level
	char.BornPoint = [4]int{char.WorkMaxHp, char.WorkFixStr, char.WorkFixTough, char.WorkFixDex}
	return char
}

// 获取优质程度 越高越好 0-5
func Enemy_getRank(ebno int) int {
	ranktbl := []int{100, 95, 90, 85, 80, 0}
	tp := GetEnemyBase(ebno)
	if tp == nil {
		return 0
	}
	paramsum := tp.BaseVital + tp.BaseStr + tp.BaseTgh + tp.BaseDex
	for i, w := range ranktbl {
		if paramsum >= w {
			return i
		}
	}
	return 0
}
