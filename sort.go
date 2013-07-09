package main

import (
    "fmt"
)



type Sorter interface{
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}



type IntSlice []int

func (data IntSlice) Len() int {
    return len(data)
}

func (data IntSlice) Less(i, j int) bool {
    return data[i] < data[j]
}

func (data IntSlice) Swap(i, j int) {
    data[i], data[j] = data[j], data[i]
}



type StringSlice []string

func (data StringSlice) Len() int {
    return len(data)
}

func (data StringSlice) Less(i, j int) bool {
    return data[i] < data[j]
}

func (data StringSlice) Swap(i, j int) {
    data[i], data[j] = data[j], data[i]
}



func InsertionSort(data Sorter) {
    fmt.Println("InsertionSort process data:")
    
    var j int
    for i := 1; i < data.Len(); i++ {
        for j = i - 1; (j >= 0) && data.Less(j + 1, j); j-- {
            data.Swap(j + 1, j)
        }
        fmt.Println(data)
    }
}



func BubbleSort(data Sorter) {
    fmt.Println("BubbleSort process data:")
    
    for i := 0; i < data.Len() - 1; i++ {
        for j := i + 1; j < data.Len(); j++ {
            if data.Less(j, i) {
                data.Swap(j, i)
            }
        }
        fmt.Println(data)
    }
}



func SlectionSort(data Sorter) {
    fmt.Println("SlectionSort process data:")
    
    for i := 0; i < data.Len() - 1; i++ {
        key := i
        for j := i + 1; j < data.Len(); j++ {
            if data.Less(j, key) {
                key = j
            }
        }
        data.Swap(i, key)
        fmt.Println(data)
    }
}



func QuickSort(data Sorter) {
    fmt.Println("QuickSort process data:")
    
    QuickSortRecurrence(data, 0, data.Len() - 1)
}

func QuickSortRecurrence(data Sorter, b, e int) {
    key, i, j := b, b + 1, e
    for i < j {
        for (j > key) && data.Less(key, j) {
            j--
        }
        if j > key {
            data.Swap(key, j)
            key = j
        }
        
        for (i < key) && !data.Less(key, i) {
            i++
        }
        if i < key {
            data.Swap(i, key)
            key = i
        }
    }
    
    fmt.Println(data)
    
    if (key - b) > 1 {
        QuickSortRecurrence(data, b, key - 1)
    }
    
    if (e - key) > 1 {
        QuickSortRecurrence(data, key + 1, e)
    }
}



func main() {
    data1 := IntSlice{5, 2, 4, 6, 1, 3}
    fmt.Println("Input:", data1)
    InsertionSort(data1)
    fmt.Println("Output:", data1, "\n")
    
    data2 := IntSlice{5, 2, 4, 6, 1, 3}
    fmt.Println("Input:", data2)
    BubbleSort(data2)
    fmt.Println("Output:", data2, "\n")
    
    data3 := IntSlice{5, 2, 4, 6, 1, 3}
    fmt.Println("Input:", data3)
    SlectionSort(data3)
    fmt.Println("Output:", data3, "\n")
    
    data4 := IntSlice{5, 2, 4, 6, 1, 3}
    fmt.Println("Input:", data4)
    QuickSort(data4)
    fmt.Println("Output:", data4, "\n")
    
    info1 := StringSlice{"nut", "ape", "elephant", "zoo", "go"}
    fmt.Println("Input:", info1)
    InsertionSort(info1)
    fmt.Println("Output:", info1, "\n")
    
    info2 := StringSlice{"nut", "ape", "elephant", "zoo", "go"}
    fmt.Println("Input:", info2)
    BubbleSort(info2)
    fmt.Println("Output:", info2, "\n")
    
    info3 := StringSlice{"nut", "ape", "elephant", "zoo", "go"}
    fmt.Println("Input:", info3)
    SlectionSort(info3)
    fmt.Println("Output:", info3, "\n")
    
    info4 := StringSlice{"nut", "ape", "elephant", "zoo", "go"}
    fmt.Println("Input:", info4)
    QuickSort(info3)
    fmt.Println("Output:", info3, "\n")
}