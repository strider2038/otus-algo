#!/bin/sh

rm ./var/*

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
