import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type Service = {
  id: string
  address: string
  port: number
  createdAt: string
}

export default class Domena {
  public readonly axios: NuxtAxiosInstance
  public baseURL: string

  constructor (axios: NuxtAxiosInstance, baseURL: string) {
    this.axios = axios
    this.baseURL = baseURL
  }

  public getServices = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/services`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Service[]
    })
}
