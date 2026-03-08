export function reverseAlphabet(str: string): string {
  const alphabets = str.replace(/[0-9]/g, "").split("");
  const numbers = str.replace(/[^0-9]/g, "");
  
  return alphabets.reverse().join("") + numbers;
}
