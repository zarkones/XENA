<template>
  <v-card
    tile
  >
    <v-tabs
      v-model = 'selectedTab'
      class = '
        setting-tabs
      '
    >
      <v-tab
        v-for = 'item in tabItems'
        :key = 'item.tab'
      >
        {{ item.tab }}
      </v-tab>
    </v-tabs>

    <v-tabs-items v-model = 'selectedTab'>
      <v-tab-item
        v-for = 'item in tabItems'
        :key = 'item.tab'
      >
        <v-card
          v-if = 'item.content == `CONNECTIONS`'
          flat
        >
          <v-card-text>
            Here you'll be able to configure your client's connections.
          </v-card-text>
          
          <div
            class = '
              pt-4
              mx-4
              mb-4
            '
          >
            <p
              class = '
                service-label
              '
            >
              Address of Xena-Atila
            </p>

            <v-text-field
              v-model = 'atilaHost'
              dense
              outlined
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveAtilaHost'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Token for Xena-Atila
            </p>

            <v-text-field
              v-model = 'atilaToken'
              dense
              outlined
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveAtilaToken'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Address of Xena-Pyramid
            </p>
            <v-text-field
              v-model = 'pyramidHost'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'savePyramidHost'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Token for Xena-Pyramid
            </p>

            <v-text-field
              v-model = 'pyramidToken'
              dense
              outlined
              color = 'rgba(189, 147, 249, 1)'
              @change = 'savePyramidToken'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Address of Xena-Ra
            </p>
            <v-text-field
              v-model = 'raHost'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveRaHost'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Token for Xena-Ra
            </p>

            <v-text-field
              v-model = 'raToken'
              dense
              outlined
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveRaToken'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Address of Xena-Domena
            </p>
            <v-text-field
              v-model = 'domenaHost'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveDomenaHost'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Token for Xena-Domena
            </p>

            <v-text-field
              v-model = 'domenaToken'
              dense
              outlined
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveDomenaToken'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Address of Xena-Sensi
            </p>
            <v-text-field
              v-model = 'sensiHost'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveSensiHost'
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Token for Xena-Sensi
            </p>

            <v-text-field
              v-model = 'sensiToken'
              dense
              outlined
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveSensiToken'
            ></v-text-field>
          </div>
        </v-card>

        <v-card
          v-if = 'item.content == `IDENTITY`'
          flat
        >
          <v-card-text>
            Here you'll be able to configure your identity.
            This settings play an important role when it comes to authorization and authentication.
            Don't share private keys, except if you know what you're doing.
          </v-card-text>

          <div
            class = '
              pt-4
              mx-4
            '
          >
            <p
              class = '
                service-label
              '
            >
              Username: 
              <v-btn
                tile
                text
                small
                color = 'rgba(189, 147, 249, 1)'
              >
                {{ getUsername }}
              </v-btn>
            </p>

            <p
              class = '
                service-label
                mt-4
              '
            >
              Your private key used for signing of messages.
            </p>
            <v-text-field
              v-model = 'privateKey'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'setPrivateKey'
            ></v-text-field>

            <p
              class = '
                service-label
                mt-4
              '
            >
              Generate a token using your private key, for authorization of calls to services.
            </p>
            
            <v-btn
              @click = 'generateNewToken'
              tile
              small
              outlined
              color = 'rgba(189, 147, 249, 1)'
              class = '
              '
              width = '100%'
            >
              Generate a Token
            </v-btn>

            <p
              class = '
                service-label
                mt-4
              '
            >
              {{ newToken }}
            </p>
          </div>

        </v-card>
      </v-tab-item>
    </v-tabs-items>

  </v-card>
</template>
<script lang = 'ts'>
import Vue from 'vue'

import { mapActions, mapGetters } from 'vuex'

import * as Service from '@/src/services'

export default Vue.extend({
  components: {
  },

  data: () => ({
    // Tab's configuration.
    selectedTab: 0,
    tabItems: [
      { tab: 'Connections', content: 'CONNECTIONS' },
      { tab: 'Identity', content: 'IDENTITY' },
    ],

    atilaHost: '',
    atilaToken: '',

    pyramidHost: '',
    pyramidToken: '',
    
    raHost: '',
    raToken: '',

    domenaHost: '',
    domenaToken: '',

    sensiHost: '',
    sensiToken: '',

    // Tab: Identity.
    privateKey: '',
    newToken: '',
  }),

  methods: {
    generateNewToken () {
      this.newToken = Service.Crypto.sign(this.getPrivateKey, {})
    },

    setPrivateKey () {
      this.setPrivateKey(this.privateKey)
    },

    saveAtilaHost () {
      this.setAtilaHost(this.atilaHost)
    },
    saveAtilaToken () {
      this.setAtilaToken(this.atilaToken)
    },

    savePyramidHost () {
      this.setPyramidHost(this.pyramidHost)
    },
    savePyramidToken () {
      this.setPyramidToken(this.pyramidToken)
    },

    saveRaHost () {
      this.setRaHost(this.raHost)
    },
    saveRaToken () {
      this.setRaToken(this.raToken)
    },

    saveDomenaHost () {
      this.setDomenaHost(this.domenaHost)
    },
    saveDomenaToken () {
      this.setDomenaToken(this.domenaToken)
    },

    saveSensiHost () {
      this.setSensiHost(this.sensiHost)
    },
    saveSensiToken () {
      this.setSensiToken(this.sensiToken)
    },

    ...mapActions([
      'setPrivateKey',
      'setAtilaHost',
      'setAtilaToken',
      'setPyramidHost',
      'setPyramidToken',
      'setRaHost',
      'setRaToken',
      'setDomenaHost',
      'setDomenaToken',
      'setSensiHost',
      'setSensiToken',
    ])
  },

  computed: {
    ...mapGetters([
      'getUsername',
      'getPrivateKey',
      'getAtilaHost',
      'getAtilaToken',
      'getPyramidHost',
      'getPyramidToken',
      'getRaHost',
      'getRaToken',
      'getDomenaHost',
      'getDomenaToken',
      'getSensiHost',
      'getSensiToken',
    ])
  },

  mounted () {
    this.privateKey = this.getPrivateKey

    this.atilaHost = this.getAtilaHost
    this.atilaToken = this.getAtilaToken

    this.pyramidHost = this.getPyramidHost
    this.pyramidToken = this.getPyramidToken

    this.raHost = this.getRaHost
    this.raToken = this.getRaToken

    this.domenaHost = this.getDomenaHost
    this.domenaToken = this.getDomenaToken

    this.sensiHost = this.getSensiHost
    this.sensiToken = this.getSensiToken
  },
})
</script>

<style lang = 'css' scoped>
.service-label {
  font-size: 18px;
  margin-bottom: 4px;
}

.setting-tabs {
  color: #bd93f9 !important;
  caret-color: #bd93f9 !important;
}
</style>
