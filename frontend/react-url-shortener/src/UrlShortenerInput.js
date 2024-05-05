import { useState } from 'react';
import axios from 'axios'


function UrlShortenerInput({updateShortUrlList}) {
    const [longUrl, setLongUrl] = useState('')
    const [description, setDescription] = useState('')
  
    function handleClick() {
      axios.post(process.env.REACT_APP_GO_BACKEND_HOST + "/shortUrls", {
        long_url: longUrl,
        description: description
      }).then(function(response) {
        axios.get(process.env.REACT_APP_GO_BACKEND_HOST + "/shortUrls").then((response) => {
            updateShortUrlList(response.data)
        })
        setLongUrl('')
        setDescription('')
      }).catch(function(error){
      })
    }
  
    return (
        <>
        <div className="mx-auto w-8/12 mt-6 p-1.5">
            <p className="text-slate-600 text-xl font-sans font-semibold">
                Create a Shortened URL
            </p>
        </div>
        <div className="h-auto w-full block">
            <div className="mx-auto w-8/12 flex flex-col items-start p-6 mb-8">
                <div className='my-4 w-full'>
                <label htmlFor="input-label" className="block text-sm font-medium mb-2">Description</label>
                <input type="text" id="input-label" className="py-3 px-6 block w-full border border-gray-400 rounded-lg text-sm disabled:opacity-50 disabled:pointer-events-none" 
                placeholder="Short URL for YouTube" value={description} onChange={(e) => setDescription(e.target.value)}/>
                </div>
                <div className='mb-8 w-full'>
                <label htmlFor="input-label" className="block text-sm font-medium mb-2">URL</label>
                <input type="text" id="input-label" className="py-3 px-6 block w-full border border-gray-400 rounded-lg text-sm disabled:opacity-50 disabled:pointer-events-none" 
                placeholder="www.youtube.com" value={longUrl} onChange={(e) => setLongUrl(e.target.value)}/>
                </div>
                <button type="button"  onClick={handleClick} className="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Create Short URL</button>
            </div>
        </div>
        </>
    )
  }

  export default UrlShortenerInput
