import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class buildProfilesController {
  public get = async ({ request, response }: HttpContextContract) => {
    const { id } = await request.validate(Validator.BuildProfile.Get)

    const maybeBuildProfile = await Repo.BuildProfile.get({ id })
    if (!maybeBuildProfile)
      return response.noContent()

    const buildProfile = Domain.BuildProfile.fromJSON(maybeBuildProfile)
    
    return response.ok(buildProfile)
  }

  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { page, status } = await request.validate(Validator.BuildProfile.GetMultiple)

    const maybeBuildProfiles = await Repo.BuildProfile.getMultiple({ page, status })
    if (!maybeBuildProfiles.length)
      return response.noContent()

    const buildProfiles = maybeBuildProfiles.map(buildProfile => Domain.BuildProfile.fromJSON(buildProfile))

    return response.ok(buildProfiles)
  }

  public insert = async ({ request, response }: HttpContextContract) => {
    const {
      name,
      description,
      gitUrl,
      config,
      status,
    } = await request.validate(Validator.BuildProfile.Insert)

    const buildProfile = await Repo.BuildProfile.insert(Domain.BuildProfile.fromJSON({
      name,
      description,
      gitUrl,
      config,
      status,
    })).then(buildProfile => Domain.BuildProfile.fromJSON(buildProfile))

    return response.ok(buildProfile)
  }

  public update = async ({}: HttpContextContract) => {
    // todo
  }

  public delete = async ({}: HttpContextContract) => {
    // todo
  }
}
