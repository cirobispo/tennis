package main

import (
	"cbtennis/internal/turning25"
	"cbtennis/internal/turning25/countingturn"
	"cbtennis/internal/turning25/timingturn"
	"cbtennis/internal/turning25/turn"
	"fmt"
)

func main() {
	t := turn.New(turning25.TSA)
	ct := countingturn.New(t)
	tt := timingturn.New(ct.Turn)
	tt.Do()
	tt.Do()
	tt.Do()
	var ect countingturn.Counting = ct
	var ett timingturn.Timing = tt
	fmt.Println(ect.Count())
	fmt.Println(ett.SlapsedTime())
}
