const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(
        '/api/account',
        createProxyMiddleware({
            target: 'http://localhost:5001',
            changeOrigin: true,
        })
    );
    app.use(
        '/api/products',
        createProxyMiddleware({
            target: 'http://localhost:5002',
            changeOrigin: true,
        })
    );
    app.use(
        '/api/orders',
        createProxyMiddleware({
            target: 'http://localhost:5003',
            changeOrigin: true,
        })
    );
    app.use(
        '/api/reviews',
        createProxyMiddleware({
            target: 'http://localhost:5004',
            changeOrigin: true,
        })
    );
};