module.exports = {
  development: {
    client: 'postgresql',
    connection: {
      database: 'ad_server'
    }
  },

  features: {
    client: 'postgresql',
    connection: {
      database: 'ad_server_features'
    },
    migrations: {
      directory: './migrations',
    },
    seeds: {
      directory: './seeds',
    },
  }
};
