package controller

const numDataTypes = 8

var DataChan chan interface{}

func Setup() {
	DataChan = make(chan interface{}, numDataTypes)
}
