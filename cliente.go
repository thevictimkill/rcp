package main

import (
	"fmt"
	"net/rpc"
)

type Materia struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func cliente() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var op int64
	for {
		fmt.Println("1.- Agregar calificacion\n2.- Promedio de alumno\n3.-Promedio general\n4.-Promedio de materia\n5.-Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var nombre string
			fmt.Print("Alumno: ")
			fmt.Scanln(&nombre)
			var materia string
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			var cal float64
			fmt.Print("Calificacion: ")
			fmt.Scanln(&cal)
			mat := &Materia{nombre, materia, cal}
			var result string
			err = c.Call("Servidor.AgregarCalificacion", mat, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Servidor.AgregarCalificacion: ", result)
			}
		case 2:
			var nombre string
			fmt.Print("Alumno: ")
			fmt.Scanln(&nombre)

			var result float64
			err = c.Call("Servidor.Alumnoprom", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Servidor.Alumnoprom", nombre, " : ", result)
			}
		case 3:
			var prom float64
			err = c.Call("Servidor.PromedioGRL", prom, &prom)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Servidor.PromedioGRL: ", prom)
			}
		case 4:
			var materia string
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)

			var result float64
			err = c.Call("Servidor.PromedioMat", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Servidor.PromedioMat", materia, " : ", result)
			}
		case 5:
			return
		}
	}
}

func main() {
	cliente()
}
