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
            ></v-text-field>

            <p
              class = '
                service-label
              '
            >
              Address of Xena-Pramid.
            </p>
            <v-text-field
              v-model = 'pyramidHost'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
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
              Your private key used for signing of messages.
            </p>
            <v-text-field
              v-model = 'privateKey'
              outlined
              dense
              color = 'rgba(189, 147, 249, 1)'
              @change = 'setPrivateKey'
            ></v-text-field>
          </div>
        </v-card>
      </v-tab-item>
    </v-tabs-items>

  </v-card>
</template>
<script lang = 'ts'>
import Vue from 'vue'

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

    // Tab: Connections.
    atilaHost: process.env.XENA_ATILA_HOST as string,
    pyramidHost: process.env.XENA_PYRAMID_HOST as string,
    raHost: process.env.XENA_RA_HOST as string,

    // Tab: Identity.
    privateKey: '',
  }),

  methods: {
    setPrivateKey () {
      this.$store.dispatch('setPrivateKey', this.privateKey)
    }
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>
.service-label {
  font-size: 18px;
  margin-bottom: 4px;
}
</style>
