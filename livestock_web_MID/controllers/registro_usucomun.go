package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/models"
	"github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/services"
)

// Registro_usucomunController operations for Registro_usucomun
type Registro_usucomunController struct {
	beego.Controller
}

// URLMapping ...
func (c *Registro_usucomunController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Registro_usucomun
// @Param	body		body 	models.Registro_usucomun	true		"body for Registro_usucomun content"
// @Success 201 {object} models.Registro_usucomun
// @Failure 403 body is empty
// @router / [post]
func (c *Registro_usucomunController) Post() {
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
// @Description get Registro_usucomun by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Registro_usucomun
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Registro_usucomunController) GetOne() {
	fmt.Println("este es el get")
}

// GetAll ...
// @Title GetAll
// @Description get Registro_usucomun
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Registro_usucomun
// @Failure 403
// @router / [get]
func (c *Registro_usucomunController) GetAll() {
	fmt.Println("este es el get")

}

// Put ...
// @Title Put
// @Description update the Registro_usucomun
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Registro_usucomun	true		"body for Registro_usucomun content"
// @Success 200 {object} models.Registro_usucomun
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Registro_usucomunController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Registro_usucomun
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Registro_usucomunController) Delete() {

}
