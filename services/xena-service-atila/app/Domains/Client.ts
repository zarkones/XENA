import { validString, validEnum } from './Validators'

type ClientStatus = 'ALIVE' | 'DEAD' | 'BANNED'

export default class Client {
  public readonly id: string
  public readonly identificationKey: string
  public readonly status: ClientStatus

  constructor (
    id: string,
    identificationKey: string,
    status: ClientStatus,
  ) {
    this.id = validString(id, 'BAD_CLIENT_ID', 'NON_EMPTY')
    this.identificationKey = validString(identificationKey, 'BAD_CLIENT_IDENTIFICATION_KEY', 'NON_EMPTY')
    this.status = validEnum(status, ['ALIVE', 'DEAD', 'BANNED'], 'BAD_CLIENT_STATUS')
  }

  public static fromJSON = (json) => new Client(
      json.id,
      json.identificationKey,
      json.status,
    )
}