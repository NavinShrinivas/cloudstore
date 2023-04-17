const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
   app.use(
      '/api/account',
      createProxyMiddleware({
         target: 'http://hackframe.navinxyz.com:8081',
         changeOrigin: true,
      })
   );
   app.use(
      '/api/products',
      createProxyMiddleware({
         target: 'http://hackframe.navinxyz.com:8081',
         changeOrigin: true,
      })
   );
   app.use(
      '/api/orders',
      createProxyMiddleware({
         target: 'http://hackframe.navinxyz.com:8081',
         changeOrigin: true,
      })
   );
   // app.use(
   //     '/api/reviews',
   //     createProxyMiddleware({
   //         target: 'http://hackframe.navinxyz.com:8081',
   //         changeOrigin: true,
   //     })
   // );
};
