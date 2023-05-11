package infra

import (
	"context"
	"database/sql"
	"github.com/junyaU/mimi/testdata/layer/adapter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type db struct {
	conn *gorm.DB
}

var txKey = struct{}{}

func NewDB() adapter.DataHandler {
	dsn := "dummy:dummy@tcp(meacle_database:3306)/meacle?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return &db{
		conn: d,
	}
}

func (d *db) Create(ctx context.Context, value interface{}) error {
	tx := d.conn.WithContext(ctx)
	if err := tx.Create(value).Error; err != nil {
		return err
	}

	return nil
}

func (d *db) Update(ctx context.Context, value interface{}) error {
	tx := d.conn.WithContext(ctx)
	if err := tx.Save(value).Error; err != nil {
		return err
	}

	return nil
}

func (d *db) Delete(ctx context.Context, value interface{}) error {
	tx := d.conn.WithContext(ctx)
	if err := tx.Delete(value).Error; err != nil {
		return err
	}

	return nil
}

func (d *db) Query(value interface{}, sql string, params ...interface{}) error {
	if err := d.conn.Raw(sql, params...).Scan(value).Error; err != nil {
		return err
	}

	return nil
}

func (d *db) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	db, err := d.conn.DB()
	if err != nil {
		return nil, err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	v, err := f(context.WithValue(ctx, &txKey, tx))
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return v, nil
}
