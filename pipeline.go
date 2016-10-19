package zd

type Packet struct {
	Value interface{}
	Error error
}

type Step interface{
	Process(in <-chan Packet) chan Packet
}
