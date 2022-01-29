import KeywordExtractor from 'keyword-extractor'
import urlParse from 'url-parse'

import { validNumber, validString } from './Validators'
import { CheerioAPI, load } from 'cheerio'
import { findPhoneNumbersInText } from 'libphonenumber-js'

type Headers = Record<string, string>

export default class WebPage {
  private readonly $: CheerioAPI
  
  constructor (
    public readonly url: string,
    public readonly headers: Headers,
    public readonly method: string,
    public readonly source: string,
    public readonly status: number,
  ) {
    this.url = validString(url, 'BAD_WEB_PAGE_URL', 'NON_EMPTY')
    this.headers = headers ?? {}
    this.method = validString(method, 'BAD_WEB_PAGE_METHOD', 'NON_EMPTY')
    this.source = validString(source, 'BAD_WEB_PAGE_SOURCE', 'NON_EMPTY')
    this.status = validNumber(status, 'BAD_WEB_PAGE_STATUS', true)
    this.$ = load(this.source)
  }

  public static fromJson = (json) => new WebPage(
    json.url,
    json.headers,
    json.method,
    json.source,
    json.status,
  )

  public get asJSON () {
    return {
      url: this.url,
      status: this.status,
      method: this.method,
      headers: this.headers,
      source: this.source,
      keywords: this.keywords(),
      phoneNumbers: this.phoneNumbers(),
    }
  }

  public keywords = (withHtml?: boolean) => {
    return KeywordExtractor.extract(withHtml ? this.source : this.$.text(), {
      language: 'english',
      remove_digits: true,
      return_changed_case:true,
      remove_duplicates: true,
    })
  }

  public phoneNumbers = () => findPhoneNumbersInText(this.source).map(phone => phone.number)

  public static parseUrl = (url: string) => new urlParse(url)
}