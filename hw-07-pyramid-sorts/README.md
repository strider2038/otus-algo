# Домашнее задание №7 "Пирамидальные сортировки"

## Цель

Создание, тестирование и сравнений алгоритмов пирамидальной сортировки.

## Задание

* [x] Реализовать алгоритм SelectionSort.
* [x] Реализовать алгоритм HeapSort.
* [x] Выполнить комплексное тестирование алгоритмов сортировки.

Протестировать алгоритмы на массивах размеров:
1, 10, 100, 1.000, 10.000, 100.000, 1.000.000, 10.000.000 (этот по желанию)

И с различным характером данных:

* random - массив из случайных чисел
* digits - массив из случайных цифр
* sorted - на 99% отсортированный массив
* reverse - обратно-отсортированный массив

## Сравнение алгоритмов

### SelectionSort

| N          | random | digits | sorted   | reverse  |
|------------|--------|--------|----------|----------|
| 1          | 258 ns | 210 ns | 243 ns   | 169 ns   |
| 10         | 362 ns | 410 ns | 317 ns   | 319 ns   |
| 100        | 10 µs  | 6.6 µs | 7.7 µs   | 5.6 µs   |
| 1 000      | 614 µs | 575 µs | 595.3 µs | 531.2 µs |
| 10 000     | 58 ms  | 60 ms  | 59.9 ms  | 45.5 ms  |
| 100 000    | 5.9 s  | 6 s    | 6 s      | 4.7 s    |
| 1 000 000  | ~      | ~      | ~        | ~        |
| 10 000 000 | ~      | ~      | ~        | ~        |

### HeapSort

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 236 ns   | 227 ns   | 201 ns   | 306 ns   |
| 10         | 622 ns   | 760 ns   | 543 ns   | 374 ns   |
| 100        | 5.5 µs   | 6.1 µs   | 4.6 µs   | 4.1 µs   |
| 1 000      | 68.7 µs  | 78.5 µs  | 52.8 µs  | 50.4 µs  |
| 10 000     | 1.1 ms   | 940.1 µs | 645.7 µs | 635.8 µs |
| 100 000    | 11.4 ms  | 7.8 ms   | 8.2 s    | 7.6 s    |
| 1 000 000  | 162.7 ms | 101 ms   | 106.6 ms | 107.4 ms |
| 10 000 000 | 2.7 s    | 1.2 s    | 1.4 s    | 1.3 s    |

## Выводы

Особенности алгоритмов

* SelectionSort
    * аналогичен BubbleSort
    * на малых N < 1000 практически не отличается от других по производительности
    * на отсортированных данных работает примерно так же как на неотсортированных (но примерно в 2 раза быстрее)
* HeapSort
    * по производительности сравним с ShellSort
    * производительность примерно одинакова для различных типов данных
    * в отличие от ShellSort на обратно-отсортированных данных работает медленнее
    * на малых N < 100 работает медленнее, чем простые алгоритмы

## Запуск кода

Для запуска тестов необходимо распаковать архив с данными для тестов в директорию `testdata/sortdata`

Запуск тестов

```shell
go test -v ./...
```