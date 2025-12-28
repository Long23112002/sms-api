<template>
  <div>
    <nuxt-link to="/" class="text-decoration-none d-flex">
      <v-avatar tile size="33" class="mt-1">
        <v-img contain :src="require('@/assets/img/logo.svg')"></v-img>
      </v-avatar>
      <h3 class="text-h4 text--primary ml-1">httpSMS</h3>
    </nuxt-link>
    <p>
      httpSMS là một ứng dụng
      <a
        class="text-decoration-none"
        href="https://github.com/NdoleStudio/httpsms"
        >mã nguồn mở</a
      >
      chuyển đổi điện thoại Android của bạn thành cổng SMS để bạn
      có thể gửi và nhận tin nhắn SMS bằng HTTP API đơn giản.
    </p>
    <v-btn :href="$store.getters.getAppData.documentationUrl">
      <v-icon left>{{ mdiBookOpenVariant }}</v-icon>
      Tài Liệu
    </v-btn>
  </div>
</template>

<script lang="ts">
import { mdiBookOpenVariant } from '@mdi/js'
import Vue from 'vue'

export default Vue.extend({
  name: 'BlogInfo',
  data() {
    return {
      mdiBookOpenVariant,
    }
  },
})
</script>
