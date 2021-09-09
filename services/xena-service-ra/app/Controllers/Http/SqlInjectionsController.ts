import * as Validator from 'App/Validators'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class SqlInjectionsController {
  public scan = async ({ request, response }: HttpContextContract) => {
    const { url } = await request.validate(Validator.SqlInjections.Scan)
  }
}
