package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DatabaseUrl = databaseUrl
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err, "invalid config")
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			_, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				return
			}
		}
		s.Close()
	}
}
