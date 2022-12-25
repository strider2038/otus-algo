package linearsort

func RadixSort[T Sortable](items []T, max T, bits int) []T {
	// количество разрядов максимального числа
	// служит для итерации по разрядам
	radixCount := 0
	for i := int(max); i > 0; i >>= bits {
		radixCount++
	}

	radix := 1 << bits

	// битовая маска для нахождения остатка от деления на разряд
	mask := T(1)
	for i := 1; i < bits; i++ {
		mask = mask<<1 | 1
	}

	// подсчет количества чисел по разрядам
	// двумерная матрица [radix][radixCount]
	counts := make([]int, radix*radixCount)

	for i := 0; i < len(items); i++ {
		item := items[i]
		for r := 0; r < radixCount; r++ {
			offset := r * radix
			// индекс счетчика вычисляется нахождением остатка от деления
			// с помощью битовой маски
			x := int(item&mask) + offset
			counts[x]++ // эквивалентно операции в матрице counts[r][i]++
			// делим число с помощью битового сдвига
			item >>= bits
		}
	}

	// расчет позиций для размещения чисел поразрядно
	for r := 0; r < radixCount; r++ {
		offset := r * radix
		for i := 1; i < radix; i++ {
			counts[offset+i] += counts[offset+i-1] // эквивалентно counts[r][i] += counts[r][i-1]
		}
	}

	// итеративная сортировка алгоритмом подсчета
	sorted := make([]T, len(items))
	for r := 0; r < radixCount; r++ {
		offset := r * radix
		for i := len(items) - 1; i >= 0; i-- {
			t := items[i]
			// индекс счетчика находится делением на текущий номер разряда
			// (с помощью битового смещения), а затем нахождением остатка
			// от деления (с помощью битовой маски)
			index := offset + int((t>>(bits*r))&mask)
			counts[index]--
			sorted[counts[index]] = t
		}
		// после сортировки по одному разряду меняем местами массивы
		items, sorted = sorted, items
	}

	return items
}
