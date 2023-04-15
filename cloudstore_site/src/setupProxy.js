const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(
        '/api/account',
        createProxyMiddleware({
            target: 'http://192.168.49.2:80',
            changeOrigin: true,
        })
    );
    app.use(
        '/api/products',
        createProxyMiddleware({
            target: 'http://192.168.49.2:80',
            changeOrigin: true,
        })
    );
    app.use(
        '/api/orders',
        createProxyMiddleware({
            target: 'http://192.168.49.2:80',
            changeOrigin: true,
        })
    );
    app.use(
        '/api/reviews',
        createProxyMiddleware({
            target: 'http://192.168.49.2:80',
            changeOrigin: true,
        })
    );
};
