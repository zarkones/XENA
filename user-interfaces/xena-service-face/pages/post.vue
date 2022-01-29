<template>
  <v-card
    v-if = 'post.id'
  >
    <v-card-title>
      {{ post.title }}

      <v-spacer></v-spacer>

      <v-btn
        class = '
          ml-4
        '
        color = 'rgba(189, 147, 249, 1)'
        text
      >
        {{ post.author.name }}
      </v-btn>
    </v-card-title>
    
    <v-card-text>
      {{ post.description }}
    </v-card-text>

    <v-card-actions>
      <v-text-field
        dense
        outlined
        v-model = 'replyMessage'
        label = 'Reply'
        color = 'rgba(189, 147, 249, 1)'
      ></v-text-field>
    </v-card-actions>

    <!--v-divider
      class = '
        mx-2
      '
    ></v-divider-->
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { Post } from '@/src/services/Xerum'
import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    post: {} as Post,
  }),

  computed: {
    ...mapGetters([
      'getXerumHost',
    ]),
  },

  methods: {
    async fetchPost () {
      const post = await new Service.Xerum(this.$axios, this.getXerumHost).getPost(this.$route.query.id)

      if (!post)
        return

      this.post = post
    }
  },

  mounted () {
    this.fetchPost()
  }
})
</script>