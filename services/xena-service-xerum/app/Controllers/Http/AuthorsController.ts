import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class AuthorsController {
  public insert = async ({ request, response }: HttpContextContract) => {
    const { name, publicKey } = await request.validate(Validator.Authors.Insert)

    const id = uuidv4()

    const author = await Repo.Authors.insert({ id, name, publicKey })

    return response.ok(author)
  }

  public get = async ({ request, response }: HttpContextContract) => {
    const { authorId } = await request.validate(Validator.Authors.Get)

    const author = await Repo.Authors.get({ id: authorId })
    if (!author)
      return response.notFound({ success: false, message: 'User not found.' })

    return response.ok(author)
  }
}
