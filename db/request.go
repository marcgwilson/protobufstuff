package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"bytes"
	"fmt"
	"io"
)

type Request struct {
	ResourceID uint64 `db:"resource_id"`
	URL        string `db:"url"`
	HeadersRaw string `db:"headers"`
	Body       string `db:"body"`
	resource   *Resource
	response   *Response
	AutoIncr
}

func NewRequest(resource *Resource, url string, headers string, body string) *Request {
	return &Request{
		resource.ID,
		url,
		headers,
		body,
		resource,
		nil,
		NewAutoIncrZero(),
	}
}

func GetRequest(q QueryArgs) *Request {
	var r Request
	err := GetQuery("request", q, &r)
	if err != nil {
		return nil
	}
	return &r
}

func GetRequestList(q QueryArgs) ([]*Request, error) {
	whereClause, whereValues := QueryString(q)
	query := fmt.Sprintf(`SELECT * FROM %s %s;`, "request", whereClause)

	if stmt, err := database.Session.Preparex(query); err != nil {
		return nil, err
	} else {
		r := []*Request{}
		err := stmt.Select(&r, whereValues...)
		return r, err
	}
}

// func QueryRequest(q QueryArgs) *RequestIterator {
// 	return NewRequestIterator(q)
// }

func (this *Request) GetResponse() *Response {
	q := make(QueryArgs)
	q["request_id"] = this.ID
	return GetResponse(q)
}

func (this *Request) Responses() *ResponseQuery {
	return NewResponseQuery(this)
}

func (this *Request) GetResource() *Resource {
	if this.resource == nil {
		this.resource = GetResource(NewIDArgs(this.ResourceID))
	}
	return this.resource
}

func (this *Request) Save() error {
	q := make(QueryArgs)
	q["resource_id"] = this.ResourceID
	q["url"] = this.URL
	q["headers"] = this.HeadersRaw
	q["body"] = this.Body
	return this.AutoIncr.Save("request", q)
}

func (this *Request) Delete() error {
	return del("request", this.ID)
}

func (this *Request) Equal(that *Request) bool {
	return this.ID == that.ID
}

func (this *Request) GetResponses() (*ResponseIterator, error) {
	q := make(QueryArgs)
	q["request_id"] = this.ID
	return NewResponseIterator(q)
}

func (this *Request) GetHeaderBytes() []byte {
	return []byte(this.HeadersRaw)
}

func (this *Request) GetHeader() (*fasthttp.RequestHeader, error) {
	return NewRequestHeader([]byte(this.HeadersRaw))
}

func (this *Request) GetBodyBuffer() *bytes.Buffer {
	return bytes.NewBufferString(this.Body)
}

func (this *Request) GetBody() []byte {
	if len(this.Body) > 0 {
		return []byte(this.Body)
	}

	path := database.Data.MakeBodyPath("request", this.ID)
	return database.Data.GetBytes(path)
}

func (this *Request) TestString() string {
	return fmt.Sprintf("Request: %#v", this)
}

func (this *Request) GetPath() (string, bool) {
	path := database.Data.MakeBodyPath("request", this.ID)
	_, exists := database.Data.FileExists(path)
	return path, exists
}

func (this *Request) GetFile() (io.Reader, int64, error) {
	p := database.Data.MakeBodyPath("request", this.ID)
	return database.Data.GetReader(p)
}

func (this *Request) WriteBytes(b []byte) (int64, error) {
	buf := bytes.NewBuffer(b)
	path := database.Data.MakeBodyPath("request", this.ID)
	log.Tracef("Request.WriteBytes: path=%s\n", path)
	writer := database.Data.GetWriteCloser(path)
	return buf.WriteTo(writer)
}

func (this *Request) WriteStream(reader io.ReadCloser) io.ReadCloser {
	path := database.Data.MakeBodyPath("request", this.ID)
	return database.Data.Tee(path, reader)
}

// func (this *Request) ReadStream(reader io.ReadCloser) io.ReadCloser {
// 	path := database.MakeBodyPath("request", this.ID)
// 	return database.Tee(path, reader)
// }

type RequestQuery struct {
	resource *Resource
}

func NewRequestQuery(resource *Resource) *RequestQuery {
	return &RequestQuery{resource}
}

func (this *RequestQuery) All() (*RequestIterator, error) {
	q := make(QueryArgs)
	q["resource_id"] = this.resource.ID
	return NewRequestIterator(q)
}

func (this *RequestQuery) Query(q QueryArgs) (*RequestIterator, error) {
	q["resource_id"] = this.resource.ID
	return NewRequestIterator(q)
}

func (this *RequestQuery) Get(q QueryArgs) *Request {
	q["resource_id"] = this.resource.ID
	return GetRequest(q)
}

func (this *RequestQuery) Count() int {
	q := make(QueryArgs)
	q["resource_id"] = this.resource.ID

	count, _ := CountQuery("request", q)
	return int(count)
}

func (this *RequestQuery) TestString() string {
	return fmt.Sprintf("RequestQuery: %#v\n", this)
}

type RequestIterator struct {
	Items []*Request
	Index int
}

func (r *RequestIterator) Count() int {
	return len(r.Items)
}

func (r *RequestIterator) Next() (*Request, error) {
	if r.Index < len(r.Items) {
		item := r.Items[r.Index]
		r.Index += 1
		return item, nil
	}

	return nil, Done
}

func NewRequestIterator(q QueryArgs) (*RequestIterator, error) {
	if list, err := GetRequestList(q); err != nil {
		return nil, err
	} else {
		return &RequestIterator{list, 0}, nil
	}
}
