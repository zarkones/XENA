<template>
  <v-card
    tile
  >
    <v-card-title
      class = '
        purple-color
      '
    >
      Cloud Build
    </v-card-title>

    <v-card-text>
      Xena provides you with cloud build functionality. Meaning that you can create build profiles, which are used
      for creation of bot clients and other software. 
    </v-card-text>

    <div
      class = '
        ml-4
        mr-4
        mb-4
      '
    >
      <v-expansion-panels
        tile
      >
        <v-expansion-panel>
          <v-expansion-panel-header>
            Create a build profile.
          </v-expansion-panel-header>

          <v-expansion-panel-content>
            <div>
              <v-text-field
                outlined
                dense
                v-model = 'build.name'
                label = 'Name'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-text-field
                outlined
                dense
                v-model = 'build.description'
                label = 'Description'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-text-field
                outlined
                dense
                v-model = 'build.configHost'
                label = 'C2 Host'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-text-field
                outlined
                dense
                v-model = 'build.configPublicKey'
                label = 'Public Key'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <v-text-field
                outlined
                dense
                v-model = 'build.gitUrl'
                label = 'Git URL'
                :placeholder = 'build.gitUrl'
                color = 'rgba(189, 147, 249, 1)'
              ></v-text-field>

              <!--v-select
                v-model = 'buildTemplate'
                :items = 'buildTemplates.map(t => t.replaceAll("_", " "))'
                label = 'Build templates'
                outlined
                dense
              ></v-select-->

              <v-btn
                @click = 'insertBuildProfile'
                tile
                small
                outlined
                color = 'rgba(189, 147, 249, 1)'
                class = '
                  mb-4
                '
                width = '100%'
                :disabled = '!buildTemplate'
              >
                Create {{ buildTemplate.split('_')[2] }}
              </v-btn>

              <BotSelection />

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

    <!--v-divider></v-divider-->

    <div
      class = '
        ml-4
        mr-4
        mb-4
      '
    >
      <BuildProfiles />
    </div>

  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import BuildProfiles from '@/components/author/build-profiles.vue'
import BotSelection from '@/components/author/bot-selection.vue'

import * as Service from '@/src/services'

import EventBus from '@/src/EventBus'

import { mapGetters } from 'vuex'

export default Vue.extend({
  components: {
    BuildProfiles,
    BotSelection,
  },

  data: () => ({
    build: {
      name: '',
      description: null,
      gitUrl: 'https://github.com/zarkones/XENA.git',
      encoding: '',
      encodingIterations: 1,
      configHost: '',
      configPublicKey: '',
    },

    encodings: [
      {
        key: 'SHIKATA_GA_NAI',
        name: 'Shikata Ga Nai'
      },
    ] as const,

    buildTemplate: '',
    buildTemplates: [
      'XENA_BOT_ANACONDA',
      'XENA_BOT_APEP',
      'XENA_BOT_VARVARA',
    ],
  }),

  computed: {
    ...mapGetters([
      'getRaHost',
    ]),
  },

  methods: {
    async insertBuildProfile () {
      await new Service.Pyramid(this.$axios, this.getPyramidHost).insertBuildProfile(
        this.build.name,
        this.build.description?.length ? this.build.description : null,
        this.build.gitUrl,
        this.buildTemplate,
      ).then(() => EventBus.$emit('updateBuildProfiles'))
    }
  },

  mounted () {
    EventBus.$on('updateBuildTemplate', (template) => {
      this.buildTemplate = template
    })
  },
})
</script>

<style lang = 'css' scoped>
</style>
