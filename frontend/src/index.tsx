import React, { Suspense } from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import '@notus-pro/react/tailwind.min.css'
import App from './App'
import { Provider } from 'react-redux'
import * as serviceWorker from './serviceWorker'
import '@fortawesome/fontawesome-free/css/all.min.css'
import 'react-toastify/dist/ReactToastify.css'
import './i18n'

ReactDOM.render(
  <React.StrictMode>
    <Suspense fallback={<div>Loading...</div>}>
      <App />
    </Suspense>
  </React.StrictMode>,
  document.getElementById('root'),
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.register()
