import axios from 'axios'
// import {useAppDispatch} from '../app/hooks'
// // import cookie from "cookie";

// window.axios.interceptors.response.use(
//   function (response) {
//     // Call was successful, don't do anything special.
//     return response;
//   },
//   function (error) {
//     switch (error.response.status) {
//       case 401: // Not logged in
//       case 419: // Session expired
//       case 503: // Down for maintenance
//         // Bounce the user to the login screen with a redirect back
//         window.location.reload();
//         break;
//       case 500:
//         alert("Oops, something went wrong!");
//         break;
//       default:
//         // Allow individual requests to handle other errors
//         return Promise.reject(error);
//     }
//   }
// );

export default axios.create({
  baseURL: 'http://localhost:8000/',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})
