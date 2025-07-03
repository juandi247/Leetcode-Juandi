package main

import (
	"fmt"
	"time"
)

/*
Crea una aplicación que reciba un archivo (puede ser simulado con una cadena o estructura),
y permita comprimirlo utilizando distintos algoritmos (Zip, Gzip, Tar, etc.).
Usa el patrón Strategy para cambiar el algoritmo de compresión en tiempo de ejecución.
*/

type Compresor interface{
	compressFile(File) 
	decompressFile(File) 
}



type GzipAlgorithm struct{}

func (a *GzipAlgorithm) compressFile(f File){
	fmt.Println("vamos a compirmirlo con gzip")
}
func (a *GzipAlgorithm) decompressFile(f File){
	fmt.Println("vamos a DESCROMPIRMIR con gzip")
}


type TarAlgo struct{}

func (a *TarAlgo) compressFile(f File){
	fmt.Println("vamos a compirmirlo con TARR")
}
func (a *TarAlgo) decompressFile(f File){
	fmt.Println("vamos a DESCROMPIRMIR con TARRR")
}


type File struct{}

func main() {

	tar:= TarAlgo{}
	gzip:= GzipAlgorithm{}

	value:=2
	myFile:= File{}

	var myCompresor Compresor

	time.Sleep(2*time.Second)

	if value==1{
		myCompresor=&tar
	}else{
		myCompresor=&gzip
	}

	func(compresor Compresor){
		myCompresor.compressFile(myFile)
		myCompresor.decompressFile(myFile)
	}(myCompresor)
	
}