package model

type SsRegion struct {
	ID         int64   `gorm:"column:id" json:"id"`
	Name       string  `gorm:"column:name" json:"name"`
	ParentId   int64   `gorm:"column:parent_id" json:"parentId"`
	ShortName  string  `gorm:"column:short_name" json:"shortName"`
	LevelType  int64   `gorm:"column:level_type" json:"levelType"`
	CityCode   string  `gorm:"column:city_code" json:"cityCode"`
	PostCode   string  `gorm:"column:post_code" json:"postCode"`
	MergerName string  `gorm:"column:merger_name" json:"mergerName"`
	Lng        float64 `gorm:"column:lng" json:"lng"`
	Lat        float64 `gorm:"column:lat" json:"lat"`
	Pinyin     string  `gorm:"column:pinyin" json:"pinyin"`
}

func (region SsRegion) TableName() string {
	return "ss_region"
}
