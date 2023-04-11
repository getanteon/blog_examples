# Performance Benchmarking Popular JS Web Frameworks - Express.js vs. Koa.js vs. Fastify 

## Initialize the Project

```bash
npm init
```

## Install the Dependincies

### Express

```bash
npm install express 
```

### Koa

```bash
npm install koa koa-router 
```

### Fastify

```bash
npm install fastify
```


## Create Index.js file for Frameworks and Add Endpoints

```bash
touch index.js
```

After creating the index.js file, paste the code below for the framework you wish to use.

### Express

```javascript
const express = require('express');
const { fibonacciSequence } = require('../fibonacci_sequence.js');
const app = express();
const port = 8000;

app.get('/', (req, res) => {
  res.end();
});

app.get('/:number', (req, res) => {
  const num = req.params.number;  
  res.status(200).send({result: fibonacciSequence(num)});
});

app.listen(port, () => {
  console.log(`Express server listening at http://0.0.0.0:${port}`);
});
```

### Koa

```javascript
const Koa = require('koa');
const Router = require('koa-router');
const { fibonacciSequence } = require('../fibonacci_sequence');
const app = new Koa();
const router = new Router();
const port = 8000;

router.get('/', async (ctx, next) => {  
  ctx.body = '';
});

router.get('/:number', async (ctx) => {
  const num = ctx.params.number; 
  ctx.body = fibonacciSequence(num);
});

app.use(router.routes()).use(router.allowedMethods());

app.listen(port, "0.0.0.0", () => {
  console.log(`Koa server listening at http://localhost:${port}`);
});
```

### Fastify

```javascript
const fastify = require('fastify')();
const { fibonacciSequence } = require('../fibonacci_sequence.js');

const port = 8000;

fastify.get('/', function (request, reply) {
  return '';
});

fastify.get('/:number', (request, reply) => {
  const num = request.params.number;
  return fibonacciSequence(num);
});

fastify.listen({port: port, host: '0.0.0.0'}, (err, address) => {
  if (err) {
    fastify.log.error(err);
    process.exit(1);
  }
  fastify.log.info(`Fastify server listening at ${address}`);
});
```

### Fibonacci Sequence Function
Add the fibonacci sequence calculation function in either the server code file or in another file.

```javascript
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
```

If you place the function in another file (as we did) do not forget to export it to outside with using:

```javascript
module.exports = { fibonacciSequence }
```

And you should also import it to index.js or to the file that you want to use with:

```javascript
const { fibonacciSequence } = require('../fibonacci_sequence.js');
```

## Run the server

You can either start the server directly from the command line using node:

```bash
node index.js
```

Or it is also possible to change the package.json file in your project and set a start script.
```json
"main": "index.js",
"scripts": {
  "start": "node index.js"
}
```

After setting up the package.json file, now you can start your server with:

```bash
npm start
```


## Ddosify Engine

Download Ddosify Engine from [here](https://github.com/ddosify/ddosify)

📌 The general run format of Ddosify Engine
```bash
ddosify -t http://localhost:8000/ -d <Duration> -n <IterationNum>
```

📌 Change the duration and iteration number depending on the test case

#### - 10 requests per second
```bash
ddosify -t http://localhost:8000/ -d 10 -n 100
```

#### - 100 requests per second
```bash
ddosify -t http://localhost:8000/ -d 10 -n 1000
```

#### - 1000 requests per second
```bash
ddosify -t http://localhost:8000/ -d 10 -n 10000
```

📌 Usage of random number generator for Fibonacci Sequence endpoint

```bash
ddosify -t http://localhost:8000/{{_randomInt}} -d <Duration> -n <IterationNum>
```

The iteration number and the duration can be again changed depending on the test case.
