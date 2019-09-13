package redis

type Record struct {
	AddressField string `json:"Address"`
	TextField    string `json:"Text"`
	IDField      int64  `json:"ID"`
}

func (r *Record) ID() int64 {
	return r.IDField
}

func (r *Record) Text() string {
	return r.TextField
}

func (r *Record) Address() string {
	return r.AddressField
}

func newRecord(id int64, text, address string) *Record {
	return &Record{
		IDField:      id,
		TextField:    text,
		AddressField: address,
	}
}
