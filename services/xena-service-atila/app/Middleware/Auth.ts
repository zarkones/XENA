import jwt from 'jsonwebtoken'
import Env from '@ioc:Adonis/Core/Env'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class Auth {
  public async handle ({ request, response }: HttpContextContract, next: () => Promise<void>) {
    const authHeader = request.header('Authorization')
    if (!authHeader)
      return response.unprocessableEntity({ success: false, message: 'Supply the auth. header.' })
    
    try {
      jwt.verify(authHeader.split('Bearer ')[1], Env.get('TRUSTED_PUBLIC_KEY'), { algorithms: ['RS512'] })
      await next()
    } catch {
      return response.unauthorized({ success: false })
    }

    return response.unauthorized({ success: false })
  }
}
