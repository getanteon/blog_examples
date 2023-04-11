function fibonacciSequence(num) {
    let result = [];
    let num1 = 0;
    let num2 = 1;
    let nextTerm;
    for (let i = 1; i <= num; i++) {
        result.push(num1);
        nextTerm = num1 + num2;
        num1 = num2;
        num2 = nextTerm;
  }
  
  return result;
}

module.exports = { fibonacciSequence }