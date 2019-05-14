const path = require('path');

module.exports = {
	entry: path.resolve(__dirname, 'app.jsx'),
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: 'bundle.js'
	},
	resolve: {
		extensions: ['.js', '.jsx']
	},
	module: {
		rules: [
			{
				test: /\.jsx/,
				use: {
					loader: 'babel-loader'
				}
			}
		]
	}
};