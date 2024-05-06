
const STATUS_CONFLICT = 409
const BAD_REQUEST = 400

export default function handleError(errCode, msg, setErrorMsg) {
    switch (errCode) {
        case STATUS_CONFLICT:
          setErrorMsg('Long URL has been shortened before. Please reuse the same shortened URL.')
          break
        case BAD_REQUEST:
          setErrorMsg('Bad Request: ' + msg)
          break
        default:
          setErrorMsg('Something went wrong during the creation of the shortened URL. Please try again later.')
      }
}