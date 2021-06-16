import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class MachinesController {
  public details = async ({ response }: HttpContextContract) => {
    const details = {
      time: await Domain.Machine.time(),
      curentSpeed: await Domain.Machine.time(),
      battery: await Domain.Machine.battery(),
      cpu: await Domain.Machine.cpu(),
    }

    return response.ok(details)
  }
}
