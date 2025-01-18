// 'Hello world hello World people'
// arrange the string in way that the output contains the first occurence of the word,it should not be duplicate.also the case should be same

function ArrangeString(string) {
  let words = string.split(" ");
  let word_map = new Set();
  let result = [];
  console.log(words);
  for (const word of words) {
    let lowercase_word = word.toLowerCase();
    if (!word_map.has(lowercase_word)) {
      word_map.add(lowercase_word);
      result.push(word);
    }
  }
  //   console.log(result);
  //   return result.join(" ").toString();
  return result
    .sort((a, b) => b.localeCompare(a))
    .join(" ")
    .toString();
}
console.log(ArrangeString("Hello world hello World people"));
