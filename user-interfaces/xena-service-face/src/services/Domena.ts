import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type Service = {
  id: string
  address: string
  port: number
  createdAt: string
}

export default class Domena {
  constructor (
    private readonly axios: NuxtAxiosInstance,
    private readonly baseURL: string,
    private readonly token: string,
  ) {
    this.axios = axios
    this.baseURL = baseURL
    this.token = token
  }

  public getServices = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/services`,
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Service[]
    })
}
