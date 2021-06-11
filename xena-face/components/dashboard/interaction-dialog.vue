<template>
  <v-card
    tile
  >
    <div>
      <v-toolbar
        dense
      >
        <v-toolbar-title>
          <v-btn
            text
            tile
            small
            color = 'rgba(189, 147, 249, 1)'
          >
            Shell
          </v-btn>

          <v-btn
            text
            tile
            small
            color = 'rgba(189, 147, 249, 1)'
          >
            File System
          </v-btn>

          <v-btn
            text
            tile
            small
            color = 'rgba(189, 147, 249, 1)'
          >
            Gallery
          </v-btn>

          <v-btn
            text
            tile
            small
            color = 'rgba(189, 147, 249, 1)'
          >
            System Monitor
          </v-btn>

          <v-btn
            text
            tile
            small
            color = 'rgba(189, 147, 249, 1)'
          >
            Details
          </v-btn>
        </v-toolbar-title>
      </v-toolbar>

      <v-card-text
        v-if = '!clients.length'
        class = '
          ma-4
        '
      >
        Please, select at least one client.
      </v-card-text>

      <!--
        We should ask the user to select clients.
      -->
      <div
        v-if = 'clients.length'
        class = '
          ml-4
          mr-4
          mt-4
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
    </div>

  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import EventBus from '@/src/EventBus'

import MessageDisplay from '@/components/interaction/message-display.vue'
import Shell from '@/components/interaction/shell.vue'

import * as Service from '@/src/services'

export default Vue.extend({
  components: {
    MessageDisplay,
    Shell,
  },

  data: () => ({
    enabled: false,
    clients: [] as any[],
    messages: [] as any[],
    selectedClient: null,
  }),

  methods: {
    interactionDialog () {
      this.enabled = !this.enabled
    },

    async fetchMessagesWithReplies () {

    },
  },

  mounted () {
    EventBus.$on('interactionDialogUpdateClients', (clients: any[]) => this.clients = clients)

    EventBus.$on('interactionDialogUpdateSelectedClient', async (clientId: string) => {
      this.selectedClient = this.clients.filter(client => client.id == clientId)[0]
      const messages = await Service.Atila.fetchMessages(this.$axios, this.selectedClient.id, true)
      this.messages = messages.length
        ? messages.map(message => ({ ...message, content: Buffer.from(message.content, 'base64') }))
        : []
    })
  },
})
</script>

<style scoped>
.zombie-panels-grid {
  padding: 0px !important;
}

.v-application p {
  margin-bottom: 0px !important;
}

.v-text-field.v-text-field--enclosed .v-text-field__details {
  margin-bottom: 0px !important;
}
</style>
