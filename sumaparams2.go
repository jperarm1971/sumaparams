package main

import "os"
import "fmt"
import "text/template"
import "github/jperarm1971/utilsumas"

type Operacionsuma struct {
	Operandop int
	Operandos int
	Sumaop    int
}

func main() {

	const templ = `La suma de {{.Operandop}} y {{.Operandos}} es {{.Sumaop}}
	`

	var er error
	n1 := utilsum.ValidaOperando(os.Args[1])
	n2 := utilsum.ValidaOperando(os.Args[2])
	if n2 == 0 || n1 == 0 {
		os.Exit(1)
	} else {
		fmt.Println("Hola. Los valores son: n1=", n1, " y n2=", n2, "\n")
	}

	sumasencilla := Operacionsuma{n1, n2, utilsum.Operacionsuma(n1, n2)}

	t, er := template.New("plantillasuma").Parse(templ)
	checkError(er)

	er = t.Execute(os.Stdout, sumasencilla)
	checkError(er)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error plantilla", err.Error())
		os.Exit(1)
	}
}
