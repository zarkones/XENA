<template>
  <v-card
    tile
    flat
  >
    <v-card-title
      class = '
        purple-color
      '
    >
      Reconnaissance
    </v-card-title>

    <v-card-text>
      Set of security tools.
    </v-card-text>

    <div class = 'ma-4'>
      <v-expansion-panels
        if = 'tools.length'
        accordion
        tile
      >
        <v-expansion-panel
          v-for = '(tool, toolIndex) in tools'
          :key = 'toolIndex'
        >
          <v-expansion-panel-header>
            {{ tool.name }}
          </v-expansion-panel-header>

          <v-expansion-panel-content>
            <Nmap
              v-if = 'tool.name == "Nmap"'
            />
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>
  </v-card>
</template>


<script lang = 'ts'>
import Vue from 'vue'

import Nmap from '@/components/tools/nmap.vue'

import EventBus from '@/src/EventBus'

import * as Service from '@/src/services'

export default Vue.extend({
  components: {
    Nmap,
  },

  data: () => ({
    tools: [
      {
        name: 'Nmap',
        description: 'Network Mapper.',
      },
      {
        name: 'Sublist3r',
        description: 'Passive subdomain mapper.',
      },
      {
        name: 'Subdomain Brute-Forcer',
        description: 'Enumerate reachable domains with worldlist.',
      },
    ] as const
  }),

  methods: {
  },

  mounted () {
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
