package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"stocky/db"
	"stocky/models"
	"stocky/services"
)

func PostReward(c *gin.Context) {
	var reward models.Reward
	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reward.Timestamp = time.Now()

	_, err := db.DB.Exec(`INSERT INTO rewards (user_id, stock, shares, timestamp) VALUES ($1, $2, $3, $4)`,
		reward.UserID, reward.Stock, reward.Shares, reward.Timestamp)
	if err != nil {
		logrus.Error("Error inserting reward: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reward recorded successfully!"})
}

func GetStats(c *gin.Context) {
	userId := c.Param("userId")

	rows, err := db.DB.Query(`SELECT stock, SUM(shares) FROM rewards WHERE user_id=$1 AND DATE(timestamp)=CURRENT_DATE GROUP BY stock`, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats"})
		return
	}
	defer rows.Close()

	type StockSummary struct {
		Stock  string  `json:"stock"`
		Shares float64 `json:"shares"`
	}

	var totalValue float64
	var stocks []StockSummary

	for rows.Next() {
		var s StockSummary
		rows.Scan(&s.Stock, &s.Shares)
		stocks = append(stocks, s)
		price := services.GetPrice(s.Stock)
		totalValue += price * s.Shares
	}

	c.JSON(http.StatusOK, gin.H{
		"today_rewards": stocks,
		"portfolio_value_in_inr": totalValue,
	})
}
