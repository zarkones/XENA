<template>
  <v-card
    tile
  >
    <v-tabs
      v-model = 'selectedTab'
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
          v-if = 'item.content == "SHELL"'
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
        </v-card>

        <v-card
          v-if = 'item.content == "SYSTEM_MONITOR"'
          flat
        >
          Feature is disabled. Development is progressing. Consider opening a pull-request for the implementation of system monitoring.
        </v-card>
      </v-tab-item>
    </v-tabs-items>

  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import EventBus from '@/src/EventBus'

import MessageDisplay from '@/components/terminal/message-display.vue'
import Shell from '@/components/terminal/shell.vue'

import * as Service from '@/src/services'

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

  mounted () {
    EventBus.$on('interactionDialogUpdateClients', (clients: any[]) => this.clients = clients)

    EventBus.$on('interactionDialogUpdateSelectedClient', async (clientId: string) => {
      this.selectedClient = this.clients.filter(client => client.id == clientId)[0]
      const messages = await Service.Atila.fetchMessages(this.$axios, this.selectedClient.id, true)
      this.messages = messages.length
        ? messages.map(message => ({ ...message, content: Buffer.from(message.content.split('.')[1], 'base64') }))
        : []
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
