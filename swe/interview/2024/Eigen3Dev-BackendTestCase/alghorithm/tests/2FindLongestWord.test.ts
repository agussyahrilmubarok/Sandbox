import { findLongestWord } from "../src/2FindLongestWord";

test("findLongestWord harus mengembalikan kata terpanjang dalam kalimat", () => {
  expect(findLongestWord("Saya sangat senang mengerjakan soal algoritma")).toBe("mengerjakan");
});

test("findLongestWord harus mengembalikan kata terpanjang dalam kalimat", () => {
  expect(findLongestWord("Pemrograman adalah seni memecahkan masalah")).toBe("Pemrograman");
});

test("findLongestWord harus mengembalikan kata terpanjang dalam kalimat", () => {
  expect(findLongestWord("Vue.js adalah framework JavaScript yang progresif")).toBe("JavaScript");
});

test("findLongestWord harus mengembalikan kata terpanjang dalam kalimat", () => {
  expect(findLongestWord("TypeScript adalah superset dari JavaScript")).toBe("TypeScript");
});

test("findLongestWord harus mengembalikan kata terpanjang dalam kalimat", () => {
  expect(findLongestWord("Algoritma dan struktur data adalah dasar dari pemrograman")).toBe("pemrograman");
});