<template>
  <v-dialog
    v-model = 'enabled'
    fullscreen
    hide-overlay
    transition = 'dialog-bottom-transition'
  >
    <template v-slot:activator>
      <v-btn
        @click = 'interactionDialog'
        text
        class = '
          purple-color
        '
        width = '100%'
        tile
      >
        Interaction
      </v-btn>
    </template>

    <v-card>
      <ClientsDisplay
        v-if = 'clients.length'
        :clients = 'clients'
      />

      <div
        style = '
          margin-right: 57px !important;
        '
      >
        <v-toolbar
          dense
        >
          <v-toolbar-title>
            Interaction Dialog
            {{ selectedClient ? ` - ${selectedClient.id}` : '' }}
          </v-toolbar-title>

          <v-spacer></v-spacer>

          <v-btn
            icon
            dark
            @click = 'interactionDialog'
          >
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-toolbar>

        <h1
          v-if = '!clients.length'
          class = '
            ma-4
          '
        >
          Please, select at least one client.
        </h1>

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
  </v-dialog>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import EventBus from '@/src/EventBus'

import ClientsDisplay from '@/components/interaction/clients-display.vue'
import MessageDisplay from '@/components/interaction/message-display.vue'
import Shell from '@/components/interaction/shell.vue'

import * as Service from '@/src/services'

export default Vue.extend({
  components: {
    ClientsDisplay,
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
