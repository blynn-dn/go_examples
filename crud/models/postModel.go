package models

type TestTable struct {
	ID   int64 `gorm:"primaryKey"`
	Name string
}

// TableName by default gorm will pluralize the table's name but in this case the table name is `test_table`
func (TestTable) TableName() string {
	return "test_table"
}
