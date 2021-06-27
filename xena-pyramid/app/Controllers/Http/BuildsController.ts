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
    
    return response.ok(build.toBinary)
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

    switch (buildProfile.config.template) {
      case 'XENA_RA':
        // const ra = await this.buildRa(buildId, buildProfileId)
        // return ra == 'ERROR'
        //   ? response.internalServerError({ success: false, message: 'Failed to build.' })
        //   : response.ok(ra)
        return response.internalServerError({ success: false, message: 'Not yet implemented.' })
      case 'XENA_APEP':
        const apep = await this.buildApep(buildId, buildProfileId)
        return !apep
          ? response.internalServerError({ success: false, message: 'Failed to build.' })
          : response.ok(apep)
      default:
        return response.unprocessableEntity({ success: false, message: 'Unrecognized build template.' })
    }
  }

  public update = async ({}: HttpContextContract) => {
    // todo
  }

  public delete = async ({}: HttpContextContract) => {
    // todo
  }

  private buildApep = async (buildId: string, buildProfileId: string) => {
    // Build the binary.
    const buildOutput = (() => {
      try {
        return Helper.Shell.exe(`go build -o ${Env.get('BUILD_DESTINATION')}${buildId}_BUILD ${Service.Git.pathPrefix}${buildId}/xena-apep`)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()
    
    if (buildOutput == 'ERROR')
      throw Error('Unable to build.')

    // Repo cleaning.
    try {
      Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    } catch (e) {
      console.warn(e)
      return 'ERROR'
    }

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(`${Env.get('BUILD_DESTINATION')}${buildId}_BUILD`)

    // Store the build.
    const build = await Repo.Build.insert( Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })).then(build => Domain.Build.fromJSON(build))

    // Return the build binary.
    return build.toBinary
  }

  private buildRa = async (buildId: string, buildProfileId: string) => {
    // Build the binary.
    const buildOutput = (() => {
      try {
        return Helper.Shell.exe(`cd ${Service.Git.pathPrefix}${buildId}/xena-ra && yarn && yarn build`)
      } catch (e) {
        console.warn('Unable to build the binary output.')
        console.warn(e)
        return 'ERROR'
      }
    })()
  
    if (buildOutput == 'ERROR')
      throw Error('Unable to clone Git repository.')
    
    // Repo cleaning.
    // try {
    //   Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    // } catch (e) {
    //   console.warn(e)
    //   return 'ERROR'
    // }

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(`${Env.get('BUILD_DESTINATION')}${buildId}`)

    // Store the build.
    const build = await Repo.Build.insert( Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })).then(build => Domain.Build.fromJSON(build))

    // Return the build binary.
    return build.toBinary
  }
}
