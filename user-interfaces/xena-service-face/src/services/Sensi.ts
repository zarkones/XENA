import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type Statement = {
  id: string
  input: string
  output: string | null
}

export type Dialog = Statement[]

export default class Sensi {
  constructor (
    private readonly axios: NuxtAxiosInstance,
    private readonly baseURL: string,
    private readonly token: string,
  ) {
    this.axios = axios
    this.baseURL = baseURL
    this.token = token
  }
  
  public getDialog = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/dialogs`,
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    })
    .then(({ data }) => {
      if (data)
        return data as Dialog
    })
    .catch(e => console.error(e))
  
  public insert = (prompt: string) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/dialogs`,
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
      data: {
        prompt,
      }
    })
    .then(({ data }) => {
      if (data)
        return data as Statement
    })
    .catch(e => console.error(e))
}
