<template>
  <div>
    <v-container>
      <v-row :class="{ 'py-16': $vuetify.breakpoint.lgAndUp }">
        <v-col
          cols="12"
          md="6"
          :class="{
            'text-center pt-16': $vuetify.breakpoint.mdAndDown,
            'py-16': $vuetify.breakpoint.lgAndUp,
          }"
        >
          <h1
            class="text-h2 font-weight-bold pb-1 gradient-header"
            :class="{
              'mt-16': $vuetify.breakpoint.lgAndUp,
              'mt-10': $vuetify.breakpoint.mdOnly,
              'mt-8': $vuetify.breakpoint.smAndDown,
            }"
          >
            Chuyển đổi điện thoại Android của bạn thành cổng SMS.
          </h1>
          <h2 class="text--secondary text-h5 mt-8 mb-8">
            <span class="gradient-underline">Tiết kiệm tiền</span> bằng cách sử dụng
            điện thoại của bạn để gửi và nhận tin nhắn SMS thông qua một API lập trình đơn giản
            với mã hóa end-to-end.
          </h2>
          <div :class="{ 'text-center': $vuetify.breakpoint.mdAndDown }">
            <v-btn
              color="primary"
              large
              exact-path
              class="mt-4 mb-4"
              :to="{ name: 'login' }"
            >
              <v-icon v-if="$vuetify.breakpoint.lgAndUp" left>{{
                mdiSend
              }}</v-icon>
              Bắt Đầu
            </v-btn>
            <v-btn
              large
              color="secondary"
              class="mt-4 mb-4 ml-4"
              href="https://sandbox.httpsms.com"
            >
              <v-icon v-if="$vuetify.breakpoint.lgAndUp" left color="#ffe500">
                {{ mdiCreation }}
              </v-icon>
              Demo Trực Tiếp
            </v-btn>
          </div>
        
          <div v-if="$vuetify.breakpoint.xl" class="mt-4">
            <a href="https://www.uneed.best/tool/httpsmscom">
              <img
                src="https://www.uneed.best/POTD1A.png"
                style="width: 250px"
                alt="Uneed POTD1 Badge"
              />
            </a>
          </div>
          <v-divider
            v-if="$vuetify.breakpoint.mdAndDown"
            class="mt-6 mr-16 success"
            :class="{ 'ml-16': $vuetify.breakpoint.mdAndDown }"
          ></v-divider>
        </v-col>
        <v-col
          v-if="$vuetify.breakpoint.mdAndUp"
          cols="12"
          md="6"
          class="py-16"
        >
          <v-img
            class="mx-auto mb-4"
            max-height="700"
            max-width="98%"
            contain
            :src="require('assets/img/writing-code-phone.svg')"
          ></v-img>
        </v-col>
      </v-row>
    </v-container>
    <v-sheet class="py-16">
      <v-container>
        <v-row class="mb-16">
          <v-col cols="12" md="6" class="d-flex align-center" :order-lg="2">
            <div>
              <h3 class="text-h3 mb-1">
                SMS Hàng Loạt
                <v-chip class="ma-2" color="pink" label text-color="white">
                  <v-icon left>{{ mdiLabel }}</v-icon>
                  Không cần code
                </v-chip>
              </h3>
              <h5 class="text-h6 font-weight-light">
                Điền vào
                <a
                  class="text-decoration-none"
                  download
                  href="/templates/httpsms-bulk.csv"
                  >mẫu CSV</a
                >
                hoặc
                <a
                  class="text-decoration-none"
                  download
                  href="/templates/httpsms-bulk.xlsx"
                  >mẫu Excel</a
                >
                của chúng tôi và tải lên httpSMS để gửi tin nhắn SMS cho tối đa 1,000
                người nhận cùng lúc mà không cần viết code.
              </h5>
              <v-btn
                to="/blog/how-to-send-sms-messages-from-excel"
                class="primary mt-4"
              >
                <v-icon left>{{ mdiMicrosoftExcel }}</v-icon>
                Hướng Dẫn Tích Hợp
              </v-btn>
            </div>
          </v-col>
          <v-col cols="12" md="6" :order-lg="1">
            <v-img
              class="mb-4"
              max-height="400"
              contain
              :src="require('assets/img/bulk-sms-template.png')"
            ></v-img>
          </v-col>
        </v-row>
        <v-row class="mb-16 mt-16">
          <v-col cols="12" md="6" class="d-flex align-center" :order-lg="2">
            <div>
              <h3 class="text-h3 mb-1">Webhooks</h3>
              <h5 class="text-h6 font-weight-light">
                Nếu bạn muốn xây dựng tích hợp nâng cao, chúng tôi hỗ trợ callback
                URLs. Nền tảng httpSMS có thể chuyển tiếp tin nhắn SMS nhận được trên
                điện thoại Android của bạn đến máy chủ của bạn bằng callback URL mà bạn
                cung cấp.
              </h5>
              <v-btn
                target="_blank"
                href="https://docs.httpsms.com/webhooks/introduction"
                class="primary mt-4"
              >
                <v-icon left>{{ mdiWebhook }}</v-icon>
                Tài Liệu
              </v-btn>
            </div>
          </v-col>
          <v-col cols="12" md="6" :order-lg="1">
            <v-img
              class="mb-4"
              max-height="300"
              contain
              :src="require('assets/img/connection.svg')"
            ></v-img>
          </v-col>
        </v-row>
        <v-row class="mb-16 mt-16">
          <v-col cols="12" md="6" class="d-flex align-center" :order-lg="1">
            <div>
              <h3 class="text-h3 mb-1">Kiểm Soát Gửi</h3>
              <h5 class="text-h6 font-weight-light">
                Gửi tin nhắn SMS mà không vượt quá giới hạn của nhà mạng di động
                của bạn. Nếu bạn đặt tốc độ ví dụ 3 tin nhắn mỗi phút, chúng tôi
                sẽ xếp hàng tin nhắn của bạn và gửi chúng với tốc độ 1 tin nhắn
                mỗi 20 giây.
              </h5>
              <v-btn
                href="https://docs.httpsms.com/features/control-sms-send-rate"
                class="primary"
              >
                <v-icon left>{{ mdiArrowRightThin }}</v-icon
                >Tài Liệu
              </v-btn>
            </div>
          </v-col>
          <v-col cols="12" md="6" :order-lg="2">
            <v-img
              class="mb-4"
              max-height="300"
              contain
              :src="require('assets/img/queue.svg')"
            ></v-img>
          </v-col>
        </v-row>
        <v-row class="mb-16 mt-16">
          <v-col cols="12" md="6" class="d-flex align-center" :order-lg="2">
            <div>
              <h3 class="text-h3 mb-1">Giám Sát</h3>
              <h5 class="text-h6 font-weight-light">
                Nếu điện thoại Android của bạn bị ngắt kết nối vì lý do nào đó và không thể
                gửi tin nhắn SMS, chúng tôi sẽ gửi thông báo cho bạn ngay lập tức.
              </h5>
            </div>
          </v-col>
          <v-col cols="12" md="6" :order-lg="1">
            <v-img
              class="mb-4"
              max-height="300"
              contain
              :src="require('assets/img/alert.svg')"
            ></v-img>
          </v-col>
        </v-row>
        <v-row class="mt-16">
          <v-col cols="12" md="6" class="d-flex align-center" :order-lg="2">
            <div>
              <h3 class="text-h3 mb-1">Mã Hóa 🔐</h3>
              <h5 class="text-h6 mb-3 font-weight-light">
                Kiểm soát quyền riêng tư của bạn với tính năng SMS mã hóa end-to-end
                của chúng tôi. Bảo vệ tin nhắn của bạn khỏi những con mắt tò mò, đảm bảo
                tính bảo mật tuyệt đối bằng thuật toán mã hóa cấp quân sự
                <a
                  class="text-decoration-none"
                  href="https://en.wikipedia.org/wiki/Advanced_Encryption_Standard"
                  >AES-256</a
                >.
              </h5>
              <v-btn
                to="/blog/end-to-end-encryption-to-sms-messages"
                class="primary"
                >Thiết lập mã hóa end-to-end</v-btn
              >
            </div>
          </v-col>
          <v-col cols="12" md="6" :order-lg="1">
            <v-img
              class="mb-4"
              max-height="300"
              contain
              :src="require('assets/img/mobile-encryption.svg')"
            ></v-img>
          </v-col>
        </v-row>
        <v-row class="mt-16">
          <v-col cols="12" md="6" class="d-flex align-center">
            <div>
              <h3 class="text-h3 mb-1">Nhiều Điện Thoại</h3>
              <h5 class="text-h6 mb-3 font-weight-light">
                Thiết lập ứng dụng gateway Android httpSMS trên nhiều điện thoại
                độc lập và an toàn mà không chia sẻ dữ liệu dưới một
                tài khoản bằng cách tạo phone API keys duy nhất.
              </h5>
              <v-btn
                href="https://docs.httpsms.com/features/phone-api-keys"
                class="primary"
              >
                <v-icon left>{{ mdiCellphoneKey }}</v-icon
                >Tài Liệu
              </v-btn>
            </div>
          </v-col>
          <v-col cols="12" md="6">
            <v-img
              class="mb-4"
              max-height="300"
              contain
              :src="require('assets/img/manage-phones.svg')"
            ></v-img>
          </v-col>
        </v-row>
        <v-row class="mt-16">
          <v-col cols="12" md="6" class="d-flex align-center" :order-lg="2">
            <div>
              <h3 class="text-h3 mb-1">Lên Lịch Tin Nhắn</h3>
              <h5 class="text-h6 mb-3 font-weight-light">
                Kiểm soát thời điểm SMS của bạn sẽ đến người nhận, cho phép bạn
                thời gian hoàn hảo cho khuyến mãi, cảnh báo quan trọng, v.v. bằng cách lên lịch
                tin nhắn của bạn trước.
              </h5>
              <v-btn
                href="https://docs.httpsms.com/features/scheduling-sms-messages"
                class="primary"
                ><v-icon left>{{ mdiClockOutline }}</v-icon
                >Tài Liệu</v-btn
              >
            </div>
          </v-col>
          <v-col cols="12" md="6" :order-lg="1">
            <v-img
              class="mb-4"
              max-height="300"
              contain
              :src="require('assets/img/schedule-messages.svg')"
            ></v-img>
          </v-col>
        </v-row>
      </v-container>
    </v-sheet>
    <v-container class="pb-16">
      <v-row class="mt-16">
        <v-col>
          <h2 class="text-h3 text-center">Bắt Đầu</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-row>
            <v-col cols="12" md="5" class="pr-4">
              <v-timeline dense class="mt-10 ml-n4">
                <v-timeline-item color="primary" :icon="mdiTallyMark1">
                  <v-card class="elevation-2">
                    <v-card-title class="text-h5"> Bước 1 </v-card-title>
                    <v-card-text class="subtitle-1">
                      <NuxtLink
                        class="font-weight-bold text-decoration-none"
                        :to="{ name: 'login' }"
                      >
                        Tạo tài khoản
                      </NuxtLink>
                      trên httpsms.com và lấy API key của bạn trên trang
                      cài đặt.
                    </v-card-text>
                  </v-card>
                </v-timeline-item>
                <v-timeline-item color="primary" :icon="mdiTallyMark2">
                  <v-card class="elevation-2">
                    <v-card-title class="text-h5"> Bước 2 </v-card-title>
                    <v-card-text class="subtitle-1">
                      <a
                        download
                        class="font-weight-bold text-decoration-none"
                        :href="$store.getters.getAppData.appDownloadUrl"
                        @click="
                          $store.dispatch('addNotification', {
                            type: 'info',
                            message: 'Đang tải ứng dụng Android httpSMS',
                          })
                        "
                        >Tải xuống</a
                      >
                      và cài đặt ứng dụng Android đồng hành trên
                      điện thoại của bạn và đăng nhập bằng API Key của bạn.
                    </v-card-text>
                  </v-card>
                </v-timeline-item>
                <v-timeline-item color="primary" :icon="mdiTallyMark3">
                  <v-card class="elevation-2">
                    <v-card-title class="text-h5"> Bước 3 </v-card-title>
                    <v-card-text class="subtitle-1">
                      Bắt đầu gửi và nhận tin nhắn SMS bằng HTTP API phong phú
                      của chúng tôi. Bạn có thể tìm tài liệu tại
                      <a
                        class="text-decoration-none"
                        :href="$store.getters.getAppData.documentationUrl"
                      >
                        {{ $store.getters.getAppData.documentationUrl }}
                      </a>
                    </v-card-text>
                  </v-card>
                </v-timeline-item>
              </v-timeline>
            </v-col>
            <v-col
              cols="12"
              md="7"
              :class="{ 'pt-16 mb-16mb-16': $vuetify.breakpoint.mdAndUp }"
            >
              <div class="w-full mt-4">
                <v-tabs v-model="selectedTab" show-arrows>
                  <v-tab href="#javascript">
                    <v-icon color="#efd81d" class="mr-1">{{
                      mdiLanguageJavascript
                    }}</v-icon>
                    Javascript
                  </v-tab>
                  <v-tab href="#php">
                    <v-icon color="#777bb3" class="mr-2">{{
                      mdiLanguagePhp
                    }}</v-icon>
                    PHP
                  </v-tab>
                  <v-tab href="#python">
                    <v-icon color="#ffffff" class="mr-2">{{
                      mdiLanguagePython
                    }}</v-icon>
                    Python
                  </v-tab>
                  <v-tab href="#go">
                    <v-icon color="#00aed8" class="mr-2">{{
                      mdiLanguageGo
                    }}</v-icon>
                    Go
                  </v-tab>
                  <v-tab href="#java">
                    <v-icon color="#0c89c7" class="mr-2">{{
                      mdiLanguageJava
                    }}</v-icon>
                    Java
                  </v-tab>
                  <v-tab href="#curl">
                    <v-icon color="primary" class="mr-2">{{
                      mdiPowershell
                    }}</v-icon>
                    cURL
                  </v-tab>
                  <v-tab href="#c-sharp">
                    <v-icon color="#68217a" class="mr-2">{{
                      mdiLanguageCsharp
                    }}</v-icon>
                    c-sharp
                  </v-tab>
                </v-tabs>
                <v-tabs-items v-model="selectedTab">
                  <v-tab-item value="javascript">
                    <pre v-highlight class="javascript w-full mt-n2 mb-n13">
