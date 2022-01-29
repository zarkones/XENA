import jwt from 'jsonwebtoken'

export default new class Crypto {
  public sign = (signingKey: string, data: any) =>
    jwt.sign(data, signingKey, { algorithm: 'RS512', expiresIn: '32d', notBefore: 0 })

  public verify = (verificationKey: string, data: any) =>
    jwt.verify(data, verificationKey, { algorithms: ['RS512'] })
}