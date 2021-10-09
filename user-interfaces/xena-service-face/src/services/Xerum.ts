import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type Post = {
  id: string
  authorId: string
  topicId: string
  title: string
  description: string
  author: {
    name: string
  }
}

export default new class Xerum {
  public readonly axios: NuxtAxiosInstance

  constructor (axios?: NuxtAxiosInstance) {
    this.axios = axios
  }
  
  public getPosts = (axios: NuxtAxiosInstance) => axios({
    method: 'GET',
    url: `${process.env.XENA_XERUM_HOST}/posts`
    })
    .then(({ data }) => {
      if (data)
        return data as Post[]
    })
    .catch(e => console.error(e)) 
}
