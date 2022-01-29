import { validString } from './Validators'
import { v4 as uuidv4 } from 'uuid'
import { readFile } from 'fs/promises'

export default class Build {
  constructor (
    public readonly id: string,
    public readonly buildProfileId: string,
    public readonly data: string,
  ) {
    this.id = validString(id ?? uuidv4(), 'BAD_BUILD_ID', 'NON_EMPTY')
    this.buildProfileId = validString(buildProfileId, 'BAD_BUILD_PROFILE_ID', 'NON_EMPTY')
    this.data = data // todo Create a validation method.
  }

  public static fromJSON = (json) => new Build(
      json.id,
      json.buildProfileId,
      json.data,
    )

  get toJSON () {
    return {
      id: this.id,
      buildProfileId: this.buildProfileId,
    }
  }

  public static getBinary = (buildPath: string) => readFile(buildPath, { encoding: 'base64' })

  get toBinary () {
    return Buffer.from(this.data, 'base64')
  } 
}