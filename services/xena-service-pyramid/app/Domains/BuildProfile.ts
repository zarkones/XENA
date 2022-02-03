import { validString, validEnum } from './Validators'
import { v4 as uuidv4 } from 'uuid'

type ProfileStatus = 'ENABLED' | 'DISABLED' | 'DELETED'

export const buildTemplates = ['XENA_BOT_RA', 'XENA_BOT_APEP', 'XENA_BOT_ANACONDA', 'XENA_BOT_VARVARA', 'XENA_BOT_MONOLITIC'] as const
type BuildTemplate = keyof typeof buildTemplates

type Configuration = {
  template: BuildTemplate
  atilaHost: string | null
  trustedPublicKey: string | null
  maxLoopWait: string | null
  minLoopWait: string | null
  gettrProfileName: string | null
  dgaSeed: string | null
  dgaAfterDays: string | null
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
      template: validEnum(config.template, ['XENA_BOT_RA', 'XENA_BOT_APEP', 'XENA_BOT_ANACONDA', 'XENA_BOT_VARVARA', 'XENA_BOT_MONOLITIC'], 'BAD_BUILD_PROFILE_CONFIG_TEMPLATE'),
      atilaHost: config.atilaHost ? validString(config.atilaHost, 'BAD_BUILD_PROFILE_ATILA_HOST', 'NON_EMPTY') : null,
      trustedPublicKey: config.trustedPublicKey ? validString(config.trustedPublicKey, 'BAD_BUILD_PROFILE_TRUSTED_PUBLIC_KEY', 'NON_EMPTY') : null,
      maxLoopWait: config.maxLoopWait ? validString(config.maxLoopWait, 'BAD_BUILD_PROFILE_MAX_LOOP_WAIT', 'NON_EMPTY') : null,
      minLoopWait: config.minLoopWait ? validString(config.minLoopWait, 'BAD_BUILD_PROFILE_MIN_LOOP_WAIT', 'NON_EMPTY') : null,
      gettrProfileName: config.gettrProfileName ? validString(config.gettrProfileName, 'BAD_BUILD_PROFILE_GETTR_PROFILE_NAME', 'NON_EMPTY') : null,
      dgaSeed: config.dgaSeed ? validString(config.dgaSeed, 'BAD_BUILD_PROFILE_DGA_SEED', 'NON_EMPTY') : null,
      dgaAfterDays: config.dgaAfterDays ? validString(config.dgaAfterDays, 'BAD_BUILD_PROFILE_DGA_TIMEOUT', 'NON_EMPTY') : null,
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