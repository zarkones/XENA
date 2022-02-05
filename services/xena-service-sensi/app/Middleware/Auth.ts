import jwt from 'jsonwebtoken'
import Env from '@ioc:Adonis/Core/Env'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class Auth {
  public async handle ({ request, response, logger }: HttpContextContract, next: () => Promise<void>) {
    const authHeader = request.header('Authorization')
    if (!authHeader)
      return response.unprocessableEntity({ success: false, message: 'Supply the auth. header.' })
    
    try {
      jwt.verify(authHeader.split('Bearer ')[1], Env.get('TRUSTED_PUBLIC_KEY').replace(/\\n/g, '\n'), { algorithms: ['RS512'] })
      await next()
    } catch (e) {
      logger.warn(e)
      return response.unauthorized({ success: false })
    }
  }
}
