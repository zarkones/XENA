import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class ServicesController {
  public getMultiple = async ({ response }: HttpContextContract) => {
    const services = await Repo.Service.getMultiple()
    if (!services.length)
      return response.noContent()

    return response.ok(services)
  }

  public insert = async ({ request, response }: HttpContextContract) => {
    const { address, port } = await request.validate(Validator.Service.Insert)

    await Repo.Service.insert({ id: uuidv4(), address, port })

    return response.noContent()
  }
}
