const redis = require('redis');
const { promisify } = require('util');

const client = redis.createClient({ url: process.env['REDIS_HOST'] });

const get = promisify(client.get).bind(client);
const set = promisify(client.set).bind(client);

module.exports = { get, set }
