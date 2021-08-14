import * as Domain from 'App/Domains'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class ReadersController {
  public get = async ({ request, response }: HttpContextContract) => {
    const { data, headers, method, status, } = await request.all()

    const webPage = Domain.WebPage.fromJson({
      source: data,
      headers: headers,
      method: method,
      status: status,
    })

    return response.ok({
      keywords: webPage.keywords(),
      phoneNumbers: webPage.phoneNumbers(),
    })
  }
}
