import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type BuildTemplate = 'XENA_BOT_RA' | 'XENA_BOT_APEP' | 'XENA_BOT_ANACONDA'
type Configuration = {
  template: BuildTemplate
  atilaHost?: string
  trustedPublicKey?: string
}

export default class Pyramid {
  public readonly axios: NuxtAxiosInstance
  public baseURL: string

  constructor (axios: NuxtAxiosInstance, baseURL: string) {
    this.axios = axios
    this.baseURL = baseURL
  }

  public getBuilldProfiles = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/build-profiles`,
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })

  public insertBuildProfile = (
    name: string,
    description: string | null,
    gitUrl: string,
    config: Configuration,
  ) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/build-profiles`,
      data: {
        name,
        description,
        gitUrl,
        config,
        status: 'ENABLED'
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })
}
