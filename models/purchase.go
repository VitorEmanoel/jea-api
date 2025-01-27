package models

// Purchase model
type Purchase struct {
	Model
	Seller			*Client				`json:"seller,omitempty" gorm:"foreignkey:ClientID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ClientID		int64				`json:"-"`
	Products		[]*PurchaseProduct	`json:"products,omitempty" gorm:"foreignkey:PurchaseID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	Company			*Company			`json:"company,omitempty" gorm:"foreignkey:CompanyID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	CompanyID		int64				`json:"-"`
	Total			float64				`json:"total"`
	Discount		float64				`json:"discount"`
	Increase		float64				`json:"increase"`
	PaymentType		int					`json:"paymentType"`
	CargoMaps		[]*CargoMap			`json:"cargoMaps,omitempty" gorm:"foreignkey:PurchaseID;association_foreignkey:ID;association_autocreate:true;association_autoupdate:false"`
}

func (p *Purchase) GetFilters() Filters {
	return Filters {
		CompanyFilter,
		ClientFilter,
	}
}

// PurchaseProduct model 
type PurchaseProduct struct  {
	Model
	Purchase		*Purchase	`json:"sale,omitempty" gorm:"foreignkey:PurchaseID;association_foreignkey:ID;association_autoupdate:false;assoaciation_autocreate:false"`
	PurchaseID		int64		`json:"-"`
	Product			*Product	`json:"product,omitempty" gorm:"foreignkey:ProductID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ProductID		int64		`json:"-"`
	Quantity		int64		`json:"quantity"`
	Increase		float64		`json:"increase"`
	Discount		float64		`json:"discount"`

}