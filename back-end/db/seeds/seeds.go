package seeds

import (
	"../schema"
	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	user := createUser("bobbyrossy@ross.ross", "Admin", "Bob", "Ross")
	db.Create(&user)

	customer := createCustomer("customer@email.com", "Dan", "Harmon", "19026652353", "")
	customer1 := createCustomer("customer1@email.com", "George", "Martin", "19026652776", "")
	db.Create(&customer)
	db.Create(&customer1)

	customerOrder := createCustomerOrder(user, customer, "")
	customerOrder1 := createCustomerOrder(user, customer1, "")
	db.Create(&customerOrder)
	db.Create(&customerOrder1)

	vendor := createVendor("vendor@company.biz", "Biz Company", "15555555555", "")
	db.Create(&vendor)

	item := createItem(vendor, "Bruschetta, but it's actually just olives", "You heard me.", "Warehouse", 52599)
	db.Create(&item)

	orderItem := createOrderItem(customerOrder, item, 52599, 1, 0)
	db.Create(&orderItem)

	vendorOrder := createVendorOrder(user, vendor, "Need more fake Bruschetta.")
	db.Create(&vendorOrder)

	address := createAddress(customer, "Bob", "Ross", "1440 Painter Ave.", "PaintScape", "12235658747", "Home", true)
	db.Create(&address)
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

func createCustomerOrder(user schema.Users, customer schema.Customers, note string) schema.CustomerOrders {
	return schema.CustomerOrders{
		User:     user,
		Customer: customer,
		Note:     note,
	}
}

func createItem(vendor schema.Vendors, name, description, vType string, price int) schema.Items {
	return schema.Items{
		Vendor:      vendor,
		Name:        name,
		Description: description,
		Type:        vType,
		Price:       price,
	}
}

func createOrderItem(order schema.CustomerOrders, item schema.Items, purchasePrice, quantity, discount int) schema.OrderItems {
	return schema.OrderItems{
		Order:         order,
		Item:          item,
		PurchasePrice: purchasePrice,
		Quantity:      quantity,
		Discount:      discount,
	}
}

func createVendor(email, name, phone, note string) schema.Vendors {
	return schema.Vendors{
		Email: email,
		Name:  name,
		Phone: phone,
		Note:  note,
	}
}

func createVendorOrder(user schema.Users, vendor schema.Vendors, note string) schema.VendorOrders {
	return schema.VendorOrders{
		User:   user,
		Vendor: vendor,
		Note:   note,
	}
}

func createAddress(customer schema.Customers, firstName, lastName, address, city, phone, note string, isActive bool) schema.Addresses {
	return schema.Addresses{
		Customer:  customer,
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		City:      city,
		Phone:     phone,
		Note:      note,
		IsActive:  isActive,
	}
}
