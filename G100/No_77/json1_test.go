package No_77

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJson1(t *testing.T) {
	t := time.Now()
	event1 := event{Time: t}
	b, err := json.Marshal(event1)

	var event2 event
	err = json.Unmarshal(b, &event2)
	fmt.Println(event1 == event2)
}

type event struct {
	Time time.Time
}
