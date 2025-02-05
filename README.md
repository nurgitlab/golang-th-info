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
Для чего нужен select?

Select позволяет находиться в ожидании нескольких каналов.

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
N17
Что такое указатель?

Это переменная котора содержит адрес объекта в памяти.
___
N18
Что такое метод?

Метод это функция для работы в структуре.

```go
type Student struct {
	name string
	age  int
}

func (s Student) changeAge(newAge int) {
	s.age = newAge
	fmt.Println("in func", s)
}
```

Имя метода является модификатором доступа - 
с маленькой буквы метод доступен только в пакете где обьявлен типа, с большой везде.
___
N19
Что такое интерфейс?

Интерфейс - специальный тип, определяющий наборы методов.

Пустой интерфейс = any из Js.
В пустой интерфейс можно положить любое значение.
```go
var currentIterface interface {} = &Animal{age: 10, say: "Nothing", name: "Бобик"}
```
___
N20
Полиморфизм?

Полиморфизм - способность функции обрабатывать данные разных типов.
___
N21
Как проверить на тип?

```go
if human, isOk := p.(*Animal); isOk {
	fmt.Println("Это животное", human.name)
} else {
	fmt.Println("Это не животное")
}
```
___
N22
Встраивание (Embedding) - что такое?

Встраивание - это встраивание свойства в определенное поле структуры.

Это не наследование!

В наследовании мы расширяем свойства и поля родителя. Также наследуемый тип является подтипом своего родителя 
(Если тип ребенка B, а родителя A, то у ребенка также будет тип B).

В композиции же мы просто получаем свойства без наследования типа.
___
N23
Как копировать слайс?

```go
arr := []int{1,2,3,4,5,6,7}
a := make([]int, len(arr))
copy(a, arr) //to from
```
___
N24
Как удалить элемент из слайса?

```go
a := append(arr[:2], arr[3:]...) //ломает исходный слайс
fmt.Println(a)
```
___
N25 Что такое defer функции?

defer функции - это функции с отложенным выполнением. Они складываются в стек, и 
___
N26 panic

Рассмотрим панику и выход из неё
```go
func myPanic() {
	defer func() {
		panicValue := recover()
		fmt.Println("->", panicValue)
	}()
	panic("something bad happened")
}
```
___
N27
WaitGroup???
___
N28
Mutex???
___
N29
RWMutex

То же самое, что и Mutex, но с блокировкой на чтение.
___
N30 <-> N13
Что такое канал?

Канал - это инструмент обмена данными между горутинами.

```go
type chan struct {
	mx sync.mutex
	buffer T //опционально
	readers []Goroutines
	writers []Goroutines
}
```
Необходимо две горутины для работы канала - на чтение и на запись:
```go
    channel := make(chan int)//обьявление интовой горутины
	go func(write chan<- int) {
		time.Sleep(time.Second)

		write <- 1
	}(channel)

	go func(read <-chan int) {
		time.Sleep(time.Second)
		fmt.Println(<-read)
	}(channel)
```
___
N31
Направленность канала?

[Источник](https://www.youtube.com/watch?v=ngEj0ZwQEn0&list=PLc2Vkg57qmuRNHp6NNvYRVgg3OP-b5E_v&index=21)

```go
writeChan := make(chan<- int) // можем только записывать, записи нет
```
```go
writeChan := make(<-chan int) // канал только на чтение
```
___
N32
select?


select выбирает неблокирующую операцию с каналом.

Если в select несколько неблокирующих операций, 
select выбирает рандомно.
```go
func main() {
	buffered := make(chan int, 1)
	buffered <- 1
	select {
	case str := <-buffered: //выбрал неблокирующую операцию при буфере 1
		fmt.Println("read ", str)
	case buffered <- 2:
		fmt.Println("write ", <-buffered, <-buffered)
	}
}

```
___
N33 Что такое контекст?

Context - обьект который служит для двух целей:

- хранить значения (для передачи в горутины)
- сообщить о завершении

Можно создать контекст Backgound и TODO
___
N34
___
N35
___
N36
___
N37
___
Для просмотра
- [Самые популярные вопросы на собеседовании (80%)](https://www.youtube.com/watch?v=_rTuAY7b1RE)


- [Менторинг, нужен ли?](https://www.youtube.com/watch?v=b2qJYHS_JvM) И сколько платят на рынке (500к)
- [Senior+ lvl Mock interview](https://www.youtube.com/watch?v=hpwDdACSfrI&t=307s)
