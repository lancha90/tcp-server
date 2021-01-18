package main

import (
	"dmha/tpc-server/struts"
	"fmt"
	"github.com/ideazxy/iso8583"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:8001")

	data := struts.Data{
		Pan: iso8583.NewNumeric("1234123412341234"),
		Amount: iso8583.NewNumeric("000000001200"),
		Coversion: iso8583.NewNumeric("00000011"),
		No:   iso8583.NewNumeric("001111"),
		Oper: iso8583.NewNumeric("22"),
		Ret:  iso8583.NewAlphanumeric("ok"),
		Sn:   iso8583.NewLlvar([]byte("abc001")),
		Info: iso8583.NewLllvar([]byte("Info NewLllvar")),
		Mac:  iso8583.NewBinary([]byte("a1s2d3f4")),
	}
	msg := iso8583.NewMessage("0800", data)
	msg.MtiEncode = iso8583.BCD
	b, err := msg.Bytes()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%x\n", b)

	c.Write(b)
}
