'use strict';

import config from '../config';
import express from 'express';
import proxy from 'express-http-proxy';

const app = express();
const env = process.env.NODE_ENV || 'production';


// Serve static files
// --------------------------------------------------
import path from 'path';
app.use(express.static(path.resolve(process.cwd(), 'dist')));

app.use('/favicon.ico', function (req, res) {
  res.redirect('http://github.com/favicon.ico');
});

// View engine
// --------------------------------------------------
import expressHandlebars from 'express-handlebars';
import handlebars from 'handlebars';

handlebars.registerHelper('json-stringify', ::JSON.stringify);

app.engine('hbs', expressHandlebars());
app.set('view engine', 'hbs');

// Proxy to API server
// --------------------------------------------------
import url from 'url';

app.use('/api/*', proxy(config.apiServer.urlPrefix, {
  forwardPath: req => {
    return url.parse(req.originalUrl).path.replace(/^\/api/, '');
  }
}));


// Render layout
// --------------------------------------------------
import render from '../lib/render';

app.get('/*', (req, res) => {
  // Js files
  const jsPaths = [ 'vendor', 'main' ].map(basename => {
    if (env === 'development') {
      let { webpackDevServer: { host, port } } = config;
      return `//${host}:${port}/js/${basename}.js`;
    }
    return `/js/${basename}.js`;
  });

  if (env === 'development') {
    let { browserSyncServer: { host, port } } = config;
    const BSVersion = require('browser-sync/package.json').version;
    jsPaths.push(`//${host}:${port}/browser-sync/browser-sync-client.${BSVersion}.js`);
  }

  // Render
  const layout = 'layouts/main';
  const payload = {
    jsPaths,
    initialState: {},
    body: ''
  };

  // if (env === 'development') {
  //   return res.render(layout, payload);
  // }

  render(req, res, layout, {
    payload
  });
});

const server = app.listen(config.server.port, () => {
  const { address: host, port } = server.address();
  console.log(`Front-End server is running at ${host}:${port}`); // eslint-disable-line no-console
});
