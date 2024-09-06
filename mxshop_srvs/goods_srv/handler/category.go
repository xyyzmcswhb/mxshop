package handler

import (
	"context"
	"encoding/json"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

// 商品分类
func (s *GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {
	/*
		{
			"id":xxx,
			"name":"",
			"level":1,
			"is_tab":false,
			"parent":
			"sub_category":[
			]
		}
	*/
	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

// 查询子类目
func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListRsp := proto.SubCategoryListResponse{}
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		//商品不存在
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	categoryListRsp.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategories []model.Category
	var subCategoriesRsp []*proto.CategoryInfoResponse
	//preload := "SubCategory"
	//if category.Level == 1 {
	//	preload = "SubCategory.SubCategory"
	//}
	//global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(preload).Find(&subCategories)
	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Find(&subCategories)
	for _, subCategory := range subCategories { //转换响应
		subCategoriesRsp = append(subCategoriesRsp, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListRsp.SubCategorys = subCategoriesRsp
	return &categoryListRsp, nil
}

//	func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
//		category := model.Category{}
//		cMap := map[string]interface{}{}
//		cMap["name"] = req.Name
//		cMap["level"] = req.Level
//		cMap["is_tab"] = req.IsTab
//		if req.Level != 1 {
//			//去查询父类目是否存在
//			cMap["parent_category_id"] = req.ParentCategory
//		}
//		tx := global.DB.Model(&model.Category{}).Create(cMap)
//		fmt.Println(tx)
//		return &proto.CategoryInfoResponse{Id: category.ID}, nil
//	}
func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{
		Name:  req.Name,
		Level: req.Level,
		IsTab: req.IsTab,
	}

	if req.Level != 1 {
		// 查询父类目是否存在
		category.ParentCategoryID = int32(req.ParentCategory)
	}

	// 使用模型结构体进行创建
	if err := global.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return &proto.CategoryInfoResponse{Id: category.ID}, nil
}

func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
