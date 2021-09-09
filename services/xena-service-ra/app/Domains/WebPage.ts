import KeywordExtractor from 'keyword-extractor'
import urlParse from 'url-parse'

import { validNumber, validString } from './Validators'
import { CheerioAPI, load } from 'cheerio'
import { findPhoneNumbersInText } from 'libphonenumber-js'

export default class WebPage {
  public readonly headers: Record<string, string> | null
  public readonly method: string
  public readonly source: string
  public readonly status: number
  
  private readonly $: CheerioAPI
  
  constructor (
    headers: Record<string, string>,
    method: string,
    source: string,
    status: number,
  ) {
    this.headers = headers
    this.method = validString(method, 'BAD_WEB_PAGE_METHOD', 'NON_EMPTY')
    this.source = validString(source, 'BAD_WEB_PAGE_SOURCE', 'NON_EMPTY')
    this.status = validNumber(status, 'BAD_WEB_PAGE_STATUS', true)
    this.$ = load(this.source)
  }

  public static fromJson = (json) => new WebPage(
    json.headers,
    json.method,
    json.source,
    json.status,
  )

  public keywords = (withHtml?: boolean) => {
    return KeywordExtractor.extract(withHtml ? this.source : this.$.text(), {
      language: 'english',
      remove_digits: true,
      return_changed_case:true,
      remove_duplicates: true,
    })
  }

  public phoneNumbers = () => findPhoneNumbersInText(this.source).map(phone => phone.number)

  public static parseUrl = (url: string) => {
    return new urlParse(url)
  }
}