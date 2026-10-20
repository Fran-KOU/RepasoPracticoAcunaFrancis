// Nombre: Francis Acuña ~ Paralelo: B

package main

import (
	"fmt"
)

// Capa UIDE
// Reto 1: Inscribir estadisticas y jugadores

type Estadisticas struct {
	Kills		int
	Muertes		int
	Asistencias	int
}

type Jugador struct {
	Nick	string
	Equipo	string
	Stats	Estadisticas
	Activo	bool
}

func registrar (j *Jugador, kills, muertes, asistencias int) {
	j.Stats.Kills = kills
	j.Stats.Muertes = muertes
	j.Stats.Asistencias = asistencias
}

func kda(e Estadisticas)float64{
	muertes := e.Muertes
	if muertes == 0 {
		muertes = 1
	}
	return float64(e.Kills+e.Asistencias) / float64(muertes)
}

// Ejercicio 3 (Aqui se devuelve una copia)

func mvp(roster []Jugador) Jugador {
	var mejor Jugador
	for _, j := range roster {
		if j.Activo && kda(j.Stats) > kda(mejor.Stats){
			mejor = j
		}
	}
	return mejor
}

func main() {
	fmt.Println("Hola Mundo")
}