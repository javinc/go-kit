package rethink

import (
	"fmt"
	"time"

	"github.com/javinc/go-kit/config"
	r "gopkg.in/gorethink/gorethink.v3"
)

// rethink abstraction

var (
	s *r.Session
)

// Init database
func Init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "Reconnecting...")
			time.Sleep(time.Second * 5)
			Init()
		}
	}()

	i, err := r.Connect(r.ConnectOpts{
		Address:  config.GetString("rethink.host"),
		Database: config.GetString("rethink.name"),
	})
	if err != nil {
		panic(fmt.Errorf("Fatal error rethink connection: %s", err))
	}

	// set instance
	s = i

	// create if not exists
	r.DBCreate(config.GetString("rethink.name")).Run(s)

	// enabling json tag as alternative on component Objects
	r.SetTags("gorethink", "store")

	// migrate schemas
	// migrate()
}

// Session instance
func Session() *r.Session {
	return s
}

// Run query
func Run(term r.Term) (*r.Cursor, error) {
	return term.Run(s)
}

// RunWrite query
func RunWrite(term r.Term) (r.WriteResponse, error) {
	return term.RunWrite(s)
}

// Migrate schema
func Migrate(tables []string) {
	for _, s := range tables {
		CreateTable(s)
	}
}
