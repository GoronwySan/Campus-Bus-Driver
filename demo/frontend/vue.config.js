const webpack = require('webpack');

module.exports = {
  devServer: {
    https: true,
    proxy: {
      "/api": {
        target: "http://localhost:3456", // Nginx 代理的端口
        changeOrigin: true,
        secure: false, // 忽略 SSL 错误
        // pathRewrite: { '^/api': '' }        // 将路径中的 /api 重写为空，防止后端服务器路径不一致问题
      },
    },
  },
  configureWebpack: {
    plugins: [
      new webpack.DefinePlugin({
        __VUE_PROD_DEVTOOLS__: JSON.stringify(false),
        __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: JSON.stringify(false),
      }),
    ],
  },
};
