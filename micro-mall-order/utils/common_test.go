package utils

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
)

//func TestAssignment(t *testing.T) {
//	v1 := &pms.PmsCategory{
//		CatId: 1,
//		Name:  "lzd",
//	}
//	v2 := &proto_product.UpdateCategoryRequest{}
//	ok := Assignment(v2, v1)
//	fmt.Println(ok)
//	fmt.Printf("%+v \n", v2)
//}

type Stu1 struct {
	SpuDesc string `json:"spu_desc"`
}

type Stu2 struct {
	SpuDesc string `json:"spuDesc"`
}

func TestAA(t *testing.T) {
	s1 := []*Stu2{
		&Stu2{
			SpuDesc: "aaa",
		},
		&Stu2{
			SpuDesc: "bbb",
		},
	}
	s2 := make([]Stu1, 0, 10)
	copier.CopyWithOption(&s2, &s1, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	fmt.Printf("%+v \n", s2)
}
