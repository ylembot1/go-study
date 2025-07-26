package examples

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 示例数据模型
type Product struct {
	ID          int     `json:"id" example:"1"`
	Name        string  `json:"name" example:"iPhone 15"`
	Price       float64 `json:"price" example:"999.99"`
	Description string  `json:"description" example:"最新款iPhone"`
	Category    string  `json:"category" example:"电子产品"`
	InStock     bool    `json:"in_stock" example:"true"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"iPhone 15"`
	Price       float64 `json:"price" binding:"required,min=0" example:"999.99"`
	Description string  `json:"description" example:"最新款iPhone"`
	Category    string  `json:"category" example:"电子产品"`
	InStock     bool    `json:"in_stock" example:"true"`
}

type ProductListResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total" example:"100"`
	Page     int       `json:"page" example:"1"`
	PerPage  int       `json:"per_page" example:"10"`
}

type ErrorResponse struct {
	Error   string `json:"error" example:"错误信息"`
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"详细错误描述"`
}

// 示例 1: 基本的 GET 请求
// GetProducts godoc
// @Summary      获取产品列表
// @Description  获取所有产品的列表，支持分页和搜索
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "页码"     default(1)
// @Param        per_page  query    int     false  "每页数量" default(10)
// @Param        category  query    string  false  "产品分类"
// @Param        search    query    string  false  "搜索关键词"
// @Success      200       {object} ProductListResponse
// @Failure      400       {object} ErrorResponse
// @Failure      500       {object} ErrorResponse
// @Router       /products [get]
func GetProducts(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	_ = c.Query("category") // 忽略未使用的变量
	_ = c.Query("search")   // 忽略未使用的变量

	// 模拟数据
	products := []Product{
		{ID: 1, Name: "iPhone 15", Price: 999.99, Category: "电子产品", InStock: true},
		{ID: 2, Name: "MacBook Pro", Price: 1999.99, Category: "电子产品", InStock: true},
	}

	response := ProductListResponse{
		Products: products,
		Total:    len(products),
		Page:     page,
		PerPage:  perPage,
	}

	c.JSON(http.StatusOK, response)
}

// 示例 2: 带路径参数的 GET 请求
// GetProduct godoc
// @Summary      获取单个产品
// @Description  根据产品ID获取产品详细信息
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "产品ID"
// @Success      200  {object}  Product
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /products/{id} [get]
func GetProduct(c *gin.Context) {
	_ = c.Param("id") // 忽略未使用的变量

	// 模拟数据
	product := Product{
		ID:          1,
		Name:        "iPhone 15",
		Price:       999.99,
		Description: "最新款iPhone，搭载A17 Pro芯片",
		Category:    "电子产品",
		InStock:     true,
	}

	c.JSON(http.StatusOK, product)
}

// 示例 3: POST 请求创建资源
// CreateProduct godoc
// @Summary      创建产品
// @Description  创建新产品
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      CreateProductRequest  true  "产品信息"
// @Success      201      {object}  Product
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /products [post]
func CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 模拟创建产品
	product := Product{
		ID:          1,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
		InStock:     req.InStock,
	}

	c.JSON(http.StatusCreated, product)
}

// 示例 4: PUT 请求更新资源
// UpdateProduct godoc
// @Summary      更新产品
// @Description  根据产品ID更新产品信息
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int                   true   "产品ID"
// @Param        product  body      CreateProductRequest  true   "产品信息"
// @Success      200      {object}  Product
// @Failure      400      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	_ = c.Param("id") // 忽略未使用的变量
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 模拟更新产品
	product := Product{
		ID:          1,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
		InStock:     req.InStock,
	}

	c.JSON(http.StatusOK, product)
}

// 示例 5: DELETE 请求删除资源
// DeleteProduct godoc
// @Summary      删除产品
// @Description  根据产品ID删除产品
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "产品ID"
// @Success      204  "No Content"
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	_ = c.Param("id") // 忽略未使用的变量

	// 模拟删除产品
	c.Status(http.StatusNoContent)
}

// 示例 6: 带认证的请求
// GetUserProfile godoc
// @Summary      获取用户资料
// @Description  获取当前登录用户的资料信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/profile [get]
func GetUserProfile(c *gin.Context) {
	// 模拟用户资料
	profile := map[string]interface{}{
		"id":       1,
		"username": "john_doe",
		"email":    "john@example.com",
		"name":     "John Doe",
		"avatar":   "https://example.com/avatar.jpg",
	}

	c.JSON(http.StatusOK, profile)
}

// 示例 7: 文件上传
// UploadFile godoc
// @Summary      上传文件
// @Description  上传产品图片文件
// @Tags         files
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "文件"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /upload [post]
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "FILE_UPLOAD_ERROR",
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 模拟文件上传
	response := map[string]interface{}{
		"filename": file.Filename,
		"size":     file.Size,
		"url":      "https://example.com/uploads/" + file.Filename,
	}

	c.JSON(http.StatusOK, response)
}

// 示例 8: 带查询参数的复杂请求
// SearchProducts godoc
// @Summary      搜索产品
// @Description  根据多个条件搜索产品
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        q         query    string   false  "搜索关键词"
// @Param        category  query    []string false  "产品分类" collectionFormat(multi)
// @Param        min_price query    float64  false  "最低价格"
// @Param        max_price query    float64  false  "最高价格"
// @Param        in_stock  query    bool     false  "是否有库存"
// @Param        sort      query    string   false  "排序方式" Enums(price_asc, price_desc, name_asc, name_desc)
// @Success      200       {object} ProductListResponse
// @Failure      400       {object} ErrorResponse
// @Router       /products/search [get]
func SearchProducts(c *gin.Context) {
	// 获取查询参数（示例中未使用，实际项目中会使用这些参数）
	_ = c.Query("q")             // 忽略未使用的变量
	_ = c.QueryArray("category") // 忽略未使用的变量
	_ = c.Query("min_price")     // 忽略未使用的变量
	_ = c.Query("max_price")     // 忽略未使用的变量
	_ = c.Query("in_stock")      // 忽略未使用的变量
	_ = c.Query("sort")          // 忽略未使用的变量

	// 模拟搜索结果
	products := []Product{
		{ID: 1, Name: "iPhone 15", Price: 999.99, Category: "电子产品", InStock: true},
	}

	response := ProductListResponse{
		Products: products,
		Total:    len(products),
		Page:     1,
		PerPage:  10,
	}

	c.JSON(http.StatusOK, response)
}
