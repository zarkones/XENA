<template>
  <v-expansion-panels accordion flat>
    <!-- RECON -->
    <v-expansion-panel>
      <v-expansion-panel-header
        class = '
          purple
        '
      >
        Recon
      </v-expansion-panel-header> 

      <v-expansion-panel-content
        class = '
          ma-2
        '
      >
        <v-row>
          <v-btn
            color = 'blue'
            small
            width = '100%'
            class = '
              mb-2
            '
            @click = 'tool_use("sublist3r")'
          >
            Sublist3r
          </v-btn>

          <v-btn
            color = 'blue'
            small
            width = '100%'
            class = '
              mb-2
            '
          >
            Nmap
          </v-btn>
        </v-row>
      </v-expansion-panel-content>
    </v-expansion-panel>

    <!-- Exploitation -->
    <v-expansion-panel>
      <v-expansion-panel-header
        class = '
          purple darken-2
        '
      >
        Exploitation
      </v-expansion-panel-header> 

      <v-expansion-panel-content
        class = '
          ma-2
        '
      >
        <v-row>
          <v-btn
            color = 'blue'
            small
            width = '100%'
            class = '
              mb-2
            '
          >
            SQLMap
          </v-btn>

          <v-btn
            color = 'blue'
            small
            width = '100%'
            class = '
              mb-2
            '
          >
            Metasploit
          </v-btn>
        </v-row>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script>
import EventBus from '@/src/EventBus'

export default {
  props: {
    domain: {
      required: true,
      type: String
    }
  },

  methods: {
    tool_use(name) {
      this.$axios({
        method: 'POST',
        url: `${process.env.XENA_KERNEL_HOST}/${process.env.DIR_BUSTER}/consume/tool`,
        data: {
          name,
          domain: this.domain
        },
        headers: {
  
        }
      })
      .catch(err => console.warn(err))
      .then(resp => {
        EventBus.$emit('nodeResultsAdd', resp.data)
      })
    }
  }
}
</script>