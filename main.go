package main

import "github.com/jpluduena/scripts-pc-cancellations/testuser"

func main() {
	/**
	DOCUMENTACION DE USO:
	https://docs.google.com/document/d/1b6NiUtkUfFhkY5MbC7ga2BMM-BChm4-ZmAznzzqqeN0

	CONFIGURAR VARIABLES EN /core/config.go
	*/

	// Creacion de usuarios de test
	//  - key: Sites para los q crear usuarios
	//  - value: Cantidad de usuarios a crear para cada site
	testuser.CreateTestUser(map[string]int{"MLA": 3, "MLB": 0, "MLC": 1, "MLU": 1, "MPE": 1, "MLM": 0, "MCO": 0})

	// Aplica restricciones segun el archivo input
	//restrictions.ApplyRestrictions()

	// Levanta restricciones masivamente segun archivo con IDs de users
	//  - Parametro tipo de restriccion (suspended - warned)
	//restrictions.MassiveRelease("warned")

}
