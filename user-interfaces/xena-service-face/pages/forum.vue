<template>
  <div
    class = '
      ma-4
    '
  >
    <v-expansion-panels
      accordion
      flat
      tile
    >
      <v-expansion-panel>
        <v-expansion-panel-header>
          Create a Post
        </v-expansion-panel-header>

        <v-expansion-panel-content>
          <v-text-field
            dense
            outlined
            v-model = 'newPost.title'
            label = 'Title'
            color = 'rgba(189, 147, 249, 1)'
          ></v-text-field>

          <v-text-field
            dense
            outlined
            v-model = 'newPost.description'
            label = 'Description (optional)'
            color = 'rgba(189, 147, 249, 1)'
          ></v-text-field>

          <v-btn
            @click = 'createPost'
            tile
            small
            outlined
            color = 'rgba(189, 147, 249, 1)'
            width = '100%'
          >
            Create
          </v-btn>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>

    <v-divider
      class = '
        mx-2
      '
    ></v-divider>

    <v-card
      flat
      tile
      v-for = '(post, index) in posts'
      :key = 'index'
    >
      <v-card-title
        @click = 'goToPost(post.id)'
        class = '
          panel-header
        '
      >
        {{ post.title }}

        <v-spacer></v-spacer>

        <v-btn
          color = 'rgba(189, 147, 249, 1)'
          text
        >
          {{ post.author.name }}
        </v-btn>
      </v-card-title>

      <v-card-text>
        {{ post.description }}
      </v-card-text>

    </v-card>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import jwt from 'jsonwebtoken'

import { Post } from '@/src/services/Xerum'

import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    posts: [] as Post[],

    newPost: {} as Post,
  }),

  computed: {
    ...mapGetters([
      'getXerumHost',
    ]),
  },

  methods: {
    goToPost (id: string) {
      this.$router.push(`/post?id=${id}`)
    },

    async fetchPosts () {
      const posts = await new Service.Xerum(this.$axios, this.getXerumHost).getPosts()
      if (!posts)
        return

      this.posts = posts
    },

    async createPost () {
      // alert(this.$store.state.privateKey)

      const name = this.$store.state.username

      const payload = (() => {
        try {
          return jwt.sign({
            name,
            title: this.newPost.title,
            description: this.newPost.description,
          }, this.$store.state.privateKey, {
            algorithm: 'RS512',
          })
        } catch (e) {
          console.error(e)
        }
      })

      const post = await this.$axios.post('http://127.0.0.1:60633/v1/posts', {
        name,
        payload,
      })
        .then(({ data }) => data as Post)
        .catch(e => console.error(e))


    }
  },

  mounted () {
    this.fetchPosts()
  },
})
</script>

<style scoped>
.panel-header {
  font-weight: bold;
  font-size: 20px;
  letter-spacing: 2px;
}
.panel-header:hover {
  cursor: pointer;
  color: #50fa7b;
  transition: 0.3s;
}
</style>