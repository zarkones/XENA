<template>
  <v-card
    class = '
      mx-auto
      text-center
    '
    outlined
    tile
  >
    <v-btn
      block
      text
      tile
      small
      color = 'rgba(189, 147, 249, 1)'
      @click = 'view_graph_data'
    >
      {{ title }}
    </v-btn>

    <v-divider></v-divider>

    <canvas
      :id = '`line-chart-${targetPlatform}`'
      class = '
        pl-2
        pr-2
        pb-2
      '
    ></canvas>

    <v-divider></v-divider>

    <v-card-text>
      {{ start }} - {{ end }}
    </v-card-text>
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
    chart: null,
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
    },

    view_graph_data () {
      EventBus.$emit('clients_table_update', this.targetPlatform)
    },

    chart_init (update?: boolean) {
      if (this.chart && !update)
        return

      this.chart = new Chart(document.getElementById(`line-chart-${this.targetPlatform}`), {
        type: 'line',
        data: {
          labels: this.zombies,
          datasets: [
            { 
              data: this.zombies,
              label: "Clients",
              pointBackgroundColor: '#bd93f9',
              pointBorderColor: '#bd93f9',
              borderColor: '#6272a4',
              borderWidth: 2,
              radius: 1.5,
              fill: true,
            }
          ]
        },
        options: {
          title: {
            display: false,
            text: this.title
          },
          scales: {
            xAxes: [{
              ticks: {
                display: false
              },
              gridLines: {
                display: false
              }
            }],
            yAxes: [{
              ticks: {
                display: false
              },
              gridLines: {
                display: true
              }
            }]
          },
          legend: {
            display: false
          },
          tooltips: {
            enabled: true
          },
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
          .then(() => this.chart_init(true))
        return
      }

      this.zombies = zombies.length == 1
        ? [0, Date.parse(zombies[0].created_at)]
        : zombies
          .map(z => z.platform == this.targetPlatform || !this.targetPlatform ? Date.parse(z.created_at) : 0)
          .sort()
      
      this.chart_init(true)
    })
  }
})
</script>

<style lang = 'css' scoped>

</style>