<template>
  <v-card
    tile
  >
    <v-tabs
      v-model = 'selectedTab'
      style = 'color: #bd93f9 !important;caret-color: #bd93f9 !important;'
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
            '
          >
            <p
              class = '
                service-label
              '
            >
              Address of Xena-Atila.
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
              Token for Xena-Atila.
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
              Address of Xena-Pyramid.
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
              Address of Xena-Ra.
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
              Address of Xena-Domena.
            </p>
            <v-text-field
              v-model = 'domenaHost'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'saveDomenaHost'
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

    saveRaHost () {
      this.setRaHost(this.raHost)
    },

    saveDomenaHost () {
      this.setDomenaHost(this.domenaHost)
    },

    ...mapActions([
      'setPrivateKey',
      'setAtilaHost',
      'setAtilaToken',
      'setPyramidHost',
      'setRaHost',
      'setDomenaHost',
    ])
  },

  computed: {
    ...mapGetters([
      'getUsername',
      'getPrivateKey',
      'getAtilaHost',
      'getAtilaToken',
      'getPyramidHost',
      'getRaHost',
      'getDomenaHost',
    ])
  },

  mounted () {
    this.privateKey = this.getPrivateKey
    this.atilaHost = this.getAtilaHost
    this.atilaToken = this.getAtilaToken
    this.pyramidHost = this.getPyramidHost
    this.raHost = this.getRaHost
    this.domenaHost = this.getDomenaHost
  },
})
</script>

<style lang = 'css' scoped>
.service-label {
  font-size: 18px;
  margin-bottom: 4px;
}
</style>
