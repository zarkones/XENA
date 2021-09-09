import * as Validator from 'App/Validators'
import * as Domain from 'App/Domains'
import * as Data from 'App/Data'

import axios from 'axios'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class WebFuzzersController {
  public scan = async ({ request, response }: HttpContextContract) => {
    const { url, method, regularInput } = await request.validate(Validator.WebFuzzers.Scan)

    const parsedUrl = Domain.WebPage.parseUrl(url)

    const responses = (regularInput
      ? [ ...Data.Worldlists.Basic, ...regularInput ]
      : Data.Worldlists.Basic)
      .map(payload => axios
      .request({
        url: `${url}${payload}`,
        method,
      })
      .then(data => ({
        success: true,
        payload,
        response: data.data,
        status: data.status,
      }))
      .catch(error => ({ success: false, payload, error: error.message, request: error.config }))
    )

    return response.ok(await Promise.all(responses))
  }
}
