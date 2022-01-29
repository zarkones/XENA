<template>
  <v-expansion-panels
    if = 'buildProfiles.length'
    accordion
    tile
  >
    <v-expansion-panel
      v-for = '(buildProfile, buildProfileIndex) in builldProfiles'
      :key = 'buildProfileIndex'
    >
      <v-expansion-panel-header>
        <span
          class = '
            build-profile-name
          '
        >
          {{ buildProfile.name }}
        </span>
      </v-expansion-panel-header>

      <v-expansion-panel-content>
        Git: {{ buildProfile.gitUrl }}
        <br>
        Description: {{ buildProfile.description }}
        <br>
        Build template: {{ buildProfile.config.template }} 
        <br>
        Status: {{ buildProfile.status }}
        <br>

        <v-btn
          @click = 'downloadBuild(buildProfile.id)'
          tile
          small
          outlined
          color = 'rgba(189, 147, 249, 1)'
          class = '
            mt-4
          '
          width = '100%'
        >
          Download
        </v-btn>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import BuildProfiles from '@/components/author/build-profiles.vue'

import EventBus from '@/src/EventBus'

import * as Service from '@/src/services'

import { mapGetters } from 'vuex'

export default Vue.extend({
  components: {
    BuildProfiles,
  },

  computed: {
    ...mapGetters([
      'getPyramidHost',
      'getPyramidToken',
    ])
  },

  data: () => ({
    builldProfiles: [] as any[]
  }),

  methods: {
    async getBuildProfiles () {
      const builldProfiles = await new Service.Pyramid(this.$axios, this.getPyramidHost, this.getPyramidToken).getBuilldProfiles()
      if (builldProfiles)
        this.builldProfiles = builldProfiles
    },

    downloadBuild (buildProfileId: string) {
      window.open(
        `${this.getPyramidHost}/builds?buildProfileId=${buildProfileId}`,
        '_blank',
      )
    }
  },

  async mounted () {
    await this.getBuildProfiles()

    EventBus.$on('updateBuildProfiles', async () => await this.getBuildProfiles())
  },
})
</script>

<style lang = 'css' scoped>
.build-profile-name {
  font-weight: bolder;
  color: #bd93f9;
  font-size: 18px;
}
</style>
