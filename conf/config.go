package conf

import (
	"context"
	"fmt"
	"github.com/xie392/restful-api/apps/host"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

// 全局config实例对象,
// 也就是我们程序，在内存中的配置对象
// 程序内部获取配置, 都通过读取该对象
// 该Config对象 什么时候被初始化?
//
//	配置加载时:
//	   LoadConfigFromToml
//	   LoadConfigFromEnv
//
// 为了不被程序在运行时恶意修改, 设置成私有变量
var config *Config

// 全局MySQL 客户端实例
var db *gorm.DB

// Config 应用配置
// 通过封装为一个对象, 来与外部配置进行对接
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

// App 应用配置
// 包含应用名称, 监听地址, 端口等信息
type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

// Log 日志配置
// 包含日志级别, 格式, 输出方式, 日志文件路径等信息
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的MySQL连接池, 需要池做一些规划配置
	// 控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制MySQL复用, 比如5, 最多运行5个来复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 这个和MySQL Server配置有关系, 必须小于Server配置
	// 一个连接用12h 换一个conn, 保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle 连接 最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`

	// 作为私有变量, 用户与控制GetDB
	lock sync.Mutex
}

func NewDefaultApp() *App {
	return &App{
		Name: "host",
		Host: "0.0.0.0",
		Port: "8050",
	}
}

func NewDefaultLog() *Log {
	return &Log{
		// 日志级别: debug | info | error | warn
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

func NewDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Password:    "123456",
		Database:    "host",
		MaxOpenConn: 10,
		MaxIdleConn: 5,
		MaxLifeTime: 12 * 3600,
		MaxIdleTime: 10 * 60,
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMysql(),
	}
}

// getDBConn 获取 MySQL 连接
func (m *MySQL) getDBConn() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.UserName,
		m.Password,
		m.Host,
		m.Port,
		m.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 为了更好的性能，启用预编译语句
	})

	err = host.AutoMigrateResource(db)
	err = host.AutoMigrateDescribe(db)

	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying SQL DB error, %s", err.Error())
	}

	// 设置连接池配置
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)
	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}

	return db, nil
}

// C 要想获取配置, 单独提供函数。全局Config对象获取函数
func C() *Config {
	return config
}

// GetDB 获取 MySQL 客户端实例
func (m *MySQL) GetDB() *gorm.DB {
	if db == nil {
		// 调用 getDBConn 方法获取数据库连接
		conn, err := m.getDBConn()
		if err != nil {
			// 如果获取连接出错，则抛出 panic
			panic(err)
		}
		// 将获取到的连接赋值给全局变量 db
		db = conn
	}
	return db
}
