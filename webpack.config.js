var path = require('path')
const VueLoaderPlugin = require('vue-loader/lib/plugin'); // плагин для загрузки кода Vue

module.exports =
[
    {
        entry: './src/components/byRoute/index/app.js',
        output: {
            path: path.resolve(__dirname, './static/js/index/'),
            publicPath: '/static/',
            filename: 'build.js'
        },
        module: {
            rules: [
                {
                    test: /\.vue$/,
                    loader: 'vue-loader'
                }, {
                    test: /\.css$/,
                    use: [
                        'vue-style-loader',
                        'css-loader'
                    ]
                }
            ]
        },
        plugins: [
            new VueLoaderPlugin()
        ]
    },
    {
        entry: './src/components/byRoute/litra/app.js',
        output: {
            path: path.resolve(__dirname, './static/js/litra/'),
            publicPath: '/static/',
            filename: 'build.js'
        },
        module: {
            rules: [
                {
                    test: /\.vue$/,
                    loader: 'vue-loader'
                }, {
                    test: /\.css$/,
                    use: [
                        'vue-style-loader',
                        'css-loader'
                    ]
                }
            ]
        },
        plugins: [
            new VueLoaderPlugin()
        ]
    }
];