package main

import (
	pk "curso/src/mypackage"
	"fmt"
	"log"
	"strconv"
	"sync"
)

type car struct {
	brand string
	year int
}

func main() {
	//constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	println("Area cuadrado:", areaCuadrado)

	//Suma
	x := 10
	y := 50

	result := x + y
	fmt.Println("suma:", result)

	//Resta
	result = y - x
	fmt.Println("resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	//modulo - residuo
	result = y % x
	fmt.Println("Residuo:", result)

	//Incremental
	x++
	fmt.Println("x:", x)
	//Decremental
	x--
	fmt.Println("x:", x)

	//Reto
	//Area rectangulo
	largo := 6
	ancho := 2
	areaRectangulo := largo * ancho
	fmt.Println("Area rectangulo:", areaRectangulo)
	//Area Trapecio
	baseA := 4
	baseB := 10
	alturaTrapecio := 4
	areaTrapecio := ((baseA + baseB) * alturaTrapecio) / 2
	fmt.Println("Area trapecio:", areaTrapecio)
	//Area circulo
	const PI float64 = 3.14
	var radioCirculo float64 = 10
	areaCirculo := PI * radioCirculo * radioCirculo
	fmt.Println("El Area del Circulo es :", areaCirculo)

	//FMT
	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("tipo %T\n", nombre)
	fmt.Printf("tipo: %T\n", cursos)

	plus := plus(2, 2)
	println("suma: ", plus)

	_ , sencondValue := plusDoubleReturn(8, 8)
	fmt.Println(sencondValue)

	//Iteraciones - ciclos
	//For
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//For while
	counter := 0
	for counter <= 5 {
		println(counter)
		counter++
	}

	//Condicionales
	valor1 := 1

	if valor1  == 1 {
		fmt.Println("Verdadero")
	}

	//Parsear valores
	value, err := strconv.Atoi("8")
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(value)

	//Array y Slices
	var array [4] int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	slice := [] int {0,1,2}
	//append
	newSlice := []int{8,9,10}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

	stringSlice := [] string {"hola", "mundo"}

	for i, valor := range stringSlice {
		fmt.Println(i, valor)
	}

	isPalindromo("amor a roma")

	//Structs
	myCar := car{brand: "Ford", year: 2023}
	fmt.Println(myCar)

	var otherCar pk.CarPublic
	otherCar.Brand = "Ferrari"
	otherCar.Year = 2023
	fmt.Println(otherCar)

	//Puntero
	r := 50
	n := &r
	fmt.Println(r)
	fmt.Println(*n)

	myPc := pc{ram: 16, disk: 200, brand: "HP"}
	fmt.Println(myPc)
	myPc.duplicateRam()
	myPc.string()

	cuadrado := cuadrado{base: 2}
	rectangulo := rectangulo{base: 2, altura: 4}

	calcular(cuadrado)
	calcular(rectangulo)

	myInterface := []interface{}{"hola", 12, 4.90}
	fmt.Println(myInterface...)

	var wg sync.WaitGroup
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()

	channel := make(chan string, 1)
	fmt.Println("Hello")
	go channelSay("Bye", channel)
	fmt.Println(<-channel)

	fmt.Printf("2, %T", 2)
}

type figuras interface {
	area() float64
}

type rectangulo struct {
	base float64
	altura float64
}

type cuadrado struct {
	base float64
}

func (cuadrado cuadrado) area() float64 {
	return cuadrado.base * cuadrado.base
}

func (rectangulo rectangulo) area() float64 {
	return rectangulo.base * rectangulo.altura
}

func calcular(figuras figuras) {
	fmt.Println("Area: ", figuras.area())
}

type pc struct {
	ram int
	disk int
	brand string
}

func (myPc *pc) duplicateRam () {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) string() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func plus(a, b int) int {
	return a + b
}

func plusDoubleReturn(a, b int) (c, d int) {
	return a, a + b
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es")
	}

	//Maps
	m := make(map[string]int)
	m["jose"] = 14
	m["pepito"] = 20

	for i,  v := range m {
		fmt.Println(i, v)
	}
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func channelSay(text string, channel chan<- string) {
	channel <- text
}