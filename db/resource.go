package db

import (
	"github.com/davecgh/go-spew/spew"

	"fmt"
	"net/url"
)

type Resource struct {
	URL      string `db:"url" json:"url"`
	requests *RequestManager
	AutoIncr
}

func NewResource(url string) *Resource {
	return &Resource{
		url,
		nil,
		NewAutoIncrZero(),
	}
}

func QueryResource(q QueryArgs) (*ResourceIterator, error) {
	return NewResourceIterator(q)
}

func GetResourceList(q QueryArgs) ([]*Resource, error) {
	whereClause, whereValues := QueryString(q)
	query := fmt.Sprintf(`SELECT * FROM %s %s;`, "resource", whereClause)

	if stmt, err := database.Session.Preparex(query); err != nil {
		return nil, err
	} else {
		r := []*Resource{}
		err := stmt.Select(&r, whereValues...)
		return r, err
	}
}

func GetResource(q QueryArgs) *Resource {
	r := &Resource{}
	var err error

	if _, ok := q["id"]; ok {
		err = GetQuery("resource", q, r)
	} else if _, ok := q["url"]; ok {
		err = GetQuery("resource", q, r)
	} else if val, ok := q["request_id"]; ok {
		err = database.Session.Get(r, innerJoinRequest, val)
	} else if val, ok := q["response_id"]; ok {
		err = database.Session.Get(r, innerJoinResponse, val)
	}

	if err != nil {
		// fmt.Printf("ERROR: %s\n", err)
		return nil
	}
	return r
}

var innerJoinRequest = `SELECT
  resource.id,
  resource.url,
  resource.created_at,
  resource.updated_at
FROM
  resource
INNER JOIN request ON request.resource_id = resource.id
WHERE
  request.id = $1;`

func GerResourceWithRequest(id uint64) (*Resource, error) {
	r := Resource{}
	err := database.Session.Get(&r, innerJoinRequest, id)
	return &r, err
}

var innerJoinResponse = `SELECT
  resource.id,
  resource.url,
  resource.created_at,
  resource.updated_at
FROM
  resource
INNER JOIN request ON request.resource_id = resource.id
INNER JOIN response ON response.request_id = request.id
WHERE
  response.id = $1;`

func GetResourceWithResponse(id uint64) (*Resource, error) {
	r := Resource{}
	err := database.Session.Get(&r, innerJoinResponse, id)
	return &r, err
}

// var innerJoin := `SELECT c.company_id,c.company_name, c.company_phone,m.member_id
// FROM company c
// INNER JOIN member m
// ON member.company_id = company .company_id;`

func (this *Resource) Requests() *RequestManager {
	if this.ID > 0 {
		if this.requests == nil {
			this.requests = NewRequestManager(this)
		}
		return this.requests
	}
	return nil
}

func (this *Resource) TestString() string {
	return spew.Sdump(this)
}

func (this *Resource) Save() error {
	if _, err := url.Parse(this.URL); err != nil {
		return err
	}

	q := make(QueryArgs)
	q["url"] = this.URL
	return this.AutoIncr.Save("resource", q)
}

func (this *Resource) Delete() error {
	return del("resource", this.ID)
}

func (this *Resource) Equal(that *Resource) bool {
	return this.ID == that.ID
}

func (this *Resource) String() string {
	return fmt.Sprintf("Resource[%d, %s]", this.ID, this.URL)
}

type ResourceIterator struct {
	Items []*Resource
	Index int
}

func (r *ResourceIterator) Count() int {
	return len(r.Items)
}

func (r *ResourceIterator) Next() (*Resource, error) {
	if r.Index < len(r.Items) {
		item := r.Items[r.Index]
		r.Index += 1
		return item, nil
	}

	return nil, Done
}

func NewResourceIterator(q QueryArgs) (*ResourceIterator, error) {
	if list, err := GetResourceList(q); err != nil {
		return nil, err
	} else {
		return &ResourceIterator{list, 0}, nil
	}
}

// type ResourceIterator struct {
// 	Iterator
// }

// func (this *ResourceIterator) Next() (*Resource, error) {
// 	item := &Resource{}
// 	if err := this.next(item); err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }

// func NewResourceIterator(q QueryArgs) *ResourceIterator {
// 	it := NewIterator("resource", q)
// 	return &ResourceIterator{*it}
// }
