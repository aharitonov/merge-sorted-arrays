package main;

// Условие задачи
// Нужно слить 2 отсортированных массива в один отсортированный
// массив

// Входные параметры
// 2 отсортированных (по возрастанию) массива A и B длины M и N
// Вывод
// Отсортированный (по возрастанию) массив, состоящий из элементов первых двух

// Примеры
// Ввод [1, 2, 5], [1, 2, 3, 4, 6]
// Вывод [1, 1, 2, 2, 3, 4, 5, 6]

import (
    "fmt"
)


const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
type Numeric int;

func main() {
	fmt.Println("Hello")

	a := []Numeric{1,2,5}
	b := []Numeric{1,2,3,4,6}
	c := MergeSortedArrays(a, b);

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}


func MergeSortedArrays(a, b []Numeric) []Numeric {

	if len(a) == 0 && len(b) == 0 {
		return make([]Numeric, 0)
	}
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	c := make([]Numeric, 0)

	lenA := len(a)
	lenB := len(b)
	iMaxA := a[0]
	iMaxB := b[0]
	indexA := 0
	indexB := 0

	for indexA < lenA || indexB < lenB {
		for indexA < lenA && iMaxA < iMaxB {
			v := a[indexA]
			c = append(c, v)
			indexA++
			if indexA < lenA {
				iMaxA = a[indexA]
			} else {
				iMaxA = Numeric(MaxInt) // block for "for"
			}
		}
		for indexB < lenB && iMaxB <= iMaxA {
			v := b[indexB]
			c = append(c, v)
			indexB++
			if indexB < lenB {
				iMaxB = b[indexB]
			} else {
				iMaxB = Numeric(MaxInt) // block for "for"
			}
		}
	}
	return c
}
