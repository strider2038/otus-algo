# Домашнее задание №19 "Алгоритм Кнутта-Морриса-Пратта"

## Цель

Реализовать алгоритм Кнутта-Морриса-Пратта

## Задание

* Составить конечный автомат и прохождение по нему для поиска шаблона в строке.
* Самостоятельно написать функцию вычисления префикс-функции, медленный вариант.
* Переписать алгоритм быстрого вычисления префикс-функции и разобраться в нём.
* Реализовать алгоритм Кнута-Морриса-Пратта.
* Протестировать работу алгоритмов для разных начальных данных.
* Составить сравнительную таблицу по тестам и написать вывод.
* Для точного расчёта времени можно прогонять тест T раз и результат делить на T, где T = 10, 100, 1000 или ещё больше.

## Сравнение производительности

Для сравнения производительности алгоритмов использовались нативные инструменты benchmark.

В качестве строки и шаблона поиска было использовано три варианта.

1. Строка "ABCAABAABBAAABAABAABAAABABAABBAAABAAAB", паттерн "AABAABAAABA".
2. Строка с небольшой фразой на естественном языке.
3. Большая строка и маленький шаблон с совпадением текста посередине.

В таблице приведено среднее время работы алгоритма. В последнем варианте приведено время работы
КМП алгоритма с предварительным парсингом шаблона и составлением префиксов шаблонов.

| Алгоритм              | Строка "aab" |     Фраза | Большая строка |
|-----------------------|-------------:|----------:|---------------:|
| Полный перебор        |       116 ns |    785 ns |     100 547 ns |
| Префикс шаблона       |        37 ns |    134 ns |       6 669 ns |
| Суффикс текста        |        36 ns |    116 ns |       5 168 ns |
| Конечный автомат      |    22 148 ns | 25 650 ns |     253 737 ns |
| КМП, медленный шаблон |       342 ns |    448 ns |       8 307 ns |
| КМП, быстрый шаблон   |       128 ns |    227 ns |       5 984 ns |
| КМП, предрасчет       |        34 ns |    146 ns |       5 474 ns |

## Выводы

Любопытно отметить, что в простых случаях эффективность алгоритма Кнутт-Морриса-Пратта
не отличается сильно от простых алгоритмов поиска по префиксу шаблона и суффиксу текста.
Заметный выигрыш можно обнаружить только в случае сложных шаблонов с повторяющимися
последовательностями.

## Запуск кода

Запуск тестов

```shell
go test -v ./...
```
