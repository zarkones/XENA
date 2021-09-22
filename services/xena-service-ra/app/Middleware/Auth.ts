import Env from '@ioc:Adonis/Core/Env'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class Auth {
  public async handle ({ request, response }: HttpContextContract, next: () => Promise<void>) {
    const authHeader = request.header('Authorization')
    if (authHeader !== Env.get('API_SECRET'))
      return response.unauthorized()

    await next()
  }
}
