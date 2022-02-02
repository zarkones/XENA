import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type BuildTemplate = 'XENA_BOT_RA' | 'XENA_BOT_APEP' | 'XENA_BOT_ANACONDA'
type Configuration = {
  template: BuildTemplate
  atilaHost?: string
  trustedPublicKey?: string
  dgaSeed?: number
  dgaAfterDays?: number
  maxLoopWait?: number
  minLoopWait?: number
  gettrProfileName?: string
}

export default class Pyramid {
  constructor (
    private readonly axios: NuxtAxiosInstance,
    private readonly baseURL: string,
    private readonly token: string,
  ) {
    this.axios = axios
    this.baseURL = baseURL
    this.token = token
  }

  public getBuilldProfiles = () => this.axios({
      method: 'GET',
      url: `${this.baseURL}/build-profiles`,
      headers: {
        Authorization: `Bearer ${this.token}`,
      },
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
