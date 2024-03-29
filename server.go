package main

import (
	"dmha/tpc-server/struts"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp4", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	rand.Seed(time.Now().Unix())

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

	defer conn.Close()

	buf := make([]byte, 1024)
	size, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	println(string(buf))
	println(size)

	/*iso := iso8583.NewMessage("", &struts.Data{
		Pan: iso8583.NewLlnumeric(""),
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

	//hx := hex.EncodeToString(buf)
	//println(hx)

	println("--------------------")

	if err != nil {
		fmt.Println("ISO Decode error:", err)
	}

	resultFields := iso.Data.(*struts.Data)

	printResponse(*resultFields)

	res, _ := iso.Bytes()

	println(len(res))

	// Send a response back to person contacting us.

	conn.Write(res)*/
}

func printResponse(data struts.Data){
	fmt.Println("PAN: ", data.Pan.Value)
	fmt.Println("Amount: ", data.Amount.Value)
	fmt.Println("Conversion Rate: ", data.Coversion.Value)
	fmt.Println("Info: ", string(data.Info.Value))
	fmt.Println("Oper: ", data.Oper.Value)
	fmt.Println("Ret: ", string(data.Ret.Value))
}