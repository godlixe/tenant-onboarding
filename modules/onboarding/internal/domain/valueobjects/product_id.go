package valueobjects

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"tenant-onboarding/modules/onboarding/internal/errorx"

	"github.com/google/uuid"
)

var ErrInvalidProductID = errorx.ErrInvalidProductID

type ProductID struct {
	ID string
}

func NewProductID(id string) (ProductID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return ProductID{}, ErrInvalidProductID
	}

	return ProductID{id}, nil
}

func GenerateProductID() ProductID {
	return ProductID{uuid.NewString()}
}

func (i ProductID) String() string {
	return i.ID
}

func (i ProductID) Equals(other ProductID) bool {
	return strings.EqualFold(i.ID, other.ID)
}

func (i *ProductID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		fmt.Println("am here 1:", v)
		res, err := NewProductID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		fmt.Println("am here 2:", v)
		return ErrInvalidProductID
	}

	return nil
}

func (i *ProductID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}
