<template>
  <v-card
    flat
  >
    <v-card-title>
      <v-btn
        v-if = '!actions'
        @click = 'actions = !actions'
        color = 'blue'
        small
        width = '100%'
      >
        Open Tools
      </v-btn>

      <v-btn
        v-if = 'actions'
        @click = 'actions = !actions'
        class = '
          blue darken-4
        '
        small
        width = '100%'
      >
        Close Tools
      </v-btn>
    </v-card-title>

    <v-card-actions>
      <ActionPanel
        v-if = 'actions'
        :domain = 'url'
      ></ActionPanel>
    </v-card-actions>

    <v-card-text>
      <v-expansion-panels accordion>
        <v-expansion-panel
          v-for = '(result, i) in data'
          :key = 'i'
        >
          <v-expansion-panel-header>
            {{ result.messages.join(' ') }}
          </v-expansion-panel-header>

        <v-expansion-panel-content>
          <v-simple-table dark>
            <template v-slot:default>
              <thead>
                <tr>
                  <th class = 'text-left'>
                    Domains
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for = '(host, i) in result.data.result'
                  :key = 'i'
                >
                  <td>{{ host }}</td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
        </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card-text>
  </v-card>
</template>

<script>
import ActionPanel from '@/components/nodes/action-panel.vue'
import EventBus from '@/src/EventBus'

export default {
  components: {
    ActionPanel
  },

  props: {
    url: {
      required: true,
      type: String
    }
  },

  data () {
    return {
      data: [],
      actions: false,
    }
  },

  mounted () {
    EventBus.$on('nodeResultsAdd', (result) => {
      this.data.push(result)
    })
  }
}
</script>