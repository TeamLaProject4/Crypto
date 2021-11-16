import './App.css'
import { BrowserRouter as Router, Route, Switch, Link } from 'react-router-dom'
import { PrivateRoute } from './utils/auth'
import Navigation from './components/Navigation.jsx'
import { ToastContainer } from 'react-toastify'
import Home from './pages/Home'

function App() {
  return (
    <Router>
      <main className={`min-h-screen bg-gray-100`}>
        <ToastContainer />
        <div className="">
          <Navigation />
        </div>

        <Switch>
          {/* public routes */}
          <Route path="/">
            <Home />
          </Route>

          {/* private routes */}
          <PrivateRoute exact path="/private">
            <h1>test</h1>
          </PrivateRoute>
        </Switch>
      </main>
    </Router>
  )
}

export default App
