package entities

type Category struct {
	Id          uint   `gorm:"column:id;type:BIGSERIAL;primaryKey"`
	Name        string `gorm:"column:name;type:VARCHAR;not null"`
	Description string `gorm:"column:description;type:VARCHAR;not null"`
}
