package gormate

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NamedTable interface {
	TableName() string
	PK() uint64
}

type Operator struct{}

func (o *Operator) db(ctx context.Context) (db *gorm.DB, err error) {
	db, err = FromContext(ctx)
	if err != nil {
		return
	}
	db = db.Omit(clause.Associations)
	return
}

func (o *Operator) Create(ctx context.Context, i NamedTable) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Create(i).Error
	return err
}

func (o *Operator) CreateAll(ctx context.Context, i interface{}) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Create(i).Error
	return err
}

func (o *Operator) GetByID(ctx context.Context, i NamedTable, id uint64) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Preload(clause.Associations).Where("id = ?", id).First(i).Error
	return err
}

func (o *Operator) GetByMultiID(ctx context.Context, i interface{}, ids []uint64) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Preload(clause.Associations).Where("id IN (?)", ids).Find(i).Error
	return err
}

func (o *Operator) GetByCondition(ctx context.Context, i, cond interface{}, args ...interface{}) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Preload(clause.Associations).Where(cond, args...).Find(i).Error
	return err
}

func (o *Operator) Save(ctx context.Context, i NamedTable, selected ...string) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	db = db.Model(i)
	if len(selected) > 0 {
		v := make([]interface{}, len(selected))
		for index := range selected {
			v[index] = selected[index]
		}
		db = db.Select(v[0], v[1:]...)
	}
	err = db.Save(i).Error
	return err
}

func (o *Operator) Updates(ctx context.Context, i NamedTable, values map[string]interface{}) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Where("id = ?", i.PK()).Updates(values).Error
	return err
}

func (o *Operator) UpdatesAll(ctx context.Context, i NamedTable, values map[string]interface{}, cond ...interface{}) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Where(cond, cond[1:]).Updates(values).Error
	return err
}

func (o *Operator) Remove(ctx context.Context, i NamedTable) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Delete(i, "id = ?", i.PK()).Error
	return err
}

func (o *Operator) RemoveAll(ctx context.Context, i NamedTable, cond ...interface{}) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Delete(i, cond...).Error
	return err
}
