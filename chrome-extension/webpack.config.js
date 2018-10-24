var path = require("path");
const VueLoaderPlugin = require('vue-loader/lib/plugin')

var env = {
    NODE_ENV: (process.env.NODE_ENV || "development"),
    PORT: (process.env.PORT || 3000)
};

var options = {
    context: __dirname,
    entry: {
        "scryfall-content": path.join(__dirname, "src", "entry", "scryfall-content.ts"),
        "edhrec-content": path.join(__dirname, "src", "entry", "edhrec-content.ts"),
        "tappedout-content": path.join(__dirname, "src", "entry", "tappedout-content.ts")
    },
    output: {
        path: path.join(__dirname, "dist", "js"),
        filename: "[name].js"
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                loader: 'ts-loader',
                exclude: /node_modules/,
                options: { appendTsSuffixTo: [/\.vue$/] }
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.scss$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    'sass-loader'
                ]
            },
            {
                test: /\.css$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                ]
            }
        ]
    },
    resolve: {
        extensions: ['.ts', '.js', '.vue'],
        alias: {
            src: path.resolve(__dirname, 'src'),
            "@": path.resolve(__dirname, 'src', 'ts')
        }
    },
    plugins: [
        // make sure to include the plugin!
        new VueLoaderPlugin()
    ]
};

if (env.NODE_ENV === "development") {
    options.devtool = "cheap-module-eval-source-map";
}

module.exports = options;