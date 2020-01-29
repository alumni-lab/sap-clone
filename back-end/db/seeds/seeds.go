package seeds

import (
	"../schema"
	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	user := createUser("bobbyrossy@ross.ross", "Admin", "Bob", "Ross")
	db.Create(&user)
}

func createUser(email, role, firstName, lastName string) schema.Users {
	return schema.Users{
		Email:     email,
		Role:      role,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func createCustomer(email, firstName, lastName, phone, note string) schema.Customers {
	return schema.Customers{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Note:      note,
	}
}

func createCustomerOrder(user schema.Users, customer schema.Customers, userID, customerID uint, note string) schema.CustomerOrders {
	return schema.CustomerOrders{
		User:       user,
		UserID:     userID,
		Customer:   customer,
		CustomerID: customerID,
		Note:       note,
	}
}

func createItem(vendor schema.Vendors, vendorID uint, name, description, vType string, price int) schema.Items {
	return schema.Items{
		Vendor:      vendor,
		VendorID:    vendorID,
		Name:        name,
		Description: description,
		Type:        vType,
		Price:       price,
	}
}

func createOrderItem(order schema.CustomerOrders, item schema.Items, orderID, itemID uint, purchasePrice, quantity, discount int) schema.OrderItems {
	return schema.OrderItems{
		Order:         order,
		OrderID:       orderID,
		Item:          item,
		ItemID:        itemID,
		PurchasePrice: purchasePrice,
		Quantity:      quantity,
		Discount:      discount,
	}
}

func createVendor(user schema.Users, customer schema.Customers, userID, customerID uint, note string) schema.Vendors {
	return schema.Vendors{
		User:       user,
		UserID:     userID,
		Customer:   customer,
		CustomerID: customerID,
		Note:       note,
	}
}

func createVendorOrder(user schema.Users, customer schema.Customers, userID, customerID uint, note string) schema.VendorOrders {
	return schema.VendorOrders{
		User:       user,
		UserID:     userID,
		Customer:   customer,
		CustomerID: customerID,
		Note:       note,
	}
}

func createAddress(customer schema.Customers, customerID uint, firstName, lastName, address, city, phone, note string, isActive bool) schema.Addresses {
	return schema.Addresses{
		Customer:   customer,
		CustomerID: customerID,
		FirstName:  firstName,
		LastName:   lastName,
		Address:    address,
		City:       city,
		Phone:      phone,
		Note:       note,
		IsActive:   isActive,
	}
}
