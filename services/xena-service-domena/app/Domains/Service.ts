import { validString, validNumber } from './Validators'

type ServiceDetails = {
  telnetUsername?: string
  telnetPassword?: string
}

export default class Service {
  constructor (
    public readonly id: string,
    public readonly address: string,
    public readonly port: number,
    public readonly details?: ServiceDetails,
  ) {
    this.id = validString(id, 'BAD_SERVICE_ID', 'NON_EMPTY')
    this.address = validString(address, 'BAD_SERVICE_ADDRESS', 'NON_EMPTY')
    this.port = validNumber(port, 'BAD_SERVICE_PORT', true)
    this.details = details
      ? {
        telnetUsername: details.telnetUsername ? validString(details.telnetUsername, 'BAD_SERVICE_DETAILS_TELNET_USERNAME', 'NON_EMPTY') : undefined,
        telnetPassword: details.telnetPassword ? validString(details.telnetPassword, 'BAD_SERVICE_DETAILS_TELNET_PASSWORD', 'NON_EMPTY') : undefined,
      }
      : undefined
  }

  public static fromJSON = (json: any) => new Service(
    json.id,
    json.address,
    json.port,
    json.details,
  )
}