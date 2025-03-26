package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

func Metodo_get(nombre_servicio, parametro string) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio) + parametro + parametro
	resp, err := http.Get(url)
	fmt.Println("esta es la url:", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, fmt.Errorf("URL del servicio no encontrada en la configuración")
}

func ProcessarJsonArreglos(datos []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := json.Unmarshal(datos, &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func Metodo_post(nombre_servicio string, data []byte) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ProcessarJson(datos []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(datos, &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func Metodo_put(nombre_servicio string, id string, data []byte) ([]byte, error) {

	baseURL := beego.AppConfig.String(nombre_servicio)

	url := fmt.Sprintf("%s/%s", baseURL, id)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ObtenerJSONDesdeAPI(url string) ([]byte, error) {
	// Hacer la petición HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	// Verificar si la respuesta está vacía
	if len(body) == 0 {
		return nil, fmt.Errorf("respuesta vacía del servidor")
	}

	return body, nil
}

func main() {

	// Obtener los datos JSON
	jsonData, err := ObtenerJSONDesdeAPI("localhost:8083/v1/ganado")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Imprimir los datos en formato JSON
	fmt.Println("Datos JSON obtenidos:", string(jsonData))
}
func EliminarGanado(id string) error {
	// Definir la consulta SQL para eliminar el recurso con el ID especificado
	query := "DELETE FROM ganado WHERE id = ?"

	// Ejecutar la consulta
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta de eliminación: %v", err)
	}

	// Verificar cuántas filas fueron afectadas (debería ser 1 si el recurso fue eliminado correctamente)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		// Si no se encontró el ID en la base de datos
		return fmt.Errorf("ningún registro encontrado con el ID proporcionado: %v", id)
	}

	// Retornar nil si la eliminación fue exitosa
	return nil
}
