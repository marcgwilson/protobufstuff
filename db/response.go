package db

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

type Response struct {
	RequestID  uint64 `db:"request_id"`
	HeadersRaw string `db:"headers"`
	Body       string `db:"body"`
	Temporal   string `db:"temporal"`
	request    *Request
	AutoIncr
}

func NewResponse(request *Request, headers string, body string, temporal string) *Response {
	return &Response{
		request.ID,
		headers,
		body,
		temporal,
		request,
		NewAutoIncrZero(),
	}
}

func GetResponse(q QueryArgs) *Response {
	var r Response
	err := GetQuery("response", q, &r)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		return nil
	}
	return &r
}

func GetResponseList(q QueryArgs) ([]*Response, error) {
	whereClause, whereValues := QueryString(q)
	query := fmt.Sprintf(`SELECT * FROM %s %s;`, "response", whereClause)

	if stmt, err := database.Session.Preparex(query); err != nil {
		return nil, err
	} else {
		r := []*Response{}
		err := stmt.Select(&r, whereValues...)
		return r, err
	}
}

func QueryResponse(q QueryArgs) (*ResponseIterator, error) {
	return NewResponseIterator(q)
}

func (this *Response) GetHeaderBytes() []byte {
	return []byte(this.HeadersRaw)
}

func (this *Response) GetHeader() (*fasthttp.ResponseHeader, error) {
	return NewResponseHeader([]byte(this.HeadersRaw))
}

func (this *Response) GetBodyBuffer() *bytes.Buffer {
	return bytes.NewBufferString(this.Body)
}

func (this *Response) GetBody() []byte {
	if len(this.Body) > 0 {
		return []byte(this.Body)
	}

	path := database.Data.MakeBodyPath("response", this.ID)
	return database.Data.GetBytes(path)
}

func (this *Response) GetRequest() *Request {
	if this.request == nil {
		this.request = GetRequest(NewIDArgs(this.RequestID))
	}
	return this.request
}

func (this *Response) GetTemporal() string {
	return this.Temporal
}

func (this *Response) SetTemporal(temporal string) {
	this.Temporal = temporal
}

func (this *Response) TestString() string {
	return fmt.Sprintf("Response: %#v\n", this)
}

func (this *Response) GetPath() (string, bool) {
	path := database.Data.MakeBodyPath("response", this.ID)
	_, exists := database.Data.FileExists(path)
	return path, exists
}

func (this *Response) GetFile() (io.Reader, int64, error) {
	p := database.Data.MakeBodyPath("response", this.ID)
	return database.Data.GetReader(p)
}

func (this *Response) WriteBytes(b []byte) (int64, error) {
	buf := bytes.NewBuffer(b)
	path := database.Data.MakeBodyPath("response", this.ID)
	log.Tracef("WriteBytes: path=%s\n", path)
	writer := database.Data.GetWriteCloser(path)
	return buf.WriteTo(writer)
}

func (this *Response) WriteStream(reader io.ReadCloser) io.ReadCloser {
	path := database.Data.MakeBodyPath("response", this.ID)
	return database.Data.Tee(path, reader)
}

func (this *Response) SetTemporalBase64(temporal string) error {
	if decoded, err := base64.URLEncoding.DecodeString(temporal); err != nil {
		return err
	} else {
		this.Temporal = string(decoded)
		return nil
	}
}

func (this *Response) Save() error {
	q := make(QueryArgs)
	q["request_id"] = this.RequestID
	q["headers"] = this.HeadersRaw
	q["body"] = this.Body
	q["temporal"] = this.Temporal
	return this.AutoIncr.Save("response", q)
}

func (this *Response) Delete() error {
	return del("response", this.ID)
}

func (this *Response) Equal(that *Response) bool {
	return this.ID == that.ID
}

type ResponseQuery struct {
	request *Request
}

func (this *ResponseQuery) All() (*ResponseIterator, error) {
	q := make(QueryArgs)
	q["request_id"] = this.request.ID
	return NewResponseIterator(q)
}

func (this *ResponseQuery) Query(q QueryArgs) (*ResponseIterator, error) {
	q["request_id"] = this.request.ID
	return NewResponseIterator(q)
}

func (this *ResponseQuery) Get(q QueryArgs) *Response {
	q["request_id"] = this.request.ID
	return GetResponse(q)
}

func (this *ResponseQuery) Count() int {
	q := make(QueryArgs)
	q["request_id"] = this.request.ID
	count, _ := CountQuery("response", q)
	return int(count)
}

func (this *ResponseQuery) TestString() string {
	return spew.Sdump(this)
}

func NewResponseQuery(request *Request) *ResponseQuery {
	return &ResponseQuery{request}
}

type ResponseIterator struct {
	Items []*Response
	Index int
}

func (r *ResponseIterator) Count() int {
	return len(r.Items)
}

func (r *ResponseIterator) Next() (*Response, error) {
	if r.Index < len(r.Items) {
		item := r.Items[r.Index]
		r.Index += 1
		return item, nil
	}

	return nil, Done
}

func NewResponseIterator(q QueryArgs) (*ResponseIterator, error) {
	if list, err := GetResponseList(q); err != nil {
		return nil, err
	} else {
		return &ResponseIterator{list, 0}, nil
	}
}

// type ResponseIterator struct {
// 	Iterator
// }

// func (this *ResponseIterator) Next() (*Response, error) {
// 	item := &Response{}
// 	err := this.next(item)
// 	return item, err
// }

// func NewResponseIterator(q QueryArgs) *ResponseIterator {
// 	it := NewIterator("response", q)
// 	return &ResponseIterator{*it}
// }
