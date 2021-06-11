<template>
  <v-container>
    <v-card>
      <v-card-title>
        Author Studio
      </v-card-title>

      <v-card-text>
        Here you'll be able to manage encoding and building of software binaries.
      </v-card-text>

      

      <div
        class = '
          pt-4
          mx-4
        '
      >
        <v-text-field
          dense
          v-model = 'build.name'
          label = 'Name'
          color = 'rgba(189, 147, 249, 1)'
        ></v-text-field>

        <v-list dense>
          <v-subheader> Encoding Type </v-subheader>

          <v-list-item-group
            v-model = 'item'
            class = '
              mb-4
            '
          >
            <v-list-item
              v-for = '(encoding, i) in encodings'
              :key = 'i'
            >
              <v-list-item-content>
                <v-list-item-title v-text = 'encoding.name'></v-list-item-title>
              </v-list-item-content>
            </v-list-item>

          </v-list-item-group>
        </v-list>

        <v-text-field
          dense
          v-model = 'build.encodingIterations'
          label = 'Encoding Iterations'
          color = 'rgba(189, 147, 249, 1)'
        ></v-text-field>
      </div>

    </v-card>
  </v-container>
</template>

<script lang = 'ts'>
import Vue from 'vue'

export default Vue.extend({
  components: {
  },

  data: () => ({
    build: {
      name: '',
      encoding: '',
      encodingIterations: 1,
    },

    encodings: [
      {
        key: 'SHIKATA_GA_NAI',
        name: 'Shikata Ga Nai'
      },
    ] as const,
  }),

  methods: {
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>
</style>
