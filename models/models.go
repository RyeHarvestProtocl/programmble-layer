package models

import (
	"context"
	"time"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/RyeHarvestProtocol/programmable-layer/pkg/database/postgresql"
	"gorm.io/gorm"
)

// BaseModel contains common columns for all tables.
type BaseModel struct {
	ID        uint `gorm:"primarykey" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type IRepository[T any] interface {
	Create(item *T) error
	GetAll() ([]T, error)
	GetById(id uint) (*T, error)
	Find(conds ...interface{}) ([]T, error)
	// First(conds ...interface{}) (*T, error)
	Update(conds interface{}, updates interface{}) error
	CreateAndGetID(item *T) (uint, error)
}

type DB struct {
	Pg *gorm.DB
}

var DBInstance *DB

func New(conf *config.Config) *DB {
	// logger.Println("===new db")
	var connectString = ""
	if conf.Postgresql.Url != "" {
		connectString = conf.Postgresql.Url
	} else {
		connectString = conf.Postgresql.String()
	}
	postgres, err := postgresql.Dial(connectString)
	if err != nil {
		panic(err)
	}

	d := &DB{
		Pg: postgres,
	}

	// create tables if not exists
	d.Pg.AutoMigrate(&BTCNetworkInfo{})
	d.Pg.AutoMigrate(&MultisigBTCFunding{})
	d.Pg.AutoMigrate(&MultisigRuneFunding{})

	return d
}

func (d *DB) Ping(ctx context.Context) error {
	PgDB, err := d.Pg.DB()
	if err != nil {
		return err
	}
	if err := PgDB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (d *DB) Close(ctx context.Context) error {
	Pg, err := d.Pg.DB()
	if err != nil {
		return err
	}
	if err := Pg.Close(); err != nil {
		return err
	}

	return nil
}

func (d *DB) Begin(ctx context.Context) (*gorm.DB, error) {
	return d.Pg.WithContext(ctx).Begin(), nil
}

func (d *DB) TruncateTable(tableName string) error {
	tx := d.Pg.Exec("TRUNCATE TABLE " + tableName)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
