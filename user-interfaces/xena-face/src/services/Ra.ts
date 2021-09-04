import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type BruteForcedSubdomains = {
  alive: string[],
  dead: string[],
}

export default new class Ra {
  public readonly axios: NuxtAxiosInstance

  constructor (axios?: NuxtAxiosInstance) {
    this.axios = axios
  }

  public sublist3r = (axios: NuxtAxiosInstance, domain: string) => axios({
      method: 'POST',
      url: `${process.env.XENA_RA_HOST}/recon/sublist3r`,
      data: {
        domain,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as string[]
    })

  public subdomainBruteforce = (axios: NuxtAxiosInstance, domain: string, dict: string[]) => axios({
      method: 'POST',
      url: `${process.env.XENA_RA_HOST}/recon/subdomain-bruteforce`,
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

  public nmap = (axios: NuxtAxiosInstance, address: string) => axios({
      method: 'POST',
      url: `${process.env.XENA_RA_HOST}/recon/nmap`,
      data: {
        address,
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data as string
    })
}
