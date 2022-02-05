import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type BruteForcedSubdomains = {
  alive: string[],
  dead: string[],
}

export default class Ra {
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

  constructor (
    private readonly axios: NuxtAxiosInstance,
    private readonly baseURL: string,
    private readonly token: string,
  ) {
    this.axios = axios
    this.baseURL = baseURL
    this.token = token
  }

  public sublist3r = (domain: string) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/recon/sublist3r`,
      data: {
        domain,
      },
      headers: {
        Authorization: `Bearer ${this.token}`,
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
      headers: {
        Authorization: `Bearer ${this.token}`,
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
      headers: {
        Authorization: `Bearer ${this.token}`,
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
      headers: {
        Authorization: `Bearer ${this.token}`,
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
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })
}
