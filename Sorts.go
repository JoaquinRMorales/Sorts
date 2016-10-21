package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

//-----------------------------------------------------------------------------
//Input from console
func Scan(str string) string {
	print(str, "  ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return input
}

//-------------------------------------------------------------------------------
//Folder reader
func readD() ([]string, string) {
	orden := Scan("Direccion de carpeta ... EJEMPLO /Home/Informatica/tarea2/")
	files, err := ioutil.ReadDir(orden)
	if err != nil {
		log.Fatal(err)
	}

	a := make([]string, 16)

	for dex, file := range files {
		a[dex] = file.Name()
	}

	return a, orden
}

//-------------------------------------------------------------------------------
//Return auxs to the Analizer()
func HeaderReader(str string) ([]string, string) {
	a := make([]string, 20)
	var acumulator string
	for i, c := range str {
		a[i] = string(c)
		acumulator = acumulator + a[i]
	}
	return a, acumulator
}

//-----------------------------------------------------------------------------
//Return the final directory to open and the len of the data in string (1k, 10k, 100k, 1m)
func Analizer() (string, string) {

	var analizado string

	a, orden := readD()
	b := Scan("Duplicates = du | Random = ra | Reversed = re | Sorted = so ... ")
	b2 := Scan("Ingrese: 1k | 10k | 100k | 1m ... ")

	for count := 0; count <= 15; count++ {
		c, d := HeaderReader(a[count]) 
		//cAux para "du", "ra"...
		cAux := c[0] + c[1]

		if b == "du" && cAux == b {
			str := strings.Replace(a[count], "duplicates", "", -1)
			str = strings.Replace(str, ".txt", "", -1)
			if str == b2 {
				analizado = orden + d
			}
		}
		if b == "ra" && cAux == b {
			str := strings.Replace(a[count], "random", "", -1)
			str = strings.Replace(str, ".txt", "", -1)
			if str == b2 {
				analizado = orden + d
			}
		}
		if b == "re" && cAux == b {
			str := strings.Replace(a[count], "reversed", "", -1)
			str = strings.Replace(str, ".txt", "", -1)
			if str == b2 {
				analizado = orden + d
			}
		}
		if b == "so" && cAux == b {
			str := strings.Replace(a[count], "sorted", "", -1)
			str = strings.Replace(str, ".txt", "", -1)
			if str == b2 {
				analizado = orden + d
			}
		}

	}

	return analizado, b2
}

//------------------------------------------------------------------------------
//Open the file, then split the data and convert it to Int type, return the array of numbers and the leng of the array
func Stringer() ([]int, uint) {
	a, cant := Analizer()

	var largo int

	if cant == "1k" {
		largo = 1000
	}
	if cant == "10k" {
		largo = 10000
	}
	if cant == "100k" {
		largo = 100000
	}
	if cant == "1m" {
		largo = 1000000
	}

	f, err := os.Open(a)
	interer := make([]int, largo)

	if err != nil {
		log.Fatalln("Broked", err.Error())
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Broked 2")
	}
	str := string(bs)
	str = strings.Replace(str, "\n", " ", -1)
	str = strings.Replace(str, "-", "", -1)
	split := strings.Split(strings.TrimSpace(str), " ")

	for i, k := range split {
		interer[i], _ = strconv.Atoi(k)
	}
	l := len(interer)
	u := uint(l)

	return interer, u
}

//-----------------------------------------------------------------------------
func Swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

//-----------------------------------------------------------------------------
func Swap2(x *int, y *int) {
	temp := *y
	*y = *x
	*x = temp
}

//------------------------------------------------------------------------------
func BubbleSort(arrayzor []int) []int {

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				Swap(arrayzor, i, i+1)
				swapped = true
			}
		}
	}

	return arrayzor
}

//------------------------------------------------------------------------------
func SelectionSort(arrayzor []int) []int {

	for i := 0; i < len(arrayzor)-1; i++ {
		min := i
		for j := i + 1; j < len(arrayzor)-1; j++ {
			if arrayzor[j] < arrayzor[min] {
				min = j
			}
		}
		Swap(arrayzor, i, min)
	}
	return arrayzor
}

//------------------------------------------------------------------------------
func InsertionSort(arrayzor []int) []int {

	for i := 1; i < len(arrayzor); i++ {
		for j := i; j > 0 && arrayzor[j] < arrayzor[j-1]; j-- {
			Swap(arrayzor, j, j-1)
		}
	}
	return arrayzor
}

//------------------------------------------------------------------------------
func MergeSort(A []int) []int {

	N := len(A)
	if N == 1 {
		return A
	}
	left := MergeSort(A[:N/2])
	right := MergeSort(A[N/2:])
	return Merge(left, right)
}
func Merge(left, right []int) []int {

	i := 0
	j := 0
	k := 0
	r := make([]int, len(left)+len(right))
	for i < len(left) || j < len(right) {
		if i >= len(left) {
			r[k] = right[j]
			k++
			j++
		} else if j >= len(right) {
			r[k] = left[i]
			k++
			i++
		} else if left[i] < right[j] {
			r[k] = left[i]
			k++
			i++
		} else {
			r[k] = right[j]
			k++
			j++
		}
	}
	return r
}

//-------------------------------------------------------------------------------

func Partition(array []int, p uint, q uint, pivotLocation uint) uint {

	pivot := array[pivotLocation]
	Swap2(&array[pivotLocation], &array[q])
	i := p
	for j := p; j < q; j++ {
		if array[j] <= pivot {
			Swap2(&array[i], &array[j])
			i++
		}
	}
	Swap2(&array[q], &array[i])
	return i
}

func QuickSort(array []int, start uint, end uint) []int {
	if start < end {
		pivot := (end + start) / 2
		r := Partition(array, start, end, pivot)
		if r > start {
			QuickSort(array, start, r-1)
		}
		QuickSort(array, r+1, end)
	}
	return array
}

//------------------------------------------------------------------------------

func Menu() {
	ar, l := Stringer()

	option := true

	for option {
		str := Scan("1 Bubble Sort \n2 Insertion Sort\n3 Selection Sort\n4 Merge Sort\n5 Quick Sort\n6 Cambiar directorio y archivo \n7   Exit")
		str = strings.ToLower(str)
		if str == "1" {
			fmt.Println(BubbleSort(ar))
		} else if str == "2" {
			fmt.Println(InsertionSort(ar))
		} else if str == "3" {
			fmt.Println(SelectionSort(ar))
		} else if str == "4" {
			fmt.Println(MergeSort(ar))
		} else if str == "5" {
			fmt.Println(QuickSort(ar, 0, l-1))
		} else if str == "6" {
			ar, l = Stringer()
		} else if str == "7" || str == "exit" {
			fmt.Println("Programa finalizado")
			break
		} else {
			fmt.Println("Input invalido...")
		}
	}

}

//-------------------------------------------------------------------------------
func main() {
	Menu()
}
