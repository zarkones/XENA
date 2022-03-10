import { NuxtAxiosInstance } from '@nuxtjs/axios'
import Crypto from './Crypto'

export type Message = {
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
  constructor (
    private readonly axios: NuxtAxiosInstance,
    private readonly baseURL: string,
    private readonly token: string,
  ) {
    this.axios = axios
    this.baseURL = baseURL
    this.token = token
  }

  public deleteMessage = (id: string) => this.axios({
    method: 'DELETE',
    url: `${this.baseURL}/messages`,
    headers: {
      Authorization: `Bearer ${this.token}`,
    },
    data: {
      id,
    }
  })

  public getCount = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/clients/stats/count`,
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as number[]
    })
  
  public getDemographic = (os?: string) => this.axios({
    method: 'GET',
    url: `${this.baseURL}/clients/stats/demographic${os ? `?os=${os}` : ''}`,
    headers: {
      Authorization: `Bearer ${this.token}`,
    },
  })
  .catch(err => console.warn(err))
  .then(resp => {
    if (resp)
      return resp.data
  })

  public getClient = (id: string) => this.axios({
      method: 'GET',
      url: `${this.baseURL}/clients/${id}`,
      params: {
        status: 'ALIVE',
      },
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Client
    })

  public getClients = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/clients`,
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
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
      },
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
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
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Message
    })
}
