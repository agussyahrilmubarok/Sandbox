import { reverseAlphabet } from "../src/1ReverseAlphabet";

test("reverseAlphabet harus membalikkan bagian alfabet dan menjaga angka di akhir", () => {
  expect(reverseAlphabet("NEGIE1")).toBe("EIGEN1");
});

test("reverseAlphabet harus membalikkan bagian alfabet dan menjaga angka di akhir", () => {
  expect(reverseAlphabet("AGU5")).toBe("UGA5");
});
