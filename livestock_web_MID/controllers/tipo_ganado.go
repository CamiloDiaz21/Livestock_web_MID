package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/services"
)

// Tipo_ganadoController operations for Tipo_ganado
type Tipo_ganadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *Tipo_ganadoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Tipo_ganado
// @Param	body		body 	models.Tipo_ganado	true		"body for Tipo_ganado content"
// @Success 201 {object} models.Tipo_ganado
// @Failure 403 body is empty
// @router / [post]
func (c *Tipo_ganadoController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Tipo_ganado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tipo_ganado
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Tipo_ganadoController) GetOne() {
	fmt.Println("Funcion Get")
	id_ingreso := c.Ctx.Input.Param(":id") // para capturar el parametro del url /id

	//asignacion de datos al body
	body, _ := services.Metodo_get("Variable_api_TG", id_ingreso)

	resultado, _ := services.ProcessarJson(body)
	//----------------------------------------------------------------------------------------

	var result2 map[string]interface{} // El JSON que esperas es un array de objetos

	result2 = map[string]interface{}{
		"Raza":     result2["NombreTipoGanado"],

	}

	//------------------------------------------------

	//informacion de estado
	fmt.Println("Los datos son:", len(result2))

	c.Data["json"] = map[string]interface{}{
		"Succes":          true,
		"Status":          200,
		"Message":         "Consulta existosa",
		"Cantidad Cartas": len(resultado),
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Tipo_ganado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Tipo_ganado
// @Failure 403
// @router / [get]
func (c *Tipo_ganadoController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Tipo_ganado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tipo_ganado	true		"body for Tipo_ganado content"
// @Success 200 {object} models.Tipo_ganado
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Tipo_ganadoController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Tipo_ganado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Tipo_ganadoController) Delete() {

}
