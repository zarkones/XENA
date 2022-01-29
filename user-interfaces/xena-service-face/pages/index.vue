<template>
  <v-container
    fill-height
    fluid
  >
    <v-row
      align = 'center'
      justify = 'center'
    >
      <v-col>
        <!-- Logo -->
        <v-img
          src = 'xena-logo.png'
          contain
          max-height = '186px'
          class = '
            logo
          '
        ></v-img>

        <v-card
          max-width = '800px'
          class = '
            mx-auto
          '
          justify = 'center'
          outlined
          tile
        >
          <v-card-title>
            Please, login.
          </v-card-title>

          <v-divider></v-divider>

          <v-text-field
            dense
            outlined
            v-model = 'username'
            label = 'Username'
            color = 'rgba(189, 147, 249, 1)'
            class = '
              pt-4
              mx-4
            '
          ></v-text-field>

          <v-text-field
            dense
            outlined
            type = 'password'
            v-model = 'privateKey'
            label = 'Private Key'
            color = 'rgba(189, 147, 249, 1)'
            class = '
              mx-4
            '
          ></v-text-field>

          <v-btn
            @click = 'login'
            tile
            small
            outlined
            color = 'rgba(189, 147, 249, 1)'
            class = '
            '
            width = '100%'
          >
            Login
          </v-btn>
        </v-card>
      </v-col>
    </v-row>

    <v-snackbar
      v-model = 'auth_msg'
      :timeout = '2000'
    >
      {{ auth_msg_payload }}

      <template v-slot:action = '{ attrs }'>
        <v-btn
          color='rgba(189, 147, 249, 1)'
          text
          v-bind = 'attrs'
          @click = 'auth_msg = false'
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import { mapActions } from 'vuex'

export default Vue.extend({
  data: () => ({
    username: '',
    privateKey: '',

    privateKey_start: 0,

    auth_msg: false,
    auth_msg_payload: 'Wrong username or privateKey.',
  }),

  methods: {
    async login () {
      this.setPrivateKey(this.privateKey)
      this.setUsername(this.username)

      this.$router.push('/dashboard')
    },

    ...mapActions([
      'setPrivateKey',
      'setUsername',
    ]),
  },
})
</script>

<style scoped>
.logo {
  margin-bottom: 56px;
}
</style>