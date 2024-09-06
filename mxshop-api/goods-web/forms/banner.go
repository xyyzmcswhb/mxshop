package forms

type BannerForm struct {
	Image string `form:"image" json:"image" binding:"url"` //图片url
	Index int    `form:"index" json:"index" binding:"required"`
	Url   string `form:"url" json:"url" binding:"url"` //跳转商品的url
}
