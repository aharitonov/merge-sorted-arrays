<?php

// Условие задачи
// Нужно слить 2 отсортированных массива в один отсортированный
// массив

// Входные параметры
// 2 отсортированных (по возрастанию) массива A и B длины M и N
// Вывод
// Отсортированный (по возрастанию) массив, состоящий из элементов первых двух

// Примеры
// Ввод [1, 2, 5], [1, 2, 3, 4, 6]
// Вывод [1, 1, 2, 2, 3, 4, 5, 6]

function merge_sorted_arrays(array $a, array $b): array
{
	if (empty($a) && empty($b)) {
		return [];
	}
	elseif (empty($a)) {
		return $b;
	}
	elseif (empty($b)) {
		return $a;
	}

	$c = [];
	$lenA = count($a);
	$lenB = count($b);
	$iMaxA = $a[0];
	$iMaxB = $b[0];
	$indexA = 0;
	$indexB = 0;

	while ($indexA < $lenA || $indexB < $lenB) {
		while ($indexA < $lenA && $iMaxA < $iMaxB) {
			$v = $a[$indexA];
			$c[] = $v;
			$iMaxA = $a[++$indexA] ?? PHP_INT_MAX;
		}
		while ($indexB < $lenB && $iMaxB <= $iMaxA) {
			$v = $b[$indexB];
			$c[] = $v;
			$iMaxB = $b[++$indexB] ?? PHP_INT_MAX;
		}
	}
	return $c;
}

// Test
// $c = merge_sorted_arrays($a = [1,2,5], $b = [1,2,3,4,6]);
// echo ($c === [1, 1, 2, 2, 3, 4, 5, 6]) ? 'TEST OK' : 'TEST FAIL';
// echo "\n\n";
