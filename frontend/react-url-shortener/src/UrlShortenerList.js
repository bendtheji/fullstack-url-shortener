
function UrlShortenerList({shortUrlList}) {
  
    const rows = []
    for (let i = 0; i < shortUrlList.length; i++) {
      const shortUrl = process.env.REACT_APP_GO_BACKEND_HOST + "/shortUrls/" + shortUrlList[i].ShortUrl
      rows.push(
      <tr className="hover:bg-gray-100 ">
        <td className="w-7/12 px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-800 text-ellipsis overflow-hidden">{shortUrlList[i].Description}</td>
        <td className="w-5/12 px-6 py-4 whitespace-nowrap text-ellipsis overflow-hidden"> <a href={shortUrl}>{shortUrl}</a></td>
      </tr>
      )
    }
    return (
        shortUrlList.length !== 0 ? 
        <>
        <hr className="mx-auto w-8/12"></hr>
        <div className="mx-auto w-8/12 mt-6 p-1.5">
            <p className="text-slate-600 text-xl font-sans font-semibold">
                List of Shortened URLs
            </p>
        </div>
      <div className="mt-8 w-full block">
        <div className="mx-auto w-8/12">
        <div class="flex flex-col">
            <div class="-m-1.5 w-full">
            <div class="p-1.5 w-full inline-block align-middle">
            <table class="w-full divide-y divide-gray-200 table-fixed">
            <thead>
                <tr>
                <th scope="col" class="w-7/12 px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Description</th>
                <th scope="col" class="w-5/12 px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Short URL</th>
                </tr>
            </thead>
          <tbody class="divide-y divide-gray-200 w-full">
            {rows}
          </tbody>
        </table>
      </div>
  </div>
</div>
        </div>
      </div>
      </> : ""
    )
  }

export default UrlShortenerList