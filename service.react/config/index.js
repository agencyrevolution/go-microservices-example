'use strict';

import _ from 'lodash';

const { PORT, API_ADDR } = process.env;

export default {
  /**
   * Front-End Server
   */
  server: {
    host: 'localhost',
    port: PORT || 3003
  },

  /**
   * API Server
   */
  apiServer: {
    urlPrefix: API_ADDR || 'http://localhost:3004'
  },

  /**
   * WebpackDevServer
   */
  webpackDevServer: {
    host: 'localhost',
    port: 9091
  },

  /**
   * browserSync
   */
  browserSyncServer: {
    host: 'localhost',
    port: 9092
  }
};
