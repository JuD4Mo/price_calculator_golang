package iiomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
