package queries

import (
	"errors"

	"github.com/alumni-lab/sap-clone/db/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ​type Vendors struct {
// 	gorm.Model
// 	Email string `gorm:"NOT NULL"`
// 	Name  string `gorm:"NOT NULL"`
// 	Phone string `gorm:"NOT NULL"`
// 	Note  string
// }
// schema.Vendors
//test
func (user *schema.Vendors) FindVendorByID(db *gorm.DB, uid uint32) (*schema.Vendors, error) {
	// The Take() function on this line will assign the user to the
	//	user object this is being called on if it exists.
	err := db.Debug().Model(schema.Vendors{}).Where("id = ?", uid).Take(&user).Error

	if err != nil {
		return &schema.Vendors{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &schema.Vendors{}, errors.New("User not found")
	}

	return user, err
}
func (user *schema.Vendors) FindUsersByName(db *gorm.DB, firstName, lastName string) ([]schema.Vendors, error) {
	err := db.Debug().Model(schema.Vendors{}).Where("firstname = ? AND lastname = ?", firstName, lastName).Take(&user).Error

	if err != nil {
		return []schema.Vendors{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return []schema.Vendors{}, errors.New("User not found")
	}

	return []schema.Vendors{}, err
}
