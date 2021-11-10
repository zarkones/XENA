import * as Validator from 'App/Validators'
import * as Domain from 'App/Domains'
import * as Repo from 'App/Repos'

import NodeRSA from 'node-rsa'
import Env from '@ioc:Adonis/Core/Env'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class FilesController {
  public insert = async ({ request, response }: HttpContextContract) => {
    const { id, originId, originPath, data: rawData } = await request.validate(Validator.Files.Insert)

    const data = new NodeRSA(Env.get('PUBLIC_KEY')).encrypt(rawData, 'binary')

    const newFile = Domain.File.fromJSON({ id, originId, originPath, data }).asJSON

    const insertedFile = Repo.File.insert(newFile)
      .then(file => Domain.File.fromJSON(file).asJSON)

    return response.ok(insertedFile)
  }
}
