import { countOccurrences } from "../src/3CountOccurrences";

test("countOccurrences harus mengembalikan jumlah kata query dalam array input", () => {
  const input = ["xc", "dz", "bbb", "dz"];
  const query = ["bbb", "ac", "dz"];
  expect(countOccurrences(input, query)).toEqual([1, 0, 2]);
});

test("countOccurrences harus mengembalikan jumlah kata query dalam array input", () => {
  const input = ["xc", "dz", "bbb", "dz"];
  const query = ["ac", "bbb", "dz"];
  expect(countOccurrences(input, query)).toEqual([0, 1, 2]);
});

test("countOccurrences harus mengembalikan jumlah kata query dalam array input", () => {
  const input = ["xc", "dz", "bbb", "dz"];
  const query = ["dz", "bbb", "xc"];
  expect(countOccurrences(input, query)).toEqual([2, 1, 1]);
});

test("countOccurrences harus mengembalikan jumlah kata query dalam array input", () => {
  const input = ["xc", "dz", "bbb", "dz"];
  const query = ["dz", "dz", "dz"];
  expect(countOccurrences(input, query)).toEqual([2, 2, 2]);
});

test("countOccurrences harus mengembalikan jumlah kata query dalam array input", () => {
  const input = ["xc", "dz", "bbb", "dz"];
  const query = ["xc", "ac", "xc"];
  expect(countOccurrences(input, query)).toEqual([1, 0, 1]);
});