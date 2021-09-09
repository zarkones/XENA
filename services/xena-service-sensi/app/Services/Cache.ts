import Redis from '@ioc:Adonis/Addons/Redis'

class Cache {
  public set = (key: string, data: string) => Redis.set(key, data)
    .catch(error => console.warn(error))
    .then(data => data)

  public get = (key: string) => Redis.get(key)
    .catch(error => console.warn(error))
    .then(data => data)
}

export default new Cache()