import React, { ReactElement, useState, useEffect } from 'react'
import NavbarLinks from '@notus-pro/react/NavbarLinks'
import axios from '../utils/axios'
import { mutate } from 'swr'
import Cookies from 'js-cookie'
import { useHistory } from 'react-router-dom'
import { useUser } from '../utils/dataFetching'

const lefLinksLoggedIn = [
  { to: '/protected', children: 'Protected route' },
  {
    dropdown: true,
    text: 'Demo Pages',
    items: [
      { disabled: true, children: 'Group 1' },
      { to: '#pablo', children: 'Demo Page 1' },
      { to: '#pablo', children: 'Demo Page 2' },
      { disabled: true, children: 'Group 2' },
      { to: '#pablo', children: 'Demo Page 3' },
      { to: '#pablo', children: 'Demo Page 4' },
      { disabled: true, children: 'Group 3' },
      { to: '#pablo', children: 'Demo Page 5' },
      { to: '#pablo', children: 'Demo Page 6' },
      { disabled: true, children: 'Group 4' },
      { to: '#pablo', children: 'Demo Page 7' },
      { to: '#pablo', children: 'Demo Page 8' },
    ],
  },
]
const initialLinks = {
  color: '',
  type: 'fixed',
  // logoImage:
  //   "https://raw.githubusercontent.com/creativetimofficial/public-assets/master/creative-tim/logo.png",
  logoText: 'Crypto',
  logoLink: { to: '/' },
  leftLinks: [
    { to: '#pablo', children: 'Elements' },
    {
      dropdown: true,
      text: 'Demo Pages',
      items: [
        { disabled: true, children: 'Group 1' },
        { to: '#pablo', children: 'Demo Page 1' },
        { to: '#pablo', children: 'Demo Page 2' },
        { disabled: true, children: 'Group 2' },
        { to: '#pablo', children: 'Demo Page 3' },
        { to: '#pablo', children: 'Demo Page 4' },
        { disabled: true, children: 'Group 3' },
        { to: '#pablo', children: 'Demo Page 5' },
        { to: '#pablo', children: 'Demo Page 6' },
        { disabled: true, children: 'Group 4' },
        { to: '#pablo', children: 'Demo Page 7' },
        { to: '#pablo', children: 'Demo Page 8' },
      ],
    },
  ],
  rightLinks: [{ to: '/login', children: 'Login' }],
}
export default function Navigation() {
  const [links, setLinks] = useState(initialLinks)
  const { user, loggedOut } = useUser()
  let history = useHistory()
  async function logout() {
    // if (!loggedOut) {
    const response = await axios.post('/api/logout')
    localStorage.setItem('logout response', JSON.stringify(response))
    document.cookie =
      'crypto_session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
    document.cookie =
      'XSRF_TOKEN=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
    mutate('/api/user')
    history.push('/login')
  }

  useEffect(() => {
    if (!loggedOut) {
      setLinks({
        ...initialLinks,
        rightLinks: [
          {
            dropdown: true,
            text: user ? user.name : 'Settings',
            items: [
              { disabled: true, children: 'Personal' },
              { to: '/user/settings', children: 'Settings' },
              { to: '#pablo', children: 'Demo Page 2' },
              { disabled: true, children: 'Organisation' },
              { to: '#pablo', children: 'Demo Page 3' },
              { to: '#pablo', children: 'Demo Page 4' },
            ],
          },
          { children: 'Logout', onClick: () => logout() },
        ],
        leftLinks: lefLinksLoggedIn,
      })
    } else {
      setLinks(initialLinks)
    }
  }, [loggedOut, user])
  return <NavbarLinks {...links} />
}
