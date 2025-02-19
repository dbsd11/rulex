package model

import (
	"database/sql/driver"
	"time"

	"gopkg.in/square/go-jose.v2/json"
)

type RulexModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
}
type stringList []string

/*
*
* 给GORM用的
*
 */
func (f stringList) Value() (driver.Value, error) {
	b, err := json.Marshal(f)
	return string(b), err
}

/*
*
* 给GORM用的
*
 */
func (f *stringList) Scan(data interface{}) error {
	return json.Unmarshal([]byte(data.(string)), f)
}

func (f stringList) String() string {
	b, _ := json.Marshal(f)
	return string(b)
}
func (f stringList) Len() int {
	return len(f)
}

type MRule struct {
	RulexModel
	UUID        string     `gorm:"not null"`
	Name        string     `gorm:"not null"`
	Type        string     // 脚本类型，目前支持"lua"和"expr"两种
	FromSource  stringList `gorm:"not null type:string[]"`
	FromDevice  stringList `gorm:"not null type:string[]"`
	Expression  string     `gorm:"not null"` // Expr脚本
	Actions     string     `gorm:"not null"`
	Success     string     `gorm:"not null"`
	Failed      string     `gorm:"not null"`
	Description string
}

type MInEnd struct {
	RulexModel
	// UUID for origin source ID
	UUID      string     `gorm:"not null"`
	Type      string     `gorm:"not null"`
	Name      string     `gorm:"not null"`
	BindRules stringList `json:"bindRules"` // 与之关联的规则表["A","B","C"]

	Description string
	Config      string
	XDataModels string
}

func (md MInEnd) GetConfig() map[string]interface{} {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(md.Config), &result)
	if err != nil {
		return map[string]interface{}{}
	}
	return result
}

type MOutEnd struct {
	RulexModel
	// UUID for origin source ID
	UUID        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Config      string
}

func (md MOutEnd) GetConfig() map[string]interface{} {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(md.Config), &result)
	if err != nil {
		return map[string]interface{}{}
	}
	return result
}

type MUser struct {
	RulexModel
	Role        string `gorm:"not null"`
	Username    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Description string
}

// 设备元数据
type MDevice struct {
	RulexModel
	UUID        string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Config      string
	BindRules   stringList `json:"bindRules"` // 与之关联的规则表["A","B","C"]
	Description string
}

func (md MDevice) GetConfig() map[string]interface{} {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(md.Config), &result)
	if err != nil {
		return map[string]interface{}{}
	}
	return result
}

//
// 外挂
//

type MGoods struct {
	RulexModel
	UUID        string     `gorm:"not null"`
	Addr        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Args        stringList `gorm:"not null"`
}

/*
*
* LUA应用
*
 */
type MApp struct {
	RulexModel
	UUID        string `gorm:"not null"` // 名称
	Name        string `gorm:"not null"` // 名称
	Version     string `gorm:"not null"` // 版本号
	AutoStart   *bool  `gorm:"not null"` // 允许启动
	Filepath    string `gorm:"not null"` // 文件路径, 是相对于main的apps目录
	Description string `gorm:"not null"` // 文件路径, 是相对于main的apps目录
}

/*
*
* 用户上传的AI数据[v0.5.0]
*
 */
type MAiBase struct {
	RulexModel
	UUID        string `gorm:"not null"` // 名称
	Name        string `gorm:"not null"` // 名称
	Type        string `gorm:"not null"` // 类型
	IsBuildIn   bool   `gorm:"not null"` // 是否内建
	Version     string `gorm:"not null"` // 版本号
	Filepath    string `gorm:"not null"` // 文件路径, 是相对于main的apps目录
	Description string `gorm:"not null"`
}

// MModbusPointPosition modbus数据点位
type MModbusPointPosition struct {
	RulexModel
	DeviceUuid   string `json:"deviceUuid"    gorm:"not null"`
	Tag          string `json:"tag"           gorm:"not null"`
	Function     int    `json:"function"      gorm:"not null"`
	SlaverId     byte   `json:"slaverId"      gorm:"not null"`
	StartAddress uint16 `json:"startAddress"  gorm:"not null"`
	Quality      uint16 `json:"quality"       gorm:"not null"`
}
