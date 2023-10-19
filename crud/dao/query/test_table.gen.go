// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/blynn-dn/go_examples/crud/dao/model"
)

func newTestTable(db *gorm.DB, opts ...gen.DOOption) testTable {
	_testTable := testTable{}

	_testTable.testTableDo.UseDB(db, opts...)
	_testTable.testTableDo.UseModel(&model.TestTable{})

	tableName := _testTable.testTableDo.TableName()
	_testTable.ALL = field.NewAsterisk(tableName)
	_testTable.ID = field.NewInt32(tableName, "id")
	_testTable.Name = field.NewString(tableName, "name")

	_testTable.fillFieldMap()

	return _testTable
}

type testTable struct {
	testTableDo testTableDo

	ALL  field.Asterisk
	ID   field.Int32
	Name field.String

	fieldMap map[string]field.Expr
}

func (t testTable) Table(newTableName string) *testTable {
	t.testTableDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t testTable) As(alias string) *testTable {
	t.testTableDo.DO = *(t.testTableDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *testTable) updateTableName(table string) *testTable {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Name = field.NewString(table, "name")

	t.fillFieldMap()

	return t
}

func (t *testTable) WithContext(ctx context.Context) *testTableDo {
	return t.testTableDo.WithContext(ctx)
}

func (t testTable) TableName() string { return t.testTableDo.TableName() }

func (t testTable) Alias() string { return t.testTableDo.Alias() }

func (t testTable) Columns(cols ...field.Expr) gen.Columns { return t.testTableDo.Columns(cols...) }

func (t *testTable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *testTable) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
}

func (t testTable) clone(db *gorm.DB) testTable {
	t.testTableDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t testTable) replaceDB(db *gorm.DB) testTable {
	t.testTableDo.ReplaceDB(db)
	return t
}

type testTableDo struct{ gen.DO }

func (t testTableDo) Debug() *testTableDo {
	return t.withDO(t.DO.Debug())
}

func (t testTableDo) WithContext(ctx context.Context) *testTableDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t testTableDo) ReadDB() *testTableDo {
	return t.Clauses(dbresolver.Read)
}

func (t testTableDo) WriteDB() *testTableDo {
	return t.Clauses(dbresolver.Write)
}

func (t testTableDo) Session(config *gorm.Session) *testTableDo {
	return t.withDO(t.DO.Session(config))
}

func (t testTableDo) Clauses(conds ...clause.Expression) *testTableDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t testTableDo) Returning(value interface{}, columns ...string) *testTableDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t testTableDo) Not(conds ...gen.Condition) *testTableDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t testTableDo) Or(conds ...gen.Condition) *testTableDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t testTableDo) Select(conds ...field.Expr) *testTableDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t testTableDo) Where(conds ...gen.Condition) *testTableDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t testTableDo) Order(conds ...field.Expr) *testTableDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t testTableDo) Distinct(cols ...field.Expr) *testTableDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t testTableDo) Omit(cols ...field.Expr) *testTableDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t testTableDo) Join(table schema.Tabler, on ...field.Expr) *testTableDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t testTableDo) LeftJoin(table schema.Tabler, on ...field.Expr) *testTableDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t testTableDo) RightJoin(table schema.Tabler, on ...field.Expr) *testTableDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t testTableDo) Group(cols ...field.Expr) *testTableDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t testTableDo) Having(conds ...gen.Condition) *testTableDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t testTableDo) Limit(limit int) *testTableDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t testTableDo) Offset(offset int) *testTableDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t testTableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *testTableDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t testTableDo) Unscoped() *testTableDo {
	return t.withDO(t.DO.Unscoped())
}

func (t testTableDo) Create(values ...*model.TestTable) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t testTableDo) CreateInBatches(values []*model.TestTable, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t testTableDo) Save(values ...*model.TestTable) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t testTableDo) First() (*model.TestTable, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestTable), nil
	}
}

func (t testTableDo) Take() (*model.TestTable, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestTable), nil
	}
}

func (t testTableDo) Last() (*model.TestTable, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestTable), nil
	}
}

func (t testTableDo) Find() ([]*model.TestTable, error) {
	result, err := t.DO.Find()
	return result.([]*model.TestTable), err
}

func (t testTableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TestTable, err error) {
	buf := make([]*model.TestTable, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t testTableDo) FindInBatches(result *[]*model.TestTable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t testTableDo) Attrs(attrs ...field.AssignExpr) *testTableDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t testTableDo) Assign(attrs ...field.AssignExpr) *testTableDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t testTableDo) Joins(fields ...field.RelationField) *testTableDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t testTableDo) Preload(fields ...field.RelationField) *testTableDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t testTableDo) FirstOrInit() (*model.TestTable, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestTable), nil
	}
}

func (t testTableDo) FirstOrCreate() (*model.TestTable, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestTable), nil
	}
}

func (t testTableDo) FindByPage(offset int, limit int) (result []*model.TestTable, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t testTableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t testTableDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t testTableDo) Delete(models ...*model.TestTable) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *testTableDo) withDO(do gen.Dao) *testTableDo {
	t.DO = *do.(*gen.DO)
	return t
}