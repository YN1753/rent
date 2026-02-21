package api

import (
	"rent/internal/model"
	"rent/internal/service"
	"rent/pkg/common"
	"rent/pkg/utils/location"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HouseHandler struct {
	HouseService *service.HouseService
}

func NewHouseHandler(mongo *mongo.Database) *HouseHandler {
	return &HouseHandler{HouseService: service.NewHouseService(mongo)}
}
func (h *HouseHandler) Create(c *gin.Context) {
	var house model.House
	err := c.ShouldBindJSON(&house)
	if err != nil {
		common.Error(c, 400, "参数错误")
	}
	house.CreatedAt = time.Now()
	id := c.GetInt("id")
	house.Owner = id
	house.Id = primitive.NewObjectID()
	house.SellStatus = "在售"
	err = h.HouseService.Create(house)
	if err != nil {
		common.Error(c, 400, err.Error())
	}
	common.Success(c, 200, "创建成功", nil)
}

func (h *HouseHandler) GetInputTips(c *gin.Context) {
	var tips model.LocationTipsRes
	var req model.LocationTipReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "参数错误")
	}
	tips, err := location.GetLocationTips(req.Keywords)
	if err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	common.Success(c, 200, "获取成功", tips)
}

func (h *HouseHandler) GetLocationByPOI(c *gin.Context) {
	var req model.LocationPOIReq
	var res []model.LocationPOIRes
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "参数错误")
	}
	res, err := location.GetLocationByPOI(req.Keywords)
	if err != nil {
		common.Error(c, 400, "获取地址失败")
		return
	}
	common.Success(c, 200, "获取成功", res)
}
