package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AutoIncr struct {
	ID      uint64    `db:"id" json:"id"`
	Created time.Time `db:"created_at" json:"created_at"`
	Updated time.Time `db:"updated_at" json:"updated_at"`
}

func (this *AutoIncr) UpdateAttr(that *AutoIncr) {
	this.ID = that.ID
	this.Created = that.Created
	this.Updated = that.Updated
}

func NewAutoIncrZero() AutoIncr {
	return AutoIncr{ID: 0, Created: time.Unix(0, 0), Updated: time.Unix(0, 0)}
}

func (this *AutoIncr) IDString() string {
	return strconv.FormatUint(this.ID, 10)
}

func (this *AutoIncr) CreatedString() string {
	return this.Created.Format(time.RFC3339)
}

func (this *AutoIncr) UpdatedString() string {
	return this.Updated.Format(time.RFC3339)
}

func (this *AutoIncr) Save(name string, q map[string]interface{}) error {
	var ai *AutoIncr
	var err error

	if this.ID > 0 {
		q["id"] = this.ID
		ai, err = update(name, q)
	} else {
		ai, err = insert(name, q)
	}

	if err == nil {
		this.UpdateAttr(ai)
	}

	return err
}

func updateString(q map[string]interface{}) (string, []interface{}, error) {
	var id uint64
	var err error
	if val, ok := q["id"]; !ok {
		return "", nil, errors.New("Missing id in query map")
	} else {
		switch v := val.(type) {
		case string:
			if id, err = strconv.ParseUint(v, 10, 64); err != nil {
				return "", nil, err
			}
		case int:
			id = uint64(v)
		case int64:
			id = uint64(v)
		case uint64:
			id = v
		default:
			return "", nil, fmt.Errorf("Unknown type for id in query map: %s", v)
		}
	}

	delete(q, "id")

	query := []string{}
	values := make([]interface{}, 0, 0)
	index := 1

	for k, v := range q {
		query = append(query, fmt.Sprintf("%s = $%d", k, index))
		values = append(values, v)
		index++
	}

	values = append(values, id)
	return fmt.Sprintf("%s WHERE id = $%d", strings.Join(query, ", "), len(values)), values, nil
}

func insertString(q map[string]interface{}) string {
	into := []string{}
	values := []string{}

	for k, _ := range q {
		into = append(into, k)
		values = append(values, fmt.Sprintf(":%s", k))
	}

	return fmt.Sprintf("(%s) VALUES (%s)", strings.Join(into, ", "), strings.Join(values, ", "))
}

func insert(name string, q map[string]interface{}) (*AutoIncr, error) {
	tx := database.Session.MustBegin()
	str := insertString(q)

	// if name == "request" {
	// 	fmt.Printf("REQUEST: %q", str)
	// }

	sql := fmt.Sprintf(`INSERT INTO %s %s RETURNING id, created_at, updated_at;`, name, str)
	stmt, err := database.Session.PrepareNamed(sql)
	if err != nil {
		return nil, err
	}

	ai := AutoIncr{}
	err = stmt.Get(&ai, q)
	if err != nil {
		return nil, err
	}

	return &ai, tx.Commit()
}

func update(name string, q map[string]interface{}) (*AutoIncr, error) {
	tx := database.Session.MustBegin()
	str, values, err := updateString(q)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`UPDATE %s SET %s RETURNING id, created_at, updated_at;`, name, str)
	stmt, err := database.Session.Preparex(query)

	if err != nil {
		return nil, err
	}

	ai := AutoIncr{}
	err = stmt.Get(&ai, values...)
	if err != nil {
		return nil, err
	}

	return &ai, tx.Commit()
}
