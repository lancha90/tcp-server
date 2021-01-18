package main

import (
	"dmha/tpc-server/struts"
	"fmt"
	"github.com/ideazxy/iso8583"
	"net"
)

func main() {

	listener, err := net.Listen("tcp4", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn)  {

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	iso := iso8583.NewMessage("", &struts.Data{
		Amount: iso8583.NewNumeric(""),
		Coversion: iso8583.NewNumeric(""),
		No:   iso8583.NewNumeric(""),
		Oper: iso8583.NewNumeric(""),
		Ret:  iso8583.NewAlphanumeric(""),
		Sn:   iso8583.NewLlvar([]byte("")),
		Info: iso8583.NewLllvar([]byte("")),
		Mac:  iso8583.NewBinary([]byte("")),
	})
	iso.MtiEncode = iso8583.BCD

	err = iso.Load(buf)

	if err != nil {
		fmt.Println("ISO Decode error:", err)
	}

	resultFields := iso.Data.(*struts.Data)

	println(resultFields.Amount.Value)
	println(resultFields.Coversion.Value)
	println(string(resultFields.Info.Value))

	res, _ := iso.Bytes()
	// Send a response back to person contacting us.

	conn.Write(res)
}