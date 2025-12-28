<template>
  <v-container>
    <v-row class="mt-16">
      <v-col cols="12" md="9">
        <v-row>
          <v-col cols="12" class="mt-4">
            <h1 class="text-h2">Blog</h1>
            <h2 class="text--secondary mt-2 mb-n4 title font-weight-light">
              Tìm hiểu thêm về httpSMS qua blog của chúng tôi!
            </h2>
          </v-col>
        </v-row>
        <v-row>
          <v-col v-for="post in blogPosts" :key="post.route" cols="12" xl="6">
            <v-hover v-slot="{ hover }">
              <vue-glow
                color="#329ef4"
                mode="hex"
                elevation="7"
                :intensity="hover ? 0.7 : 0"
                intense
              >
                <v-card :to="post.route">
                  <v-card-title class="text-h4 text-break">{{
                    post.title
                  }}</v-card-title>
                  <v-card-subtitle>
                    <span class="text-uppercase blue--text">{{
                      post.date
                    }}</span>
                    • <span class="text-uppercase">{{ post.readTime }}</span>
                  </v-card-subtitle>
                  <v-card-text class="mt-n2">
                    <p class="text--primary subtitle-1">
                      {{ post.description }}
                    </p>
                  </v-card-text>
                </v-card>
              </vue-glow>
            </v-hover>
          </v-col>
        </v-row>
      </v-col>
      <v-col v-if="$vuetify.breakpoint.mdAndUp" md="3">
        <blog-info></blog-info>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { mdiGithub, mdiTwitter } from '@mdi/js'

export default {
  name: 'BlogIndex',
  layout: 'website',
  data() {
    return {
      mdiTwitter,
      mdiGithub,
      blogPosts: [
        {
          route: '/blog/grant-send-and-read-sms-permissions-on-android',
          title: 'Cách cấp quyền SMS trên Android 15+',
          date: '18 Tháng 2, 2025',
          readTime: '5 phút đọc',
          description:
            'Kiểm soát quyền riêng tư của bạn bằng cách mã hóa tin nhắn SMS end-to-end. Bảo vệ tin nhắn của bạn khỏi những con mắt tò mò, đảm bảo tính bảo mật tuyệt đối.',
        },
        {
          route: '/blog/end-to-end-encryption-to-sms-messages',
          title:
            'Bảo mật cuộc trò chuyện của bạn với mã hóa end-to-end cho tin nhắn SMS',
          date: '21 Tháng 1, 2024',
          readTime: '10 phút đọc',
          description:
            'Kiểm soát quyền riêng tư của bạn bằng cách mã hóa tin nhắn SMS end-to-end. Bảo vệ tin nhắn của bạn khỏi những con mắt tò mò, đảm bảo tính bảo mật tuyệt đối.',
        },
        {
          route:
            '/blog/send-sms-when-new-row-is-added-to-google-sheets-using-zapier',
          title:
            'Gửi SMS khi một hàng mới được thêm vào Google Sheets bằng Zapier',
          date: '29 Tháng 11, 2023',
          readTime: '7 phút đọc',
          description:
            'Tự động hóa việc gửi tin nhắn SMS cá nhân hóa mỗi khi một hàng mới được thêm vào tài liệu Google Sheets của bạn bằng Zapier mà không cần viết code.',
        },
        {
          route: '/blog/how-to-send-sms-messages-from-excel',
          title:
            'Cách gửi tin nhắn SMS đến nhiều số điện thoại từ Excel',
          date: '29 Tháng 10, 2023',
          readTime: '5 phút đọc',
          description:
            'Gửi nhiều tin nhắn SMS đến khách hàng của bạn từ file Excel bằng cách sử dụng điện thoại Android của bạn làm cổng SMS.',
        },
        {
          route: '/blog/send-bulk-sms-from-csv-file-with-no-code',
          title: 'Gửi tin nhắn SMS hàng loạt từ file CSV mà không cần code',
          date: '29 Tháng 10, 2023',
          readTime: '7 phút đọc',
          description:
            'Gửi tin nhắn SMS đến nhiều người nhận trong file CSV của bạn mà không cần viết một dòng code nào.',
        },
        {
          route: '/blog/send-sms-from-android-phone-with-python',
          title: 'Gửi SMS từ điện thoại Android của bạn bằng Python',
          date: '03 Tháng 6, 2023',
          readTime: '6 phút đọc',
          description:
            'Bài viết này sẽ hướng dẫn bạn cách cấu hình điện thoại Android của bạn làm cổng SMS để tự động hóa việc gửi tin nhắn văn bản bằng ngôn ngữ lập trình Python.',
        },
        {
          route: '/blog/forward-incoming-sms-from-phone-to-webhook',
          title:
            'Chuyển tiếp tin nhắn SMS từ điện thoại Android vào webhook của bạn',
          date: '08 Tháng 4, 2023',
          readTime: '5 phút đọc',
          description:
            'Lập trình điện thoại Android của bạn để chuyển tiếp tin nhắn nhận được trên điện thoại đến máy chủ của bạn và kích hoạt các tự động hóa mạnh mẽ với các công cụ như Zapier và IFTTT.',
        },
      ],
    }
  },
  head() {
    return {
      title: 'Bài Viết Blog - httpSMS',
    }
  },
}
</script>
