<template>
  <v-container class="pt-8">
    <v-row class="mt-16">
      <v-col cols="12" md="9">
        <h1
          class="mt-1"
          :class="{
            'text-h2': $vuetify.breakpoint.mdAndUp,
            'text-h3': !$vuetify.breakpoint.mdAndUp,
          }"
        >
          Bảo mật cuộc trò chuyện của bạn bằng cách mã hóa tin nhắn SMS end-to-end
        </h1>
        <p class="subtitle-2 mt-2">
          <span class="text-uppercase blue--text">{{ postDate }}</span>
          • <span class="text-uppercase">{{ readTime }}</span>
        </p>
        <p class="text--secondary subtitle-1 mt-2">
          Chúng tôi đã thêm hỗ trợ mã hóa end-to-end cho tin nhắn SMS để
          không ai có thể xem nội dung tin nhắn bạn gửi bằng httpSMS
          ngoài bạn.
        </p>
        <p>
          Bạn thiết lập một khóa mã hóa mà bạn sử dụng để mã hóa tin nhắn của mình
          trước khi thực hiện yêu cầu API đến httpSMS và bạn cũng sử dụng cùng khóa
          để giải mã các tin nhắn bạn nhận từ httpSMS thông qua
          <a
            class="text-decoration-none"
            href="https://docs.httpsms.com/webhooks/introduction"
            >sự kiện webhook</a
          > của chúng tôi. Chúng tôi đang sử dụng thuật toán mã hóa
          <a
            class="text-decoration-none"
            href="https://en.wikipedia.org/wiki/Advanced_Encryption_Standard"
            >AES 256</a
          >
          để mã hóa và giải mã tin nhắn.
        </p>
        <h3 class="text-h4 mt-8 mb-2">Thiết lập khóa mã hóa của bạn</h3>
        <p>
          <a
            target="_blank"
            class="text-decoration-none"
            href="https://github.com/NdoleStudio/httpsms/releases/latest/download/HttpSms.apk"
            >⬇️ Tải xuống và cài đặt</a
          >
          ứng dụng Android httpSMS trên điện thoại của bạn và đặt khóa mã hóa trong
          trang <b>Cài Đặt Ứng Dụng</b> của ứng dụng.
        </p>
        <v-img
          style="border-radius: 4px"
          alt="httpsms android app"
          height="800"
          contain
          :src="
            require('@/static/img/blog/end-to-end-encryption-to-sms-messages/encryption-key-android.png')
          "
        ></v-img>
        <h3 class="text-h4 mb-4 mt-16">Mã hóa tin nhắn SMS của bạn</h3>
        <p>
          Chúng tôi sử dụng thuật toán mã hóa AES-256 để mã hóa tin nhắn SMS.
          Thuật toán này yêu cầu một khóa mã hóa 256 bit. Để giải quyết điều này,
          chúng tôi sẽ hash bất kỳ khóa mã hóa nào bạn đặt trên ứng dụng di động
          bằng thuật toán sha-256 để nó luôn tạo ra một khóa
          256 bit.
        </p>
        <p>
          Thuật toán AES cũng có tham số initialization vector (IV)
          được sử dụng để đảm bảo rằng cùng một giá trị được mã hóa nhiều lần
          sẽ không tạo ra cùng một giá trị mã hóa. IV là 16 bit và nó được
          thêm vào tin nhắn đã mã hóa trước khi mã hóa nó bằng base64.
        </p>
        <p>
          Khi bạn sử dụng các thư viện client của chúng tôi, nó sẽ tự động xử lý
          việc mã hóa tin nhắn của bạn để bạn không phải lo lắng về việc tạo
          initialization vector và mã hóa payload của riêng bạn.
        </p>
        <v-tabs v-model="selectedTab" show-arrows>
          <v-tab href="#javascript">
            <v-icon color="#efd81d" class="mr-1">{{
              mdiLanguageJavascript
            }}</v-icon>
            Javascript
          </v-tab>
          <v-tab href="#go">
            <v-icon color="#00aed8" class="mr-1">{{ mdiLanguageGo }}</v-icon>
            Go
          </v-tab>
        </v-tabs>
        <v-tabs-items v-model="selectedTab">
          <v-tab-item value="javascript">
            <pre v-highlight class="javascript w-full mb-n12">
<code>import HttpSms from "httpsms"

const client = new HttpSms("" /* API Key from https://httpsms.com/settings */);

