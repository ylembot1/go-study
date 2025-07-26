package handler

import (
	"go-tools/swagger/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary      获取用户列表
// @Description  获取所有用户的列表，支持分页
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page     query    int  false  "页码"     default(1)
// @Param        per_page query    int  false  "每页数量" default(10)
// @Success      200      {object} model.UserListResponse
// @Failure      400      {object} map[string]interface{}
// @Failure      500      {object} map[string]interface{}
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	// page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	// 模拟数据
	users := []model.User{
		{ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 25},
		{ID: 2, Name: "李四", Email: "lisi@example.com", Age: 30},
	}

	response := model.UserListResponse{
		Users: users,
		Total: len(users),
	}

	c.JSON(http.StatusOK, response)
}

// GetUser godoc
// @Summary      获取单个用户
// @Description  根据用户ID获取用户详细信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
	// id := c.Param("id")

	// 模拟数据
	user := model.User{
		ID:    1,
		Name:  "张三",
		Email: "zhangsan@example.com",
		Age:   25,
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary      创建用户
// @Description  创建新用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.CreateUserRequest  true  "用户信息"
// @Success      201   {object}  model.User
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟创建用户
	user := model.User{
		ID:    1,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary      更新用户
// @Description  根据用户ID更新用户信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true   "用户ID"
// @Param        user  body      model.UpdateUserRequest true   "用户信息"
// @Success      200   {object}  model.User
// @Failure      400   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Router       /users/{id} [put]
func UpdateUser(c *gin.Context) {
	// id := c.Param("id")
	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟更新用户
	user := model.User{
		ID:    1,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      删除用户
// @Description  根据用户ID删除用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功", "id": id})
}
