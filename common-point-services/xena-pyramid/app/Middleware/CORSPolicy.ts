import Env from '@ioc:Adonis/Core/Env'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class CorsPolicy {
  public async handle ({ request, response, logger }: HttpContextContract, next: () => Promise<void>) {
    const originHeader = request.headers().origin

    const allowedOrigins = Env.get('CORS_POLICY_ALLOWED_ORIGINS').split(',')

    /**
     * Check if the caller has a valid origin header.
     * 
     * If no origin header is present, it means that we're outside of the browser context.
     * Or that the caller's browser has been exploited in some way.
     * Since the origin header is supposed to be controled only by the browser.
     */

    if (originHeader && !allowedOrigins.includes(originHeader)) {
      logger.warn(`${request.ips().join(', ')} have requested a resource at ${request.url()} with a non-allowed origin of '${originHeader}'.`)
      return response.badRequest({ success: false, message: 'Origin header value not allowed.' })
    }

    await next()
  }
}
