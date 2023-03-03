package database

import (
	"fmt"
	pb "github.com/A-Chidalu/forwardbiddingservice/api/proto"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

type Bid struct {
	ID            uint32    `gorm:"primary_key"`
	UserID        uint32    `gorm:"not null"`
	ItemID        uint32    `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	Timestamp     time.Time `gorm:"autoCreateTime"`
	IsTerminating bool      `gorm:"not null;default:false"`
}

func ConnectDB() error {
	// Connect to the database.
	dsn := "root:password@tcp(db:3306)/BidDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set the maximum number of idle connections.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL database: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum number of open connections.
	sqlDB.SetMaxOpenConns(100)

	// Auto-migrate the database schema.
	if err := db.AutoMigrate(&Bid{}); err != nil {
		return fmt.Errorf("failed to migrate database schema: %w", err)
	}

	// Store the database connection in the global variable.
	dbConn = db

	log.Println("Connected to database")

	return nil
}

func CloseDB() error {
	// Close the database connection.
	sqlDB, err := dbConn.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL database: %w", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}

	log.Println("Closed database connection")

	return nil
}

func SaveForwardBid(request *pb.ForwardBidRequest) (*Bid, error) {

	bid := Bid{
		UserID:        request.UserId,
		ItemID:        request.ItemId,
		Amount:        request.Amount,
		Timestamp:     time.Now(),
		IsTerminating: request.IsTerminating,
	}

	result := dbConn.Create(&bid) //create a record in the database

	if result.Error != nil {
		return nil, fmt.Errorf("failed to save bid: %v", result.Error)
	}
	return &bid, nil

}

func GetAllBids() ([]Bid, error) {
	var bids []Bid
	result := dbConn.Find(&bids)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get all bids: %v", result.Error)
	}
	return bids, nil
}

func GetBidById(id uint) (*Bid, error) {
	var bid Bid
	result := dbConn.Find(&bid, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get all bids: %v", result.Error)
	}
	return &bid, nil
}
