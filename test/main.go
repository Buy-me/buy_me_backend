package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"
// )

// type Model struct {
// 	Name   string `json:"name"`
// 	Millis int64  `json:"lastModified"`
// }

// func (m Model) Lastmodified() time.Time {
// 	return time.Unix(0, m.Millis*int64(time.Millisecond))
// }

// func main() {
// 	modelVar := Model{}
// 	err := json.Unmarshal([]byte(`{ "name" : "hello", "lastModified" : 564483600000 }`), &modelVar)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(modelVar.Lastmodified())
// }
