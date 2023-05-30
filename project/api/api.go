package api

type Api interface {
	ValidateStruct(any interface{}) (err error)
}

type FillDataInterface interface {
	FillData() (any interface{}, err error)
}
