import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'
import * as Service from 'App/Services'
import * as Helper from 'App/Helpers'

import Env from '@ioc:Adonis/Core/Env'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class BuildsController {
  public get = async ({ request, response }: HttpContextContract) => {
    const { id } = await request.validate(Validator.Build.Get)

    const maybeBuild = await Repo.Build.get({ id })
    if (!maybeBuild)
      return response.noContent()

    const build = Domain.Build.fromJSON(maybeBuild)
    
    return response.ok(build)
  }

  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { page } = await request.validate(Validator.Build.GetMultiple)

    const maybeBuilds = await Repo.Build.getMultiple({ page })
    if (!maybeBuilds.length)
      return response.noContent()

    const builds = maybeBuilds.map(build => Domain.Build.fromJSON(build).toJSON)

    return response.ok(builds)
  }

  public insert = async ({ request, response }: HttpContextContract) => {
    const { buildProfileId } = await request.validate(Validator.Build.Insert)

    const maybeBuildProfile = await Repo.BuildProfile.get({ id: buildProfileId })
    if (!maybeBuildProfile)
      return response.notFound({ success: false, message: 'Build profile not found.' })

    const buildProfile = Domain.BuildProfile.fromJSON(maybeBuildProfile)

    const buildId = uuidv4()

    // Clone the repo.
    const repositoryStatus = (() => {
      try {
        return Service.Git.clone(buildProfile.gitUrl, buildId)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()

    if (repositoryStatus != 'CLONED')
      return response.internalServerError({ success: true, message: 'Failed to clone the repository.' })

    // Shell instruction is the only way at the moment.
    if (!buildProfile.config.shell)
      return response.unprocessableEntity({ success: false, message: 'Build configuration requires the shell instruction.' })

    // Build the binary.
    const buildOutput = (() => {
      try {
        return Helper.Shell.exe(`${buildProfile.config.shell} -o ${Env.get('BUILD_DESTINATION')}${buildId} ${Service.Git.pathPrefix}${buildId}/xena-apep/main.go`)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()
    
    // Repo cleaning.
    try {
      Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    } catch (e) {
      console.warn(e)
      return 'ERROR'
    }

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(buildId)

    // Store the build.
    const build = await Repo.Build.insert( Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })).then(build => Domain.Build.fromJSON(build))

    // Return the build binary.
    return response.ok(build.toBinary)
  }

  public update = async ({}: HttpContextContract) => {
    // todo
  }

  public delete = async ({}: HttpContextContract) => {
    // todo
  }
}
