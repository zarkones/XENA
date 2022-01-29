import { DateTime } from 'luxon'
import { BaseModel, column, hasOne, HasOne } from '@ioc:Adonis/Lucid/Orm'
import Author from './Author'

export default class Post extends BaseModel {
  @column({ isPrimary: true })
  public id: string

  @column()
  public topicId: string

  @column()
  public authorId: string

  @column()
  public title: string

  @column()
  public description: string | null

  @hasOne(() => Author, {
    foreignKey: 'id',
    localKey: 'authorId'
  })
  public author: HasOne<typeof Author>

  @column.dateTime({ autoCreate: true })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true })
  public updatedAt: DateTime
}
