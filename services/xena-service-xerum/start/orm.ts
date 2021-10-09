import { LucidModel } from '@ioc:Adonis/Lucid/Model'
import { BaseModel } from '@ioc:Adonis/Lucid/Orm'

BaseModel.namingStrategy.serializedName = (_model: LucidModel, attributeName: string): string => {
  return attributeName
}