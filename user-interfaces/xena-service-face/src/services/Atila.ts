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

export default new class Atila {
  public readonly axios: NuxtAxiosInstance

  constructor (axios?: NuxtAxiosInstance) {
    this.axios = axios
  }

  public getCount = (axios: NuxtAxiosInstance) => axios({
      method: 'GET',
      url: `${process.env.XENA_ATILA_HOST}/clients/stats/count`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as number[]
    })

  public getClient = (axios: NuxtAxiosInstance, id: string) => axios({
      method: 'GET',
      url: `${process.env.XENA_ATILA_HOST}/clients/${id}`,
      params: {
        status: 'ALIVE',
      }
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Client
    })

  public getClients = (axios: NuxtAxiosInstance) => axios({
      method: 'GET',
      url: `${process.env.XENA_ATILA_HOST}/clients`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as Client[]
    })

  public fetchMessages = (axios: NuxtAxiosInstance, clientId: string, verificationKey: string, clientVerificationKey: string, withReplies?: boolean) => axios({
      method: 'GET',
      url: `${process.env.XENA_ATILA_HOST}/messages`,
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

  public publishMessage = (axios: NuxtAxiosInstance, clientId: string, subject: string, content: string) => axios({
      method: 'POST',
      url: `${process.env.XENA_ATILA_HOST}/messages`,
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
