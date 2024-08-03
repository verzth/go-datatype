package main

import (
	"encoding/json"
	"fmt"
	"git.verzth.work/go/datatype"
)

func main() {
	//s := "[1,3]"
	//s := "[false,true]"
	//s := "[\"1\",\"2\"]"
	/*var x datatype.ArrayMutatedValue
	_ = json.Unmarshal([]byte(s), &x)
	fmt.Printf("%+v %v\n", x, x[1])
	b, _ := json.MarshalIndent(x, "", "    ")
	fmt.Println(string(b))*/

	y := datatype.MapJSON{}
	json.Unmarshal([]byte(`{"a": 2.3, "b": ["2"]}`), &y)
	fmt.Println(y.String(), y["b"])

}
