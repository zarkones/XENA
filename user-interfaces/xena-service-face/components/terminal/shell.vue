<template>
  <div>
    <v-text-field
      :disabled = '!getPrivateKey'
      dense
      v-model = 'shellCode'
      outlined
      label = 'Enter a command...'
      color = 'rgba(189, 147, 249, 1)'
      @change = 'issueMessages'
    ></v-text-field>

    <h3
      v-if = '!getPrivateKey'
    >
      Please, go to Settings -> Identity and enter your private key in order to sign the messages.
    </h3>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import jwt from 'jsonwebtoken'

import { EncryptJWT, importJWK, importPKCS8, importSPKI, importX509 } from 'jose'

import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    shellCode: '',
    encryptPayload: true,
  }),

  props: {
    clients: {
      required: true,
    }
  },

  computed: {
    ...mapGetters([
      'getPrivateKey',
    ])
  },

  methods: {
    async craftMessage (key: string, data: any, encrypt: boolean = true) {
      return encrypt
        ? new EncryptJWT({ 'urn:example:claim': true })
          .setProtectedHeader({ alg: 'RSA', enc: 'A256GCM' })
          .setIssuedAt()
          .setIssuer('urn:example:issuer')
          .setAudience('urn:example:audience')
          .setExpirationTime('32d')
          .encrypt(await importPKCS8(key, 'RS512'))
        : jwt.sign(data, key, { algorithm: 'RS512', expiresIn: '32d', notBefore: 0 })
    },

    async issueMessages () {
      const message = await this.craftMessage(this.getPrivateKey, { shell: this.shellCode }, this.encryptPayload)

      console.log(message)

      await Service.Atila.publishMessage(this.$axios, this.clients[0].id, 'instruction', message)
      // const createdMessages = await Promise.all(this.clients.map(client => {
      //   return Service.Atila.publishMessage(this.$axios, client.id, 'instruction', message)
      // }))

      this.shellCode = ''
    }
  },
})
</script>

<style scoped>
</style>