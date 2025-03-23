package controllers

import (
	"github.com/astaxie/beego"
)

// Hacienda_lotesController operations for Hacienda_lotes
type Hacienda_lotesController struct {
	beego.Controller
}

// URLMapping ...
func (c *Hacienda_lotesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Hacienda_lotes
// @Param	body		body 	models.Hacienda_lotes	true		"body for Hacienda_lotes content"
// @Success 201 {object} models.Hacienda_lotes
// @Failure 403 body is empty
// @router / [post]
func (c *Hacienda_lotesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Hacienda_lotes by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Hacienda_lotes
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Hacienda_lotesController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Hacienda_lotes
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Hacienda_lotes
// @Failure 403
// @router / [get]
func (c *Hacienda_lotesController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Hacienda_lotes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Hacienda_lotes	true		"body for Hacienda_lotes content"
// @Success 200 {object} models.Hacienda_lotes
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Hacienda_lotesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Hacienda_lotes
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Hacienda_lotesController) Delete() {

}
