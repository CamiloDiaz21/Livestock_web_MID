package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/models"
	"github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/services"
)

// GanadoController operations for Ganado
type GanadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *GanadoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Ganado
// @Param	body		body 	models.Ganado	true		"body for Ganado content"
// @Success 201 {object} models.Ganado
// @Failure 403 body is empty
// @router / [post]
func (c *GanadoController) Post() {
	var body_ingresa []map[string]interface{}
	var alerta models.Alert
	var temporal []byte
	var temporal_producto []byte

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("body que ingresa: ", body_ingresa)

		jsonData, err := json.MarshalIndent(body_ingresa, "", " ")
		if err != nil {
			fmt.Println("Error al convertir Json", err)
		}

		json_usuario := body_ingresa[0]
		json_producto := body_ingresa[1]
		fmt.Println("Body usuario:", json_usuario)
		fmt.Println("Body producto:", json_producto)

		json_usuario_byte, err := json.Marshal(json_usuario)
		if err != nil {
			fmt.Println("Error al convertir usuario a JSON:", err)
			return
		}
		response_usuario, err := services.Metodo_post("servicio_post", json_usuario_byte)
		if err != nil {
			fmt.Println("Error en Metodo_post para usuario:", err)
			return
		}

		json_producto_byte, err := json.Marshal(json_producto)
		if err != nil {
			fmt.Println("Error al convertir producto a JSON:", err)
			return
		}
		response_producto, err := services.Metodo_post("servicio_post", json_producto_byte)
		if err != nil {
			fmt.Println("Error en Metodo_post para producto:", err)
			return
		}

		temporal = response_usuario
		temporal_producto = response_producto

		fmt.Println("Esto responde el post de producto:", string(response_producto))
		fmt.Println("Esto responde el post de usuario", string(response_usuario))
		fmt.Println("Body de ingreso en Json:", string(jsonData))
	}

	var temporal2 map[string]interface{}
	var temporal3 map[string]interface{}

	if err := json.Unmarshal(temporal, &temporal2); err != nil {
		fmt.Println("Error al deserializar response_usuario:", err)
		return
	}
	if err := json.Unmarshal(temporal_producto, &temporal3); err != nil {
		fmt.Println("Error al deserializar response_producto:", err)
		return
	}

	alerta.Code = "201"
	alerta.Type = "Post"
	c.Data["json"] = alerta
	c.ServeJSON()
}


// GetOne ...
// @Title GetOne
// @Description get Ganado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	id_2		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ganado
// @Failure 403 :id is empty
// @router /:id/:id_2 [get]
func (c *GanadoController) GetOne() {
	fmt.Println("Funcion Get")
	id_ingreso:= c.Ctx.Input.Param(":id")
	//fmt.Println("este es el segundo parametro:", id_ingreso2)

	// Obtener datos del servicio
	body, _ := services.Metodo_get("Variable_api_Ganado",id_ingreso)

	fmt.Println("body:", string(body))


	fmt.Println("body que ingresa: ", string(body))

	// Procesar los datos JSON
	resultado1, err := services.ProcessarJsonArreglos(body)
	if err != nil {
		fmt.Println("Error al procesar los datos:", err)
		c.Data["json"] = map[string]interface{}{
			"Succes":  false,
			"Status":  500,
			"Message": err.Error(),
		}
		c.ServeJSON()
		return
	}

	fmt.Println("este es el resultado: ", resultado1)

	for i := range resultado1 {
		resultado1[i] = map[string]interface{}{
			"Datos":     resultado1[i]["DatosGanado"],
			"Vendedor":  resultado1[i]["DatosVendedor"],
			"categoria": resultado1[i]["CategoriaGanado"],
			"Ubicacion": resultado1[i]["UbicacionGanando"],
		}
	}

	fmt.Println("Los datos son:", len(resultado1))

	c.Data["json"] = map[string]interface{}{
		"Succes":            true,
		"Status":            200,
		"Message":           "Consulta exitosa",
		"Cantidad de datos": len(resultado1),
		"Datos":             body,
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Ganado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Ganado
// @Failure 403
// @router / [get]
func (c *GanadoController) GetAll() {
	fmt.Println("Funcion GetAll")
    
    // Llamada a un servicio (puedes cambiar "Variable_api_Ganado" por el nombre adecuado)
    body, err := services.Metodo_get("Variable_api_Ganado", "")
    if err != nil {
        fmt.Println("Error al obtener los datos:", err)
        c.Data["json"] = map[string]interface{}{
            "Succes":  false,
            "Status":  500,
            "Message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    fmt.Println("body que ingresa: ", string(body))

    // Procesar los datos obtenidos
    resultado, err := services.ProcessarJsonArreglos(body)
    if err != nil {
        fmt.Println("Error al procesar los datos:", err)
        c.Data["json"] = map[string]interface{}{
            "Succes":  false,
            "Status":  500,
            "Message": err.Error(),
        }
        c.ServeJSON()
        return
    }

    fmt.Println("este es el resultado:", resultado)

    // Procesar la respuesta para devolver solo los datos necesarios
    for i := range resultado {
        resultado[i] = map[string]interface{}{
            "Datos":     resultado[i]["DatosGanado"],
            "Vendedor":  resultado[i]["DatosVendedor"],
            "categoria": resultado[i]["CategoriaGanado"],
            "Ubicacion": resultado[i]["UbicacionGanando"],
        }
    }

    // Devolver los datos procesados
    c.Data["json"] = map[string]interface{}{
        "Succes":            true,
        "Status":            200,
        "Message":           "Consulta exitosa",
        "Cantidad de datos": len(resultado),
        "Datos":             resultado,
    }

    c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Ganado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Ganado	true		"body for Ganado content"
// @Success 200 {object} models.Ganado
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GanadoController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Ganado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GanadoController) Delete() {
	fmt.Println("Funcion Delete")
	id_ingreso := c.Ctx.Input.Param(":id") // para capturar el parametro del url /id

	// Llamamos al servicio para eliminar el recurso con el ID especificado
	// Aquí deberías incluir la lógica para eliminar el recurso
	// (por ejemplo, una llamada a un servicio o base de datos para eliminar el dato)

	err := services.EliminarGanado(id_ingreso) // Ejemplo de función para eliminar el ganado
	if err != nil {
		// Si ocurre un error durante la eliminación
		c.Data["json"] = map[string]interface{}{
			"Succes":  false,
			"Status":  500,
			"Message": "Error al eliminar el recurso",
		}
		c.ServeJSON()
		return
	}

	// Respuesta de éxito
	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Recurso eliminado correctamente",
	}
	c.ServeJSON()
}
