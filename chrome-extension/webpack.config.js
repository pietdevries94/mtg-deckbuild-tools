var path = require("path");

var env = {
    NODE_ENV: (process.env.NODE_ENV || "development"),
    PORT: (process.env.PORT || 3000)
};

var options = {
    context: __dirname,
    entry: {
        "scryfall-content": path.join(__dirname, "src", "entry", "scryfall-content.ts")
    },
    output: {
        path: path.join(__dirname, "dist", "js"),
        filename: "[name].js"
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/
            }
        ]
    },
    resolve: {
        extensions: ['.ts', '.js']
    }
};

if (env.NODE_ENV === "development") {
    options.devtool = "cheap-module-eval-source-map";
}

module.exports = options;