package httprouter

import "encoding/json"

type Response struct {
	data []byte
}

func (resp *Response) SetPlainBody(body string) {
	resp.data = []byte(body)
}

func (resp *Response) SetJsonBody(body interface{}) (err error) {
	resp.data, err = json.Marshal(body)

	return err
}
