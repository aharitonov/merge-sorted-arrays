<?php
require "merge_sorted_arrays.php";

class TestCase extends PHPUnit\Framework\TestCase
{
	/**
	 * @dataProvider dataProviderFunc
	 */
	public function testRun(array $a, array $b, array $expected, bool $isPositive=true)
	{
		if ($isPositive) {
			self::assertEquals($expected, merge_sorted_arrays($a, $b));
		} else {
			self::assertNotEquals($expected, merge_sorted_arrays($a, $b));
		}
	}

	public function dataProviderFunc(): array
	{
		return [
			// $a       $b           $expected
			[  [],      [],          []                 ],
			[  [],      [3],         [3]                ],
			[  [3,4],   [],          [3,4]              ],
			[  [1,2,5], [1,2,3,4,6], [1,1,2,2,3,4,5,6]  ],

			[  [1],     [1],         [1,1,1],    false  ],
		];
	}
}
