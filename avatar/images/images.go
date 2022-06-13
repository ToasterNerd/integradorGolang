package images

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

//struct publica porque se va a llamar desde service.
type GeneratorImageStruct struct {
}

//metodo publico porque se va a llamar desde service para "engancharlo" desde la interfaz"
func (m *GeneratorImageStruct) BuildAndSaveImage(encodedInfo uint64) error {

	//voy a la api de memes y traigo el id de los memes
	response, err := http.Get("https://api.imgflip.com/get_memes")
	if err != nil {
		//log.Fatal(err)

		//panic(err)
		panic("NO TIENE CONEXION, POR FAVOR VUELVA A CONECTARSE A INTERNET")

	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// log.Fatal(err)
		//panic(err)
		panic("NO PUDO LEER EL BODY")
	}

	var info Info
	err = json.Unmarshal([]byte(responseData), &info)
	if err != nil {
		// fmt.Println(err)

		//panic(err)
		panic("NO PUDO PARSEAR")
	}

	//guardo todos los memes en un map.
	mapita := make(map[int]string)
	for i := 0; i < len(info.Data.Memes)-1; i++ {

		mapita[i] = info.Data.Memes[i].Name
	}

	//traigo el numero codeado, sumo sus ultimo 7 digitos y a ese resultado le sumo sus digitos
	//ese resultado final lo comparo con el map, y si son igual, le asigno ese meme
	var numEncodeado uint64 = encodedInfo

	//sumo todos los digitos que tiene
	var cont uint64
	var suma uint64 = 0
	for cont = 1; cont < 7; cont++ {
		sumando := digit(uint64(numEncodeado), cont)
		suma = suma + sumando
	}
	sumandoDenuevo := digit(uint64(suma), 1)
	sumandoDenuevo = sumandoDenuevo + digit(uint64(suma), 2)
	fmt.Println("el numero final del encodeado es ", sumandoDenuevo)

	for i := 0; i < len(mapita)-1; i++ {
		if sumandoDenuevo == uint64(i) {
			fmt.Println("el personaje asignado para este usuario es: ", mapita[i])

			fileName := "sample.jpg"
			URL := info.Data.Memes[i].Url
			err := downloadFile(URL, fileName)
			if err != nil {
				// log.Fatal(err)
				panic(err)

			}
			fmt.Printf("File %s downlaod in current working directory", fileName)
		}
	}

	return nil
}

//estructura que define el json que se recibe por GET
type Info struct {
	//Success string `json:"success"`
	Data struct {
		Memes []struct {
			Box_count uint64 `json:"box_count"`
			Height    uint64 `json:"height"`
			Id        string `json:"id"`
			Name      string `json:"name"`
			Url       string `json:"url"`
		} `json:"memes"`
	} `json:"data"`
}

//funcion para extraer los digitos del numero
func digit(num, place uint64) uint64 {
	r := num % uint64(math.Pow(10, float64(place)))
	return r / uint64(math.Pow(10, float64(place-1)))
}

//funcion que descargar un archivo
func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
