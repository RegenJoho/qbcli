package qbcli

type DirectCallback struct {
	callback   QuestionCallback
	name       string
	shallClear bool
}

func NewDirectCallback(name string, callback QuestionCallback) *DirectCallback {
	return &DirectCallback{
		name:     name,
		callback: callback,
	}
}

func (d *DirectCallback) Clear(shallClear bool) *DirectCallback {
	d.shallClear = shallClear
	return d
}

func (d *DirectCallback) GetOptions() []string {
	return []string{}
}

func (d *DirectCallback) GetName() string {
	return d.name
}

func (d *DirectCallback) Ask() error {
	d.callback("")
	return nil
}
