<template>
  <v-card
    class = 'mx-auto text-center'
    outlined
    tile
  >
    <v-card-title>
      {{ title }} 
      <span
        class = '
          graph-number
        '
      > 
        ({{ zombies.filter(z => z > 0).length }})
      </span>
    </v-card-title>

    <v-card-text>
      <v-sheet>
        <v-sparkline
          :value = 'zombies'
          color = 'rgba(255, 121, 198, 1)'
          height = '100'
          line-width = '1.0'
          stroke-linecap = 'round'
          smooth
          auto-draw
          :auto-draw-duration = '1000'
        >
        </v-sparkline>
      </v-sheet>
    </v-card-text>

    <v-card-text>
      {{ start }} - {{ end }}
    </v-card-text>

    <v-divider></v-divider>

    <v-card-actions
      class = '
        justify-center
        pa-0
      '
    >
      <v-btn
        block
        text
        tile
        small
        color = 'rgba(189, 147, 249, 1)'
        @click = 'view_graph_data'
      >
        View
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import { ZombieRepo, Zombie } from '@/src/api/zombie'
import EventBus from '@/src/EventBus'

export default Vue.extend({
  data: () => ({
    zombies: [] as number[],
    updated_at: '',
    start: '',
    end: '',
  }),

  props: {
    targetPlatform: {
      required: false,
      type: String,
    },
    title: {
      required: true,
      type: String,
    }
  },

  methods: {
    async update_graph () {
      const analytics = await ZombieRepo.get_zombies_analytics(this.$axios, this.targetPlatform)
      this.zombies = analytics.zombies
      this.start = analytics.start
      this.end = analytics.end

      this.updated_at = new Date(Date.now())
        .toISOString()
        .split('T')
        .join(' ')
        .split('.')[0]
    },

    view_graph_data () {
      EventBus.$emit('clients_table_update', this.targetPlatform)
    },
  },
  
  mounted () {
    this.update_graph()

    EventBus.$on('interactionDialog_update_zombies', (zombies: Zombie[]) => {
      if (!zombies.length) {
        this.update_graph()
        return
      }

      this.zombies = zombies.length == 1
        ? [0, Date.parse(zombies[0].created_at)]
        : zombies
          .map(z => z.platform == this.targetPlatform || !this.targetPlatform ? Date.parse(z.created_at) : 0)
          .sort()
    })
  }
})
</script>

<style lang = 'css' scoped>

</style>