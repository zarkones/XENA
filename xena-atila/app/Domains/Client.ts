import { validString, validEnum } from './Validators'

type ClientStatus = 'ALIVE' | 'DEAD' | 'BANNED'

export default class Client {
  public readonly id: string
  public readonly status: ClientStatus

  constructor (id: string, status: ClientStatus) {
    this.id = validString(id, 'BAD_CLIENT_ID', 'NON_EMPTY')
    this.status = validEnum(status, ['ALIVE', 'DEAD', 'BANNED'], 'BAD_CLIENT_STATUS')
  }

  public static fromJSON = (json) => {
    return new Client(
      json.id,
      json.status,
    )
  }
}