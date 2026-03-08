export function countOccurrences(input: string[], query: string[]): number[] {
  return query.map((q) => input.filter((i) => i === q).length);
}
