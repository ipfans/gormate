package gormate

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NamedTable for name and PK.
type NamedTable interface {
	TableName() string
	PK() uint64
}

// Creator interface.
type Creator interface {
	Create(ctx context.Context, i NamedTable) error
	CreateAll(ctx context.Context, i interface{}) error
}

// Fetcher interface.
type Fetcher interface {
	GetByID(ctx context.Context, i NamedTable, id uint64) error
	GetByMultiID(ctx context.Context, i interface{}, ids []uint64) error
	GetByCondition(ctx context.Context, i, cond interface{}, args ...interface{}) error
}

// Updater interface.
type Updater interface {
	Save(ctx context.Context, i NamedTable, selected ...string) error
	Updates(ctx context.Context, i NamedTable, values map[string]interface{}) error
	UpdatesAll(ctx context.Context, i NamedTable, values map[string]interface{}, cond ...interface{}) error
}

// Remover interface.
type Remover interface {
	Remove(ctx context.Context, i NamedTable) error
	RemoveAll(ctx context.Context, i NamedTable, cond ...interface{}) error
}

// Operator collections.
type Operator struct{}

func (o *Operator) db(ctx context.Context) (db *gorm.DB, err error) {
	db, err = FromContext(ctx)
	if err != nil {
		return
	}
	db = db.Omit(clause.Associations)
	return
}

// Create item.
func (o *Operator) Create(ctx context.Context, i NamedTable) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Create(i).Error
	return err
}

// CreateAll items.
func (o *Operator) CreateAll(ctx context.Context, i interface{}) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Create(i).Error
	return err
}

// GetByID query item by id.
func (o *Operator) GetByID(ctx context.Context, i NamedTable, id uint64) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Preload(clause.Associations).Where("id = ?", id).First(i).Error
	return err
}

// GetByMultiID query items by id.
func (o *Operator) GetByMultiID(ctx context.Context, i interface{}, ids []uint64) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Preload(clause.Associations).Where("id IN (?)", ids).Find(i).Error
	return err
}

// GetByCondition query items by conditions.
func (o *Operator) GetByCondition(ctx context.Context, i, cond interface{}, args ...interface{}) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Preload(clause.Associations).Where(cond, args...).Find(i).Error
	return err
}

// Save item.
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

// Updates given fields.
func (o *Operator) Updates(ctx context.Context, i NamedTable, values map[string]interface{}) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Where("id = ?", i.PK()).Updates(values).Error
	return err
}

// UpdatesAll items.
func (o *Operator) UpdatesAll(ctx context.Context, i NamedTable, values map[string]interface{}, cond ...interface{}) error {
	db, err := FromContext(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Where(cond, cond[1:]).Updates(values).Error
	return err
}

// Remove item.
func (o *Operator) Remove(ctx context.Context, i NamedTable) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Delete(i, "id = ?", i.PK()).Error
	return err
}

// RemoveAll items.
func (o *Operator) RemoveAll(ctx context.Context, i NamedTable, cond ...interface{}) error {
	db, err := o.db(ctx)
	if err != nil {
		return err
	}
	err = db.Table(i.TableName()).Delete(i, cond...).Error
	return err
}
