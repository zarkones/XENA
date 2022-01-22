import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'
import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { v4 as uuidv4 } from 'uuid'

export default class ClientsController {
  // Returns a single bot client.
  public get = async ({ request, response }: HttpContextContract) => {
    const { id, status } = await request.validate(Validator.Client.Get)

    const maybeClient = await Repo.Client.get({ id, status })
    if (!maybeClient)
      return response.noContent()

    const client = Domain.Client.fromJSON(maybeClient)
    
    return response.ok(client)
  }

  // Returns multiple bot clients.
  public getMultiple = async ({ request, response }: HttpContextContract) => {
    const { page, status } = await request.validate(Validator.Client.GetMultiple)

    const maybeClients = await Repo.Client.getMultiple({ page, status })
    if (!maybeClients.length)
      return response.noContent()

    const clients = maybeClients.map(client => Domain.Client.fromJSON(client))

    return response.ok(clients)
  }

  // Bot reaches out to this endpoint to get identified.
  public insert = async ({ request, response }: HttpContextContract) => {
    const { id, os, arch, publicKey, status } = await request.validate(Validator.Client.Insert)

    const maybeClient = await Repo.Client.get({ id })
    if (maybeClient)
      return response.conflict({ success: false, message: 'Client ID has been taken.' })
    
    const system = await (async () => {
      const maybeSystem = await Repo.System.get({ name: os })
      if (maybeSystem)
        return Domain.System.fromJSON(maybeSystem)
      return Domain.System.fromJSON(await Repo.System.insert(Domain.System.fromJSON({ id: uuidv4(), name: os, arch, count: 0 })))
    })()

    await Repo.Client.insert(Domain.Client.fromJSON({ id, osId: system.id, ip: request.ip(), publicKey, status, system }))
    
    await Repo.System.update({ id: system.id, count: system.count + 1 })
    
    return response.created()
  }

  public update = async ({}: HttpContextContract) => {
    // todo
  }

  public delete = async ({}: HttpContextContract) => {
    // todo
  }

  // Returns timestamps as an array of when clients where created.
  public getCount = async ({ response }: HttpContextContract) => {
    // const { os } = await request.validate(Validator.Client.Count)

    const count = await Repo.Client.count()
    if (!count.length)
      return response.noContent()
    return response.ok(count)
  }

  // Returns stats of clients available on a operating system.
  public demographic = async ({ request, response }: HttpContextContract) => {
    const { os: name } = await request.validate(Validator.Client.Demographic)

    if (name) {
      const maybeSystem = await Repo.System.get({ name })
      if (!maybeSystem)
        return response.noContent()
      
      return response.ok(Domain.System.fromJSON(maybeSystem))
    }

    const maybeSystems = await Repo.System.getMultiple()
    if (!maybeSystems.length)
      return response.noContent()
    
    return response.ok(maybeSystems.map(s => Domain.System.fromJSON(s)))
  }
}
