<template>
  <v-card
    flat
  >
    <v-card-text
      v-if = '!clients.length'
    >
      Please, select at least one client.
    </v-card-text>

    <!--
      We should ask the user to select clients.
    -->
    <div
      v-if = 'clients.length'
      class = '
        ma-4
      '
    >
      <div
        v-for = 'message in messages'
        :key = 'message.id'
      >
        <MessageDisplay
          :message = 'message'
        />
      </div>

      <Shell
        :clients = 'clients'
      />
    </div>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import EventBus from '@/src/EventBus'

import MessageDisplay from '@/components/terminal/message-display.vue'
import Shell from '@/components/terminal/shell.vue'

import * as Service from '@/src/services'

import { mapGetters } from 'vuex'

export default Vue.extend({
  components: {
    MessageDisplay,
    Shell,
  },

  data: () => ({
    selectedTab: 0,
    tabItems: [
      { tab: 'Terminal', content: 'SHELL' },
      { tab: 'System Monitor', content: 'SYSTEM_MONITOR' },
    ],

    enabled: false,
    clients: [] as any[],
    messages: [] as any[],
    selectedClient: null,
  }),

  methods: {
    interactionDialog () {
      this.enabled = !this.enabled
    },

    selectTab (tabName: string) {
      this.selectedTab = tabName
    },

    async fetchMessagesWithReplies () {

    },
  },

  computed: {
    ...mapGetters([
      'getPrivateKey',
      'getAtilaHost',
      'getAtilaToken',
    ])
  },

  mounted () {
    EventBus.$on('interactionDialogUpdateClients', (clients: any[]) => this.clients = clients)

    EventBus.$on('interactionDialogUpdateSelectedClient', async (clientId: string) => {
      // Without this the rendering engine won't update. This needs to be fixed somehow.
      this.messages = []

      this.selectedClient = this.clients.filter(client => client.id == clientId)[0]

      const messages = await new Service.Atila(this.$axios, this.getAtilaHost, this.getAtilaToken)
        .fetchMessages(this.selectedClient.id, this.getPrivateKey, this.selectedClient.publicKey, true)
      
      this.messages = messages?.length ? messages : this.messages
    })
  },
})
</script>

<style scoped>
.v-application p {
  margin-bottom: 0px !important;
}

.v-text-field.v-text-field--enclosed .v-text-field__details {
  margin-bottom: 0px !important;
}
</style>
