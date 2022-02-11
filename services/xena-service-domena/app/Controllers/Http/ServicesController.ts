import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class ServicesController {
  public getMultiple = async ({ response }: HttpContextContract) => {
    const maybeService = await Repo.Service.getMultiple()
    if (!maybeService.length)
      return response.noContent()

    const services = maybeService.map(service => Domain.Service.fromJSON(service))

    return response.ok(services)
  }

  public insert = async ({ request, response }: HttpContextContract) => {
    const { address, port, details } = await request.validate(Validator.Service.Insert)

    const newService = Domain.Service.fromJSON({ id: uuidv4(), address, port, details })

    await Repo.Service.insert(newService)

    return response.created()
  }
}
