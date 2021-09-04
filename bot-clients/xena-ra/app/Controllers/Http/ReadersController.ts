import * as Domain from 'App/Domains'
import * as Validator from 'App/Validators'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import axios from 'axios'

export default class ReadersController {
  public webSearch = async ({ request, response }: HttpContextContract) => {
    const { term } = await request.validate(Validator.Readers.WebSearch)

    const searchResult = await axios.post(`https://lite.duckduckgo.com/lite?q=${term}`, {}, {
      headers: {
        'content-type': 'application/x-www-form-urlencoded',
      },
    }).catch(error => console.log(error))

    if (!searchResult)
      return response.internalServerError({ success: false, message: 'Failed to duck it.' })

    return response.ok(searchResult.data)
  }

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
