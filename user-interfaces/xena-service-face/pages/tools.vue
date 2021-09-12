<template>
  <div>
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
        <!-- RECON -->
        <v-card
          tile
          flat
          v-if = 'item.content == "RECON"'
        >
          <!--v-card-title
            class = '
              purple-color
            '
          >
            Reconnaissance
          </v-card-title-->

          <v-card-text>
            Set of tools meant to gather intelligence.
          </v-card-text>

          <div class = 'ma-4'>
            <Recon />
          </div>
        </v-card>

        <!-- SCANNERS -->
        <v-card
          tile
          flat
          v-if = 'item.content == "SCANNERS"'
        >
          <!--v-card-title
            class = '
              purple-color
            '
          >
            Scanners
          </v-card-title-->

          <v-card-text>
            Set of web scanners, fuzzers and such...
          </v-card-text>

          <div class = 'ma-4'>
            <Scanners />
          </div>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>


<script lang = 'ts'>
import Vue from 'vue'

import Recon from '@/components/tool-set/recon.vue'
import Scanners from '@/components/tool-set/scanners.vue'

export default Vue.extend({
  components: {
    Recon,
    Scanners,
  },

  data: () => ({
    selectedTab: 0,
    tabItems: [
      { tab: 'Recon', content: 'RECON' },
      { tab: 'Scanners', content: 'SCANNERS' },
    ],
  }),

  methods: {
  },

  mounted () {
  },
})
</script>