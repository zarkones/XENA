import { validString } from './Validators'

export default class File {
  constructor (
    public readonly id: string,
    public readonly originId: string,
    public readonly originPath: string,
    public readonly data: string,
  ) {
    id = validString(id, 'BAD_FILE_ID', 'NON_EMPTY')
    originId = validString(originId, 'BAD_FILE_ORIGIN_ID', 'NON_EMPTY')
    originPath = validString(originPath, 'BAD_FILE_ORIGIN_PATH', 'NON_EMPTY')
    data = validString(data, 'BAD_FILE_DATA', 'NON_EMPTY')
  }

  public static fromJSON = (json: any) => new File(
    json.id,
    json.originId,
    json.originPath,
    json.data,
  )

  public get asJSON () {
    return {
      id: this.id,
      originId: this.originId,
      originPath: this.originPath,
      data: this.data,
    }
  }
}