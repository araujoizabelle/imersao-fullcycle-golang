package model
import (
	"time"
	uuid "github.com/satoring/go.uuid"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// instead of classes we use structs

type Bank struct {
	Base `valid:"required"`
	Code       string     `json:"code" valid:"notnull"`
	Name       string     `json:"name" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-`
}


// that's a method
func (bank *Bank) isValid() error {
	// it was validator, err but validator is not used
	_, err := govalidator.ValidateStruct(bank)
	
	if err != nil {
		return err
	}
	
	return nil
}



// thats basically a constructor
// all golang functions returns two values, your result and error
// the * is a passage by pointer, cool 

func NewBank (code string, name string) (*Bank, error) {
	// := is how we create instances
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err:= bank.isValid()
	if err != nil {
		return nil, err
	}

	// return the reference and not the pointer to the memory
	return &bank, nil // nil is a blank error
} 

