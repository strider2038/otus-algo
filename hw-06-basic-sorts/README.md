# Домашнее задание №6 "Простые сортировки"

## Цель

Создание, тестирование и сравнений алгоритмов сортировки.

## Задание

* [x] Реализовать алгоритм BubbleSort.
* [x] Реализовать алгоритм InsertionSort.
* [x] Реализовать алгоритм ShellSort.
* [x] Оптимизировать алгоритм BubbleSort.
* [x] Оптимизировать алгоритм InsertionSort, сделать сдвиг элементов вместо обмена.
* [x] Оптимизировать алгоритм InsertionSort, сделать бинарный поиск места вставки.
* [x] Оптимизировать алгоритм ShellSort:
    * выбрать альтернативный вариант выбора шагов;
    * добавить внутреннюю сортировку вставкой.
* [x] Выполнить комплексное тестирование алгоритмов сортировки.

Протестировать алгоритмы на массивах размеров:
1, 10, 100, 1.000, 10.000, 100.000, 1.000.000, 10.000.000 (этот по желанию)

И с различным характером данных:

* random - массив из случайных чисел
* digits - массив из случайных цифр
* sorted - на 99% отсортированный массив
* reverse - обратно-отсортированный массив

## Сравнение алгоритмов

### BubbleSort

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 310 ns   | 98 ns    | 200 ns   | 115 ns   |
| 10         | 298 ns   | 336 ns   | 364 ns   | 302 ns   |
| 100        | 11.7 µs  | 11.6 µs  | 13.1 µs  | 8.7 µs   |
| 1 000      | 796 µs   | 782 µs   | 648.1 µs | 694.4 µs |
| 10 000     | 116.2 ms | 114.3 ms | 64.9 ms  | 74.8 ms  |
| 100 000    | 13.8 s   | 13.7 s   | 6.7 s    | 7.4 s    |
| 1 000 000  | ~        | ~        | ~        | ~        |
| 10 000 000 | ~        | ~        | ~        | ~        |

### BubbleSort с оптимизацией

| N          | random  | digits   | sorted   | reverse  |
|------------|---------|----------|----------|----------|
| 1          | 99 ns   | 111 ns   | 167 ns   | 162 ns   |
| 10         | 219 ns  | 233 ns   | 209 ns   | 192 ns   |
| 100        | 8.3 µs  | 7.9 µs   | 3.2 µs   | 5.4 µs   |
| 1 000      | 563 µs  | 578.6 µs | 460.7 µs | 511.2 µs |
| 10 000     | 88.6 ms | 87.1 ms  | 43.2 ms  | 47.8 ms  |
| 100 000    | 11 s    | 10.9 s   | 4.5 s    | 4.7 s    |
| 1 000 000  | ~       | ~        | ~        | ~        |
| 10 000 000 | ~       | ~        | ~        | ~        |

### InsertionSort

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 144 ns   | 127 ns   | 150 ns   | 188 ns   |
| 10         | 195 ns   | 153 ns   | 138 ns   | 288 ns   |
| 100        | 2.5 µs   | 2.7 µs   | 299 ns   | 31.3 µs  |
| 1 000      | 201.1 µs | 170.7 µs | 6.2 µs   | 638.7 µs |
| 10 000     | 18 ms    | 15.6 ms  | 600.2 µs | 41 ms    |
| 100 000    | 1.8 s    | 1.6 s    | 46.8 ms  | 3.6 s    |
| 1 000 000  | ~        | ~        | ~        | ~        |
| 10 000 000 | ~        | ~        | ~        | ~        |

### InsertionSort со сдвигом вместо обмена

| N          | random   | digits   | sorted  | reverse  |
|------------|----------|----------|---------|----------|
| 1          | 358 ns   | 178 ns   | 159 ns  | 188 ns   |
| 10         | 164 ns   | 196 ns   | 212 ns  | 288 ns   |
| 100        | 1.7 µs   | 2.2 µs   | 3.2 µs  | 31.3 µs  |
| 1 000      | 109.8 µs | 138.3 µs | 312 µs  | 638.7 µs |
| 10 000     | 10.7 ms  | 9.5 ms   | 21.2 ms | 41 ms    |
| 100 000    | 1.1 s    | 973.3 ms | 2.2 s   | 3.6 s    |
| 1 000 000  | ~        | ~        | ~       | ~        |
| 10 000 000 | ~        | ~        | ~       | ~        |

