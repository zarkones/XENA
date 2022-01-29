import * as Validator from 'App/Validators'
import * as Domain from 'App/Domains'
import * as Data from 'App/Data'

import axios from 'axios'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class WebFuzzersController {
  public scan = async ({ request, response }: HttpContextContract) => {
    const { url, method, wordlist } = await request.validate(Validator.WebFuzzers.Scan)

    const parsedUrl = Domain.WebPage.parseUrl(url)

    const completeWordlist = wordlist?.length ? wordlist : Data.Worldlists.Basic

    const responses = await Promise.all(
      completeWordlist.map(payload => axios
        .request({
          url: `${url}${payload}`,
          method,
        })
        .then(data => ({
          success: true,
          payload,
          ...data,
        }))
        .catch(error => ({ success: false, payload, error: error.message, request: error.config }))
    ))

    const pages = responses.map(resp => {
      if (resp['error'])
        return resp

      return Domain.WebPage.fromJson({
        method,
        url: resp['config']['url'],
        headers: resp['headers'],
        source: resp['data'],
        status: resp['status'],
      }).asJSON
    })

    return response.ok(pages)
  }
}
