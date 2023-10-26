package manager

import (
	"fmt"
	"mnc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", i.cfg.User, i.cfg.Password, i.cfg.Host, i.cfg.Port, i.cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
