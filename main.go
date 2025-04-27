package main

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

var node *snowflake.Node

func init() {
	var err error
	// 初始化一个 Node (你可以根据不同服务实例，设置不同的node number)
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

// Product模型
type Product struct {
	Code      string `gorm:"primaryKey"`
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 插入前自动生成Code
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Code == "" {
		p.Code = node.Generate().String()
	}
	return
}

func main() {
	// 读取配置
	cfg := loadConfig("config/db/db.yaml")

	// URL编码密码
	//encodedPassword := url.QueryEscape(cfg.Database.Password)
	//fmt.Println(encodedPassword)

	// 拼接DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	// 设置打印日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second * 10,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	// 设置全局的logger,这个logger在我们执行每个sql语句时都会打印出来

	fmt.Println("数据库连接成功！", db)
	// 定义表结构，将表结构直接生成对应的表 自动建表
	db.AutoMigrate(&Product{})
	// 创建一条记录测试
	product := Product{Price: 100}
	result := db.Create(&product)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Product Created:", product)
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("failed to read config file:", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal("failed to parse config file:", err)
	}

	return cfg
}
