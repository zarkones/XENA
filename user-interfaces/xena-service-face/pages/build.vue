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
                label = 'Description (optional)'
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
                v-model = 'build.trustedPublicKey'
                label = 'Trusted Public Key Used For Communication Integrity'
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

              <div
                v-if = 'buildTemplate == "XENA_BOT_APEP"'
              >
                <v-text-field
                  outlined
                  dense
                  v-model = 'build.dgaSeed'
                  label = 'Domain Generation ALgorithm (DGA) Seed'
                  :placeholder = 'build.dgaSeed'
                  color = 'rgba(189, 147, 249, 1)'
                ></v-text-field>

                <v-text-field
                  outlined
                  dense
                  v-model = 'build.dgaAfterDays'
                  label = 'Domain Generation ALgorithm (DGA) Should Take Place After How Many Days With No Contact From CNC?'
                  :placeholder = 'build.dgaAfterDays'
                  color = 'rgba(189, 147, 249, 1)'
                ></v-text-field>

                <v-text-field
                  outlined
                  dense
                  v-model = 'build.maxLoopWait'
                  label = 'The Most Amount Of Time A Bot Will Go Without Reaching Out to CNC (seconds)'
                  :placeholder = 'build.maxLoopWait'
                  color = 'rgba(189, 147, 249, 1)'
                ></v-text-field>

                <v-text-field
                  outlined
                  dense
                  v-model = 'build.minLoopWait'
                  label = 'The Least Amount Of Time A Bot Will Go Without Reaching Out to CNC (seconds)'
                  :placeholder = 'build.minLoopWait'
                  color = 'rgba(189, 147, 249, 1)'
                ></v-text-field>

                <v-text-field
                  outlined
                  dense
                  v-model = 'build.gettrProfileName'
                  label = 'Username Of A Gettr Profile Which Is Going To Be Used As A Fallback Channel Via The "website" Column In The Description Of The Profile (optional)'
                  :placeholder = 'build.gettrProfileName'
                  color = 'rgba(189, 147, 249, 1)'
                ></v-text-field>
              </div>

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
      trustedPublicKey: '',
      dgaSeed: '123',
      dgaAfterDays: '7',
      maxLoopWait: '60',
      minLoopWait: '10',
      gettrProfileName: '',
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
      'XENA_BOT_MONOLITIC',
    ],
  }),

  computed: {
    ...mapGetters([
      'getPyramidHost',
      'getPyramidToken',
    ]),
  },

  methods: {
    async insertBuildProfile () {
      console.log(this.build)
      await new Service.Pyramid(this.$axios, this.getPyramidHost, this.getPyramidToken).insertBuildProfile(
        this.build.name,
        this.build.description?.length ? this.build.description : null,
        this.build.gitUrl,
        {
          template: this.buildTemplate,
          atilaHost: this.build.configHost,
          trustedPublicKey: this.build.trustedPublicKey,
          dgaSeed: this.build.dgaSeed,
          dgaAfterDays: this.build.dgaAfterDays,
          maxLoopWait: this.build.maxLoopWait,
          minLoopWait: this.build.minLoopWait,
          gettrProfileName: this.build.gettrProfileName,
        },
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
