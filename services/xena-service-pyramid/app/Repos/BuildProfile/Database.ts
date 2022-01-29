import BuildProfile from 'App/Models/BuildProfile'

import { BuildProfileId } from '../Types'

type ProfileStatus = 'ENABLED' | 'DISABLED' | 'DELETED'

type Get = {
  id?: BuildProfileId
}

type GetMultiple = {
  page?: number
  status?: ProfileStatus
}

type Insert = {
  id: BuildProfileId
  name: string
  description: string | null
  gitUrl: string
  config: any
  status: ProfileStatus
}

class Database {
  public get = ({ id }: Get) => BuildProfile.query()
    .select('*')
    .if(id, builder => builder.where('id', id as BuildProfileId))
    .first()
    .then(buildProfile => buildProfile?.serialize())
  
  public getMultiple = ({ page }: GetMultiple) => BuildProfile.query()
    .select('*')
    .if(page, builder => builder.offset(page as number * 10))
    .if(page, builder => builder.limit(page as number * 10))
    .exec()
    .then(buildProfiles => buildProfiles.map(buildProfile => buildProfile.serialize()))
  
  public insert = (payload: Insert) => BuildProfile.create(payload).then(buildProfile => buildProfile.serialize())
}

export default new Database()