### InsertionSort со сдвигом и бинарным поиском

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 243 ns   | 188 ns   | 607 ns   | 157 ns   |
| 10         | 458 ns   | 320 ns   | 333 ns   | 442 ns   |
| 100        | 5 µs     | 4.6 µs   | 2.1 µs   | 6.8 µs   |
| 1 000      | 225.9 µs | 177.8 µs | 35.8 µs  | 260.4 µs |
| 10 000     | 12.2 ms  | 12.1 ms  | 541.8 µs | 21.2 ms  |
| 100 000    | 1.1 s    | 1.18 s   | 32 ms    | 2.2 s    |
| 1 000 000  | ~        | ~        | ~        | ~        |
| 10 000 000 | ~        | ~        | ~        | ~        |

### ShellSort

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 151 ns   | 270 ns   | 108 ns   | 110 ns   |
| 10         | 208 ns   | 364 ns   | 134 ns   | 204 ns   |
| 100        | 3 µs     | 2.9 µs   | 924 ns   | 1 µs     |
| 1 000      | 46 µs    | 25.5 µs  | 20.8 µs  | 10.9 µs  |
| 10 000     | 683 µs   | 305 µs   | 500.5 µs | 138.4 µs |
| 100 000    | 10.5 ms  | 3.4 ms   | 6.1 ms   | 1.8 ms   |
| 1 000 000  | 135.1 ms | 41.4 ms  | 86.2 ms  | 22.6 ms  |
| 10 000 000 | 1.8 s    | 486.9 ms | 1.2 s    | 309.6 ms |

### ShellSort, формула Frank and Lazarus

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 193 ns   | 253 ns   | 241 ns   | 165 ns   |
| 10         | 233 ns   | 251 ns   | 178 ns   | 214 ns   |
| 100        | 2.9 µs   | 1.9 µs   | 973 ns   | 1.1 µs   |
| 1 000      | 44.7 µs  | 48.3 µs  | 20.7 µs  | 10.2 µs  |
| 10 000     | 680 µs   | 384.4 µs | 396.3 µs | 168.6 µs |
| 100 000    | 9.6 ms   | 3 ms     | 6.5 ms   | 1.7 ms   |
| 1 000 000  | 132.4 ms | 33 ms    | 85.8 ms  | 21.2 ms  |
| 10 000 000 | 1.9 s    | 399.6 ms | 1.2 s    | 281.6 ms |

### ShellSort, внутренняя сортировка вставками

| N          | random   | digits   | sorted   | reverse  |
|------------|----------|----------|----------|----------|
| 1          | 177 ns   | 225 ns   | 301 ns   | 189 ns   |
| 10         | 231 ns   | 262 ns   | 200 ns   | 300 ns   |
| 100        | 2.8 µs   | 2.1 µs   | 930 ns   | 1.6 µs   |
| 1 000      | 57.4 µs  | 25.7 µs  | 20.8 µs  | 15.2 µs  |
| 10 000     | 799 µs   | 319.9 µs | 418.8 µs | 208.4 µs |
| 100 000    | 9.1 ms   | 3.8 ms   | 6.1 ms   | 1.8 ms   |
| 1 000 000  | 119.7 ms | 39.9 ms  | 85.1 ms  | 21.2 ms  |
| 10 000 000 | 1.6 s    | 501.6 ms | 1.2 s    | 262.1 ms |

## Выводы

Интересно отметить, что различные алгоритмы обладают своими достоинствами и недостатками
и ведут себя по-разному на различных видах данных.

Особенности алгоритмов

* BubbleSort
  * самый медленный и самый простой
  * на малых N < 1000 практически не отличается от других по производительности
  * на отсортированных данных работает примерно так же как на неотсортированных (но примерно в 2 раза быстрее)
* InsertionSort
  * обладает хорошей производительностью при N < 10 000
  * быстро работает на отсортированных данных
  * если обмен заменить на сдвиг, то на отсортированных данных работает медленнее
  * со сдвигом и бинарным поиском так же быстро работает на отсортированных данных
* ShellSort
  * наиболее производительный среди всех рассмотренных
  * на малых N < 100 работает медленнее, чем другие алгоритмы

## Запуск кода

Для запуска тестов необходимо распаковать архив с данными для тестов в директорию `testdata/sortdata`

Запуск тестов

```shell
go test -v ./...
```