<code>import HttpSms from 'httpsms'

const client = new HttpSms('' /* Get the API Key from https://httpsms.com/settings */);

client.messages.postSend({
    content:   'This is a sample text message',
    from:      '+18005550199', // Put the correct phone number here
    to:        '+18005550100', // Put the correct phone number here
})
.then((message) => {
    console.log(message.id); // log the ID of the sent message
})
</code>
                    </pre>
                  </v-tab-item>
                  <v-tab-item value="php">
                    <pre v-highlight class="php w-full mt-n2 mb-n13">
<code>&#60;?php
$apiKey = "Get API Key from https://httpsms.com/settings";

$options = array(
  'http' => array(
    'method'  => 'POST',
    'content' => json_encode( [
        'content' => 'This is a sample text message',
        'from'    => "+18005550199", // Put the correct phone number here
        'to'      => "+18005550100"  // Put the correct phone number here
    ]),
    'header'=>  "Content-Type: application/json\r\n" .
                "Accept: application/json\r\n" .
                "x-api-key: $apiKey\r\n"
    )
);

$context  = stream_context_create( $options );
$result = file_get_contents( "https://api.httpsms.com/v1/messages/send", false, $context );

echo $result;
</code>
                    </pre>
                  </v-tab-item>
                  <v-tab-item value="python">
                    <pre v-highlight class="python w-full mt-n2 mb-n13">
