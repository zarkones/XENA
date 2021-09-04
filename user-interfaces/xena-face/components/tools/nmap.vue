<template>
  <v-card>
    <v-card-content>
        <v-text-field
        dense
        v-model = 'shellCode'
        outlined
        label = 'Enter a command...'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'issueMessages'
      >
        Address
      </v-text-field>
    </v-card-content>
  </v-card>
</template>