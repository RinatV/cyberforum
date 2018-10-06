package main

import (
	"testing"
)

func TestTask(t *testing.T) {
	var arr [20]int
	var mem [20]int

	for i, _ := range arr {
		mem[i] = arr[i]
	}

	/*Заполнить числовой массив из 20 элементов случайными значениями от -50 до 100*/
	Fill(&arr)

	t.Run("массив заполняется", func(t *testing.T) {
		for i, _ := range arr {
			if mem[i] != arr[i] {
				return
			}
		}
		t.Error("нет изменений в массиве")
	})

	t.Run("значения от -50 до 100", func(t *testing.T) {
		for i, _ := range arr {
			if arr[i] > 100 || arr[i] < -50 {
				t.Errorf("выход за пределы диапазона: %d", arr[i])
			}
		}
	})

	/*Вывести массив.*/
	/*Отсортировать элементы массива при помощи алгоритма сортировки «Выборкой» по возрастанию.*/
	Sort(&arr)

	t.Run("массив сортируется по возростанию", func(t *testing.T) {
		for i, _ := range arr {
			if i > 0 && arr[i-1] > arr[i] {
				t.Errorf("нарушение порядка arr[%d]=%d > arr[%d]=%d", i-1, arr[i-1], i, arr[i])
			}
		}
	})
	/*Вывести отсортированный массив.*/

}
