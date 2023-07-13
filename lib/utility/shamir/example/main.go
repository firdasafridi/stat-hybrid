package main

import (
	"encoding/base64"
	"fmt"

	"github.com/firdasafridi/stat-hybrid/lib/utility/shamir"
)

func main() {
	encode()
	decode()
}

func encode() {
	plain := []byte(`Sample data yg di split`)

	secrets, err := shamir.Split(plain, 5, 3)
	if err != nil {
		panic(err)
	}

	for _, part := range secrets {
		sEnc := base64.StdEncoding.EncodeToString(part)
		fmt.Println(sEnc)
	}
}

func decode() {

	s1, _ := base64.StdEncoding.DecodeString("Hw/Uni3QpqFN4oRvSlMtxjqpvmS+ZM1k")
	s2, _ := base64.StdEncoding.DecodeString("DRhq1LDXD3y7XFjvS1V64pFLymkKUMYP")
	s3, _ := base64.StdEncoding.DecodeString("m/vJrxSghSV82OXWrpTKUtLybMxkYeFJ")

	comb := [][]byte{
		s1,
		s2,
		s3,
	}

	text, err := shamir.Combine(comb)
	if err != nil {
		panic(err)
	}
	fmt.Println("OUTPUT:", string(text))
}