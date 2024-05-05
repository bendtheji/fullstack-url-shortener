import './App.css';
import UrlShortenerList from './UrlShortenerList';
import UrlShortenerInput from './UrlShortenerInput';
import { useEffect, useState } from 'react';
import axios from 'axios'

function App() {
  return (
    <UrlShorternerApp />
  );
}

export default App;

function UrlShorternerApp() {
  const [shortUrlList, setShortUrlList] = useState([])

  useEffect(() => {
    axios.get(process.env.REACT_APP_GO_BACKEND_HOST + "/shortUrls").then((response) => {
      setShortUrlList(response.data)
    })
  }, [])
  return (
    <div className="flex flex-col">
    <UrlShortenerInput updateShortUrlList={setShortUrlList}/>
    <UrlShortenerList shortUrlList={shortUrlList}/>
    </div>
  )
}