<template>
  <v-container fluid fill-height>
    <v-row align="center" justify="center">
      <v-col
        cols="12"
        md="4"
        xl="3"
        :class="{ 'mt-n16': $vuetify.breakpoint.lgAndUp }"
      >
        <div class="text-center mb-5">
          <v-avatar tile size="45" class="mt-n8 mr-4">
            <v-img contain :src="require('@/assets/img/logo.svg')"></v-img>
          </v-avatar>
          <span class="text-h3">Chào Mừng</span>
        </div>
        <p class="subtitle-1 text-center text--secondary mt-1 mb-4">
          Tham gia cùng <b>13,195+</b> người dùng đã gửi hoặc
          <br />
          nhận hơn <b>5,263,593+</b> tin nhắn SMS
        </p>
        <v-card max-width="360" class="mx-auto">
          <v-card-text class="px-0">
            <client-only>
              <custom-auth :to="to"></custom-auth>
            </client-only>
          </v-card-text>
        </v-card>
        <div class="text-center mt-4">
          <back-button></back-button>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'LoginIndex',
  middleware: ['guest'],
  data() {
    return {
      to: '/threads',
    }
  },
  head() {
    return {
      title: 'Đăng Nhập Tài Khoản - Http SMS',
    }
  },
  mounted() {
    if (this.$route.query.to) {
      this.to = this.$route.query.to as string
    }
  },
})
</script>
