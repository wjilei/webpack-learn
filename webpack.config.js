const path = require('path');
const MiniCssPlugin = require("mini-css-extract-plugin");
const HtmlWebpackPlugin = require('html-webpack-plugin');
const VueLoaderPlugin = require('vue-loader/lib/plugin');
module.exports = {
    entry: {
        'main': './src/main.js',
        'page1/index': './src/page1/page1.js',
    },
    mode: 'development',
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname,'./build')
    },
    resolve: {
		
        alias: {
            'vue$': 'vue/dist/vue.esm.js' 
        }
         
    },
    module:{
        rules: [
            {
                test: /\.vue$/,
                loader: 'vue-loader',
            },
            {
                test: /\.css$/,
                use: [MiniCssPlugin.loader, 'css-loader'],
            },
            {
                test: /\.(gif|png|jpg)$/,
                use: [
                    {
                        loader: 'file-loader',
                        options: {
                            name: "images/[name].[ext]",
                            esModule: false,
                            publicPath: './',
                        }
                    }
                ],
            },
            {
                test: /\.(js|jsx)$/,
                use: 'babel-loader',
                exclude: /node_modules/
            }
        ],
    },
    // 开发服务器设置  
    devServer: {
        static: {
            directory: path.resolve(__dirname,  "build"),
            watch: true,
        },
        compress: true,
        liveReload: true,
        port: 8090,
    },
    devtool: "source-map",
    plugins: [
        new MiniCssPlugin({
            filename: "css/[name].css",
        }),
        new HtmlWebpackPlugin({
            chunks: ['main'],
            // 该 html 文件引用的 chunks 数组，chunk 的概念将在稍后的内容中介绍
            filename: 'index.html', // 构建后的文件名
            template: './src/index.html' // 源文件名
        }),
        new HtmlWebpackPlugin({
            chunks: ['page1/index'],
            // 该 html 文件引用的 chunks 数组，chunk 的概念将在稍后的内容中介绍
            filename: 'page1/index.html', // 构建后的文件名
            template: './src/page1/index.html' // 源文件名
        }),
        new VueLoaderPlugin(),
    ]
};