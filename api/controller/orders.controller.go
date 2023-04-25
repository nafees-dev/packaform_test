package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nafees-dev/packaform_test/entity"
	"gorm.io/gorm"
)

type iOrderController interface {
	GetAllOrders() []entity.Order
}

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(DB *gorm.DB) OrderController {
	return OrderController{DB}
}

func (pc *OrderController) GetAllOrders(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var fromDate, toDate time.Time
	createdDate := ctx.DefaultQuery("createdAt", "")
	searchStr := ctx.DefaultQuery("search", "")

	var orders []entity.Order
	var results *gorm.DB

	if createdDate != "" {
		dateRange := make([]string, 2)
		err := json.Unmarshal([]byte(createdDate), &dateRange)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid date range format"})
			return
		}
		fromDate, err = time.Parse(time.RFC3339, dateRange[0])
		toDate, err = time.Parse(time.RFC3339, dateRange[1])
		fmt.Printf("fromDate: %s, toDate: %s\n", fromDate, toDate)

		fromDateString := fromDate.Format("2006-01-02 15:04:05")
		toDateString := toDate.Format("2006-01-02 15:04:05")

		results = pc.DB.Preload("Customers").Preload("Customers.Customer_Companies").Preload("OrderItems").Preload("OrderItems.Deliveries").
			Where("created_at BETWEEN ? AND ?", fromDateString, toDateString).Limit(intLimit).Offset(offset).Find(&orders)
	} else if searchStr != "" {
		searchStr = strings.ToLower(searchStr)
		results = pc.DB.
			Preload("Customers").
			Preload("Customers.Customer_Companies").
			// Joins("JOIN order_items ON order_items.order_id = orders.id").
			Where("orders.id IN (SELECT order_id FROM order_items WHERE LOWER(product) LIKE ?) OR LOWER(order_name) LIKE ?", "%"+searchStr+"%", "%"+searchStr+"%").
			// Preload("OrderItems", "LOWER(order_items.product) LIKE ?", "%"+searchStr+"%").
			// Where("LOWER(orders.order_name) LIKE ?", "%"+searchStr+"%").
			Preload("OrderItems.Deliveries").
			Find(&orders)

	} else {
		results = pc.DB.Preload("Customers").Preload("Customers.Customer_Companies").Preload("OrderItems").Preload("OrderItems.Deliveries").Limit(intLimit).Offset(offset).Find(&orders)
	}

	// search := ctx.DefaultQuery("search", "")
	var count int64
	pc.DB.Table("orders").Count(&count)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(orders), "data": orders, "total": count})
}
