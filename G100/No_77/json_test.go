package No_77

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	event := Event{
		ID:   1,
		Time: time.Now(),
	}
	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

type Event struct {
	ID int
	time.Time
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			ID   int
			Time time.Time
		}{
			ID:   e.ID,
			Time: e.Time,
		},
	)
}
