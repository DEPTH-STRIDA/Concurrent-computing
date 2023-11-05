package main

import (
	"fmt"
	"math/rand"
)

var L, n, nMax, S, finish, eventTime = 0, 0, 40, "free", false, 0
var eventMode int
var typeEvent = []int{1, 3}
var eventTimeM = []int{0, 120}

func main() {
	fmt.Println("Стратовала программа")
	for finish != true {
		fmt.Println("Работает цикл")
		eventMode = typeEvent[0]
		eventTime = eventTimeM[0]

		fmt.Printf("Режим: %d. Время %d. L: %d. n: %d\n", eventMode, eventTime, L, n)
		fmt.Println(typeEvent)
		fmt.Println(eventTimeM)
		fmt.Println("-----------------\n")
		if finish == true || n > nMax || eventMode == 3 {
			A3()
			break
		}
		switch eventMode {
		case 1:
			A1()
		case 2:
			A2()
		case 3:
			A3()
			finish = true
		}
	}

}

func plan(E int, T int) {
	if E == 1 {
		//fmt.Println("plan(1)")
		//fmt.Println(typeEvent)
		//fmt.Println(eventTimeM)
		addTo(&typeEvent, 1)
		addTo(&eventTimeM, T+getTask(1))
		quickSort(eventTimeM, typeEvent, 0, len(typeEvent)-1)
	}
	if E == 2 {
		//fmt.Println("plan(2)")
		//fmt.Println(typeEvent)
		//fmt.Println(eventTimeM)
		addTo(&typeEvent, 2)
		addTo(&eventTimeM, T+getTask(2))
		quickSort(eventTimeM, typeEvent, 0, len(typeEvent)-1)
		//typeEvent, eventTimeM = quickSort(eventTimeM, typeEvent, 0, len(typeEvent)-1)
	}
	if E == 3 {
		A3()
	}
}
func A1() {
	//fmt.Println("А1")
	if len(typeEvent) != 1 {
		remove(&typeEvent, &eventTimeM, typeEvent, eventTimeM)
	}
	//fmt.Println("После удаление элемента")
	//fmt.Println(typeEvent)
	//fmt.Println(eventTimeM)
	plan(1, eventTime)
	//fmt.Println("free" == S)

	if S == "free" {
		//	fmt.Println("Процессор свободен, поэтому plan(2)")
		S = "Busy"
		plan(2, eventTime)
	} else {
		L++
	}

}
func A2() {
	//fmt.Println("А2")
	n++
	//fmt.Println(n)
	if n >= nMax {
		fmt.Println("Количество потоков превышено.")
		plan(3, eventTime)
	} else {
		if L > 0 {
			//fmt.Println("l>0")
			plan(2, eventTime)
			L--
			if len(typeEvent) != 1 {
				remove(&typeEvent, &eventTimeM, typeEvent, eventTimeM)
			}
		} else {
			//fmt.Println("L<0 -->  s=free")
			S = "free"
			//fmt.Println(S)
		}
	}
}
func A3() {

	fmt.Println("Программа завершила работу.")
	finish = false
}

func addTo(origSlice *[]int, number int) {
	//fmt.Println("В слайс ", *origSlice, " добавляется", number)
	*origSlice = append(*origSlice, number)
	//fmt.Println(*origSlice)
}
func remove(o_1rigSlice *[]int, o_2rigSlice *[]int, o_11rigSlice []int, o_22rigSlice []int) {
	a := *o_1rigSlice
	copy(a[:], a[0+1:])
	a[len(a)-1] = 0
	a = a[:len(a)-1]
	*o_1rigSlice = a
	a = *o_2rigSlice
	copy(a[:], a[0+1:])
	a[len(a)-1] = 0
	a = a[:len(a)-1]
	*o_2rigSlice = a
}
func getTask(E int) int {
	if E == 1 {
		return 9 + rand.Intn(13-9+1)
	} else {
		return 9 + rand.Intn(15-9+1)
	}

}
func sort() {
	quickSort(eventTimeM, typeEvent, 0, len(typeEvent)-1)
}
func quickSort(arr []int, arr1 []int, low, high int) ([]int, []int) {
	if low < high {
		var p int
		arr, arr1, p = partition(arr, arr1, low, high)
		arr, arr1 = quickSort(arr, arr1, low, p-1)
		arr, arr1 = quickSort(arr, arr1, p+1, high)
	}
	return arr, arr1
}
func partition(arr []int, arr1 []int, low, high int) ([]int, []int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr1[i], arr1[j] = arr1[j], arr1[i]
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	arr1[i], arr1[high] = arr1[high], arr1[i]
	return arr, arr1, i
}
