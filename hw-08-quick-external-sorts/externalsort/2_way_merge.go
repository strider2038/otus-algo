package externalsort

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/sort"
)

var ErrCompleted = errors.New("completed")

// TwoWayMerge - сортирует файл filename с целыми числами с помощью алгоритма
// внешней сортировки с помощью двух промежуточных файлов. Считывает из исходного файла
// по limit чисел, сортирует их и записывает в исходные чанки.
// Затем итеративно считывает последовательности из исходных чанков, и с помощью
// алгоритма слияния записывает возрастающие последовательности в два файла.
// На последней итерации в первом чанке получается полностью отсортированный массив.
// См. https://en.wikipedia.org/wiki/External_sorting.
func TwoWayMerge(filename string, limit int) error {
	if err := makeTwoSortedChunks(filename, limit); err != nil {
		return fmt.Errorf("make sorted chunks: %w", err)
	}

	for {
		// итеративно объединяем отсортированные списки
		if err := mergeSort(filename); err != nil {
			if errors.Is(err, ErrCompleted) {
				break
			}
			return err
		}
	}

	// в первом чанке находится отсортированный массив
	sortedFilename := fmt.Sprintf("%s/chunk_0.txt", filepath.Dir(filename))

	// подмена старого файла отсортированным
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("remove origin file: %w", err)
	}
	if err := os.Rename(sortedFilename, filename); err != nil {
		return fmt.Errorf("replace origin file by sorted: %w", err)
	}

	return nil
}

func makeTwoSortedChunks(filename string, limit int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	chunkDir := filepath.Dir(filename)

	// создаем два чанка
	chunk1, err := os.Create(fmt.Sprintf("%s/chunk_0.txt", chunkDir))
	if err != nil {
		return fmt.Errorf("create chunk 1: %w", err)
	}
	defer chunk1.Close()
	chunk2, err := os.Create(fmt.Sprintf("%s/chunk_1.txt", chunkDir))
	if err != nil {
		return fmt.Errorf("create chunk 2: %w", err)
	}
	defer chunk2.Close()

	// последовательно сканируем исходный файл и создаем 2 чанка
	scanner := bufio.NewScanner(file)
	for {
		numbers, size, err := readChunk(scanner, limit)
		if err != nil {
			return err
		}
		if size == 0 {
			break
		}

		numbers = sort.QuickMiddle(numbers)
		if err := writeToFile(chunk1, numbers, size); err != nil {
			return fmt.Errorf("write to chunk: %w", err)
		}

		// меняем местами чанки
		chunk1, chunk2 = chunk2, chunk1
	}

	return nil
}

func mergeSort(filename string) error {
	chunkDir := filepath.Dir(filename)

	input1filename := fmt.Sprintf("%s/chunk_0.txt", chunkDir)
	input2filename := fmt.Sprintf("%s/chunk_1.txt", chunkDir)
	output1filename := fmt.Sprintf("%s/sorted_chunk_0.txt", chunkDir)
	output2filename := fmt.Sprintf("%s/sorted_chunk_1.txt", chunkDir)

	// открываем исходные чанки
	input1, err := ReadIntegersFromFile(input1filename)
	if err != nil {
		return fmt.Errorf("create chunk 1: %w", err)
	}
	defer input1.Close()
	input2, err := ReadIntegersFromFile(input2filename)
	if err != nil {
		return fmt.Errorf("create chunk 2: %w", err)
	}
	defer input2.Close()

	// создаем чанки для слияния
	output1, err := os.Create(output1filename)
	if err != nil {
		return fmt.Errorf("create chunk 1: %w", err)
	}
	defer output1.Close()
	output2, err := os.Create(output2filename)
	if err != nil {
		return fmt.Errorf("create chunk 2: %w", err)
	}
	defer output2.Close()

	// объединяем последовательности из входных файлов в выходные
	if err := mergeFiles(input1, input2, output1, output2); err != nil {
		return err
	}

	if err := os.Rename(output1filename, input1filename); err != nil {
		return fmt.Errorf("swap sorted chunks: %w", err)
	}
	if err := os.Rename(output2filename, input2filename); err != nil {
		return fmt.Errorf("swap sorted chunks: %w", err)
	}

	return nil
}

func mergeFiles(input1 *IntReader, input2 *IntReader, output1 *os.File, output2 *os.File) error {
	// количество открытых файлов
	n := 2

	value1, err := input1.Next()
	// если первый файл пуст, то входных данных не было
	if errors.Is(err, ErrEndOfList) {
		return ErrCompleted
	}
	if err != nil {
		return err
	}

	value2, err := input2.Next()
	// если второй файл пуст, то сортировка завершена
	if errors.Is(err, ErrEndOfList) {
		return ErrCompleted
	}
	if err != nil {
		return err
	}

	// максимумы нужны для отслеживания начала следующей отсортированной последовательности
	max1 := value1
	max2 := value2
	// флаг управляющий возможностью менять значения местами
	canSwap := true

	for {
		// если первое число больше второго, то меняем местами вместе со списками
		if value1 > value2 && canSwap {
			value1, value2 = value2, value1
			input1, input2 = input2, input1
			max1, max2 = max2, max1
		}

		// записываем минимальное число
		if _, err := io.WriteString(output1, strconv.Itoa(value1)+"\n"); err != nil {
			return fmt.Errorf(`write to file "%s": %w`, output1.Name(), err)
		}

		value1, err = input1.Next()
		// если список закончился, то меняем местами и уменьшаем n
		if errors.Is(err, ErrEndOfList) {
			n--
			// оба списка закончились
			if n == 0 {
				break
			}
			value1, value2 = value2, value1
			input1, input2 = input2, input1
			max1, max2 = max2, max1
			// больше менять местами нельзя
			canSwap = false
		} else if err != nil {
			return err
		}

		// увеличиваем максимумы
		if value1 >= max1 {
			max1 = value1
		}
		if value2 > max2 {
			max2 = value2
		}

		// на первом проходе: если первый список начался с начала, то меняем местами
		// со вторым и запрещаем им меняться;
		// если и второй список начался с начала, то разрешаем спискам меняться местами
		if n > 1 && value1 < max1 {
			canSwap = !canSwap
			value1, value2 = value2, value1
			input1, input2 = input2, input1
			max1, max2 = max2, max1
		}

		// если оба списка начались с начала
		// или остался один список и он начался с начала,
		// то меняем местами выходные файлы
		if value1 < max1 && (value2 < max2 || n == 1) {
			output1, output2 = output2, output1
			max1 = value1
			max2 = value2
		}
	}

	return nil
}
