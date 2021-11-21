import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type BruteForcedSubdomains = {
  alive: string[],
  dead: string[],
}

export default class Ra {
  public readonly axios: NuxtAxiosInstance
  public baseURL: string

  public readonly webMethods = [
    'GET',
    'DELETE',
    'HEAD',
    'OPTIONS',
    'POST',
    'PUT',
    'PATCH',
    'PURGE',
    'LINK',
    'UNLINK',
  ] as const

  constructor (axios: NuxtAxiosInstance, baseURL: string) {
    this.axios = axios
    this.baseURL = baseURL
  }

  public sublist3r = (domain: string) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/recon/sublist3r`,
      data: {
        domain,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as string[]
    })

  public subdomainBruteforce = (domain: string, dict: string[]) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/recon/subdomain-bruteforce`,
      data: {
        domain,
        dict,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data 

      console.log(resp)
    })

  public nmap = (address: string) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/recon/nmap`,
      data: {
        address,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as string
    })

  public sqlmap = (url: string) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/scans/sql-injection`,
      data: {
        url,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as string[]
    })

  public webFuzzer = (url: string, method: string, wordlist?: string[]) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/scans/web-fuzzer`,
      data: {
        url,
        method,
        wordlist,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })
}
