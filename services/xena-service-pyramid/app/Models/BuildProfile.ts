import { DateTime } from 'luxon'
import { BaseModel, column } from '@ioc:Adonis/Lucid/Orm'

type ProfileStatus = 'ENABLED' | 'DISABLED' | 'DELETED'

export default class BuildProfile extends BaseModel {
  @column({ isPrimary: true })
  public id: string
  
  @column()
  public name: string
  
  @column()
  public description: string | null
  
  @column()
  public gitUrl: string
  
  @column()
  public config: any

  @column()
  public status: ProfileStatus

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
