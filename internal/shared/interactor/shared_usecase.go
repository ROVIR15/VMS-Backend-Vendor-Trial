package interactor

import (
	"reflect"

	reflectionHelper "github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/reflection/reflection_helper"
)

type sharedUsecaseContainer struct {
	// list of shared usecase
}

func NewSharedUsecaseContainer() *sharedUsecaseContainer {
	return &sharedUsecaseContainer{}
}

func (i *sharedUsecaseContainer) InjectUsecase(uc interface{}) error {
	container := reflect.ValueOf(i).Elem()
	usecase := reflect.ValueOf(uc)

	for i := 0; i < container.NumField(); i++ {
		field := container.Field(i)
		if field.Type() == usecase.Type() {
			reflectionHelper.SetFieldValue(field, uc)
		}
	}
	return nil
}

// implement shared usecase
