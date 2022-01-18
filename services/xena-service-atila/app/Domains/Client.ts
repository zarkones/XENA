import { validString, validEnum } from './Validators'

type ClientStatus = 'ALIVE' | 'DEAD' | 'BANNED'

export default class Client {
  constructor (
    public readonly id: string,
    public readonly ip: string,
    public readonly publicKey: string,
    public readonly status: ClientStatus,
  ) {
    this.id = validString(id, 'BAD_CLIENT_ID', 'NON_EMPTY')
    this.ip = validString(ip, 'BAD_CLIENT_IP', 'NON_EMPTY')
    this.publicKey = validString(publicKey, 'BAD_CLIENT_PUBLIC_KEY', 'NON_EMPTY')
    this.status = validEnum(status, ['ALIVE', 'DEAD', 'BANNED'], 'BAD_CLIENT_STATUS')
  }

  public static fromJSON = (json) => new Client(
    json.id,
    json.ip,
    json.publicKey,
    json.status,
  )
}