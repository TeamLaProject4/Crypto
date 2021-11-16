import { Route, Redirect } from 'react-router-dom'
import axios from './axios'
import { mutate } from 'swr'
import Cookies from 'js-cookie'
import { useUser } from './dataFetching'

// A wrapper for <Route> that redirects to the login
// screen if you're not yet authenticated.
export function PrivateRoute({ children, ...rest }) {
  //do something with roles?
  const { loggedOut } = useUser()
  return (
    <Route
      {...rest}
      render={({ location }) =>
        !loggedOut ? (
          <div className="p-20 flex justify-center items-center">
            {children}
          </div>
        ) : (
          <Redirect
            to={{
              pathname: '/login',
              state: { from: location },
            }}
          />
        )
      }
    />
  )
}

export const logOut = async () => {
  if (typeof window !== 'undefined') {
    await axios.post('/api/logout')

    Cookies.remove('XSRF-TOKEN', {
      expires: 1 / 12,
      sameSite: 'lax',
    })
    Cookies.remove('iris_session', {
      expires: 1 / 12,
      sameSite: 'lax',
    })
    mutate('/api/user')

    // router.push("/");
  }
}
