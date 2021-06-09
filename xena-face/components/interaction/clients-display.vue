<template>
  <v-navigation-drawer
    v-model = 'draw'
    :right = 'true'
    fixed
    :clipped = 'true'
    :mini-variant = 'true'
    expand-on-hover
  >
    <v-list
      dense
      nav
    >
      <v-list-item
        v-for = 'client in clients'
        :key = 'client.id'
        link
      >
        <!--v-list-item-icon>
          <v-icon>{{ client.icon }}</v-icon>
        </v-list-item-icon-->

        <v-list-item-content
          @click = 'selectClient(client.id)'
        >
          <v-list-item-title>{{ client.id }}</v-list-item-title>
        </v-list-item-content>
        
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import EventBus from '@/src/EventBus'

export default Vue.extend({
  components: {
  },

  data: () => ({
    draw: true,
  }),

  props: {
    clients: {
      required: true,
    }
  },

  methods: {
    selectClient (clientId: string) {
      EventBus.$emit('interactionDialogUpdateSelectedClient', clientId)
    }
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>

</style>
