import Env from '@ioc:Adonis/Core/Env'
import jwt from 'jsonwebtoken'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class Auth {
  public async handle ({ request, response }: HttpContextContract, next: () => Promise<void>) {
    const maybeToken = request.header('Authorization')?.split('Bearer ')[1]
    if (!maybeToken)
      return response.unprocessableEntity({ success: false })
    
    try {
      jwt.verify(maybeToken, Env.get('TRUSTED_PUBLIC_KEY'), { algorithms: ['RS512'] })
    }
    catch {
      return response.unauthorized({ success: false })
    }
    
    await next()
  }
}