const key = "Password123";

const encryptedMessage = client.cipher.encrypt(key, "This is a sample text message");

// The encrypted message looks like this, note that you will get a different encrypted message when you run this code on your computer
// Qk3XGN5+Ax38Ig01m4AqaP6Y0b0wYpCXtx59sU23uVLWUU/c7axF7LozDg==
</code>
            </pre>
          </v-tab-item>
          <v-tab-item value="go">
            <pre v-highlight class="go w-full mb-n12">
<code>import "github.com/NdoleStudio/httpsms-go"

client := htpsms.New(htpsms.WithAPIKey(""/* API Key from https://httpsms.com/settings */))

key := "Password123" // use the same key on the Android app
encryptedMessage := client.Cipher.Encrypt(key, "This is a test text message")

// The encrypted message looks like this, note that you will get a different encrypted message when you run this code on your computer
// Qk3XGN5+Ax38Ig01m4AqaP6Y0b0wYpCXtx59sU23uVLWUU/c7axF7LozDg==
</code>
        </pre>
          </v-tab-item>
        </v-tabs-items>
        <h3 class="text-h4 mt-6">Send an encrypted message</h3>
        <p>
          After generating the encrypted message payload, you can send it
          directly using the httpSMS API. Make sure to set
          <code>encrypted: true</code> in the JSON request payload so that
          httpSMS knows that the message is encrypted and it will be decoded in
          the Android app before sending to your recipient.
        </p>
        <v-tabs v-model="selectedTab" show-arrows>
          <v-tab href="#javascript">
            <v-icon color="#efd81d" class="mr-1">{{
              mdiLanguageJavascript
            }}</v-icon>
            Javascript
          </v-tab>
          <v-tab href="#go">
            <v-icon color="#00aed8" class="mr-1">{{ mdiLanguageGo }}</v-icon>
            Go
          </v-tab>
        </v-tabs>
        <v-tabs-items v-model="selectedTab">
          <v-tab-item value="javascript">
            <pre v-highlight class="javascript w-full mb-n12">
<code>import HttpSms from "httpsms"

client.messages.postSend({
    content:   encryptedMessage,
    from:      '+18005550199',
    encrypted: true,
    to:        '+18005550100',
})
.then((message) => {
    console.log(message.id); // log the ID of the sent message
});</code>
            </pre>
          </v-tab-item>
          <v-tab-item value="go">
            <pre v-highlight class="go w-full mb-n12">
<code>import "github.com/NdoleStudio/httpsms-go"

client.Messages.Send(context.Background(), &httpsms.MessageSendParams{
    Content:   encryptedMessage,
    From:      "+18005550199",
    To:        "+18005550100",
    Encrypted: true,
})
</code>
        </pre>
          </v-tab-item>
        </v-tabs-items>
        <p class="mt-4">
          When you make the API request, the message will be decrypted before
          sending to the recipient. This is a screenshot of the SMS message
          which is sent to the recipient.
        </p>
        <v-img
          style="border-radius: 4px"
          alt="httpsms android app"
          height="800"
          contain
          :src="
            require('@/static/img/blog/end-to-end-encryption-to-sms-messages/send-sms-message.png')
          "
        ></v-img>
        <h3 class="text-h4 mb-4 mt-16">Receiving an encrypted message</h3>
        <p>
          When your android phone receives a new message, it will be encrypted
          with the encryption Key on your Android phone before it is delivered
          to your server's webhook endpoint. You can configure webhooks by
          following
          <a
            href="https://httpsms.com/blog/forward-incoming-sms-from-phone-to-webhook"
            >this guide.</a
          >
        </p>
        <v-tabs v-model="selectedTab" show-arrows>
          <v-tab href="#javascript">
            <v-icon color="#efd81d" class="mr-1">{{
              mdiLanguageJavascript
            }}</v-icon>
            Javascript
          </v-tab>
          <v-tab href="#go">
            <v-icon color="#00aed8" class="mr-1">{{ mdiLanguageGo }}</v-icon>
            Go
          </v-tab>
        </v-tabs>
        <v-tabs-items v-model="selectedTab">
          <v-tab-item value="javascript">
            <pre v-highlight class="javascript w-full mb-n12">
<code>import HttpSms from "httpsms"

const client = new HttpSms("" /* API Key from https://httpsms.com/settings */);

