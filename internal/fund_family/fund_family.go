package fund

type FundFamily struct {
	FundID            string `json:"FundID,omitempty" gorm:"column:FAMILYID;primary_key" validate:"required,max=20,code"`
	FundDescription   string `json:"FundDescription,omitempty" gorm:"column:FAMILYDESCRIPTION;primary_key" validate:"required,max=20,code"`
	FundLevel         string `json:"FundLevel,omitempty" gorm:"column:FUNDLEVEL;primary_key" validate:"required,max=20,code"`
	ImmediateParentId string `json:"ImmediateParentId,omitempty" gorm:"column:IMMEDIATEPARENTID;primary_key" validate:"required,max=20,code"`
	PortFolioType     int8   `json:"PortFolioType,omitempty" gorm:"column:PORTFOLIOTYPE;primary_key" validate:"required,max=20,code"`
	UmbrellaType      int8   `json:"UmbrellaType,omitempty" gorm:"column:UMBRELLATYPE;primary_key" validate:"required,max=20,code"`
}
