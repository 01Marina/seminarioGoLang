package main

import (
	"bytes"
	"errors"
	"strconv"

	"seminarioGoLang2021.com/model"
)

func IsChar(charByte []byte) bool {
	return bytes.ContainsAny(charByte, "ABCDEFGHIJKLMNOPKRSTUVWXYZabcdefghijklmnñopqrstuvwxyz")
}
func IsNumber(charByte []byte) bool {
	return bytes.ContainsAny(charByte, "0123456789")
}
func IsAllString(s string) bool {
	ss := []byte(s)
	Ischar := true
	arrayCharB := bytes.Split(ss, []byte(""))
	for i := 0; i < len(s); i++ {
		if !IsChar(arrayCharB[i]) {
			Ischar = false
			break
		}
	}
	return Ischar
}

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

func rightLenght(length int, value string) bool {
	return length == len(value)
}

func setValuesResult(r *model.Result) {
	r.Length = 0
	r.Type = ""
	r.Value = ""
}

func checkInputStatus(input []byte) bool {
	if len(input) > 3 {
		return true
	} else {
		return false
	}
}

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
