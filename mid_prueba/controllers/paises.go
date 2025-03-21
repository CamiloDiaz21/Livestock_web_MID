package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/mid_prueba/services"
)

// PaisesController operations for Paises
type PaisesController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaisesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Paises
// @Param	body		body 	models.Paises	true		"body for Paises content"
// @Success 201 {object} models.Paises
// @Failure 403 body is empty
// @router / [post]
func (c *PaisesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Paises by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Paises
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PaisesController) GetOne() {
	fmt.Println("Funcion Get")

	pais_ingreso := c.Ctx.Input.Param(":subregion")
	fmt.Println("La subregion que ingreso es:", pais_ingreso)

	body, _ := services.Metodo_get("servicio_paises", pais_ingreso)
	resultado, _ := services.ProcessarJsonArreglos(body)

	for i := range resultado {
		resultado[i] = map[string]interface{}{
			"Nombre":    resultado[i]["name"].(map[string]interface{})["common"],
			"Fifa":      resultado[i]["fifa"],
			"Fronteras": resultado[i]["borders"],
			"Poblacion": resultado[i]["population"],
		}
	}
	c.Data["json"] = map[string]interface{}{
		"Succes":          true,
		"Status":          200,
		"Message":         "Consulta existosa",
		"Data":            resultado,
		"Cantidad paises": len(resultado),
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Paises
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Paises
// @Failure 403
// @router / [get]
func (c *PaisesController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Paises
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Paises	true		"body for Paises content"
// @Success 200 {object} models.Paises
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PaisesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Paises
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PaisesController) Delete() {

}

// GetOne ...
// @Title GetOnev2
// @Description get Paises by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	id_2	path 	string	true		"este es el segundo parametro"
// @Param	id_3 	path 	string	true		"este es el tercer parametro"
// @Param	id_4	path 	string	true		"este es el cuarto parametro"
// @Param	id_5	path 	string	true	    "este es el quinto parametro"
// @Success 200 {object} models.Paises
// @Failure 403 :id is empty
// @router /capitales/:id/:id_2/:id_3/:id_4/:id_5 [get]
func (c *PaisesController) GetOne2() {

	fmt.Println("Funcion Get")

	capitales := []string{
		c.Ctx.Input.Param(":id"), c.Ctx.Input.Param(":id_2"),
		c.Ctx.Input.Param(":id_3"), c.Ctx.Input.Param(":id_4"),
		c.Ctx.Input.Param(":id_5"),
	}

	var resultadof []map[string]interface{}

	for _, capital := range capitales {
		body, _ := services.Metodo_get("servicio_capitales", capital)
		var result []map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatal(err)
		}

		resultado, _ := services.ProcessarJsonArreglos(body)
		for _, item := range resultado {
			monedas := item["currencies"].(map[string]interface{})
			conversion := make(map[string]float64)
			for codigo := range monedas {
				tasa := tasa_cambio(codigo)
				conversion[codigo] = tasa
			}

			resultadof = append(resultadof, map[string]interface{}{
				"Nombre":        item["name"].(map[string]interface{})["official"],
				"Moneda":        monedas,
				"ConversionCOP": conversion,
				"region":        item["region"],
				"subregion":     item["subregion"],
			})
		}
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Consulta existosa",
		"Data":    resultadof,
	}
	c.ServeJSON()
}

func tasa_cambio(codigo string) float64 {
	// Simulación de tasas de cambio (se debe reemplazar con una API real)
	tasas := map[string]float64{
		"USD": 3900.0,
		"MXN": 210.0,
		"JPY": 28.0,
		"RUB": 0.020,
		"COP": 1,
	}
	if tasa, existe := tasas[codigo]; existe {
		return tasa
	}
	return 0.0
}
