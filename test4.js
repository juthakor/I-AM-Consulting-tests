function findSum(n) {
    let sum = 0;
    for (let i = 1; i <= n; i++) {
        sum += i ** i;
    }
    return sum
}
console.log("3: " + findSum(3));
