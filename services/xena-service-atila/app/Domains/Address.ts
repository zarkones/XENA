import { validNumber, validEnum } from './Validators'

type AddressStatus = 'OK' | 'BANNED' | 'UNKNOWN'

export default class Address {
  constructor (
    public readonly x: number,
    public readonly y: number,
    public readonly z: number,
    public readonly w: number,
    public readonly status: AddressStatus,
  ) {
    this.x = validNumber(x, 'BAD_ADDRESS_VALUE')
    this.y = validNumber(y, 'BAD_ADDRESS_VALUE')
    this.z = validNumber(z, 'BAD_ADDRESS_VALUE')
    this.w = validNumber(w, 'BAD_ADDRESS_VALUE')
    this.status = validEnum(status, ['OK', 'BANNED'], 'BAD_ADDRESS_STATUS')
  }

  public static fromString = (address: string, status?: AddressStatus) => {
    const addressChunks = address.split('.')

    if (addressChunks.length != 4)
      return

    return new Address(
      parseInt(addressChunks[0]),
      parseInt(addressChunks[1]),
      parseInt(addressChunks[2]),
      parseInt(addressChunks[3]),
      status ?? 'UNKNOWN',
    )
  }

  public toString = () => `${this.x}.${this.y}.${this.z}.${this.w}`
}