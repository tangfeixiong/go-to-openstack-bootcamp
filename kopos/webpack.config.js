var path = require('path');
var webpack = require('webpack');
var CleanPlugin = require('clean-webpack-plugin');
var ExtractPlugin = require('extract-text-webpack-plugin');
var AssetsPlugin = require('assets-webpack-plugin');

var argv = process.argv;
var DEV_HOT = false;

//TODO change to parse argv instead
if(process.env.DEV_HOT){
  DEV_HOT = true;
}

var production = process.env.NODE_ENV === 'production';
var plugins = [
  new webpack.NoErrorsPlugin(),
  new ExtractPlugin('[name]-[contenthash].css'),
  new AssetsPlugin({
    fullPath: false,
    prettyPrint: true,
  }),
];

if (production) {
  plugins = plugins.concat([
    new CleanPlugin('public/bundles'),
    new webpack.optimize.DedupePlugin(),
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.optimize.CommonsChunkPlugin({
      name: 'main',
      children: true,
      minChunks: 2,
    }),
    new webpack.optimize.MinChunkSizePlugin({
      minChunkSize: 51200, // ~50kb
    }),
    new webpack.optimize.UglifyJsPlugin({
      mangle: true,
      compress: {
        warnings: false, // Suppress uglification warnings
      },
    }),
    new webpack.DefinePlugin({
      __SERVER__: !production,
      __DEVELOPMENT__: !production,
      __DEVTOOLS__: !production,
      'process.env': {
        BABEL_ENV: JSON.stringify(process.env.NODE_ENV),
        NODE_ENV: JSON.stringify(process.env.NODE_ENV),
      },
    }),
  ]);
}

var config = {
  debug: !production,
  devtool: production ? false : 'eval',
  plugins: plugins,

  entry: './assets/index.js',
  output: {
    path: 'public/bundles',
    publicPath: '/assets/bundles/',
    filename: '[name]-[hash].js',
    chunkFilename: '[name]-[chunkhash].js',
  },

  resolve: {
    extensions: ['', '.js', '.jsx', '.css'],
    alias: {
      '#coms': path.join(__dirname, 'assets/components'),
    },
  },

  module: {
    loaders: [
      { test: /\.css$/, loader: ExtractPlugin.extract('style', 'css') },
      { test: /\.scss$/, loader: ExtractPlugin.extract('style', 'css!sass') },
      { test: /\.html$/, loader: 'html' },
      { test: /\.(png|gif|svg)$/, loader: 'url?name=[name]@[hash].[ext]&limit=5000' },
      { test: /\.(pdf|ico|jpg|eot|otf|woff|ttf|mp4|webm)$/, loader: 'file?name=[name]@[hash].[ext]' },
      {
        test: /\.jsx?$/,
        loader: "babel",
        query: {
          presets: ['es2015', 'react', 'stage-0'],
        },
        include: __dirname + '/assets',
        exclude: /(node_modules|bower_components)/,
      },
    ],
  },
};

//merge hot reload config

if(DEV_HOT){
  config.devServer = {
    hot: true,
    inline: true,
    debug: true,
    progress: true,
    port: 8080,
    proxy: [{
      // for all not hot-update request
      path:    /^(?!.*\.hot-update\.js)(.*)$/,
      target: 'http://localhost:'+ process.env.PORT || 9000
    }],
    // contentBase:'http://localhost:9000',
    port: process.env.DEV_HOT_PORT || 8090,
    devtool: 'eval-source-map',
    open: true,
    watchOptions: {
      aggregateTimeout: 300,
      poll: 1000
    },
    open: true,
    stats: { colors: true }
  };

  config.plugins.unshift(new webpack.HotModuleReplacementPlugin());

  var babelLoader = config.module.loaders[config.module.loaders.length - 1];
  babelLoader.query.presets.unshift('react-hmre');
  babelLoader.query.plugins = babelLoader.query.plugins || [];
  babelLoader.query.plugins.push([
    'react-transform', {
      transforms: [{
        transform : 'react-transform-hmr',
        imports   : ['react'],
        locals    : ['module']
      }]
    }
  ]);
  config.output.publicPath = "http://localhost:" + config.devServer.port + config.output.publicPath
}

module.exports = config;
