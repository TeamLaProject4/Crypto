// import axios from "axios";
import axios from '../utils/axios'
import useSWR from 'swr'
const fetcher = (url) => axios.get(url).then((res) => res.data)

export function useUser() {
  const { data, mutate, error } = useSWR('/api/user', fetcher, {
    errorRetryCount: 2,
    revalidateOnFocus: false,
  })

  const loading = !data && !error
  let loggedOut = false
  if (error) {
    loggedOut = true
  }
  return {
    loading,
    loggedOut,
    user: data,
    mutate,
  }
}
