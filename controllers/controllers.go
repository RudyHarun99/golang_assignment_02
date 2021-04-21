package controllers

import (
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"GO14/assignment02/models"
	"gorm.io/gorm"
)

func GetOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var orders []models.Orders
	db.Debug().Preload("Items").Find(&orders)

	c.JSON(200, gin.H{
		"orders": orders,
	})
}

func CreateOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	customer_name := c.PostForm("Customer_name")
	ordered_at := time.Now()
	item_code := c.PostForm("Item_code")
	description := c.PostForm("Description")
	quantity, _ := strconv.Atoi(c.PostForm("Quantity"))

	newItem := models.Item{
		Item_code: item_code,
		Description: description,
		Quantity: quantity,
	}

	data := models.Orders{
		Customer_name: customer_name,
		Ordered_at: ordered_at,
		Items: []models.Item{newItem},
	}

	db.Debug().Create(&data)

	c.JSON(200, gin.H{
		"data": data,
	})
}

func EditOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	customer_name := c.PostForm("Customer_name")
	ordered_at := time.Now()
	item_code := c.PostForm("Item_code")
	description := c.PostForm("Description")
	quantity, _ := strconv.Atoi(c.PostForm("Quantity"))

	newItem := models.Item{
		Item_code: item_code,
		Description: description,
		Quantity: quantity,
	}

	data := models.Orders{
		Customer_name: customer_name,
		Ordered_at: ordered_at,
		Items: []models.Item{newItem},
	}

	var order models.Orders
	db.Debug().Model(&order).Where("id = ?", id).Updates(&data)

	var item models.Item
	db.Debug().Model(&item).Where("order_id = ?", id).Updates(&newItem)

	db.Debug().Preload("Items").Where("id = ?", id).Find(&order)

	c.JSON(200, gin.H{
		"message": order,
	})
}

func DeleteOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")

	var item models.Item
	db.Debug().Where("order_id = ?", id).Delete(&item)

	var order models.Orders
	db.Debug().Where("id = ?", id).Delete(&order)

	c.JSON(200, gin.H{
		"message": id,
	})
}