// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTestTable = "test_table"

// TestTable mapped from table <test_table>
type TestTable struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

// TableName TestTable's table name
func (*TestTable) TableName() string {
	return TableNameTestTable
}
