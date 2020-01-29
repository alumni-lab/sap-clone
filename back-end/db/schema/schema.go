package schema

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Email     string
	Role      string
	FirstName string
	LastName  string
}

type Customers struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Phone     string
	Note      string
}

type CustomerOrders struct {
	gorm.Model
	User       Users
	UserID     uint
	Customer   Customers
	CustomerID uint
	Date       time.Time
	note       string
}

type Items struct {
	gorm.Model
	Vendor      Vendors
	VendorID    uint
	Description string
	Type        string
	Price       int
}

type OrderItems struct {
	gorm.Model
	Order         CustomerOrders
	OrderID       uint
	Item          Items
	ItemID        uint
	Date          time.Time
	PurchasePrice int
	Quantity      int
	Discount      int
}

type Vendors struct {
	gorm.Model
	User       Users
	UserID     uint
	Customer   Customers
	CustomerID uint
	Date       time.Time
	Note       string
}

type VendorOrders struct {
	gorm.Model
	User       Users
	UserID     uint
	Customer   Customers
	CustomerID uint
	Date       time.Time
	Note       string
}

type Addresses struct {
	gorm.Model
	Customer   Customers
	CustomerID uint
	FirstName  string
	LastName   string
	Address    string
	City       string
	Phone      string
	Note       string
	IsActive   bool
}
