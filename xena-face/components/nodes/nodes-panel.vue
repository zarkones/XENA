<template>
  <v-row justify = 'center'>
    <v-expansion-panels accordion>
      <v-expansion-panel
        v-for = '(host, i) in hosts'
        :key = 'i'
      >
        <v-expansion-panel-header>{{ host }}</v-expansion-panel-header>
        <v-expansion-panel-content>
          <Node
            :url = 'host'
          ></Node>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-row>
</template>

<script>
import EventBus from '@/src/EventBus'
import Node from '@/components/nodes/node.vue'

export default {
  components: {
    Node
  },

  data () {
    return {
      hosts: []
    }
  },

  mounted () {
    EventBus.$on('scopeAdd', (host) => {
      this.hosts.push(host)
    })
  }
}
</script>