<code>import requests
import json

api_key = "Get API Key from https://httpsms.com/settings"

url = 'https://api.httpsms.com/v1/messages/send'

headers = {
    'x-api-key': api_key,
    'Accept': 'application/json',
    'Content-Type': 'application/json'
}

payload = {
    "content": "This is a sample text message",
    "from": "+18005550199",
    "to": "+18005550100"
}

response = requests.post(url, headers=headers, data=json.dumps(payload))

print(json.dumps(response.json(), indent=4))
</code>
                    </pre>
                  </v-tab-item>
                  <v-tab-item value="go">
                    <pre v-highlight class="go w-full mt-n2 mb-n13">
<code>import "github.com/NdoleStudio/httpsms-go"

client := htpsms.New(htpsms.WithAPIKey(/* API Key from https://httpsms.com/settings */))

client.Messages.Send(context.Background(), &httpsms.MessageSendParams{
    Content: "This is a sample text message",
    From:    "+18005550199",
    To:      "+18005550100",
})
</code>
                    </pre>
                  </v-tab-item>
                  <v-tab-item value="java">
                    <pre v-highlight class="java w-full mt-n2 mb-n13">
<code>var client = HttpClient.newHttpClient();
var apiKey = "Get API Key from https://httpsms.com/settings";

var payload = """
        {
            "content": "This is a sample text message",
            "from": "+18005550199",
            "to": "+18005550100"
        }
        """;

