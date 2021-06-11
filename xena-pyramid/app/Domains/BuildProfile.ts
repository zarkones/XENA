import { validString, validEnum } from './Validators'

type ProfileStatus = 'ENABLED' | 'DISABLED' | 'DELETED'

type Configuration = {
  shell?: string
}

export default class BuildProfile {
  public readonly id: string
  public readonly name: string
  public readonly description: string | null
  public readonly gitUrl: string
  public readonly config: Configuration
  public readonly status: ProfileStatus

  constructor (
    id: string,
    name: string,
    description: string | null,
    gitUrl: string,
    config: Configuration,
    status: ProfileStatus,
  ) {
    this.id = validString(id, 'BAD_CLIENT_ID', 'NON_EMPTY')
    this.name = validString(name, 'BAD_BUILD_PROFILE')
    this.description = description ? validString(description, 'BAD_BUILD_PROFILE') : null
    this.gitUrl = validString(gitUrl, 'BAD_BUILD_PROFILE')
    this.config = config
    this.status = validEnum(status, ['ENABLED', 'DISABLED', 'DELETED'], 'BAD_CLIENT_STATUS')
  }

  public static fromJSON = (json) => {
    return new BuildProfile(
      json.id,
      json.name,
      json.description,
      json.gitUrl,
      json.config,
      json.status,
    )
  }
}