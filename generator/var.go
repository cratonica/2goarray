package generator

import (
	"fmt"
	"io/ioutil"
)

type Var struct {
	File string
	Name string
	Data []byte
}

func NewVar() *Var {
	v := Var{}
	v.Data = make([]byte, 0)
	return &v
}

func (v *Var) SetData(name string, data []byte) {
	var totalBytes uint64
	var tmp = ""
	for _, k := range data {
		if totalBytes%12 == 0 {
			tmp += "\n\t"
		}
		tmp += fmt.Sprintf("0x%02x, ", k)
		totalBytes++
	}
	v.Name = FilepathToStructName(name)
	v.Data = []byte(tmp)
}

func (v *Var) SetDataFromFile(fPath string) error {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return err
	}
	v.SetData(fPath, data)
	v.File = fPath
	return nil
}

func (v *Var) GenerateCode() string {
	return "// " + v.Name + " store the data of '" + v.File + "' as a byte array.\nvar " + v.Name + " []byte = []byte{" + string(v.Data) + "\n}\n\n"
}