var request = HttpRequest.newBuilder()
        .uri(URI.create("https://api.httpsms.com/v1/messages/send"))
        .header("accept", "application/json")
        .header("Content-Type", "application/json")
        .header("x-api-key", apiKey)
        .POST(HttpRequest.BodyPublishers.ofString(payload))
        .build();

var response = client.send(request, HttpResponse.BodyHandlers.ofString());
System.out.println(response.body());
</code>
                    </pre>
                  </v-tab-item>
                  <v-tab-item value="curl">
                    <pre v-highlight class="bash w-full mt-n2 mb-n13">
<code>curl --location --request POST 'https://api.httpsms.com/v1/messages/send' \
--header 'x-api-key: Get API Key from https://httpsms.com/settings' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": "+18005550199",
    "to": "+18005550100",
    "content": "This is a sample text message"
}'
</code>
                    </pre>
                  </v-tab-item>
                  <v-tab-item value="c-sharp">
                    <pre v-highlight class="c-sharp w-full mt-n2 mb-n13">
<code>var client = new HttpClient();
client.DefaultRequestHeaders.Add("x-api-key", ""/* Get API Key from https://httpsms.com/settings */);

var response = await client.PostAsync(
    "https://api.httpsms.com/v1/messages/send",
    new StringContent(
        JsonSerializer.Serialize(new {
            from = "+18005550199",
            To = "+18005550100",
            Content = "This is a sample text message",
        }),
        Encoding.UTF8,
        "application/json"
    )
);

