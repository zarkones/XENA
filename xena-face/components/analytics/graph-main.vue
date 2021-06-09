<template>
  <v-card
    outlined
    tile
  >
    <canvas
      id = 'line-chart'
    ></canvas>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import { ZombieRepo, Zombie } from '@/src/api/zombie'
import EventBus from '@/src/EventBus'

import { Chart } from 'chart.js'

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

    chart_init () {
      const graph_chart_line = new Chart(document.getElementById('line-chart'), {
        type: 'line',
        data: {
          labels: this.zombies,
          datasets: [
            { 
              data: this.zombies,
              label: "Clients",
              borderColor: "#3e95cd",
              fill: true,
            }
          ]
        },
        options: {
          title: {
            display: true,
            text: `All platforms ${this.zombies.lenght}`
          }
        }
      })
    },
  },
  
  mounted () {
    this.update_graph()
      .then(() => this.chart_init())

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