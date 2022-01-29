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

export default class Xerum {
  public readonly axios: NuxtAxiosInstance
  public baseURL: string

  constructor (axios: NuxtAxiosInstance, baseURL: string) {
    this.axios = axios
    this.baseURL = baseURL
  }
  
  public getPosts = () => this.axios({
    method: 'GET',
    url: `${this.baseURL}/posts`
    })
    .then(({ data }) => {
      if (data)
        return data as Post[]
    })
    .catch(e => console.error(e))

  public getPost = (postId: string) => this.axios({
      method: 'GET',
      url: `${this.baseURL}/posts/${postId}`
      })
      .then(({ data }) => {
        if (data)
          return data as Post
      })
      .catch(e => console.error(e))
}
