import { NuxtAxiosInstance } from '@nuxtjs/axios'
import Crypto from './Crypto'

type Message = {
  id: string
  from: string | null
  to: string | null
  subject: string
  content: string
  replyTo: string | null
  replies: Message[]
}

type Client = {
  id: string
  publicKey: string
}

export default class Atila {
  public readonly axios: NuxtAxiosInstance
  public baseURL: string

  constructor (axios: NuxtAxiosInstance, baseURL: string) {
    this.axios = axios
    this.baseURL = baseURL
  }

  public getCount = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/clients/stats/count`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as number[]
    })

  public getClient = (id: string) => this.axios({
      method: 'GET',
      url: `${this.baseURL}/clients/${id}`,
      params: {
        status: 'ALIVE',
      }
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Client
    })

  public getClients = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/clients`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Client[]
    })

  public fetchMessages = (clientId: string, verificationKey: string, clientVerificationKey: string, withReplies?: boolean) => this.axios({
      method: 'GET',
      url: `${this.baseURL}/messages`,
      params: {
        clientId,
        withReplies,
      }
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp && resp.data)
        return (resp.data as Message[]).map(message => ({
          ...message,
          content: Crypto.verify(verificationKey, message.content),
          replies: message.replies.map(reply => Crypto.verify(clientVerificationKey, reply.content)),
        }))
    })

  public publishMessage = (clientId: string, subject: string, content: string) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/messages`,
      data: {
        to: clientId,
        subject,
        content,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })
}
