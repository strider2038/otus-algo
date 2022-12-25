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

// KWayMerge - сортирует файл filename с целыми числами с помощью алгоритма
// внешней сортировки. Разбивает исходный файл на k файлов (чанков) с максимальным размером chunkSize.
// Каждый чанк сортируется алгоритмом быстрой сортировки.
// Затем все чанки объединяются с помощью алгоритма k-way merge.
// В случае успешного завершения сортировки заменяет исходный файл отсортированным.
// См. https://en.wikipedia.org/wiki/External_sorting.
func KWayMerge(filename string, chunkSize int) error {
	// разделяем входной файла на k-отсортированных файлов
	chunkCount, err := makeKSortedChunks(filename, chunkSize)
	if err != nil {
		return fmt.Errorf("make sorted chunks: %w", err)
	}

	// открываем на чтение все отсортированный файлы
	// в ОП находится только прочитанное число и файловый указатель для чтения следующего
	chunks, err := openSortedChunks(filename, chunkCount)
	defer func() {
		for i := 0; i < len(chunks); i++ {
			chunks[i].Close()
		}
	}()
	if err != nil {
		return fmt.Errorf("open sorted chunks: %w", err)
	}

	// создаем min-кучу
	heap, err := NewMinHeapFromReaders(chunks...)
	if err != nil {
		return fmt.Errorf("create heap: %w", err)
	}

	// сортируем с использованием кучи и заменяем исходный файл
	if err := kWayMergeSort(filename, heap); err != nil {
		return fmt.Errorf("merge sorted chunks: %w", err)
	}

	return nil
}

func makeKSortedChunks(filename string, chunkSize int) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	chunkDir := filepath.Dir(filename)

	// последовательно сканируем исходный файл и создаем k-чанков, каждый
	// из которых ограничен максимальным размером chunkSize
	scanner := bufio.NewScanner(file)
	chunkCount := 0
	for {
		chunk, size, err := readChunk(scanner, chunkSize)
		if err != nil {
			return 0, err
		}
		if size == 0 {
			break
		}

		chunk = sort.QuickMiddle(chunk)
		chunkName := fmt.Sprintf("%s/chunk_%d.txt", chunkDir, chunkCount)
		if err := writeChunk(chunk, size, chunkName); err != nil {
			return 0, fmt.Errorf("write chunk %d: %w", chunkCount, err)
		}

		chunkCount++
	}

	return chunkCount, nil
}

func openSortedChunks(filename string, chunkCount int) ([]*IntReader, error) {
	var err error

	chunkDir := filepath.Dir(filename)

	chunks := make([]*IntReader, chunkCount)
	for i := 0; i < chunkCount; i++ {
		chunkName := fmt.Sprintf("%s/chunk_%d.txt", chunkDir, i)
		chunks[i], err = ReadIntegersFromFile(chunkName)
		if err != nil {
			return chunks, fmt.Errorf("open chunk: %w", err)
		}
	}

	return chunks, nil
}

func kWayMergeSort(filename string, heap *MinHeap) error {
	// создаем временный файл, в который будет записываться отсортированный массив
	sortedFile, err := os.CreateTemp(".", "sorted")
	if err != nil {
		return fmt.Errorf("create temporary file: %w", err)
	}
	defer sortedFile.Close()
	defer os.Remove(sortedFile.Name())

	// основной цикл k-merge сортировки
	for {
		min, err := heap.GetMin()
		// значения закончились
		if errors.Is(err, ErrEndOfList) {
			break
		}
		// системная ошибка
		if err != nil {
			return err
		}
		if _, err := sortedFile.WriteString(strconv.Itoa(min) + "\n"); err != nil {
			return fmt.Errorf(`write to file: %w`, err)
		}
	}

	// подмена старого файла отсортированным
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("remove origin file: %w", err)
	}
	if err := os.Rename(sortedFile.Name(), filename); err != nil {
		return fmt.Errorf("replace origin file by sorted: %w", err)
	}

	return nil
}

func readChunk(scanner *bufio.Scanner, chunkSize int) ([]int, int, error) {
	chunk := make([]int, chunkSize)
	for i := 0; i < chunkSize; i++ {
		if !scanner.Scan() {
			return chunk, i, nil
		}
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, 0, fmt.Errorf("parse int: %w", err)
		}
		chunk[i] = number
	}

	return chunk, chunkSize, nil
}

func writeChunk(numbers []int, size int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	return writeToFile(file, numbers, size)
}

func writeToFile(file *os.File, numbers []int, size int) error {
	for i := 0; i < size; i++ {
		s := strconv.Itoa(numbers[i])
		if _, err := io.WriteString(file, s+"\n"); err != nil {
			return fmt.Errorf(`write to file "%s": %w`, file.Name(), err)
		}
	}

	return nil
}
