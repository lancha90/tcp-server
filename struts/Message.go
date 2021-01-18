package struts

import "github.com/ideazxy/iso8583"

type Data struct {
	Amount *iso8583.Numeric    `field:"4" length:"12" encode:"bcd"`
	Coversion *iso8583.Numeric `field:"10" length:"8" encode:"bcd"`
	No   *iso8583.Numeric      `field:"3" length:"6" encode:"bcd"` // bcd value encoding
	Oper *iso8583.Numeric      `field:"26" length:"2" encode:"ascii"` // ascii value encoding
	Ret  *iso8583.Alphanumeric `field:"39" length:"2"`
	Sn   *iso8583.Llvar        `field:"45" length:"23" encode:"bcd,ascii"` // bcd length encoding, ascii value encoding
	Info *iso8583.Lllvar       `field:"46" length:"42" encode:"bcd,ascii"`
	Mac  *iso8583.Binary       `field:"64" length:"8"`
}