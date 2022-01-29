import * as Repo from 'App/Repos'

import jwt from 'jsonwebtoken'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class Auth {
  public async handle ({ request, response }: HttpContextContract, next: () => Promise<void>) {
    const { name, payload } = request.all()
    if (!name || !payload)
      return response.unauthorized({ success: false, message: 'No authorization token supplied.' })

    const author = await Repo.Authors.getByName(name)
    if (!author)
      return response.notFound({ success: false, message: 'User not found.' })

    const decoded = jwt.verify(payload, Buffer.from(author.publicKey, 'base64').toString('utf-8'), { algorithms: ['RS512'] })

    request.updateBody({ ...decoded as jwt.JwtPayload, authorId: author.id })
    
    await next()
  }
}
