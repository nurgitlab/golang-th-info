## Подсказки по синтаксису для Leetcode
___

N1. Сортировка по функции-условию:

```go
sort.SliceStable(nums, func (i, j int) bool {
return nums[i] < nums[j]
})
```
___
N2. Как взять слайс по слайсу:

```go
arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
fmt.Println(arr[2:4]) //3 4
```
___
N3
float в int и обратно:

```go
a:= 10 //int
aFloat := float64(a) //int -> float

b:=int(aFloat) //float -> int
```
___
N4 Из строки в int и обратно:

```go
strconv.Atoi("333") // to int
strconv.Itoa(5555) // to string
```
___
N5
___