Console.WriteLine(await response.Content.ReadAsStringAsync());
</code>
                    </pre>
                  </v-tab-item>
                </v-tabs-items>
              </div>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-container>
    <v-sheet class="mt-16 pb-16">
      <v-container>
        <v-row>
          <v-col class="mt-16" md="6" offset-md="3">
            <h2
              id="pricing"
              class="text-center primary--text text-h2 mb-4 text-decoration-underline"
            >
              <span class="text--primary">Pricing</span>
            </h2>
            <h4 class="text-center text-h5 text--secondary">
              Hầu hết các tính năng của httpSMS hoàn toàn
              <span class="primary--text">miễn phí</span> nhưng nếu bạn muốn thêm một chút
              tính năng, bạn có thể nâng cấp lên pro
            </h4>
            <div class="d-flex justify-center mt-4 align-center">
              <p class="text-h6 mr-3 mt-3">Hàng Tháng</p>
              <v-switch v-model="yearlyPricing"></v-switch>
              <p class="text-h6 ml-3 mt-3">
                Hàng Năm
                <v-chip color="primary" small>
                  <v-avatar left>
                    <v-icon small>{{ mdiGift }}</v-icon>
                  </v-avatar>
                  2 tháng miễn phí</v-chip
                >
              </p>
            </div>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-slider
              v-model="pricing"
              :tick-labels="
                $vuetify.breakpoint.lgAndUp ? pricingLabelsFull : pricingLabels
              "
              :max="4"
              step="1"
              thumb-label="always"
              thumb-size="32"
              tick-size="4"
              tick
            >
              <template v-slot:thumb-label>
                {{ pricingLabels[pricing] }}
              </template>
            </v-slider>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" lg="4">
            <v-card elevation="4" color="#121212">
              <v-card-text>
                <h1 class="text-center text-h3 text--primary">Miễn Phí</h1>
                <p class="subtitle-1 text-center">
                  Thử gửi và nhận SMS trên các trang web sở thích và
                  thử nghiệm của bạn.
                </p>
                <p class="text-center text--primary">
                  <span class="text-h3">{{ formatVND(0) }}</span>
                </p>
                <p class="text-center mt-n3">Không cần thẻ tín dụng</p>
                <v-btn block :to="{ name: 'login' }" class="white--text" large
                  >Bắt Đầu</v-btn
                >
                <p class="mt-8 subtitle-1">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Gửi hoặc nhận tối đa <b>200</b> SMS/tháng
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>
                    {{ mdiCheckCircle }} </v-icon
                  >Thông báo offline cho điện thoại của bạn
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Chuyển tiếp tin nhắn đã nhận qua webhook
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Hỗ trợ email cơ bản
                </p>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="12" lg="4">
            <v-card elevation="4" color="black">
              <v-card-text>
                <h1 class="text-center text-h3 primary--text">Pro</h1>
                <p class="subtitle-1 text-center">
                  Gửi và nhận nhiều tin nhắn SMS hơn như một chuyên gia với các tính năng
                  nâng cao.
                </p>
                <p v-if="!yearlyPricing" class="text-center text--primary">
                  <span class="text-h3">{{ formatVND(10) }}</span>/tháng
                </p>
                <p v-else class="text-center text--primary">
                  <span class="text-h3">{{ formatVND(100) }}</span>/năm
                </p>
                <p v-if="!yearlyPricing" class="text-center mt-n3">
                  hoặc <b>{{ formatVND(100) }}</b> mỗi năm
                </p>
                <p v-else class="text-center mt-n3">
                  hoặc <b>{{ formatVND(8.33) }}</b> mỗi tháng
                </p>
                <v-btn block color="primary" :to="{ name: 'login' }" large
                  >Dùng Thử Miễn Phí</v-btn
                >
                <p class="mt-8 subtitle-1">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Gửi hoặc nhận tối đa <b>5,000</b> SMS/tháng
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>
                    {{ mdiCheckCircle }} </v-icon
                  >Thông báo offline cho điện thoại của bạn
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Chuyển tiếp tin nhắn đã nhận qua webhook
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Hỗ trợ ưu tiên
                </p>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="12" lg="4">
            <v-card color="#121212" border outlined>
              <v-card-text>
                <h1 class="text-center text-h3 text--primary">
                  Gói {{ pricingLabels[pricing] }}
                </h1>
                <p class="subtitle-1 text-center">
                  Gửi và nhận tối đa {{ planMessages }} tin nhắn SMS như một
                  người dùng chuyên nghiệp.
                </p>
                <p v-if="!yearlyPricing" class="text-center text--primary">
                  <span class="text-h3">{{ formatVND(planMonthlyPrice) }}</span
                  >/tháng
                </p>
                <p v-else class="text-center text--primary">
                  <span class="text-h3">{{ formatVND(planYearlyPrice) }}</span
                  >/năm
                </p>
                <p v-if="!yearlyPricing" class="text-center mt-n3">
                  hoặc <b>{{ formatVND(planYearlyPrice) }}</b> mỗi năm
                </p>
                <p v-else class="text-center mt-n3">
                  hoặc <b>{{ formatVND(planYearlyMonthlyPrice) }}</b> mỗi tháng
                </p>
                <v-btn block color="default" :to="{ name: 'login' }" large
                  >Dùng Thử Miễn Phí</v-btn
                >
                <p class="mt-8 subtitle-1">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Gửi hoặc nhận tối đa
                  <b>{{ pricingLabels[pricing] }}</b> SMS/tháng
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>
                    {{ mdiCheckCircle }} </v-icon
                  >Thông báo offline cho điện thoại của bạn
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Chuyển tiếp tin nhắn đã nhận qua webhook
                </p>
                <p class="subtitle-1 mt-n4">
                  <v-icon color="primary" class="mt-n1" left>{{
                    mdiCheckCircle
                  }}</v-icon
                  >Hỗ trợ ưu tiên
                </p>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-sheet>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import {
  mdiGithub,
  mdiCheckCircle,
  mdiSend,
  mdiGift,
  mdiLightbulbOn60,
  mdiForum,
  mdiCreation,
  mdiNumeric1,
  mdiSale,
  mdiLanguagePython,
  mdiNumeric2,
  mdiNumeric3,
  mdiCellphoneKey,
  mdiTallyMark1,
  mdiTallyMark3,
  mdiTallyMark2,
  mdiLabel,
  mdiLanguageJavascript,
  mdiLanguagePhp,
  mdiPlus,
  mdiMinus,
  mdiLanguageCsharp,
  mdiLanguageJava,
  mdiMicrosoftExcel,
  mdiWebhook,
  mdiClockOutline,
  mdiArrowRightThin,
  mdiPowershell,
  mdiLanguageGo,
} from '@mdi/js'

