package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
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