// The payload in the webhook HTTP request looks like this
/*
{
  "specversion": "1.0",
  "id": "8dca3b0a-446a-4a5d-8d2a-95314926c4ed",
  "source": "/v1/messages/receive",
  "type": "message.phone.received",
  "datacontenttype": "application/json",
  "time": "2024-01-21T12:27:29.1605708Z",
  "data": {
    "message_id": "0681b838-4157-44bb-a4ea-721e40ee7ca7",
    "user_id": "XtABz6zdeFMoBLoltz6SREDvRSh2",
    "owner": "+37253920216",
    "encrypted": true,
    "contact": "+37253920216",
    "timestamp": "2024-01-21T12:27:17.949Z",
    "content": "bdmZ7n6JVf/ST+SoNlSaOGUL1DcL5705ETw8GAB4llYBgE9HOOL+Pu/h+w==",
    "sim": "SIM1"
  }
}
*/

const encryptedMessage = "bdmZ7n6JVf/ST+SoNlSaOGUL1DcL5705ETw8GAB4llYBgE9HOOL+Pu/h+w==" // get the encrypted message from the request payload
const encryptionkey = "Password123" // use the same key on the Android app
const decryptedMessage = client.cipher.decrypt(encryptionkey, encryptedMessage)

// This is a test text message
</code>
        </pre>
          </v-tab-item>
          <v-tab-item value="go">
            <pre v-highlight class="go w-full mb-n12">
<code>import "github.com/NdoleStudio/httpsms-go"

client := htpsms.New(htpsms.WithAPIKey(/* API Key from https://httpsms.com/settings */))

// The payload in the webhook HTTP request looks like this
/*
{
  "specversion": "1.0",
  "id": "8dca3b0a-446a-4a5d-8d2a-95314926c4ed",
  "source": "/v1/messages/receive",
  "type": "message.phone.received",
  "datacontenttype": "application/json",
  "time": "2024-01-21T12:27:29.1605708Z",
  "data": {
    "message_id": "0681b838-4157-44bb-a4ea-721e40ee7ca7",
    "user_id": "XtABz6zdeFMoBLoltz6SREDvRSh2",
    "owner": "+37253920216",
    "encrypted": true,
    "contact": "+37253920216",
    "timestamp": "2024-01-21T12:27:17.949Z",
    "content": "bdmZ7n6JVf/ST+SoNlSaOGUL1DcL5705ETw8GAB4llYBgE9HOOL+Pu/h+w==",
    "sim": "SIM1"
  }
}
*/

encryptedMessage = "bdmZ7n6JVf/ST+SoNlSaOGUL1DcL5705ETw8GAB4llYBgE9HOOL+Pu/h+w==" // get the encrypted message from the request payload
encryptionkey := "Password123" // use the same key on the Android app
decryptedMessage := client.Cipher.Decrypt(encryptionkey, encryptedMessage)

// This is a test text message
</code>
        </pre>
          </v-tab-item>
        </v-tabs-items>
        <h3 class="text-h4 mt-12">Kết Luận</h3>
        <p>
          Chúc mừng, bạn đã cấu hình thành công điện thoại Android của mình
          để gửi và nhận tin nhắn SMS với mã hóa end-to-end. Đừng
          ngần ngại liên hệ với chúng tôi nếu bạn gặp bất kỳ vấn đề nào khi làm theo
          hướng dẫn này.
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
import { mdiLanguageGo, mdiTwitter, mdiLanguageJavascript } from '@mdi/js'
export default {
  name: 'EndToEndEncryptionToSmsMessages',
  layout: 'website',
  data() {
    return {
      mdiTwitter,
      mdiLanguageGo,
      mdiLanguageJavascript,
      selectedTab: 'javascript',
      postDate: '21 Tháng 1, 2024',
      readTime: '10 phút đọc',
    }
  },
  head() {
    return {
      title:
        'Bảo mật cuộc trò chuyện của bạn với mã hóa end-to-end cho tin nhắn SMS - httpSMS',
      meta: [
        {
          hid: 'og:title',
          property: 'og:title',
          content:
            'Secure your conversations with end-to-end encryption for SMS messages',
        },
        {
          hid: 'og:description',
          property: 'og:description',
          content:
            'We have added support for end-to-end encryption for SMS messages so that no one can see the content of the messages you send using httpSMS except you.',
        },
      ],
    }
  },
}
</script>
