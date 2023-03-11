# Алгоритм Run-Length Encoding (RLE)

## Цель

Создать программу сжатия файлов на основе алгоритма RLE

## Задание

* Написать функцию сжатия массива по алгоритму RLE
* Написать программу для сжатия файлов
* Написать программу для распаковки файлов
* При запуске программы без параметров она должна выводить краткую инструкцию, как её запускать для сжатия/распаковки
  файлов.
* Реализовать улучшенный алгоритм RLE: сжатие и распаковку
* Сравнить работу программы с разными типами файлов: текст, фото, аудио, zip-архив.
* Составить отчёт сравнения результата работы двух версий алгоритма с разными файлами.

## Результаты

* Реализация неоптимизированного алгоритма RLE представлена в пакете `rle`
* Реализация оптимизированного варианта в пакете `rle2`

## Тестирование приложений

Запуск тестов

```shell
go test -v ./...
```

```shell
# компиляция приложений
cd cmd/rle
go build
cd ../..

cd cmd/rle2
go build
cd ../..

# сжатие и распаковка тестовых файлов (RLE без оптимизации)
./cmd/rle/rle compress   ./testdata/real_text.txt ./var/real_text.rle
./cmd/rle/rle decompress ./var/real_text.rle ./var/rle_real_text.txt

./cmd/rle/rle compress   ./testdata/magic_square.txt ./var/magic_square.rle
./cmd/rle/rle decompress ./var/magic_square.rle ./var/rle_magic_square.txt

./cmd/rle/rle compress   ./testdata/voice.wav ./var/voice.rle
./cmd/rle/rle decompress ./var/voice.rle ./var/rle_voice.wav

./cmd/rle/rle compress   ./testdata/gopher.bmp ./var/gopher.rle
./cmd/rle/rle decompress ./var/gopher.rle ./var/rle_gopher.bmp

./cmd/rle/rle compress   ./testdata/gopher.bmp.zip ./var/gopher-zip.rle
./cmd/rle/rle decompress ./var/gopher-zip.rle ./var/rle_gopher-bmp.zip

# сжатие и распаковка тестовых файлов (RLE с оптимизацией)
./cmd/rle2/rle2 compress   ./testdata/real_text.txt ./var/real_text.rle2
./cmd/rle2/rle2 decompress ./var/real_text.rle2 ./var/rle2_real_text.txt

./cmd/rle2/rle2 compress   ./testdata/magic_square.txt ./var/magic_square.rle2
./cmd/rle2/rle2 decompress ./var/magic_square.rle2 ./var/rle2_magic_square.txt

./cmd/rle2/rle2 compress   ./testdata/voice.wav ./var/voice.rle2
./cmd/rle2/rle2 decompress ./var/voice.rle2 ./var/rle2_voice.wav

./cmd/rle2/rle2 compress   ./testdata/gopher.bmp ./var/gopher.rle2
./cmd/rle2/rle2 decompress ./var/gopher.rle2 ./var/rle2_gopher.bmp

./cmd/rle2/rle2 compress   ./testdata/gopher.bmp.zip ./var/gopher-zip.rle2
./cmd/rle2/rle2 decompress ./var/gopher-zip.rle2 ./var/rle2_gopher-bmp.zip
```

## Сравнение эффективности сжатия

| Файл                                | Исходный размер, байт | RLE, байт | RLE2, байт |
|-------------------------------------|----------------------:|----------:|-----------:|
| Текстовый файл (текст)              |                13 343 |    28 626 |     14 475 |
| Текстовый файл (магический квадрат) |                 1 364 |       824 |        774 |
| Звуковой файл WAV                   |               833 336 | 1 644 406 |    840 234 |
| Изображение PNG                     |               839 094 |   874 742 |    498 200 |
| Архив ZIP                           |                58 706 |   116 100 |     59 434 | 

## Выводы

Алгоритм RLE является одним из простейших алгоритмов сжатия.
В общем случае он малоэффективен. Подходит только для определенных видов данных,
в которых есть много повторяющихся символов.
