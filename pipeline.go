package zd

type Packet interface {
	Value() interface{}
	Error() error
}

type ProviderPacket struct{
	value Provider
	error error
}
func (p ProviderPacket) Value() interface{} {
	var i interface{} = p.value
	return i
}
func (p ProviderPacket) Error() string {
	return p.err
}
