package Protocol

type Packet struct {
	typeMessage int
	data        []byte
}

func (p *Packet) GetTypeMessage() int {
	return p.typeMessage
}

func (p *Packet) GetData() []byte {
	return p.data
}

func (p *Packet) SetData(data []byte) {
	p.data = data
}

func (p *Packet) SetTypeMessage(typeM int) {
	p.typeMessage = typeM
}
