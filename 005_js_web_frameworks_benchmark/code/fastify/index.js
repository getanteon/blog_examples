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
