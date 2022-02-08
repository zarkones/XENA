import { DateTime } from 'luxon'
import { BaseModel, column } from '@ioc:Adonis/Lucid/Orm'

type ServiceDetails = {
  telnetUsername?: string
  telnetPassword?: string
  sshUsername?: string
  sshPassword?: string
}

export default class Service extends BaseModel {
  @column({ isPrimary: true })
  public id: string

  @column()
  public address: string 
  
  @column()
  public port: number

  @column()
  public details: ServiceDetails

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
