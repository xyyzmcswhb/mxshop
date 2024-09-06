package forms

type BrandsForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	Logo string `form:"logo" json:"logo" binding:"required"`
}

type CategoryBrandsForm struct {
	CategoryId int `form:"category_id" json:"category_id" binding:"required"`
	BrandId    int `form:"brand_id" json:"brand_id" binding:"required"`
}
