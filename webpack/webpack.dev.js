const webpackMerge = require('webpack-merge');
const WebpackNotifierPlugin = require('webpack-notifier');
const commonConfig = require('./webpack.common.js');
const helpers = require('../helpers');

module.exports = webpackMerge(commonConfig, {
    devtool: 'eval-source-map',
    entry: [
        'babel-polyfill',
        './client/index.js',
    ],
    output: {
        path: helpers.root('dist'),
        filename: 'bundle.js',
    },
    plugins: [
        new WebpackNotifierPlugin({
            alwaysNotify: true,
        }),
    ],
    devServer: {
        historyApiFallback: true,
        contentBase: './client',
        proxy: {
            '/graphql': {
                target: 'http://api:3000', // `api` being the alias of the Docker container - see: Docker-compose.yml
                secure: false,
            },
        },
    },
});
