import Build from 'App/Models/Build'

import { BuildId, BuildProfileId } from '../Types'

type Get = {
  id?: BuildId
  buildProfileId?: BuildProfileId
}

type GetMultiple = {
  page?: number
}

type Insert = {
  id: BuildId
  buildProfileId: BuildProfileId
  data: string
}

class Database {
  public get = ({ id }: Get) => Build.query()
    .select('*')
    .if(id, builder => builder.where('id', id as BuildId))
    .first()
    .then(build => build?.serialize())
  
  public getMultiple = ({ page }: GetMultiple) => Build.query()
    .select('*')
    .if(page, builder => builder.offset(page as number * 10))
    .if(page, builder => builder.limit(page as number * 10))
    .exec()
    .then(builds => builds.map(build => build.serialize()))
  
  public insert = (payload: Insert) => Build.create(payload).then(build => build.serialize())
}

export default new Database()