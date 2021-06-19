import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class MachinesController {
  public details = async ({ response }: HttpContextContract) => {
    const machine = await Domain.Machine.serialize()

    return response.ok(machine)
  }
}
