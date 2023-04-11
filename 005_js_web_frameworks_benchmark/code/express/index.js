const express = require('express');
const { fibonacciSequence } = require('../fibonacci_sequence.js');
const app = express();
const port = 8000;

app.get('/', (req, res) => {
  res.end();
});

app.get('/:number', (req, res) => {
  const num = req.params.number;  
  res.send({result: fibonacciSequence(num)});
});

app.listen(port, () => {
  console.log(`Express server listening at ${port}`);
});
