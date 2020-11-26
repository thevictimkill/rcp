package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Materia struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

type Servidor struct {
	Materias    map[string]map[string]float64
	Estudiantes map[string]map[string]float64
}

func (sr *Servidor) BuscarMateria(nombre string) bool {
	for materia, _ := range sr.Materias {
		if materia == nombre {
			return true
		}
	}
	return false
}

func (sr *Servidor) BuscarAlumno(nombre string) bool {
	for alumno, _ := range sr.Estudiantes {
		if alumno == nombre {
			return true
		}
	}
	return false
}

func (sr *Servidor) AgregarCalificacion(mat *Materia, reply *string) error {
	for alumno, _ := range sr.Materias[mat.Materia] {
		if alumno == mat.Alumno {
			return errors.New("Ya tiene calificación")
		}
	}
	if sr.BuscarMateria(mat.Materia) {
		sr.Materias[mat.Materia][mat.Alumno] = mat.Calificacion
	} else {

		alumno := make(map[string]float64)
		alumno[mat.Alumno] = mat.Calificacion
		sr.Materias[mat.Materia] = alumno

	}
	if sr.BuscarAlumno(mat.Alumno) {
		sr.Estudiantes[mat.Alumno][mat.Materia] = mat.Calificacion
	} else {
		materia := make(map[string]float64)
		materia[mat.Materia] = mat.Calificacion
		sr.Estudiantes[mat.Alumno] = materia
	}
	s := fmt.Sprintf("%.2f", mat.Calificacion)
	*reply = "Se añadio la calificaion: " + mat.Materia + "de" + mat.Alumno + ": " + s
	return nil
}

func (sr *Servidor) Alumnoprom(nombre string, reply *float64) error {
	aux_cal := 0.0
	aux_mat := 0.0
	for _, calificacion := range sr.Estudiantes[nombre] {
		aux_cal += calificacion
		aux_mat += 1
	}
	if aux_mat > 0 {
		*reply = aux_cal / aux_mat
		return nil
	} else {
		return errors.New("No se encontro el alumno")
	}
}

func (sr *Servidor) PromedioMat(nombre string, reply *float64) error {
	aux_cal := 0.0
	aux_alu := 0.0
	for _, calificacion := range sr.Materias[nombre] {
		aux_cal += calificacion
		aux_alu += 1
	}
	if aux_alu > 0 {
		*reply = aux_cal / aux_alu
		return nil
	} else {
		return errors.New("No se encontro la materia")
	}
}

func (sr *Servidor) PromedioGRL(reply_value float64, reply *float64) error {
	aux_cal := 0.0
	aux_alu := 0.0
	for alumno, _ := range sr.Estudiantes {
		prom := 0.0
		err := sr.Alumnoprom(alumno, &prom)
		if err != nil {
			return err
		}
		aux_cal += prom
		aux_alu += 1
	}
	if aux_alu > 0 {
		*reply = aux_cal / aux_alu
		return nil
	} else {
		return errors.New("No hay alumnos")
	}
}

func server(Materias, Estudiantes map[string]map[string]float64) {
	new_server := new(Servidor)
	new_server.Estudiantes = Estudiantes
	new_server.Materias = Materias
	rpc.Register(new_server)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	Materias := make(map[string]map[string]float64)
	Estudiantes := make(map[string]map[string]float64)
	go server(Materias, Estudiantes)

	var input string
	fmt.Scanln(&input)
}
