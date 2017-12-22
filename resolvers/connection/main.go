package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Data struct {
	Model string
	Table string
}

func Generate(w io.Writer, d Data) error {
	t := template.New("connection")
	t, err := t.Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(w, d)
}

func main() {
	model := flag.String("model", "", "")
	table := flag.String("table", "", "")
	flag.Parse()

	outFile := fmt.Sprintf("%s_connection.go", *table)

	w, err := os.Create(filepath.Join(".", outFile))
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	data := Data{*model, *table}

	if err := Generate(w, data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("created " + outFile)
}

var tmpl = `
package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
)

//{{ .Model }}EdgeResolver resolves the fields of an edge
type {{ .Model }}EdgeResolver struct {
	db     *models.DB
	node   *models.{{ .Model }}
	cursor scalars.Cursor
}

//New{{ .Model }}EdgeResolver returns a new {{ .Model }}EdgeResolver
func New{{ .Model }}EdgeResolver(db *models.DB, r *models.{{ .Model }}, c scalars.Cursor) *{{ .Model }}EdgeResolver {
	return &{{ .Model }}EdgeResolver{
		db:     db,
		node:   r,
		cursor: c,
	}
}

//Cursor returns the cursor for the edge
func (e *{{ .Model }}EdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

//Node returns the region for the edge
func (e *{{ .Model }}EdgeResolver) Node() *{{ .Model }}Resolver {
  return New{{ .Model }}Resolver(e.db, e.node)
}

//{{ .Model }}ConnectionResolver resolves the fields of a region connection
type {{ .Model }}ConnectionResolver struct {
	db    *models.DB
	items []models.{{ .Model }}
	start scalars.Cursor
	end   scalars.Cursor
}

//New{{ .Model }}ConnectionResolver returns a new {{ .Model }}ConnectionResolver
func New{{ .Model }}ConnectionResolver(db *models.DB, items []models.{{ .Model }}, args arguments.Connection) (*{{ .Model }}ConnectionResolver, error) {
	size := len(items)

	start := scalars.NewCursor("{{ .Table }}", 0)
	if args.After != nil {
		start = *args.After
	}

	starti, err := start.IntValue()
	if err != nil {
		return nil, err
	}

	end := scalars.NewCursor("{{ .Table }}", size+starti)
	if args.First != nil {
		actualEnd := starti + int(*args.First)
		if actualEnd > size {
			end = scalars.NewCursor("{{ .Table }}", actualEnd)
		}
	}
	if err != nil {
		return nil, err
	}

	return &{{ .Model }}ConnectionResolver{
		db:    db,
		items: items,
		start: start,
		end:   end,
	}, nil
}

//TotalCount returns the total number of items in a connection
func (c {{ .Model }}ConnectionResolver) TotalCount() (int32, error) {
	count, err := c.db.Count("{{ .Table }}")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (c {{ .Model }}ConnectionResolver) PageInfo() (*PageResolver, error) {
	count, err := c.TotalCount()
	if err != nil {
		return nil, err
	}
	endi, err := c.end.IntValue()
	if err != nil {
		return nil, err
	}

	hasNext := int(count) > endi

	return NewPageResolver(c.start, c.end, hasNext), nil
}

//Edges returns the edges of a connection
func (c {{ .Model }}ConnectionResolver) Edges() (*[]*{{ .Model }}EdgeResolver, error) {
	var e []*{{ .Model }}EdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor("{{ .Table }}", cursorLocation)
		e = append(e, New{{ .Model }}EdgeResolver(c.db, &item, cursor))
	}

	return &e, nil
}
`
