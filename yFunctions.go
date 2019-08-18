package ylx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

// print struct as json format
// 以json格式打印结构体
func StructPrint(s interface{}) {
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		b, err := json.Marshal(s)
		if err != nil {
			fmt.Printf("%+v", s)
		}
		var out bytes.Buffer
		err = json.Indent(&out, b, "", "    ")
		if err != nil {
			fmt.Printf("%+v", s)
		}
		if _,err :=io.Copy(os.Stdout, bytes.NewReader(out.Bytes())); err!= nil{
			fmt.Println("error:", err.Error())
		}
	}
}

// calc unit for bytes, Kb, Mb, Gb, Tb, Pb
// 计算容量单位 func UnitCalc(容量， 以前的单位， 新单位) 加了单位的字符串
//func UnitCalc(size float64, oldUnit string) string{
//
//}
