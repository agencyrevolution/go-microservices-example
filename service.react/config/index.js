'use strict';

import _ from 'lodash';

const { addr, apiAddr } = process.env;

export default {
  /**
   * Front-End Server
   */
  server: {
    host: 'localhost',
    port: addr ? parseInt(_.last(addr.toString().split(':'))) : 8080
  },

  /**
   * API Server
   */
  apiServer: {
    urlPrefix: apiAddr || 'https://api.github.com'
  },

  /**
   * WebpackDevServer
   */
  webpackDevServer: {
    host: 'localhost',
    port: 8081
  },

  /**
   * browserSync
   */
  browserSyncServer: {
    host: 'localhost',
    port: 8082
  }
};
