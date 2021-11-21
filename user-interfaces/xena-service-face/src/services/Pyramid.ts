import { NuxtAxiosInstance } from '@nuxtjs/axios'

export type BuildTemplate = 'XENA_APEP' | 'XENA_RA'

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
    template: BuildTemplate,
  ) => this.axios({
      method: 'POST',
      url: `${this.baseURL}/build-profiles`,
      data: {
        name,
        description,
        gitUrl,
        config: {
          template,
        },
        status: 'ENABLED'
      },
    })
    .catch(err => console.warn(err))
    .then(resp => {
      if (resp)
        return resp.data
    })
}
