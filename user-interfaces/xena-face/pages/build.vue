<template>
  <v-card
    tile
    flat
  >
    <v-card-title
      class = '
        purple-color
      '
    >
      Cloud Build
    </v-card-title>

    <v-card-text>
      Xena-Pyramid provides you with cloud build functionality. Meaning that you can create build profiles, which are used
      for creation of bot clients and other software. 
    </v-card-text>

    <div class = 'ma-4'>
      <v-expansion-panels>
        <v-expansion-panel>
          <v-expansion-panel-header>
            Create a build profile.
          </v-expansion-panel-header>

          <v-expansion-panel-content>
            <div>
              <v-text-field
                dense
                v-model = 'build.name'
                label = 'Name'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-text-field
                dense
                v-model = 'build.description'
                label = 'Description'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-text-field
                dense
                v-model = 'build.gitUrl'
                label = 'Git URL'
                :placeholder = 'build.gitUrl'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-select
                v-model = 'buildTemplate'
                :items = 'buildTemplates'
                label = 'Build templates'
                outlined
                dense
              ></v-select>

              <v-btn
                @click = 'insertBuildProfile'
                tile
                small
                light
                color = 'rgba(189, 147, 249, 1)'
                class = '
                  mb-4
                '
                width = '100%'
              >
                Create
              </v-btn>

              <!-- Encoding - Not yet ready. -->
              
              <!--v-list dense>
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
              ></v-text-field-->
            </div>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>

    <v-divider></v-divider>

    <div
      class = '
        ma-4
      '
    >
      <BuildProfiles />
    </div>

  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import BuildProfiles from '@/components/author/build-profiles.vue'

import * as Service from '@/src/services'

import EventBus from '@/src/EventBus'

export default Vue.extend({
  components: {
    BuildProfiles,
  },

  data: () => ({
    build: {
      name: '',
      description: null,
      gitUrl: 'https://github.com/zarkones/XENA.git',
      encoding: '',
      encodingIterations: 1,
    },

    encodings: [
      {
        key: 'SHIKATA_GA_NAI',
        name: 'Shikata Ga Nai'
      },
    ] as const,

    buildTemplate: '',
    buildTemplates: [
      'XENA_RA',
      'XENA_APEP',
    ],
  }),

  methods: {
    async insertBuildProfile () {
      const newBuildProfile = await Service.Pyramid.insertBuildProfile(
        this.$axios,
        this.build.name,
        this.build.description?.length ? this.build.description : null,
        this.build.gitUrl,
        this.buildTemplate,
      ).then(() => EventBus.$emit('updateBuildProfiles'))
    }
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>
</style>
