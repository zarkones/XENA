import { System } from '.'
import { validString, validEnum } from './Validators'

type ClientStatus = 'ALIVE' | 'DEAD' | 'BANNED'

export default class Client {
  constructor (
    public readonly id: string,
    public readonly ip: string,
    public readonly osId: string,
    public readonly publicKey: string,
    public readonly status: ClientStatus,
    public readonly system: System, 
  ) {
    this.id = validString(id, 'BAD_CLIENT_ID', 'NON_EMPTY')
    this.ip = validString(ip, 'BAD_CLIENT_IP', 'NON_EMPTY')
    this.osId = validString(osId, 'BAD_CLIENT_OS_ID', 'NON_EMPTY')
    this.publicKey = validString(publicKey, 'BAD_CLIENT_PUBLIC_KEY', 'NON_EMPTY')
    this.status = validEnum(status, ['ALIVE', 'DEAD', 'BANNED'], 'BAD_CLIENT_STATUS')
    this.system = System.fromJSON(system)
  }

  public static fromJSON = (json) => new Client(
    json.id,
    json.ip,
    json.osId,
    json.publicKey,
    json.status,
    json.system,
  )
}