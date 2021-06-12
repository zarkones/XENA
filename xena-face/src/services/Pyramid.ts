import { NuxtAxiosInstance } from '@nuxtjs/axios'

type Message = {
  id: string
  from: string | null
  to: string | null
  subject: string
  content: string
  replyTo: string | null
  replies: Message[]
}

class Pyramid {
  public readonly axios: NuxtAxiosInstance

  constructor (axios?: NuxtAxiosInstance) {
    this.axios = axios
  }

  public getBuilldProfiles = (axios: NuxtAxiosInstance) => {
    return axios({
      method: 'GET',
      url: `${process.env.XENA_PYRAMID_HOST}/build-profiles`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })
  }
}

export default new Pyramid()