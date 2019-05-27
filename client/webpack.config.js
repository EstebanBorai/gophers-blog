/* eslint-disable */

const path = require('path');
const DefinePlugin = require('webpack').DefinePlugin;
const CleanWebpackPlugin = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const HotModuleReplacementPlugin = require('webpack').HotModuleReplacementPlugin;

module.exports = {
	entry: './src/index.js',
	output: {
		filename: 'index.js',
		path: path.resolve(__dirname, 'bundle')
	},
	module: {
		rules: [
			{
				test: /\.(js|jsx)$/,
				exclude: /node_modules/,
				use: {
					loader: 'babel-loader'
				}
			},
			{
				test: /\.css$/,
				use: [
					'style-loader',
					'css-loader'
				]
      },
      {
        test: /\.scss$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader'
        ]
      },
			{
				test: /\.(png|jpg|gif|svg)$/,
				use: [
					{
						loader: 'file-loader',
						options: {
							name: 'assets/[name].[ext]',
						}
					}
				]
			},
			{
				test: /\.(woff|woff2|eot|ttf|otf)$/,
				loader: 'file-loader',
				options: {
					name: 'assets/[name].[ext]'
				}
			}
		]
	},
	devServer: {
		contentBase: path.join(__dirname, 'src'),
		compress: true,
		port: 4200
	},
	plugins: [
		new HtmlWebpackPlugin({
			template: 'public/index.html'
		}),
    new HotModuleReplacementPlugin(),
    new CleanWebpackPlugin(),
    new DefinePlugin({
      API_URL: JSON.stringify('http://localhost:8080/')
    }),
	],
	resolve: {
		extensions: ['.js', '.jsx']
	}
};
