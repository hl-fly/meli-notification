package viewmodel

import (
	"fmt"
	"time"

	"github.com/hector-leite/meli-notification/domain/entity"
)

type NotificationInsertRequest struct {
	ProductUUID string    `json:"product_uuid"`
	Message     string    `json:"message"`
	Link        string    `json:"link"`
	ExpDate     time.Time `json:"exp_date"`
	Target      bool      `json:"target"`
}

func (vm NotificationInsertRequest) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var aggErrors []error

	if vm.Message == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "message"))
	}

	if vm.ExpDate.IsZero() {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "exp_date"))
	}

	if len(aggErrors) > 0 {
		return aggErrors
	}

	return nil
}

func (vm NotificationInsertRequest) Parse() entity.Notification {

	return entity.Notification{
		Product: entity.Product{
			UUID: vm.ProductUUID,
		},
		Message:        vm.Message,
		Link:           vm.Link,
		ExpirationDate: vm.ExpDate,
		Target:         vm.Target,
	}
}
