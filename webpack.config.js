let webpack = require('webpack');
let path = require('path');

let BUILD_DIR = path.resolve(__dirname, 'public/build');
let APP_DIR = path.resolve(__dirname, 'public/app');

let config = {
    entry: {
        index: APP_DIR + '/index.jsx',
    },
    output: {
        path: BUILD_DIR,
        filename: '[name].bundle.js'
    },
    mode: "development",
    module: {
        rules: [
            {
                test: /\.jsx?/,
                include: APP_DIR,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['env', 'react']
                    }
                }
            }
        ]
    },
    resolve: {
        extensions: ['.js', '.jsx']
    }

};

module.exports = config;