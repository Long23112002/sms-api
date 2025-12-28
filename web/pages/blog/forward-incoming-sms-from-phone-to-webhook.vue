<template>
  <v-container class="pt-8">
    <v-row class="mt-16">
      <v-col cols="12" md="9">
        <v-img
          style="border-radius: 4px"
          alt="blog post header image"
          :src="
            require('@/static/img/blog/forward-incoming-sms-from-phone-to-webhook/header.png')
          "
        ></v-img>
        <h1
          class="mt-1"
          :class="{
            'text-h2': $vuetify.breakpoint.mdAndUp,
            'text-h3': !$vuetify.breakpoint.mdAndUp,
          }"
        >
          Cách chuyển tiếp tin nhắn SMS từ điện thoại Android vào webhook của bạn
        </h1>
        <p class="subtitle-2 mt-2">
          <span class="text-uppercase blue--text">{{ postDate }}</span>
          • <span class="text-uppercase">{{ readTime }}</span>
        </p>
        <p class="text--secondary subtitle-1 mt-2">
          Bạn có thể lập trình điện thoại Android của mình để chuyển tiếp tin nhắn nhận được trên
          điện thoại đến máy chủ của bạn và kích hoạt các tự động hóa mạnh mẽ với các công cụ
          như Zapier và IFTTT. Ứng dụng mã nguồn mở
          <nuxt-link to="/" class="text-decoration-none">httpSMS</nuxt-link>
          giúp bạn thiết lập điều này một cách dễ dàng
        </p>
        <h3 class="text-h4 mt-2">Bước 1: Lấy API_KEY của bạn</h3>
        <p>
          Tạo tài khoản trên ứng dụng web httpSMS và sao chép API key của bạn
          từ trang cài đặt.
          <nuxt-link class="text-decoration-none" to="/settings"
            >https://httpsms.com/settings</nuxt-link
          >
        </p>
        <vue-glow
          color="#329ef4"
          mode="hex"
          elevation="14"
          :intensity="1.07"
          intense
        >
          <v-img
            style="border-radius: 4px"
            alt="httpsms.com settings page"
            :src="
              require('@/static/img/blog/forward-incoming-sms-from-phone-to-webhook/settings.png')
            "
          ></v-img>
        </vue-glow>
        <h3 class="text-h4 mt-12">Bước 2: Cài đặt ứng dụng Android httpSMS</h3>
        <p>
          Cài đặt ứng dụng Android
          <a
            target="_blank"
            class="text-decoration-none"
            href="https://github.com/NdoleStudio/httpsms/releases/latest/download/HttpSms.apk"
            >httpSMS</a
          >
          trên điện thoại của bạn và đăng nhập bằng API KEY mà bạn
          đã sao chép ở trên. Ứng dụng này lắng nghe tin nhắn SMS nhận được trên
          điện thoại Android của bạn. <br /><span v-if="$vuetify.breakpoint.mdAndUp"
            >👉</span
          >
          <a
            target="_blank"
            class="text-decoration-none text-break"
            href="https://github.com/NdoleStudio/httpsms/releases/latest/download/HttpSms.apk"
            >https://github.com/NdoleStudio/httpsms/releases/latest/download/HttpSms.apk</a
          >
        </p>
        <v-img
          style="border-radius: 4px"
          alt="httpsms android app"
          height="800"
          contain
          :src="
            require('@/static/img/blog/forward-incoming-sms-from-phone-to-webhook/android-app.png')
          "
        ></v-img>
        <h3 class="text-h4 mt-12">Bước 3: Thiết lập webhook</h3>
        <p>
          Sau khi ứng dụng được cài đặt, nó sẽ lắng nghe tin nhắn SMS
          nhận được trên điện thoại Android. Bạn có thể cấu hình
          ứng dụng để gửi tin nhắn SMS này đến URL máy chủ của bạn bằng
          webhook. Bạn có thể cấu hình URL này trong trang cài đặt của
          ứng dụng httpSMS
          <nuxt-link to="/settings" class="text-decoration-none"
            >https://httpsms.com/settings</nuxt-link
          >
        </p>
        <vue-glow
          color="#329ef4"
          mode="hex"
          elevation="14"
          :intensity="1.07"
          intense
        >
          <v-img
            style="border-radius: 4px"
            alt="httpSMS webhook configuration"
            :src="
              require('@/static/img/blog/forward-incoming-sms-from-phone-to-webhook/webhook.png')
            "
          ></v-img>
        </vue-glow>
        <h3 class="text-h4 mt-12">Kết Luận</h3>
        <p>
          Chúc mừng, bạn đã thiết lập thành công việc chuyển tiếp SMS từ
          điện thoại Android của bạn đến webhook! Công cụ tự động hóa mạnh mẽ này có thể giúp bạn
          tối ưu hóa quy trình làm việc của doanh nghiệp và tiết kiệm thời gian và công sức.
        </p>
        <p>
          Bạn cũng có thể kích hoạt ứng dụng httpSMS để gửi SMS bằng
          API đơn giản. Bạn có thể tìm thêm thông tin trên trang tài liệu tại
          <a
            target="_blank"
            class="text-decoration-none"
            href="https://docs.httpsms.com"
            >https://docs.httpsms.com</a
          >
        </p>
        <v-divider class="mx-16"></v-divider>
        <div class="text-center mt-8 mb-4">
          <back-button></back-button>
        </div>
      </v-col>
      <v-col v-if="$vuetify.breakpoint.mdAndUp" md="3">
        <blog-info></blog-info>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { mdiTwitter } from '@mdi/js'
export default {
  name: 'ForwardIncomingSMSFromPhoneToWebhook',
  layout: 'website',
  data() {
    return {
      mdiTwitter,
      postDate: '08 Tháng 4, 2023',
      readTime: '5 phút đọc',
    }
  },
  head() {
    return {
      title:
        'Cách chuyển tiếp tin nhắn SMS từ điện thoại Android vào webhook của bạn - httpSMS',
      meta: [
        {
          hid: 'og:title',
          property: 'og:title',
          content:
            'Cách chuyển tiếp tin nhắn SMS từ điện thoại Android vào webhook của bạn',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content:
            'Bạn có thể lập trình điện thoại Android của mình để chuyển tiếp tin nhắn nhận được trên điện thoại đến máy chủ của bạn và kích hoạt các tự động hóa mạnh mẽ với các công cụ như Zapier và IFTTT.',
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content:
            'https://httpsms.com/img/blog/forward-incoming-sms-from-phone-to-webhook/header.png',
        },
        {
          hid: 'twitter:card',
          name: 'twitter:card',
          content: 'summary_large_image',
        },
        {
          hid: 'og:url',
          property: 'og:url',
          content:
            'https://httpsms.com/blog/forward-incoming-sms-from-phone-to-webhook/',
        },
      ],
    }
  },
}
</script>
