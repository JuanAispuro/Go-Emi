package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pocketbase/pocketbase"
)

func mainv() {
	app := pocketbase.New()

	httpposturl := "http://20.186.180.168:9021/ememi-secure-service/emprendedores/registro/crear"

	id := "SazDq6Nlr1jT4CD"
	emprendimiento_record, err := app.Dao().FindFirstRecordByData("emprendimientos", "id", id)
	if err != nil {
		fmt.Println("No hay emprendimientos con ese ID.")
	}

	promotor_fk := emprendimiento_record.Get("id_promotor_fk").(string)

	usuarios_rec, err := app.Dao().FindFirstRecordByData("emi_users", "id", promotor_fk)
	if err != nil {
		fmt.Println("No hay usuarios con ese ID.")
	}
	idUsuario := usuarios_rec.Get("id_emi_web").(string)

	// ID del emprendimeinto de bonita HcFKTHOx8eI75Ee

	emprededor_fk := emprendimiento_record.Get("id_emprendedor_fk").(string)

	emprendedor, err := app.Dao().FindFirstRecordByData("emprendedores", "id", emprededor_fk)
	if err != nil {
		fmt.Println("No hay emprendedores con ese ID.")

	}

	nombreUsuario_fk := usuarios_rec.Get("nombre_usuario").(string)
	nombre_fk := emprendedor.Get("nombre_emprendedor").(string)
	apellidoP_fk := emprendedor.Get("apellido_p").(string)
	apellidoM_fk := emprendedor.Get("apellido_m").(string)
	curp_fk := emprendedor.Get("curp").(string)
	integranteFamilia_fk := emprendedor.Get("integrantes_familia").(string)
	comunidad_fk := emprendedor.Get("comunidad").(string)

	comunidad, err := app.Dao().FindFirstRecordByData("comunidades", "id", comunidad_fk)
	if err != nil {
		fmt.Println("No hay comunidades con ese ID.")

	}

	comunidad_id_emi_web := comunidad.Get("id_emi_web").(string)
	municipio_fk := comunidad.Get("id_municipio_fk").(string)
	municipio, err := app.Dao().FindFirstRecordByData("municipios", "id", municipio_fk)
	if err != nil {
		fmt.Println("No hay municipios con ese ID.")

	}
	municipio_id_emi_web := municipio.Get("id_emi_web").(string)
	estado_fk := municipio.Get("id_estado_fk").(string)
	estado, err := app.Dao().FindFirstRecordByData("estados", "id", estado_fk)
	if err != nil {
		fmt.Println("No hay estados con ese ID.")

	}

	estado_id_emi_web := estado.Get("id_emi_web").(string)
	nombre_emprendimiento := emprendimiento_record.Get("nombre_emprendimiento").(string)
	telefono_fk := emprendedor.Get("telefono").(string)
	comentarios_fk := emprendedor.Get("comentarios").(string)
	fecha_registro_fk := emprendedor.Get("created").(string)

	layout := "2006-01-02T15:04:05"
	fecha_registro_format, err := time.Parse(layout, fecha_registro_fk)
	if err != nil {
		fmt.Println("El parseo fallo.")

	}

	archivado := false

	/*
		var jsonData = []byte(` {
			"idUsuario": %s
			"nombreUsuario": %s
			"nombre" : %s
			"apellidos":
			"curp":
			"integrantesFamilia":
			"comunidad":
			"estado":
			"municipio":
			"emprendimiento":
			"telefono":
			"comentarios":
			"fechaRegistro":
			"archivado":


		}`)
	*/
	jsonData := map[string]interface{}{
		"idUsuario":          idUsuario,
		"nombreUsuario":      nombreUsuario_fk,
		"nombre":             nombre_fk,
		"apellidos":          apellidoP_fk + " " + apellidoM_fk,
		"curp":               curp_fk,
		"integrantesFamilia": integranteFamilia_fk,
		"comunidad":          comunidad_id_emi_web,
		"estado":             estado_id_emi_web,
		"municipio":          municipio_id_emi_web,
		"emprendimiento":     nombre_emprendimiento,
		"telefono":           telefono_fk,
		"comentarios":        comentarios_fk,
		"fechaRegistro":      fecha_registro_format,
		"archivado":          archivado,
	}

	data, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Fallo el cambio a bytes")

	}
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Fallo el posteo.")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer dde0f874-b720-43b5-ac93-cd009bccd04e")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println("Fallo la responsiva.")

	}
	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
}
