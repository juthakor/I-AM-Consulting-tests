function isPalindrom(word) {
    word = word.toLowerCase();
    for (let i = 0; i < Math.floor(word.length / 2); i++) {
        if (word[i] !== word[word.length - i - 1]) {
            return false
        }
    }
    return true
}
var word = 'Deleveled';
console.log(word, isPalindrom(word));
