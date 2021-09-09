import * as Service from 'App/Services'
import * as Validator from 'App/Validators'
import * as Repo from 'App/Repos'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export default class ReadersController {
  public injection = async ({ request }: HttpContextContract) => {
    const { method, url, action, options, } =
      await request.validate(Validator.Databases.Injection)

    switch (action) {
      case 'GET_ALL_DATABASES':
        return Service.WebReader.makeRequest({
          method,
          url: url.replace('SQL_INJECTION', Repo.SQL.Postgers.allDatabases({
            prefix: options?.prefix,
            suffix: options?.suffix,
          })),
        }).catch(data => data)
    }
  }
}
