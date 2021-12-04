import { validString, validEnum } from './Validators'
import { v4 as uuidv4 } from 'uuid'

type ProfileStatus = 'ENABLED' | 'DISABLED' | 'DELETED'

type Configuration = {
  template: 'XENA_BOT_RA' | 'XENA_BOT_APEP' | 'XENA_BOT_ANACONDA' | 'XENA_BOT_VARVARA'
}

export default class BuildProfile {
  constructor (
    public readonly id: string,
    public readonly name: string,
    public readonly description: string | null,
    public readonly gitUrl: string,
    public readonly config: Configuration,
    public readonly status: ProfileStatus,
  ) {
    this.id = validString(id ?? uuidv4(), 'BAD_BUILD_PROFILE_ID', 'NON_EMPTY')
    this.name = validString(name, 'BAD_BUILD_PROFILE_NAME', 'NON_EMPTY')
    this.description = description ? validString(description, 'BAD_BUILD_PROFILE_DESCRIPTION', 'NON_EMPTY') : null
    this.gitUrl = validString(gitUrl, 'BAD_BUILD_PROFILE_GIT_URL', 'NON_EMPTY')
    this.config = {
      template: validEnum(config.template, ['XENA_BOT_RA', 'XENA_BOT_APEP', 'XENA_BOT_ANACONDA', 'XENA_BOT_VARVARA'], 'BAD_BUILD_PROFILE_CONFIG_TEMPLATE')
    }
    this.status = validEnum(status, ['ENABLED', 'DISABLED', 'DELETED'], 'BAD_BUILD_STATUS')
  }

  public static fromJSON = (json) => new BuildProfile(
      json.id,
      json.name,
      json.description,
      json.gitUrl,
      json.config,
      json.status,
    )
}