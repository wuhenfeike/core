package relationDB

import (
	"context"
	"gitee.com/unitedrhino/share/stores"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
这个是参考样例
使用教程:
1. 将example全局替换为模型的表名
2. 完善todo
*/

type ExampleRepo struct {
	db *gorm.DB
}

func NewExampleRepo(in any) *ExampleRepo {
	return &ExampleRepo{db: stores.GetCommonConn(in)}
}

type ExampleFilter struct {
	//todo 添加过滤字段
}

func (p ExampleRepo) fmtFilter(ctx context.Context, f ExampleFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	//todo 添加条件
	return db
}

func (p ExampleRepo) Insert(ctx context.Context, data *SysExample) error {
	result := p.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (p ExampleRepo) FindOneByFilter(ctx context.Context, f ExampleFilter) (*SysExample, error) {
	var result SysExample
	db := p.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}
func (p ExampleRepo) FindByFilter(ctx context.Context, f ExampleFilter, page *stores.PageInfo) ([]*SysExample, error) {
	var results []*SysExample
	db := p.fmtFilter(ctx, f).Model(&SysExample{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p ExampleRepo) CountByFilter(ctx context.Context, f ExampleFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&SysExample{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}

func (p ExampleRepo) Update(ctx context.Context, data *SysExample) error {
	err := p.db.WithContext(ctx).Where("id = ?", data.ID).Save(data).Error
	return stores.ErrFmt(err)
}

func (p ExampleRepo) DeleteByFilter(ctx context.Context, f ExampleFilter) error {
	db := p.fmtFilter(ctx, f)
	err := db.Delete(&SysExample{}).Error
	return stores.ErrFmt(err)
}

func (p ExampleRepo) Delete(ctx context.Context, id int64) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&SysExample{}).Error
	return stores.ErrFmt(err)
}
func (p ExampleRepo) FindOne(ctx context.Context, id int64) (*SysExample, error) {
	var result SysExample
	err := p.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

// 批量插入 LightStrategyDevice 记录
func (p ExampleRepo) MultiInsert(ctx context.Context, data []*SysExample) error {
	err := p.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Model(&SysExample{}).Create(data).Error
	return stores.ErrFmt(err)
}

func (d ExampleRepo) UpdateWithField(ctx context.Context, f ExampleFilter, updates map[string]any) error {
	db := d.fmtFilter(ctx, f)
	err := db.Model(&SysExample{}).Updates(updates).Error
	return stores.ErrFmt(err)
}
