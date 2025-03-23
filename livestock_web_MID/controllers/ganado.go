package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
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

}

// GetOne ...
// @Title GetOne
// @Description get Ganado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ganado
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GanadoController) GetOne() {
	fmt.Println("Funcion Get")
	// id_ingreso := c.Ctx.Input.Param(":id")
	// 	fmt.Println("El id ingresado es:", id_ingreso)
	// body, _ := services.Metodo_get(  "postgres://postgres:postgres@127.0.0.1/DATOS_VENTAS?sslmode=disable&search_path=public", id_ingreso)
	// var result2 map[string]interface{}
	// err := json.Unmarshal(body, &result2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// resultado1, _ := services.ProcessarJsonArreglos(body)
	// for i := range resultado1 {
	// 	resultado1[i] = map[string]interface{}{
	// 		"id": i + 1,
	// 		"datos_ganado":  resultado1[i]["datos_ganado"],
	// 	}
	// 	fmt.Println("Response Body: ", string(body))
	// 	fmt.Println("Los datos son: ", len(resultado1))
	// 	c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Consulta correctamente", "Consulta de id": id_ingreso}
	// 	c.ServeJSON()
	// 	}

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

}
