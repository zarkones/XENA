import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class ClientsController {
  public get = async ({ request, response }: HttpContextContract) => {
    const { id, status } = await request.validate(Validator.Client.Get)

    const maybeClient = await Repo.Client.get({ id, status })
    if (!maybeClient)
      return response.noContent()

    const client = Domain.Client.fromJSON(maybeClient)
    
    return response.ok(client)
  }

  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { page, status } = await request.validate(Validator.Client.GetMultiple)

    const maybeClients = await Repo.Client.getMultiple({ page, status })
    if (!maybeClients.length)
      return response.noContent()

    const clients = maybeClients.map(client => Domain.Client.fromJSON(client))

    return response.ok(clients)
  }

  public insert = async ({ request, response }: HttpContextContract) => {
    const { id, publicKey, status } = await request.validate(Validator.Client.Insert)

    const maybeClient = await Repo.Client.get({ id })
    if (maybeClient)
      return response.conflict({ success: false, message: 'Client ID has been taken.' })
    
    const client = await Repo.Client.insert(Domain.Client.fromJSON({ id, publicKey, status }))
      .then(client => Domain.Client.fromJSON(client))

    return response.ok(client)
  }

  public update = async ({}: HttpContextContract) => {
    // todo
  }

  public delete = async ({}: HttpContextContract) => {
    // todo
  }

  public getCount = async ({ response }) => {
    const count = await Repo.Client.getCount()
    if (!count.length)
      return response.noContent()

    return response.ok(count)
  }
}
