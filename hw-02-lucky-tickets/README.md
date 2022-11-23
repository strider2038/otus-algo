# Домашнее задание №2 "Счастливые билеты"

## Задача

Билет с 2N значным номером считается счастливым, если сумма N первых 
цифр равна сумме последних N цифр.

Посчитать, сколько существует счастливых 2N-значных билетов.

* Начальные данные: число N от 1 до 10.
* Вывод результата: количество 2N-значных счастливых билетов.

## Решение

1. Наивный алгоритм грубым перебором для N=3 - функции `CountByBruteForceForN6` и `CountByBruteForce2ForN6`.
2. Рекурсивный алгоритм для произвольного N - функция `CountRecursively`.
3. Быстрый алгоритм на суммах цифр для произвольного N = функция `CountFast`.

## Запуск кода

Запуск примеров кода

```shell
go run main.go
```

Запуск тестов

```shell
go test -v ./...
```

Пример выполнения теста

```
go test -v ./luckytickets/fast_test.go 
=== RUN   TestCountFast
    fast_test.go:31: ЗАДАЧА.
         Счастливые билеты 20
        
        Билет с 2N значным номером считается счастливым,
        если сумма N первых цифр равна сумме последних N цифр.
        Посчитать, сколько существует счастливых 2N-значных билетов.
        
        Начальные данные: число N от 1 до 10.
        Вывод результата: количество 2N-значных счастливых билетов.
=== RUN   TestCountFast/test_0:_1
    datatest.go:96: elapsed time: 747ns
=== RUN   TestCountFast/test_1:_2
    datatest.go:96: elapsed time: 999ns
=== RUN   TestCountFast/test_2:_3
    datatest.go:96: elapsed time: 6.118µs
=== RUN   TestCountFast/test_3:_4
    datatest.go:96: elapsed time: 5.317µs
=== RUN   TestCountFast/test_4:_5
    datatest.go:96: elapsed time: 1.462µs
=== RUN   TestCountFast/test_5:_6
    datatest.go:96: elapsed time: 6.227µs
=== RUN   TestCountFast/test_6:_7
    datatest.go:96: elapsed time: 2.373µs
=== RUN   TestCountFast/test_7:_8
    datatest.go:96: elapsed time: 7.17µs
=== RUN   TestCountFast/test_8:_9
    datatest.go:96: elapsed time: 5.923µs
=== RUN   TestCountFast/test_9:_10
    datatest.go:96: elapsed time: 9.233µs
--- PASS: TestCountFast (0.00s)
    --- PASS: TestCountFast/test_0:_1 (0.00s)
    --- PASS: TestCountFast/test_1:_2 (0.00s)
    --- PASS: TestCountFast/test_2:_3 (0.00s)
    --- PASS: TestCountFast/test_3:_4 (0.00s)
    --- PASS: TestCountFast/test_4:_5 (0.00s)
    --- PASS: TestCountFast/test_5:_6 (0.00s)
    --- PASS: TestCountFast/test_6:_7 (0.00s)
    --- PASS: TestCountFast/test_7:_8 (0.00s)
    --- PASS: TestCountFast/test_8:_9 (0.00s)
    --- PASS: TestCountFast/test_9:_10 (0.00s)
PASS
ok      command-line-arguments  0.001s
```
