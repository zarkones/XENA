import * as Validator from 'App/Validators'

import axios from 'axios'

import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { execSync } from 'child_process'
import { quote } from 'shell-quote'

export default class SubdomainsController {
  public sublist3r = async ({ request, response }: HttpContextContract) => {
    const { domain } = await request.validate(Validator.Subdomains.Sublist3r)

    const command = quote(['sublist3r', '-d', domain])

    const rawOutput = execSync(command).toString('utf-8')

    return response.ok(rawOutput)
  }

  public bruteForce = async ({ request, response }: HttpContextContract) => {
    const { domain, dict } = await request.validate(Validator.Subdomains.BruteForce)

    const alive = [] as string[]
    const dead = [] as string[]

    for (const subdomain of dict) {
      const responseFromServer = await axios.get(`https://${subdomain}.${domain}`).catch(() => null)

      if (!responseFromServer) {
        dead.push(subdomain)
        continue
      }

      alive.push(subdomain)
    }
    return response.ok({ alive, dead })
  }
}
