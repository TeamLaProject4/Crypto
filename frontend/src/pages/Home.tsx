import React, { ReactElement, useState } from 'react'
import { mutate } from 'swr'
import axios from '../utils/axios'
import { useTestData } from '../utils/dataFetching'
interface Props {}

export default function Home({}: Props): ReactElement {
  const [text, setText] = useState('lorem')
  const { testData } = useTestData()
  return (
    <div className="pt-20 flex flex-col justify-center items-center">
      {/* <div className="grid grid-cols-2 gap-4"> */}
      <div className="flex flex-row gap-4">
        <div className="bg-white p-10 rounded-md shadow-md ">
          <div className="flex flex-col">
            <h2 className="font-bold">Add test text to db</h2>
            <input type="text" onBlur={(e) => setText(e.target.value)}></input>
            <button
              onClick={() =>
                axios
                  .post('/test_table', { text })
                  .then(() => mutate('/test_table'))
                  .catch((e) => console.log(e))
              }
              className="mt-5 p-2 bg-gray-500 rounded-md text-white"
            >
              submit
            </button>
          </div>
        </div>

        <div className="bg-white p-10 rounded-md shadow-md flex flex-col ">
          <h2 className="font-bold">Database records</h2>
          {testData &&
            testData.map((data) => (
              <p>
                {data.id} {data.text}
              </p>
            ))}
        </div>
      </div>
      <button
        className="mt-20 bg-gray-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-5"
        type="button"
        onClick={() => {
          axios
            .get('/test')
            .then((r) => alert(JSON.stringify(r.data)))
            .catch((e) => console.log(e))
        }}
      >
        Secret api call
      </button>
    </div>
  )
}
