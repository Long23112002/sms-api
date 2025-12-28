<template>
  <v-container class="pt-8">
    <v-row class="mt-16">
      <v-col cols="12" md="9">
        <v-img
          style="border-radius: 4px"
          alt="blog post header image"
          :src="
            require('@/static/img/blog/send-sms-from-android-phone-with-python/header.png')
          "
        ></v-img>
        <h1
          class="mt-1"
          :class="{
            'text-h2': $vuetify.breakpoint.mdAndUp,
            'text-h3': !$vuetify.breakpoint.mdAndUp,
          }"
        >
          Gửi SMS từ điện thoại Android của bạn bằng Python
        </h1>
        <p class="subtitle-2 mt-2">
          <span class="text-uppercase blue--text">{{ postDate }}</span>
          • <span class="text-uppercase">{{ readTime }}</span>
        </p>
        <p class="text--secondary subtitle-1 mt-2">
          Trong thời đại bị thống trị bởi mạng xã hội, ứng dụng nhắn tin tức thời và
          các công nghệ truyền thông không ngừng phát triển, rất dễ bỏ qua
          dịch vụ tin nhắn ngắn (SMS) khiêm tốn nhưng đáng chú ý. Từ khi
          ra đời vào những năm 1990, SMS đã vượt qua thử thách của thời gian, vẫn là một trong
          những phương tiện truyền thông di động được sử dụng rộng rãi và đáng tin cậy nhất.
        </p>
        <p>
          Cho dù bạn là chủ doanh nghiệp đang tìm cách tối ưu hóa chiến lược truyền thông
          của mình, một nhà phát triển đang tìm cách tích hợp chức năng SMS vào
          ứng dụng của mình, hoặc đơn giản là bị thu hút bởi sức hấp dẫn lâu dài của SMS, bài viết này
          sẽ giải thích cách thiết lập điện thoại Android của bạn để gửi tin nhắn
          SMS.
        </p>
        <h3 class="text-h4 mt-8 mb-2">Yêu Cầu</h3>
        <ul>
          <li>Hiểu biết cơ bản về Python.</li>
          <li>Một điện thoại Android.</li>
          <li>
            <a href="https://www.python.org/" class="text-decoration-none"
              >Python</a
            >
            đã được cài đặt trên máy tính của bạn.
          </li>
        </ul>
        <h3 class="text-h4 mt-8 mb-2">Bước 1: Lấy API Key của bạn</h3>
        <p>
          Create an account on
          <nuxt-link class="text-decoration-none" to="/">httpsms.com</nuxt-link>
          and copy your API key from the settings page
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
        <h3 class="text-h4 mb-4 mt-16">
          Bước 2: Cài đặt ứng dụng Android httpSMS
        </h3>
        <p>
          <a
            target="_blank"
            class="text-decoration-none"
            href="https://github.com/NdoleStudio/httpsms/releases/latest/download/HttpSms.apk"
            >⬇️ Tải xuống và cài đặt</a
          >
          ứng dụng Android httpSMS trên điện thoại của bạn và đăng nhập bằng API KEY
          mà bạn đã sao chép ở trên. Ứng dụng này lắng nghe tin nhắn SMS nhận được trên
          điện thoại Android của bạn.
        </p>
        <v-alert type="info" outlined>
          Đảm bảo nhập số điện thoại của bạn theo định dạng quốc tế, ví dụ
          +18005550199 khi xác thực với ứng dụng Android httpSMS.
        </v-alert>
        <v-img
          style="border-radius: 4px"
          alt="httpsms android app"
          height="800"
          contain
          :src="
            require('@/static/img/blog/forward-incoming-sms-from-phone-to-webhook/android-app.png')
          "
        ></v-img>
        <h3 class="text-h4 mt-12">Bước 3: Viết code</h3>
        <p>
          Bây giờ bạn đã thiết lập điện thoại Android của mình đúng cách trên httpSMS, bạn
          có thể viết code Python bên dưới trong một file mới có tên
          <code>send_sms.py</code>. Code này sẽ gửi SMS và sau
          khi chạy script qua điện thoại Android của bạn đến số điện thoại
          người nhận được chỉ định trong <code>payload</code>.
        </p>
        <v-alert type="info" outlined class="mt-2 mb-4">
          Đảm bảo sử dụng <code>api_key</code> đúng từ bước 1 và cũng
          sử dụng số điện thoại <code>to</code> và <code>from</code> đúng trong
          biến <code>payload</code>.
        </v-alert>
        <pre v-highlight class="python w-full mb-n9">
<code>import requests
import json

api_key = "" # Get API Key from https://httpsms.com/settings

url = 'https://api.httpsms.com/v1/messages/send'

headers = {
    'x-api-key': api_key,
    'Accept': 'application/json',
    'Content-Type': 'application/json'
}

payload = {
    "content": "This is a sample text message sent via python",
    "from": "+18005550199", # This is the phone number of your android phone */
    "to": "+18005550100" # This is the recipient phone number */
}

response = requests.post(url, headers=headers, data=json.dumps(payload))

print(json.dumps(response.json(), indent=4))
</code>
        </pre>
        <p>
          Chạy code trên bằng lệnh
          <code>python send_sms.py</code> và kiểm tra điện thoại được chỉ định trong
          trường <code>to</code> của <code>payload</code> để xác minh rằng
          tin nhắn đã được nhận thành công.
        </p>
        <v-img
          style="border-radius: 4px"
          alt="httpsms android app"
          height="800"
          contain
          :src="
            require('@/static/img/blog/send-sms-from-android-phone-with-python/sms-sent.png')
          "
        ></v-img>
        <h3 class="text-h4 mt-12">Kết Luận</h3>
        <p>
          Chúc mừng, bạn đã cấu hình thành công điện thoại Android của mình
          để gửi tin nhắn SMS qua Python. Bây giờ bạn có thể tái sử dụng code này để gửi
          tin nhắn SMS từ ứng dụng Python của mình.
        </p>
        <p>
          Nếu bạn cũng quan tâm đến việc chuyển tiếp SMS đến từ
          điện thoại Android của bạn đến máy chủ của bạn, hãy xem
          <nuxt-link
            to="/blog/forward-incoming-sms-from-phone-to-webhook"
            class="text-decoration-none"
            >hướng dẫn chuyển tiếp SMS.</nuxt-link
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
  name: 'SendSmsFromAndroidPhoneWithPython',
  layout: 'website',
  data() {
    return {
      mdiTwitter,
      postDate: '03 Tháng 6, 2023',
      readTime: '6 phút đọc',
    }
  },
  head() {
    return {
      title: 'Gửi SMS từ điện thoại Android của bạn bằng Python - httpSMS',
      meta: [
        {
          hid: 'og:title',
          property: 'og:title',
          content: 'Send an SMS from your Android phone with Python',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content:
            'Configure your Android phone as an SMS gateway to automate sending text messages with the Python programing language.',
        },
        {
          hid: 'og:image',
          property: 'og:image',
          content:
            'https://httpsms.com/img/blog/send-sms-from-android-phone-with-python/header.png',
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
            'https://httpsms.com/blog/send-sms-from-android-phone-with-python/',
        },
      ],
    }
  },
}
</script>
