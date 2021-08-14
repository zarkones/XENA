import axios from 'axios'

type RequestMethod =
  | 'GET'
  | 'DELETE'
  | 'HEAD'
  | 'OPTIONS'
  | 'POST'
  | 'PUT'
  | 'PATCH'
  | 'PURGE'
  | 'LINK'
  | 'UNLINK'

type RequestOptions = {
  url: string
  method: RequestMethod
}

export default new class WebReader {
  public makeRequest = (requestOptions: RequestOptions) => axios.request({
    url: requestOptions.url,
    method: requestOptions.method,
  })
}