# Golang Middle Developer

## Вопросы для собеседований
___

## Описание проекта

Данный проект предназначен для прохождение Middle+ собеседования на Golang.

Помощь в поддержании актуализации базы вопросов, исправлении багов и прочего приветствуется.
Прошу сообщаться об опечатках через pull request.
___
## Теория

___
N1. Что такое slice?

В Go массивы являются структурой данных с фиксированным размером
(например, [4]int{1,2,3,4}).

[Источник](https://www.youtube.com/watch?v=10LW7NROfOQ&t=1389s)

Slice является обёрткой над массивом, которая представляет из себя
сл структуру данных:
- pointer: *array //**указатель** на базовый массив в котором хранятся данные
- length: int //**размер слайса**
- capacity: int //**ёмкость** массива от элемента, на который смотрит указатель, до конца массива

Можем рассматривать как массив с динамической длинной.
___
N2 Что такое append?

Функция для добавления элемента в конец слайса.
___
N3
Какое нулевое значение слайса?

nil
```go
func main () {
	var list []int
	fmt.Println(list == nil) //true

	list = []int{}

	fmt.Println(list == nil) //false
}
```
В примере в первом выводе true, потому что слайс ещё не проинициализирован,
а во втором выводе false тк проинициализировали пустым слайсом.
___
N4 Как лучше всего проверить, пустой ли слайс?

```go
fmt.Println(len(slice)==0) //true
```
___
N5 Аллокация памяти под слайс.

Слайсу можно заранее определить его длину и ёмкость через функцию make.
```go
func main() {
	list:=make([]int, 5, 10)
	fmt.Println(cap(list))//10
	fmt.Println(len(list))//5
}
```
___
N6 Как скопировать один слайс в другой?

Воспользуемся функцией copy.
```go
copy(result, from)
```
___
N7 Что такое Map?

[Источник.](https://www.youtube.com/watch?v=P_SXTUiA-9Y)

Это структура данных которая умеет:
- удалять
- вставлять
- находить значение по ключу

за константное время

Разделим на "бакеты", класть в бакет будем по остатку от хэш функции.

Состоит из
- размера (количество элементов в мапе)
- log //buckets count, количество бакетов
- *Buckets //указатель на список бакетов


Что представляет из себя Bucket?

ToDo: Нормально сформулировать, как кладутся в бакет данные


[LOB (Low order bits)][Bucket]

LOB (остаток от деления)

HOB + LOB = Hash(key)

Структура самого бакета:

[HOB Hash] (8 слотов)

Key: value (8 ключей-значений)
___
N8 Что такое эвакуцаия данных из бакета?

Создаётся новый список бакетов, который в 2 раза длиннее предыдущего,
и данные переностся из старого списка бакетов в новый.
___
N9 Можно ли взяь указатель на элемент мапы?

Нет, нельзя.

Потому что мы не можем взять ссылку на бакет в мапе, тк может произойти эвакуация данных, и данные могут оказатсья в другом бакете.
___
N10 Что такое многопотоное вычисление?

Для этого рассмотрим следующие определения:

[Источник.](https://www.youtube.com/watch?v=mvUiw9ilqn8&t=205s)

- паралеллизм - возможность выполнять вычисления одновременно (на ыизическом уровне)
- конкуретность - способность выполнять несколько процессов на логическом уровне (те программно)

Для работы с конкуретностью используются потоки, но в Go также есть механизм горутин (обвязка поверх потоков).


| Потоки                                                                 | Горутины                                                                             |
|------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| Потоки ОС управляются ядром ОС                                         | Горутины управляются рантаймом Go                                                    |
| Потоки ОС имеют фиксированный размер 1-2мб                             | Горутины обычно имеют размер стека 2кб                                               |
| Размер стэка определяется во время компиляции и не может увеличиваться | Размер стэка определяется во время рантайма и может расти до 1гб (за счёт аллокации) |
| Коммуникация между потоками сложная и медленная                        | Горутины используют каналы для быстрого общения между собой с маленькой задержкой    |


___
N11 Что такое runtime?

runtime - библиотека для запуска Go кода (содержит в себе сборщик мусора, управление стеками и тд)
___
N12 Почему не всегда выводится текст "Go Go"?

```go
func main() {
	go fmt.Println("Go Go") 
	fmt.Println("Goroutine")
}
```

func main - главная горутина, c её выполнением заканчивается выполнение программы.
Если горутина ```go fmt.Println("Go Go") ``` не успеет выполниться до выполнения func main, то текст не будет выведен
___
N13 Что такое каналы?

Каналы это примитивы для синхронизации и обмена данными между горутинами.

```go
message := make(chan string) //задаётся сл образом

msg, isOpen := <-message //возвращает 2 переменные - актуальное состояние канала и то, открыт ли он
```
___
N14
Какими бывают каналы?

- Буферизированными (при создании горутины вторым параметром задаём глубтну буфера)
- Не буферизированными (каждая запись в канал блокирует выполнение горутины)
___
N15
Что такое Wait группа?

Wait группа необходима для ожидания горутины.
```go
var wg sync.WaitGroup
for i := 0; i < len(urls); i++ {
wg.Add(1)
go func() {
doHttp(urls[i])
wg.Done()
}()
}

wg.Wait()
```
___
N16
select?

Select позволяет go процедуре находиться в ожидании нескольких операций передачи данных.

Блокируется, пока не будет готов какой либо из блоков в case.
```go
for {
    select {
	    case msg1 := <-message1:{
                fmt.Println(msg1)
            }
			
	    case msg2 := <-message2:{
		fmt.Println(msg2)
	    }
	}
}
```
___

	

Для просмотра
- https://www.youtube.com/watch?v=_rTuAY7b1RE
- 
