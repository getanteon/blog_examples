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
  console.log(`Koa server listening at ${port}`);
});
