package singleton

type President struct {
	instance *President
}

func (p *President) getInstance() *President {
	if p.instance == nil {
		p.instance = &President{}
	}

	return p.instance
}
