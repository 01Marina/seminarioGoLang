# Trabajo Práctico
### Seminario GoLang 2021
Marina Caseres
## Consigna
Crear una función que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena.
[Trabajo Practico-TUDAI 2021](https://github.com/juanpablopizarro/tudai2021#trabajo-practico)

## Ejecutar resolución
Para ejecutar la resolución y comprobar su funcionamiento se debe correr desde el test
#### Se recomiendan los siguientes comandos por consola
 * go test
 * go test ./...
 * go test ./... -count=1
 * go test ./... -coverprofile=out.test
 * go tool cover -html out -o out.html
 #### cobertura de testeo: 100.0%
## Resolución
La primer función que se debe ver es la siguiente:
#### func ParseString(input []byte)
```
func ParseString(input []byte) (model.Result, error) {
	if checkInputStatus(input) {
		result, err := separateString(input)
		if err != nil {
			setValuesResult(&result)
			return result, errors.New("Los datos no son coherentes")
		}
		if rightLenght(result.Length, result.Value) {
			if result.Type == "TX" {
				if len(result.Value) > 0 {
					if IsAllString(result.Value) {
						return result, nil
					} else {
						setValuesResult(&result)
						return result, errors.New("Los valores no son todas letras")
					}
				} else {
					return result, nil
				}
			} else if result.Type == "NN" {
				if len(result.Value) > 0 {
					if IsAllNumber(result.Value) {
						return result, nil
					} else {
						setValuesResult(&result)
						return result, errors.New("Los valores no son todos numeros")
					}
				} else {
					return result, nil
				}
			} else {
				setValuesResult(&result)
				return result, errors.New("El typo de valores no es valido")
			}
		} else {
			setValuesResult(&result)
			return result, errors.New("El lenght de los valores no coincide")
		}
	} else {
		return model.NewResult("", "", 0), errors.New("Los datos no son coherentes")
	}
}
```
En un principio lo que hace es chequear que el input que recibe tenga como mínimo cuatro caracteres, los dos primeros para el Type y dos siguientes para el Lenght. Ya que se considera que son el mínimo de datos que se necesitan para comenzar con el análisis. Para esto se emplea la función **func checkInputStatus(input []byte)**
````
func checkInputStatus(input []byte) bool {
	if len(input) > 3 {
		return true
	} else {
		return false
	}
}
````

Si cumpe con el minimo, se procede a crear la estructura resultante de este input.
````
func ParseString(input []byte) (model.Result, error) {
	if checkInputStatus(input) {
		result, err := separateString(input)
		...
	} else {
		return model.NewResult("", "", 0), errors.New("Los datos no son coherentes")
	}
````
#### func separateString(input)
La función empleada para esto es la que se muestra a continuación
````
func separateString(input []byte) (model.Result, error) {
	data := bytes.Split(input, []byte(""))
	var Value string
	var s string

	Type := string(data[0]) + string(data[1])

	if len(data) >= 4 {
		for i := 4; i < len(data); i++ {
			Value += string(data[i])
		}
	}

	if bytes.Equal(data[2], []byte("0")) {
	  	s = string(data[3])
	} else {
		  s = string(data[2]) + string(data[3])
	}
	Length, err := strconv.Atoi(s)
	if err != nil {
		Result := model.NewResult("", "", 0)
		return Result, errors.New("Lenght distinto a numeros")
	} else {
		Result := model.NewResult(Type, Value, Length)
		return Result, nil
	}
}
````
Esta recibe un input del tipo byte y retorna una estructura del tipo Result. Al input lo separa en secciones ya que en las pautas especifica que determinadas posiciones del string corresponden a un valor específico (las dos primeras letras son el Type, las dos segundas son el Lenght y el resto es el Value)

De esta forma se hace la conversión de byte a string, ya que la estructura que se debe generar tiene predefinido que el Type y el Value son de tipo string y el Length es de tipo int.

La conversión de las posiciones determinadas para el Type de la estructura se hace en la siguiente línea:
````
Type := string(data[0]) + string(data[1])
````
Para el Value de la estructura primero se hace un chequeo si este existen las posiciones donde se ubicaría este, si el input no cuenta con estas posiciones, no se hace nada, se deja Value con su valor por defecto, que vacío, si este si se encuentra en el input, se genera un string a partir de la cuarta posición del byte.

````
if len(data) >= 4 {
		for i := 4; i < len(data); i++ {
			Value += string(data[i])
		}
	}
````
Para el Length primero se tiene en cuenta si la primera posición del espacio destinado a este se encuentra en cero, si es así se lo descarta, si no se encuentra en cero, se lo toma en consideración.

Se lo convierte a string y luego a int, esto se hace así ya que se consideró la forma más sencilla de convertir un byte a int.
````
if bytes.Equal(data[2], []byte("0")) {
    s = string(data[3])
} else {
   s = string(data[2]) + string(data[3])
}
Length, err := strconv.Atoi(s)
  ````
Luego se captura un posible error al momento de la conversión de string a int, esto puede suceder en caso que el string contenga signos distintos a números.

Si todo resulta sin errores se instancia el resultado y se lo retorna
````
if err != nil {
    Result := model.NewResult("", "", 0)
    return Result, errors.New("Lenght distinto a numeros")
} else {
    Result := model.NewResult(Type, Value, Length)
    return Result, nil
}
````
#### func ParseString(input []byte)
Continuando con la función principal

Una vez instanciada la estructura y guardada en una variable(result), comienza con las validaciones de esta.

Primero captuara un posible error, proveniente de la conversión de string a int que se hace con el Length en la función anterior. Si encuentra el error setea la estructura vaciando sus campos, y retorna esta vacia mas un error.

````
if err != nil {
  setValuesResult(&result)
  return result, errors.New("Los datos no son coherentes")
}
````
Si la ejecución continúa, se hace una validación si el Length de la estructura coincide con el length real del campo Value de esta misma.
En caso de que no coincidan setea la result,vaciando sus campos, la retorna acompañada de un error.
````
if rightLenght(result.Length, result.Value) {
    ...
} else {
	setValuesResult(&result)
	return result, errors.New("El lenght de los valores no coincide")
}
````
Para comparar los tamaños se ejecuta la función **rightLenght(length, value)**
````
func rightLenght(length int, value string) bool {
	return length == len(value)
}
````
En caso de que los tamaños si coinciden, se continua por validar el Type de result. Contemplando los dos casos ("TX", "NN", más el caso de que no sea ninguno de estos).
Para el caso de que no coincida con ninguno de los Type posibles, se vacía result y se lo retorna junto con un error.
````
if rightLenght(result.Length, result.Value) {
  if result.Type == "TX" {
    ...
  } else if result.Type == "NN" {
  ....
  } else {
    setValuesResult(&result)
    return result, errors.New("El typo de valores no es valido")
  }
} else {
...
}

````
Si el Type coincide con "TX", se hace la validación si es que todos los caracteres son letras. En el caso de que no lo sean, se retorna result vacío y un error.

````
if result.Type == "TX" {
	if IsAllString(result.Value) {
	return result, nil
} else {
	setValuesResult(&result)
	return result, errors.New("Los valores no son todas letras")
}
````
Para la validación si son todos caracteres se ejecuta la siguiente función **IsAllString(s string)** .

En esta lo primero que se hace es convertir el string a byte, ya que en formato byte se encontró una función que valida si un byte determinado coincide con algún carácter de una lista.
Al resultado de la conversión se lo separa y se lo recorre. Por cada posición se valida si no es un carácter, si no es un carácter letra se corta el recorrido y se guarda el valor de falso en una variable que luego es retornada. En caso de que todas sean caracteres Letra el for se recorre completo y la variable que se retorna, en un principio inicializada en true, se retorna true.
````
func IsAllString(s string) bool {
	ss := []byte(s)
	Ischar := true
	arrayCharB := bytes.Split(ss, []byte(""))
	for i := 0; i < len(ss); i++ {
		if !IsChar(arrayCharB[i]) {
			Ischar = false
			break
		}
	}
	return Ischar
}
````
Dentro del recorrido mencionado en la función anterior se hace uso de la función **Ischar(arrayCharB[i])**. 
Esta retorna el resultado de la función **ContainsAny([]byte, string)** la cual recibe un byte y lo compara con el string que se le pase por parámetro y devuelve true en caso que coincida con algunos de los caracteres del string, retorna false si no coincide.
````
func IsChar(charByte []byte) bool {
	return bytes.ContainsAny(charByte, "ABCDEFGHIJKLMNOPKRSTUVWXYZabcdefghijklmnñopqrstuvwxyz")
}
````
En caso de que Type de result coincida con "NN", se sigue el mismo proceso de validación que para "TX", lo único que cambia es la función **IsChar(arrayCharB[i])** por **!IsNumber(arrayNumberB[i])**
````
func IsAllNumber(s string) bool {
	ss := []byte(s)
	Isnumber := true
	arrayNumberB := bytes.Split(ss, []byte(""))
	for i := 0; i < len(ss); i++ {
		if !IsNumber(arrayNumberB[i]) {
			Isnumber = false
			break
		}
	}
	return Isnumber
}
````
La función  **IsNumber(arrayNumberB[i])** implementa una lógica similar a  **IsChar(arrayNumberB[i])**
````
func IsNumber(charByte []byte) bool {
	return bytes.ContainsAny(charByte, "0123456789")
}
````
#### func ParseString(input []byte)
Volviendo a la función principal la validación de los Type si estos resultan coincidir se retorna result y ningún error, en caso de que no coincidan se retorna result vacío y un error. Esto se plantea de la siguiente manera:
````
if result.Type == "TX" {
  if IsAllString(result.Value) {
    return result, nil
  } else {
    setValuesResult(&result)
    return result, errors.New("Los valores no son todas letras")
  }
} else if result.Type == "NN" {
  if IsAllNumber(result.Value) {
    return result, nil
  } else {
    setValuesResult(&result)
    return result, errors.New("Los valores no son todos numeros")
  }
} else {
  setValuesResult(&result)
  return result, errors.New("El typo de valores no es valido")
}
````
#### Struct Result

Como se viene mencionando para la resolución del ejercicio se hizo uso de una estructura la cual contiene lo siguiente:

````
type Result struct {
	Type   string
	Value  string
	Length int
}

func NewResult(Type string, Value string, Length int) Result {
	return Result{Type, Value, Length}
}
````
