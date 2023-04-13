
import React from 'react';
import {
   createBrowserRouter,
   RouterProvider,
} from "react-router-dom";

import Land from './land.js'
import Login from './login.js'
import AllProducts from './allproducts.js'

import ReactDOM from 'react-dom';

let loggedin = false;

const router = createBrowserRouter([
   {
      path: "/",
      element: <Land />,
   },
   {
      path: "/login",
      element: <Login loggedin= { loggedin } />,
   },
   {
      path: "/allproducts",
      element: <AllProducts loggedin = {loggedin} />,
   },
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
   <React.StrictMode>
      <RouterProvider router={router} />
   </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();
