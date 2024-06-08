package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    aemetAPIUrl = "https://opendata.aemet.es/api/valores/climatologicos/diarios/datos/fechaini/%s/fechafin/%s/estacion/%s"
    apiKey      = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJhbGVqYW5kcm9waW50b3NkZXZAZ21haWwuY29tIiwianRpIjoiY2FkMTQxNGQtNTA2MS00Njk2LThkMDUtZjNkOTRmODE2MmY3IiwiaXNzIjoiQUVNRVQiLCJpYXQiOjE3MTc1ODkzOTMsInVzZXJJZCI6ImNhZDE0MTRkLTUwNjEtNDY5Ni04ZDA1LWYzZDk0ZjgxNjJmNyIsInJvbGUiOiIifQ.Fnm-VEZpzWftlQCe87ZOtsUHsTurizq3PRkzK2PrYAw"
)

func main() {
    fechaIniStr := "2024-01-01"
    fechaFinStr := "2024-01-31"
    idema := "ID_ESTACION_METEOROLOGICA"

    url := fmt.Sprintf(aemetAPIUrl, fechaIniStr, fechaFinStr, idema)

    // Hacer la solicitud HTTP a la API de AEMET
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creando la solicitud HTTP:", err)
        return
    }
    req.Header.Set("api_key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error haciendo la solicitud HTTP:", err)
        return
    }
    defer resp.Body.Close()

    // Leer el cuerpo de la respuesta
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error leyendo el cuerpo de la respuesta:", err)
        return
    }

    // Analizar el cuerpo de la respuesta JSON
    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("Error analizando el JSON de la respuesta:", err)
        return
    }

    // Verificar si la respuesta contiene una URL de datos y hacer otra solicitud para obtener los datos reales
    if urlDatos, ok := data["datos"].(string); ok {
        // Hacer una segunda solicitud HTTP para obtener los datos reales
        reqDatos, err := http.NewRequest("GET", urlDatos, nil)
        if err != nil {
            fmt.Println("Error creando la solicitud HTTP para los datos:", err)
            return
        }
        reqDatos.Header.Set("api_key", apiKey)

        respDatos, err := client.Do(reqDatos)
        if err != nil {
            fmt.Println("Error haciendo la solicitud HTTP para los datos:", err)
            return
        }
        defer respDatos.Body.Close()

        // Leer el cuerpo de la respuesta con los datos reales
        bodyDatos, err := ioutil.ReadAll(respDatos.Body)
        if err != nil {
            fmt.Println("Error leyendo el cuerpo de la respuesta para los datos:", err)
            return
        }

        // Aquí puedes procesar los datos reales según sea necesario
        fmt.Println("Datos climatológicos diarios obtenidos:")
        fmt.Println(string(bodyDatos))
    } else {
        fmt.Println("La respuesta no contiene una URL de datos válida.")
    }
}

