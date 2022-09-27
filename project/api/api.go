package api

type Api interface {
	ValidateStruct(any interface{}) (err error)
}
