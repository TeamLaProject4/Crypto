import React, { ReactElement } from 'react'
import { useTranslation } from 'react-i18next'
import axios from '../utils/axios'
interface Props {}

export default function Home({}: Props): ReactElement {
  const { t, i18n } = useTranslation()

  const changeLanguage = (lng) => {
    i18n.changeLanguage(lng)
  }
  return (
    <div className="pt-20 flex justify-center items-center">
      <button
        className="mr-5 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        type="button"
        onClick={() => {
          axios
            .get('api/test')
            .then((r) => alert(JSON.stringify(r.data)))
            .catch((e) => console.log(e))
        }}
      >
        Secret api call
      </button>
    </div>
  )
}
