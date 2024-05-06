import './App.css';
import UrlShortenerList from './UrlShortenerList';
import UrlShortenerInput from './UrlShortenerInput';
import { useEffect, useState } from 'react';
import axios from 'axios'
import handleError, { errors } from './errors';


function App() {
  return (
    <UrlShorternerApp />
  );
}

export default App;

function UrlShorternerApp() {
  const [shortUrlList, setShortUrlList] = useState([])
  const [fetchListErr, setFetchListErr] = useState('')

  useEffect(() => {
    axios.get(process.env.REACT_APP_GO_BACKEND_HOST + "/shortUrls").then((response) => {
      setShortUrlList(response.data)
    }).catch((error) => {
      setFetchListErr("An error occurred while fetching list. Please try again later.")
    })
  }, [])
  return (
    <div className="flex flex-col">
    <UrlShortenerInput updateShortUrlList={setShortUrlList} setFetchListErr={setFetchListErr}/>
    <UrlShortenerList shortUrlList={shortUrlList} fetchListErr={fetchListErr}/>
    </div>
  )
}