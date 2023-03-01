package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	httpposturl := "http://20.186.180.168:9021/ememi-secure-services/api/emprendedores/registro/crear"

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// add new "GET /hello" route to the app router (echo)
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/emi_mobile/emprendimientos/:id", //variable
			Handler: func(c echo.Context) error {
				emprendimiento_record, err := app.Dao().FindFirstRecordByData("emprendimientos", "id", c.PathParam("id"))
				if err != nil {
					return apis.NewNotFoundError(" No hay emprendimientos con ese ID ", err)
				} else {
					fmt.Println(" Se logro  encontrar el emprendimiento con el ID.")
				}

				promotor_fk := emprendimiento_record.Get("id_promotor_fk").(string)

				usuarios_rec, err := app.Dao().FindFirstRecordByData("emi_users", "id", promotor_fk)
				if err != nil {
					return apis.NewNotFoundError(" No hay promotor con este id.", err)
				} else {
					fmt.Println(" Se logro  encontrar el usuario de emi con el ID.")
				}

				idUsuario := usuarios_rec.Get("id_emi_web").(string)

				// ID del emprendimeinto de bonita HcFKTHOx8eI75Ee

				emprededor_fk := emprendimiento_record.Get("id_emprendedor_fk").(string)

				emprendedor, err := app.Dao().FindFirstRecordByData("emprendedores", "id", emprededor_fk)
				if err != nil {
					return apis.NewNotFoundError(" No hay ningun emprededor con este ID.", err)
				} else {
					fmt.Println(" Se logro  encontrar el emprendedor con el ID.")
				}

				nombreUsuario_fk := usuarios_rec.Get("nombre_usuario").(string)
				fmt.Println(" Se logro  encontrar el nombre de usuario.")
				nombre_fk := emprendedor.Get("nombre_emprendedor").(string)
				fmt.Println(" Se logro  encontrar el nombre del emprendedor.")
				apellidoP_fk := usuarios_rec.Get("apellido_p").(string)
				fmt.Println(" Se logro  encontrar el apellido p del usuario.")
				apellidoM_fk := usuarios_rec.Get("apellido_m").(string)
				fmt.Println(" Se logro  encontrar el apellido m del usuario.")
				apellidoEmp := emprendedor.Get("apellidos_emp").(string)
				fmt.Println(" Se logro  encontrar los apellido del emprendedor.")
				curp_fk := emprendedor.Get("curp").(string)
				fmt.Println(" Se logro  encontrar el curp del emprendedor.")
				integranteFamilia_fk := emprendedor.Get("integrantes_familia").(float64)
				fmt.Println(" Se logro  encontrar el número de integrantes del emprendedor.")
				comunidad_fk := emprendedor.Get("id_comunidad_fk").(string)
				fmt.Println(" Se logro  encontrar el id_comunidad_fk del emprendedor  .")

				comunidad, err := app.Dao().FindFirstRecordByData("comunidades", "id", comunidad_fk)
				if err != nil {
					return apis.NewNotFoundError(" No hay ninguna comunidad con este ID.", err)
				} else {
					fmt.Println(" Se logro  encontrar la comunidad con el ID.")
				}

				comunidad_id_emi_web := comunidad.Get("id_emi_web").(string)
				municipio_fk := comunidad.Get("id_municipio_fk").(string)
				municipio, err := app.Dao().FindFirstRecordByData("municipios", "id", municipio_fk)
				if err != nil {
					return apis.NewNotFoundError(" No hay ningun municipio con este ID.", err)
				} else {
					fmt.Println(" Se logro  encontrar el municipio con el ID.")
				}
				municipio_id_emi_web := municipio.Get("id_emi_web").(string)
				estado_fk := municipio.Get("id_estado_fk").(string)
				estado, err := app.Dao().FindFirstRecordByData("estados", "id", estado_fk)
				if err != nil {
					return apis.NewNotFoundError(" No hay ningun estado con este ID.", err)
				} else {
					fmt.Println(" Se logro  encontrar el estado con el ID.")
				}

				estado_id_emi_web := estado.Get("id_emi_web").(string)
				nombre_emprendimiento := emprendimiento_record.Get("nombre_emprendimiento").(string)
				telefono_fk := emprendedor.Get("telefono").(string)
				comentarios_fk := emprendedor.Get("comentarios").(string)
				//fecha_registro_fk := emprendedor.Get("created").(types.DateTime).String()
				fecha_registro_fk := time.Now().Format("2006-01-02T15:04:05")

				//layout := "2006-01-02T15:04:05"
				/*
					fecha_registro_format, err := time.Parse(layout, fecha_registro_fk)
					if err != nil {
						return apis.NewNotFoundError(" El parseo de la fecha fallo.", err)
					} else {
						fmt.Println(" Se logro parsear la fecha.")
					}
				*/

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
					"nombreUsuario":      nombreUsuario_fk + " " + apellidoP_fk + " " + apellidoM_fk,
					"nombre":             nombre_fk,
					"apellidos":          apellidoEmp,
					"curp":               curp_fk,
					"integrantesFamilia": integranteFamilia_fk,
					"comunidad":          comunidad_id_emi_web,
					"estado":             estado_id_emi_web,
					"municipio":          municipio_id_emi_web,
					"emprendimiento":     nombre_emprendimiento,
					"telefono":           telefono_fk,
					"comentarios":        comentarios_fk,
					"fechaRegistro":      fecha_registro_fk,
					"archivado":          archivado,
				}

				infoEmprendimiento := map[string]interface{}{
					"emprendimiento": emprendimiento_record,
				}

				infoEmprendedor := map[string]interface{}{
					"emprendedor": emprendedor,
				}

				InfoTotal := map[string]interface{}{
					"info_emprendimiento": infoEmprendimiento,
					"info_emprendedor":    infoEmprendedor,
					"jsonData":            jsonData,
				}

				data, err := json.Marshal(jsonData)
				if err != nil {
					return apis.NewNotFoundError(" La conversion a bytes fallo", err)
				} else {
					fmt.Println(" Se logro convertir a bytes.")
				}
				request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(data))
				if error != nil {
					panic(error)
				} else {
					fmt.Println(" Se logro Postear.")
				}
				request.Header.Set("Content-Type", "application/json")
				request.Header.Set("Authorization", "Bearer eb5157cc-8569-46ff-bce4-ae5ac6db9617")
				client := &http.Client{}
				response, error := client.Do(request)
				if error != nil {
					panic(error)
				} else {
					fmt.Println(" Se obtuvo el response.")
				}
				defer response.Body.Close()
				fmt.Println("response Status:", response.Status)
				body, _ := ioutil.ReadAll(response.Body)
				fmt.Println("response Body:", string(body))
				return c.JSON(http.StatusOK, InfoTotal)
				//Todo ok
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
