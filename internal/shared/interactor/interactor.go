package interactor

type SharedUsecaseContainer interface {
	InjectUsecase(uc interface{}) error
	// Other exported usecase method
}
