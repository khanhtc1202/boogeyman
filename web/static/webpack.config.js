const path = require('path');

module.exports = {
	entry: path.resolve(__dirname, 'app.jsx'),
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: 'index.js'
	},
	resolve: {
		extensions: ['.js', '.jsx']
	}
};