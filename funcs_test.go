package main

import (
	"fmt"
	"testing"
)

func TestAddElementBasic(t *testing.T) {
    var array []int
    ans := addElement(array, 1)
    if len(ans) != 1 || ans[0] != 1 {
        t.Errorf("addElement(array, 1) = %v; Правильно: [1]", ans)
    }
}

func TestAddElementTableDriven(t *testing.T) {
    var tests = []struct {
        arr []int
        elem int
        want []int
    }{
        {[]int{1, 2, 3}, 4, []int{1, 2, 3, 4}},  // Массив длины 3
        {[]int{1}, 42, []int{1, 42}},            // Массив длины 1
        {[]int{}, 4, []int{4}},                  // Пустой массив
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%v, %v", tt.arr, tt.elem)
        t.Run(testname, func(t *testing.T) {
            ans := addElement(tt.arr, tt.elem)
            if len(ans) != len(tt.want) {
                t.Errorf("Получено: %v, Правильно: %v", ans, tt.want)
            } else {
                for i := 0; i < len(ans); i++ {
                    if ans[i] != tt.want[i] {
                        t.Errorf("Получено: %v, Правильно: %v", ans, tt.want)
                    }
                }
            }
        })
    }
}

func TestRemoveElementBasic(t *testing.T) {
    array := []int{1, 2}
    ans := removeElement(array)
    if len(ans) != 1 || ans[0] != 1 {
        t.Errorf("removeElement(array) = %v; Правильно: [1]", ans)
    }
}

func TestRemoveElementTableDriven(t *testing.T) {
    var tests = []struct {
        arr []int
        want []int
    }{
        {[]int{1, 2, 3}, []int{1, 2}},  // Массив длины 3
        {[]int{1}, []int{}},            // Массив длины 1
        {[]int{}, []int{}},             // Пустой массив
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%v", tt.arr)
        t.Run(testname, func(t *testing.T) {
            ans := removeElement(tt.arr)
            if len(ans) != len(tt.want) {
                t.Errorf("Получено: %v, Правильно: %v", ans, tt.want)
            } else {
                for i := 0; i < len(ans); i++ {
                    if ans[i] != tt.want[i] {
                        t.Errorf("Получено: %v, Правильно: %v", ans, tt.want)
                    }
                }
            }
        })
    }
}

func TestAddOneToArrayBasic(t *testing.T) {
    array := []int{1, 2}
    ans := addOneToArray(array)
    if len(ans) != 2 || ans[0] != 2 || ans[1] != 3 {
        t.Errorf("addOneToArray(array) = %v; Правильно: [1]", ans)
    }
}

func TestAddOneToArrayTableDriven(t *testing.T) {
    var tests = []struct {
        arr []int
        want []int
    }{
        {[]int{1, 2, 3}, []int{2, 3, 4}},  // Массив длины 3
        {[]int{1}, []int{2}},              // Массив длины 1
        {[]int{}, []int{}},                // Пустой массив
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%v", tt.arr)
        t.Run(testname, func(t *testing.T) {
            ans := addOneToArray(tt.arr)
            if len(ans) != len(tt.want) {
                t.Errorf("Получено: %v, Правильно: %v", ans, tt.want)
            } else {
                for i := 0; i < len(ans); i++ {
                    if ans[i] != tt.want[i] {
                        t.Errorf("Получено: %v, Правильно: %v", ans, tt.want)
                    }
                }
            }
        })
    }
}