export default Vue.extend({
  name: 'LandingPage',
  layout: 'website',
  data() {
    return {
      mdiMicrosoftExcel,
      mdiWebhook,
      mdiGithub,
      mdiLabel,
      mdiLightbulbOn60,
      mdiCheckCircle,
      mdiSend,
      mdiGift,
      mdiArrowRightThin,
      mdiClockOutline,
      mdiCreation,
      mdiForum,
      mdiNumeric1,
      mdiNumeric2,
      mdiNumeric3,
      mdiTallyMark1,
      mdiSale,
      mdiTallyMark2,
      mdiPlus,
      mdiMinus,
      mdiTallyMark3,
      mdiLanguageJavascript,
      mdiLanguagePhp,
      mdiLanguagePython,
      mdiLanguageCsharp,
      mdiLanguageJava,
      mdiCellphoneKey,
      mdiPowershell,
      mdiLanguageGo,
      selectedTab: 'javascript',
      yearlyPricing: false,
      faqPanel: null,
      pricing: 0,
      pricingLabels: ['10K', '20K', '50K', '100K', '200K'],
      pricingLabelsFull: ['10,000', '20,000', '50,000', '100,000', '200,000'],
    }
  },
  computed: {
    planMessages() {
      const plan = this.pricingLabels[this.pricing]
      return plan.replace('K', ',000')
    },
    planMonthlyPrice() {
      const prices = [20, 35, 89, 175, 350]
      return prices[this.pricing]
    },
    planYearlyPrice() {
      const prices = [200, 350, 1068, 2100, 4200]
      return prices[this.pricing]
    },
    planYearlyMonthlyPrice() {
      const prices = [16.66, 29.16, 89, 175, 350]
      return prices[this.pricing]
    },
  },
  methods: {
    formatVND(usd: number): string {
      const vnd = usd * 24000; // Assuming 1 USD = 24,000 VND
      return new Intl.NumberFormat('vi-VN', {
        style: 'currency',
        currency: 'VND',
      }).format(vnd);
    },
  },
})
</script>

<style lang="scss">
.fade-enter-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-active {
  opacity: 0;
}

.min-h-650 {
  min-height: 650px;
}

.gradient-header {
  color: #1ad37f;
  background-image: -webkit-linear-gradient(0deg, #1ad37f 14%, #329ef4 55%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.gradient-underline {
  color: white;
}
</style>
