# Bit manipulation

## Logic gates

| --- | AND | OR | XOR |NAND|
|-----|-----|----|-----|----|
| 0 0 | 0   | 0  | 0   | 1  |
| 0 1 | 0   | 1  | 1   | 0  |
| 1 0 | 0   | 1  | 1   | 0  |
| 1 1 | 1   | 1  | 0   | 0  |

- XOR: when inputs are different, return 1
- NAND: NOT(AND)

## Bit manipulation in Go

|Operation    |Result   |Description|
|-------------|---------|-----------|
|0011 & 0101  |0001     |AND|
|0011 \| 0101 |0111     |OR|
|0011 ^ 0101  |0110     |XOR|
|^0101        |1010     |NOT (same as 1111 ^ 0101)|
|1111 &^ 0111 |1000     |AND NOT e.g. AND_NOT(a, b) = AND(a, NOT(b)) clears all bits in position of `1` in second number |
|00110101<<2  |11010100 |Left shift by specified number (`<<1` is written as `<<`)|
|00110101>>2  |00001101 |Right shift|
