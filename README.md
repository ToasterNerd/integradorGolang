# integradorGolang
Este proyecto es el integrador resultado de la primera parte de un curso de golang APIRest. Este modulo genera un avatar mediante un input. Este input es hasheado y mediante algún algoritmo se le asigna dicho avatar. El avatar es una imagen .jpg que se extrae de una API.

## Uso del modulo
En la terminal se su IDE deben colocar la siguiente línea para obtener la ultima versión del módlo
```
go get github.com/ToasterNerd/integradorGolang/avatar@latest
```



Basta con darle un valor por teclado, en este caso, mediante scanner, y el código se encargará de generar el hash, asignar el jpg, y descargarlo en la carpeta raiz del proyecto.
```golang
package main

import (
	"fmt"

	pkgAvatar "github.com/ToasterNerd/integradorGolang/avatar"
)

func main() {
	defer func() {
		fmt.Print("EJECUCION FINALIZADA")
	}()

	var dameEmail string
	fmt.Println("Ingrese Email que desea hashear")
	fmt.Scan(&dameEmail)

	err := pkgAvatar.ServiceGenerator().GenerateAndSaveAvatar(pkgAvatar.Information{Email: dameEmail})

	if err != nil {
		//log.Fatal(err)
		fmt.Println("error imprimiendo: ", err)
	}

}
```
## Version History
* 0.1.0
    * Initial Release
