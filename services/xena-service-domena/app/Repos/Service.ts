import Service from 'App/Models/Service'

export default new class ServiceRepo {
  public getMultiple = () => Service.query()
    .select('*')
    .exec()
    .then(services => services.map(service => service.serialize()))
  
  public insert = (payload: any) => Service.create(payload)
    .then(service => service.serialize())
}