package schema
​
import (
	"errors"
​
	"github.com/jinzhu/gorm"
)
​
// Users Schema
type Users struct {
	gorm.Model
	Email     string `gorm:"NOT NULL"`
	Role      string `gorm:"NOT NULL"`
	FirstName string `gorm:"NOT NULL"`
	LastName  string `gorm:"NOT NULL"`
}
​
// We can organize queries for specific entities here, but
// we'll need to pass the database we're querying as a result.
​
func (user *Users) FindUserByID(db *gorm.DB, uid uint32) (*Users, error) {
	// The Take() function on this line will assign the user to the
	//	user object this is being called on if it exists.
	err := db.Debug().Model(Users{}).Where("id = ?", uid).Take(&user).Error
​
	if err != nil {
		return &Users{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Users{}, errors.New("User not found")
	}
​
	return user, err
}
​
func (user *Users) FindUsersByName(db *gorm.DB, firstName, lastName string) ([]Users, error) {
	err := db.Debug().Model(Users{}).Where("firstname = ? AND lastname = ?", firstName, lastName).Take(&user).Error
​
	if err != nil {
		return []Users{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return []Users{}, errors.New("User not found")
	}
​
	return []Users{}, err
}