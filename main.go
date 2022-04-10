package main

import "fmt"

type ContaCorrente = struct {
	titular string
	numeroAgencia int 
	numeroConta int 
	saldo float64 
}

func main() {
	contaDoBruno := ContaCorrente{"Bruno",123, 123456, 100.00}
	contaDoThiago := ContaCorrente{"Thiago", 123, 123457, 150.00}
	fmt.Println(contaDoBruno)
	fmt.Println(contaDoThiago)
}