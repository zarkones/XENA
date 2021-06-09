<template>
  <v-card
    width = '100%'
    height = '72px'
    class = '
    '
  >
    <v-card-actions>
      <v-text-field
        label = 'Host address'
        dense
        v-model = 'host'
        class = '
        '
      >
      </v-text-field>

      <v-btn
        small
        class = '
          blue
        '
        @click = 'add_host()'
      >
        Add
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import EventBus from '@/src/EventBus'

export default {
  components: {
  },

  data () {
    return {
      host: '',
    }
  },

  methods: {
    add_host() {
      EventBus.$emit('scopeAdd', this.host)
    }
  }
}
</script>