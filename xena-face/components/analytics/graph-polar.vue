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
      disabled
      color = 'rgba(189, 147, 249, 1)'
      @click = 'view_graph_data'
    >
      {{ title }}
    </v-btn>

    <v-divider></v-divider>

    <canvas
      :id = '`pie-chart-${targetPlatform}`'
    ></canvas>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import { ZombieRepo } from '@/src/api/zombie'
import EventBus from '@/src/EventBus'

import { Chart } from 'chart.js'

export default Vue.extend({
  data: () => ({
    platforms: {} as any,
    updated_at: '',
    chart: null,
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
      this.platforms = await ZombieRepo.get_zombies_analytics_pie(this.$axios, this.targetPlatform)
    },

    view_graph_data () {
      EventBus.$emit('clients_table_update', this.targetPlatform)
    },

    chart_init () {
      if (this.chart)
        return

      this.chart = new Chart(document.getElementById(`pie-chart-${this.targetPlatform}`), {
        type: 'polarArea',
        data: {
          labels: Object.keys(this.platforms),
          datasets: [
            { 
              data: Object.keys(this.platforms).map(p => this.platforms[p]),
              label: "Clients",
              borderColor: '#6272a4',
              hoverBorderColor: '#bd93f9',
              borderWidth: 0.9,
              fill: true,
            },
          ]
        },
        options: {
          title: {
            display: false,
            text: this.title
          },
          scale: {
            display: false
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
                display: false
              }
            }]
          },
          legend: {
            display: true
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

    // EventBus.$on('interactionDialog_update_zombies', (zombies: Zombie[]) => {
    //   if (!zombies.length) {
    //     this.update_graph()
    //     return
    //   }
    // 
    //   this.zombies = zombies.length == 1
    //     ? [0, Date.parse(zombies[0].created_at)]
    //     : zombies
    //       .map(z => z.platform == this.targetPlatform || !this.targetPlatform ? Date.parse(z.created_at) : 0)
    //       .sort()
    // })
  }
})
</script>

<style lang = 'css' scoped>

</style>