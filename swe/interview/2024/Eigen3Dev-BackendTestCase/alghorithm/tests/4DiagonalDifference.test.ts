import { diagonalDifference } from "../src/4DiagonalDifference";

test("diagonalDifference harus mengembalikan selisih absolut antara jumlah diagonal matriks", () => {
  const matrix = [
    [1, 2, 0],
    [4, 5, 6],
    [7, 8, 9],
  ];
  expect(diagonalDifference(matrix)).toBe(3);
});

test("diagonalDifference harus mengembalikan selisih absolut antara jumlah diagonal matriks", () => {
  const matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
  ];
  expect(diagonalDifference(matrix)).toBe(0);
});

test("diagonalDifference harus mengembalikan selisih absolut antara jumlah diagonal matriks", () => {
  const matrix = [
    [2, 3, 4],
    [5, 6, 7],
    [8, 9, 1],
  ];
  expect(diagonalDifference(matrix)).toBe(9);
});