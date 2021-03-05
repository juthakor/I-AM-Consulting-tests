let permArr = [], usedChars = [];
function permute(input) {
    let i, ch;
    for (i = 0; i < input.length; i++) {
        ch = input.splice(i, 1)[0];
        usedChars.push(ch);
        if (input.length == 0) {
            permArr.push(usedChars.slice().join(""));
        }
        permute(input);
        input.splice(i, 0, ch);
        usedChars.pop();
    }
    return permArr
};

function main(str) {
    let strArr = str.split("");
    let result = permute(strArr);
    result = result.filter(function (item, pos) {
        return result.indexOf(item) == pos;
    })
    return result
}

console.log(main("ABC"));