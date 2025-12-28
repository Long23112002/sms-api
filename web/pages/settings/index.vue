<template>
  <v-container
    fluid
    class="px-0 pt-0"
    :fill-height="$vuetify.breakpoint.lgAndUp"
  >
    <div class="w-full h-full">
      <v-app-bar height="60" fixed :dense="$vuetify.breakpoint.mdAndDown">
        <v-btn icon to="/threads">
          <v-icon>{{ mdiArrowLeft }}</v-icon>
        </v-btn>
        <v-toolbar-title>
          <div class="py-16">Cài Đặt</div>
        </v-toolbar-title>
      </v-app-bar>
      <v-container class="mt-16">
        <v-row>
          <v-col cols="12" md="9" offset-md="1" xl="8" offset-xl="2">
            <div v-if="$fire.auth.currentUser" class="text-center">
              <v-avatar size="100" color="indigo" class="mx-auto">
                <img
                  v-if="$fire.auth.currentUser.photoURL"
                  :src="$fire.auth.currentUser.photoURL"
                  :alt="$fire.auth.currentUser.displayName"
                />
                <v-icon v-else dark size="70">{{ mdiAccountCircle }}</v-icon>
              </v-avatar>
              <h3 v-if="$fire.auth.currentUser.displayName">
                {{ $fire.auth.currentUser.displayName }}
              </h3>
              <h4 class="text--secondary">
                {{ $fire.auth.currentUser.email }}
                <v-icon
                  v-if="$fire.auth.currentUser.emailVerified"
                  small
                  color="primary"
                >
                  {{ mdiShieldCheck }}
                </v-icon>
              </h4>
              <v-autocomplete
                v-if="$store.getters.getUser"
                dense
                outlined
                :value="$store.getters.getUser.timezone"
                class="mx-auto mt-2"
                style="max-width: 250px"
                label="Múi Giờ"
                :items="timezones"
                @change="updateTimezone"
              ></v-autocomplete>
            </div>
            <h5 class="text-h4 mb-3 mt-3">API Key</h5>
            <p class="text--secondary">
              Sử dụng API Key của bạn trong HTTP Header <code>x-api-key</code> khi
              gửi requests đến các endpoints
              <code>https://api.httpsms.com</code>.
            </p>
            <div v-if="apiKey === ''" class="mb-n9 pl-3 pt-5">
              <v-progress-circular
                :size="20"
                :width="2"
                color="primary"
                indeterminate
              ></v-progress-circular>
            </div>
            <v-text-field
              v-else
              :append-icon="apiKeyShow ? mdiEye : mdiEyeOff"
              :type="apiKeyShow ? 'text' : 'password'"
              :value="apiKey"
              readonly
              name="api-key"
              outlined
              class="mb-n2"
              @click:append="apiKeyShow = !apiKeyShow"
            ></v-text-field>
            <div class="d-flex flex-wrap">
              <copy-button
                :value="apiKey"
                color="primary"
                copy-text="Sao Chép API Key"
                notification-text="Đã sao chép API Key thành công"
              ></copy-button>
              <v-btn
                v-if="$vuetify.breakpoint.mdAndUp"
                color="primary"
                class="ml-4"
                @click="showQrCodeDialog = true"
              >
                <v-icon left>{{ mdiQrcode }}</v-icon>
                Hiển Thị QR Code
              </v-btn>
              <v-dialog
                v-model="showQrCodeDialog"
                overlay-opacity="0.9"
                max-width="400px"
              >
                <v-card>
                  <v-card-title class="justify-center"
                    >QR Code API Key</v-card-title
                  >
                  <v-card-subtitle class="mt-2 text-center"
                    >Quét mã QR này bằng ứng dụng
                    <a :href="$store.getters.getAppData.appDownloadUrl"
                      >httpSMS</a
                    >
                    trên điện thoại Android của bạn để đăng nhập.</v-card-subtitle
                  >
                  <v-card-text class="text-center">
                    <canvas ref="qrCodeCanvas"></canvas>
                  </v-card-text>
                  <v-card-actions>
                    <v-btn
                      color="primary"
                      block
                      class="mb-4"
                      @click="showQrCodeDialog = false"
                      >Đóng</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </v-dialog>
              <v-btn
                v-if="$vuetify.breakpoint.lgAndUp"
                class="ml-4"
                :href="$store.getters.getAppData.documentationUrl"
                >Documentation</v-btn
              >
              <v-spacer></v-spacer>
              <v-dialog
                v-model="showRotateApiKey"
                overlay-opacity="0.9"
                max-width="550"
              >
                <template #activator="{ on, attrs }">
                  <v-btn
                    :small="$vuetify.breakpoint.mdAndDown"
                    :text="$vuetify.breakpoint.lgAndUp"
                    color="warning"
                    v-bind="attrs"
                    v-on="on"
                  >
                    <v-icon left>{{ mdiRefresh }}</v-icon>
                    Xoay API Key
                  </v-btn>
                </template>
                <v-card>
                  <v-card-title class="text-h5 text-break">
                    Bạn có chắc chắn muốn xoay API Key của mình?
                  </v-card-title>
                  <v-card-text>
                    Bạn sẽ phải đăng xuất và đăng nhập lại trên ứng dụng
                    <b>httpSMS</b> Android với API key mới sau khi xoay.
                  </v-card-text>
                  <v-card-actions class="pb-4">
                    <v-btn
                      color="primary"
                      :loading="rotatingApiKey"
                      @click="rotateApiKey"
                    >
                      <v-icon left>{{ mdiRefresh }}</v-icon>
                      Có, Xoay Key
                    </v-btn>
                    <v-spacer></v-spacer>
                    <v-btn text @click="showRotateApiKey = false">
                      Close
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </div>
            <h5 id="webhook-settings" class="text-h4 mb-3 mt-12">Webhooks</h5>
            <p class="text--secondary">
              Webhooks cho phép chúng tôi gửi sự kiện đến server của bạn, ví dụ khi
              điện thoại Android nhận được tin nhắn SMS, chúng tôi có thể chuyển tiếp
              tin nhắn đến server của bạn.
            </p>
            <div v-if="loadingWebhooks">
              <v-progress-circular
                :size="60"
                :width="2"
                color="primary"
                class="mb-4"
                indeterminate
              ></v-progress-circular>
            </div>
            <v-simple-table v-else-if="webhooks.length" class="mb-4">
              <template #default>
                <thead>
                  <tr class="text-uppercase subtitle-2">
                    <th v-if="$vuetify.breakpoint.xlOnly" class="text-left">
                      ID
                    </th>
                    <th class="text-left text-break">URL Callback</th>
                    <th v-if="$vuetify.breakpoint.lgAndUp" class="text-center">
                      Sự Kiện
                    </th>
                    <th class="text-center">Thao Tác</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="webhook in webhooks" :key="webhook.id">
                    <td v-if="$vuetify.breakpoint.xlOnly" class="text-left">
                      {{ webhook.id }}
                    </td>
                    <td class="text-break">{{ webhook.url }}</td>
                    <td v-if="$vuetify.breakpoint.lgAndUp" class="text-center">
                      <v-chip
                        v-for="event in webhook.events"
                        :key="event"
                        small
                        >{{ event }}</v-chip
                      >
                    </td>
                    <td class="text-center">
                      <v-btn
                        :icon="$vuetify.breakpoint.mdAndDown"
                        small
                        color="info"
                        :disabled="updatingWebhook"
                        @click.prevent="onWebhookEdit(webhook.id)"
                      >
                        <v-icon small>{{ mdiSquareEditOutline }}</v-icon>
                        <span v-if="!$vuetify.breakpoint.mdAndDown">
                          Sửa
                        </span>
                      </v-btn>
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
            <div class="d-flex">
              <v-btn color="primary" @click="onWebhookCreate">
                <v-icon left>{{ mdiLinkVariant }}</v-icon>
                Thêm Webhook
              </v-btn>
              <v-btn
                v-if="$vuetify.breakpoint.lgAndUp"
                class="ml-4"
                href="https://docs.httpsms.com/webhooks/introduction"
                >Documentation</v-btn
              >
            </div>
            <h5 id="discord-settings" class="text-h4 mb-3 mt-12">
              Tích Hợp Discord
            </h5>
            <p class="text--secondary">
              Gửi và nhận tin nhắn SMS mà không cần rời khỏi server Discord của bạn
              với ứng dụng httpSMS Discord sử dụng lệnh
              <code>/httpsms</code>.
            </p>
            <div v-if="loadingDiscordIntegrations">
              <v-progress-circular
                :size="60"
                :width="2"
                color="primary"
                class="mb-4"
                indeterminate
              ></v-progress-circular>
            </div>
            <v-simple-table v-else-if="discords.length" class="mb-4">
              <template #default>
                <thead>
                  <tr class="text-uppercase subtitle-2">
                    <th class="text-left">Tên</th>
                    <th class="text-left">ID Server</th>
                    <th class="text-left">ID Kênh</th>
                    <th class="text-center">Thao Tác</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="discord in discords" :key="discord.id">
                    <td class="text-left">
                      {{ discord.name }}
                    </td>
                    <td class="text-left">
                      {{ discord.server_id }}
                    </td>
                    <td class="text-left">
                      {{ discord.incoming_channel_id }}
                    </td>
                    <td class="text-center">
                      <v-btn
                        :icon="$vuetify.breakpoint.mdAndDown"
                        small
                        color="info"
                        :disabled="updatingDiscord"
                        @click.prevent="onDiscordEdit(discord.id)"
                      >
                        <v-icon small>{{ mdiSquareEditOutline }}</v-icon>
                        <span v-if="!$vuetify.breakpoint.mdAndDown">
                          Sửa
                        </span>
                      </v-btn>
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
            <v-btn color="#5865f2" @click="onDiscordCreate">
              <v-img
                contain
                height="24"
                width="24"
                class="mr-2"
                :src="require('assets/img/discord-logo.svg')"
              ></v-img>
              Add Discord
            </v-btn>
            <h5 id="phones" class="text-h4 mb-3 mt-12">Điện Thoại</h5>
            <p class="text--secondary">
              Danh sách điện thoại di động đã được đăng ký để gửi và
              nhận tin nhắn SMS.
            </p>
            <v-simple-table>
              <template #default>
                <thead>
                  <tr class="text-uppercase subtitle-2">
                    <th v-if="$vuetify.breakpoint.xlOnly" class="text-left">
                      ID
                    </th>
                    <th class="text-left">Số Điện Thoại</th>
                    <th v-if="$vuetify.breakpoint.lgAndUp" class="text-center">
                      Thử Lại
                    </th>
                    <th class="text-center">Tốc Độ</th>
                    <th class="text-center">Cập Nhật Lúc</th>
                    <th class="text-center">Thao Tác</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="phone in $store.getters.getPhones" :key="phone.id">
                    <td v-if="$vuetify.breakpoint.xlOnly" class="text-left">
                      {{ phone.id }}
                    </td>
                    <td>{{ phone.phone_number | phoneNumber }}</td>
                    <td v-if="$vuetify.breakpoint.lgAndUp">
                      <div class="d-flex justify-center">
                        {{
                          phone.max_send_attempts ? phone.max_send_attempts : 1
                        }}
                      </div>
                    </td>
                    <td class="text-center">
                      <span v-if="phone.messages_per_minute"
                        >{{ phone.messages_per_minute }}/min</span
                      >
                      <span v-else>Không Giới Hạn</span>
                    </td>
                    <td class="text-center">
                      {{ phone.updated_at | timestamp }}
                    </td>
                    <td class="text-center">
                      <v-btn
                        :icon="$vuetify.breakpoint.mdAndDown"
                        color="info"
                        :disabled="updatingPhone"
                        @click.prevent="showEditPhone(phone.id)"
                      >
                        <v-icon small>{{ mdiSquareEditOutline }}</v-icon>
                        <span v-if="!$vuetify.breakpoint.mdAndDown">
                          Sửa
                        </span>
                      </v-btn>
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
            <h5 id="email-notifications" class="text-h4 mb-3 mt-12">
              Thông Báo Email
            </h5>
            <p class="text--secondary">
              Quản lý các thông báo email mà bạn nhận được từ httpSMS.
              Bạn có thể bật/tắt từng thông báo bất cứ lúc nào để
              không bị quá tải với email
            </p>
            <v-switch
              v-model="notificationSettings.heartbeat_enabled"
              label="Email Heartbeat"
              :disabled="updatingEmailNotifications"
              hint="Công tắc này điều khiển thông báo email chúng tôi gửi khi không nhận được heartbeat từ điện thoại của bạn trong 1 giờ."
              persistent-hint
            ></v-switch>
            <v-switch
              v-model="notificationSettings.webhook_enabled"
              label="Email Webhook và Discord"
              :disabled="updatingEmailNotifications"
              hint="Công tắc này điều khiển thông báo email chúng tôi gửi khi không thể chuyển tiếp sự kiện đến server Discord hoặc webhook của bạn."
              persistent-hint
            ></v-switch>
            <v-switch
              v-model="notificationSettings.message_status_enabled"
              label="Email Trạng Thái Tin Nhắn"
              :disabled="updatingEmailNotifications"
              hint="Công tắc này điều khiển thông báo email chúng tôi gửi khi tin nhắn của bạn thất bại hoặc hết hạn."
              persistent-hint
            ></v-switch>
            <v-switch
              v-model="notificationSettings.newsletter_enabled"
              label="Email Bản Tin"
              :disabled="updatingEmailNotifications"
              hint="Công tắc này điều khiển email bản tin về tính năng mới, cập nhật và khuyến mãi."
              persistent-hint
            ></v-switch>
            <v-btn
              color="primary"
              :loading="updatingEmailNotifications"
              class="mt-4"
              @click="saveEmailNotifications"
            >
              <v-icon left>{{ mdiContentSave }}</v-icon>
              Lưu Cài Đặt Thông Báo
            </v-btn>
            <h5 id="email-notifications" class="text-h4 error--text mb-3 mt-12">
              Xóa Tài Khoản
            </h5>
            <p v-if="hasActiveSubscription" class="text--secondary">
              Bạn không thể xóa tài khoản vì bạn có gói đăng ký đang hoạt động
              trên httpSMS.
              <router-link class="text-decoration-none" to="/billing"
                >Hủy đăng ký của bạn</router-link
              >
              trước khi xóa tài khoản.
            </p>
            <p v-else class="text--secondary">
              Bạn có thể xóa tất cả dữ liệu của mình trên httpSMS bằng cách nhấp vào nút
              bên dưới. Hành động này <b>không thể hoàn tác</b> và tất cả dữ liệu của bạn sẽ
              bị xóa vĩnh viễn khỏi cơ sở dữ liệu httpSMS ngay lập tức và
              không thể khôi phục.
            </p>
            <v-btn
              color="error"
              :loading="deletingAccount"
              class="mt-4"
              :disabled="hasActiveSubscription"
              @click="showDeleteAccountDialog = true"
            >
              <v-icon left>{{ mdiDelete }}</v-icon>
              Xóa Tài Khoản Của Bạn
            </v-btn>
            <v-dialog
              v-model="showDeleteAccountDialog"
              overlay-opacity="0.9"
              max-width="600px"
            >
              <v-card>
                <v-card-title class="justify-center text-center"
                  >Xóa tài khoản httpSMS của bạn</v-card-title
                >
                <v-card-text class="mt-2 text-center">
                  Bạn có chắc chắn muốn xóa tài khoản của mình? Hành động này
                  <b>không thể hoàn tác</b> và tất cả dữ liệu của bạn sẽ bị xóa vĩnh viễn
                  khỏi cơ sở dữ liệu httpSMS ngay lập tức.
                </v-card-text>
                <v-card-actions>
                  <v-btn
                    color="error"
                    text
                    :loading="deletingAccount"
                    @click="deleteUserAccount"
                  >
                    <v-icon v-if="$vuetify.breakpoint.lgAndUp" left>{{
                      mdiDelete
                    }}</v-icon>
                    Xóa Tài Khoản Của Tôi
                  </v-btn>
                  <v-spacer></v-spacer>
                  <v-btn
                    color="primary"
                    @click="showDeleteAccountDialog = false"
                  >
                    <span v-if="$vuetify.breakpoint.lgAndUp"
                      >Giữ Tài Khoản</span
                    >
                    <span v-else>Close</span>
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-col>
        </v-row>
      </v-container>
    </div>
    <v-dialog v-model="showPhoneEdit" overlay-opacity="0.9" max-width="700px">
      <v-card>
        <v-card-title>Sửa Điện Thoại</v-card-title>
        <v-card-text v-if="activePhone" class="mt-6">
          <v-container>
            <v-row>
              <v-col>
                <v-text-field
                  outlined
                  dense
                  disabled
                  label="ID"
                  :value="activePhone.id"
                >
                </v-text-field>
                <v-text-field
                  outlined
                  disabled
                  dense
                  label="Số Điện Thoại"
                  :value="activePhone.phone_number"
                >
                </v-text-field>
                <v-text-field
                  outlined
                  disabled
                  dense
                  label="SIM"
                  :value="activePhone.sim"
                >
                </v-text-field>
                <v-textarea
                  outlined
                  disabled
                  dense
                  label="FCM Token"
                  :value="activePhone.fcm_token"
                >
                </v-textarea>
                <v-text-field
                  v-model="activePhone.message_expiration_seconds"
                  outlined
                  type="number"
                  dense
                  label="Thời Gian Hết Hạn Tin Nhắn (giây)"
                >
                </v-text-field>
                <v-text-field
                  v-model="activePhone.messages_per_minute"
                  outlined
                  type="number"
                  dense
                  label="Tin Nhắn Mỗi Phút"
                >
                </v-text-field>
                <v-text-field
                  v-model="activePhone.max_send_attempts"
                  outlined
                  type="number"
                  dense
                  placeholder="Số lần thử lại khi gửi SMS"
                  label="Số Lần Gửi Tối Đa"
                >
                </v-text-field>
                <v-textarea
                  v-model="activePhone.missed_call_auto_reply"
                  outlined
                  dense
                  label="Tự Động Trả Lời Cuộc Gọi Nhỡ"
                  persistent-placeholder
                  persistent-hint
                  placeholder="Chúng tôi hiện đang đóng cửa, vui lòng gửi tin nhắn cho chúng tôi từ 09:00 đến 17:00"
                  hint="Ở đây bạn có thể cấu hình tin nhắn SMS tự động được gửi đến người gọi khi điện thoại này có cuộc gọi nhỡ"
                >
                </v-textarea>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions class="mt-n8">
          <v-btn small color="info" @click="updatePhone">
            <v-icon v-if="$vuetify.breakpoint.lgAndUp" small>
              {{ mdiContentSave }}
            </v-icon>
            Update
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn small color="error" text @click="deletePhone(activePhone.id)">
            <v-icon v-if="$vuetify.breakpoint.lgAndUp" small>
              {{ mdiDelete }}
            </v-icon>
            Xóa
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="showWebhookEdit" overlay-opacity="0.9" max-width="600px">
      <v-card>
        <v-card-title>
          <span v-if="!activeWebhook.id">Thêm&nbsp;</span>
          <span v-else>Sửa&nbsp;</span>
          webhook mới
        </v-card-title>
        <v-card-text>
          <v-row>
            <v-col class="pt-8">
              <v-text-field
                v-if="activeWebhook.id"
                outlined
                dense
                disabled
                label="ID"
                :value="activeWebhook.id"
              >
              </v-text-field>
              <v-text-field
                v-model="activeWebhook.url"
                outlined
                dense
                label="URL Callback"
                persistent-placeholder
                persistent-hint
                :error="errorMessages.has('url')"
                :error-messages="errorMessages.get('url')"
                hint="Một POST request sẽ được gửi đến URL này mỗi khi một sự kiện được kích hoạt trong httpSMS."
                placeholder="https://example.com/webhook"
              >
              </v-text-field>
              <v-text-field
                v-model="activeWebhook.signing_key"
                outlined
                dense
                class="mt-6"
                persistent-placeholder
                persistent-hint
                label="Signing Key (tùy chọn)"
                placeholder="******************"
                :error="errorMessages.has('signing_key')"
                :error-messages="errorMessages.get('signing_key')"
                hint="Signing key được sử dụng để xác minh webhook được gửi từ httpSMS."
              >
              </v-text-field>
              <v-select
                v-model="activeWebhook.events"
                :items="events"
                label="Sự Kiện"
                multiple
                outlined
                persistent-placeholder
                class="mt-6"
                dense
                :error="errorMessages.has('events')"
                :error-messages="errorMessages.get('events')"
                hint="Chọn nhiều sự kiện httpSMS để theo dõi"
                persistent-hint
              ></v-select>
              <v-select
                v-model="activeWebhook.phone_numbers"
                :items="phoneNumbers"
                label="Số Điện Thoại"
                multiple
                outlined
                persistent-placeholder
                class="mt-6"
                dense
                :error="errorMessages.has('phone_numbers')"
                :error-messages="errorMessages.get('phone_numbers')"
                hint="Chọn nhiều số điện thoại để theo dõi sự kiện"
                persistent-hint
              ></v-select>
            </v-col>
          </v-row>
        </v-card-text>
        <v-card-actions class="mt-n4 pb-4">
          <loading-button
            v-if="!activeWebhook.id"
            :icon="mdiContentSave"
            :loading="updatingWebhook"
            @click="createWebhook"
          >
            Lưu Webhook
          </loading-button>
          <loading-button
            v-else
            small
            color="info"
            :loading="updatingWebhook"
            @click="updateWebhook"
          >
            <v-icon v-if="$vuetify.breakpoint.lgAndUp" small>
              {{ mdiContentSave }}
            </v-icon>
            Cập Nhật Webhook
          </loading-button>
          <v-spacer></v-spacer>
          <v-btn
            v-if="activeWebhook.id"
            :disabled="updatingWebhook"
            small
            color="error"
            text
            @click="deleteWebhook(activeWebhook.id)"
          >
            <v-icon v-if="$vuetify.breakpoint.lgAndUp" small>
              {{ mdiDelete }}
            </v-icon>
            Xóa
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="showDiscordEdit" overlay-opacity="0.9" max-width="700px">
      <v-card>
        <v-card-title>
          <span v-if="!activeDiscord.id">Thêm&nbsp;</span>
          <span v-else>Sửa&nbsp;</span>
          tích hợp Discord mới
        </v-card-title>
        <v-card-text>
          <v-row>
            <v-col class="pt-8">
              <p class="mt-n4 subtitle-1">
                Nhấp vào nút bên dưới để thêm bot httpSMS vào server Discord của bạn.
                Bạn cần làm điều này để chúng tôi có quyền gửi
                và nhận tin nhắn trên server Discord của bạn.
              </p>
              <v-btn
                color="#5865f2"
                class="mb-6"
                target="_blank"
                href="https://discord.com/api/oauth2/authorize?client_id=1095780203256627291&permissions=2147485760&scope=bot%20applications.commands"
              >
                <v-icon left>{{ mdiConnection }}</v-icon>
                Thêm Bot Discord
              </v-btn>
              <v-text-field
                v-if="activeDiscord.id"
                outlined
                dense
                disabled
                label="ID"
                :value="activeDiscord.id"
              >
              </v-text-field>
              <v-text-field
                v-model="activeDiscord.name"
                outlined
                dense
                label="Tên"
                persistent-placeholder
                persistent-hint
                :error="errorMessages.has('name')"
                :error-messages="errorMessages.get('name')"
                hint="Tên của tích hợp Discord"
                placeholder="e.g Game Server"
              >
              </v-text-field>
              <v-text-field
                v-model="activeDiscord.server_id"
                outlined
                dense
                class="mt-6"
                persistent-placeholder
                persistent-hint
                label="ID Server Discord"
                placeholder="e.g 1095778291488653372"
                :error="errorMessages.has('server_id')"
                :error-messages="errorMessages.get('server_id')"
                hint="Bạn có thể lấy điều này bằng cách nhấp chuột phải vào server của bạn và nhấp vào Copy Server ID."
              >
              </v-text-field>
              <v-text-field
                v-model="activeDiscord.incoming_channel_id"
                outlined
                dense
                class="mt-6"
                persistent-placeholder
                persistent-hint
                label="ID Kênh Nhận Discord"
                placeholder="e.g 1095778291488653372"
                :error="errorMessages.has('incoming_channel_id')"
                :error-messages="errorMessages.get('incoming_channel_id')"
                hint="Bạn có thể lấy điều này bằng cách nhấp chuột phải vào kênh Discord của bạn và nhấp vào Copy Channel ID."
              >
              </v-text-field>
            </v-col>
          </v-row>
        </v-card-text>
        <v-card-actions class="mt-n4 pb-4 pl-6">
          <loading-button
            v-if="!activeDiscord.id"
            :icon="mdiContentSave"
            :loading="updatingDiscord"
            @click="createDiscord"
          >
            Lưu Tích Hợp Discord
          </loading-button>
          <loading-button
            v-else
            color="info"
            :loading="updatingDiscord"
            @click="updateDiscord"
          >
            <v-icon v-if="$vuetify.breakpoint.lgAndUp" small>
              {{ mdiContentSave }}
            </v-icon>
            Cập Nhật Tích Hợp Discord
          </loading-button>
          <v-spacer></v-spacer>
          <v-btn
            v-if="activeDiscord.id"
            :disabled="updatingDiscord"
            color="error"
            text
            @click="deleteDiscord(activeDiscord.id)"
          >
            <v-icon v-if="$vuetify.breakpoint.lgAndUp" small>
              {{ mdiDelete }}
            </v-icon>
            Xóa
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import Vue from 'vue'
import {
  mdiArrowLeft,
  mdiAccountCircle,
  mdiShieldCheck,
  mdiDelete,
  mdiContentSave,
  mdiConnection,
  mdiEye,
  mdiRefresh,
  mdiLinkVariant,
  mdiEyeOff,
  mdiSquareEditOutline,
  mdiQrcode,
} from '@mdi/js'
import QRCode from 'qrcode'
import { ErrorMessages } from '~/plugins/errors'
import LoadingButton from '~/components/LoadingButton.vue'

export default Vue.extend({
  name: 'SettingsIndex',
  components: { LoadingButton },
  middleware: ['auth'],
  data() {
    return {
      mdiEye,
      mdiEyeOff,
      mdiRefresh,
      mdiArrowLeft,
      mdiAccountCircle,
      mdiShieldCheck,
      mdiDelete,
      mdiQrcode,
      mdiLinkVariant,
      mdiContentSave,
      mdiSquareEditOutline,
      mdiConnection,
      errorMessages: new ErrorMessages(),
      apiKeyShow: false,
      showPhoneEdit: false,
      showDiscordEdit: false,
      showRotateApiKey: false,
      rotatingApiKey: false,
      showQrCodeDialog: false,
      deletingAccount: false,
      showDeleteAccountDialog: false,
      activeWebhook: {
        id: null,
        url: '',
        signing_key: '',
        phone_numbers: [],
        events: ['message.phone.received'],
      },
      activeDiscord: {
        id: null,
        name: '',
        server_id: '',
        missed_call_auto_reply: '',
        incoming_channel_id: '',
      },
      updatingEmailNotifications: false,
      notificationSettings: {
        webhook_enabled: true,
        message_status_enabled: true,
        newsletter_enabled: true,
        heartbeat_enabled: true,
      },
      updatingWebhook: false,
      loadingWebhooks: false,
      discords: [],
      webhooks: [],
      showWebhookEdit: false,
      activePhone: null,
      updatingPhone: false,
      updatingDiscord: false,
      loadingDiscordIntegrations: false,
      events: [
        'message.phone.received',
        'message.phone.sent',
        'message.phone.delivered',
        'message.send.failed',
        'message.send.expired',
        'message.call.missed',
        'phone.heartbeat.offline',
        'phone.heartbeat.online',
      ],
    }
  },
  head() {
    return {
      title: 'Cài Đặt - httpSMS',
    }
  },
  computed: {
    apiKey() {
      if (this.$store.getters.getUser === null) {
        return ''
      }
      return this.$store.getters.getUser.api_key
    },
    hasActiveSubscription() {
      if (this.$store.getters.getUser === null) {
        return true
      }
      return this.$store.getters.getUser.subscription_renews_at != null
    },
    timezones() {
      return Intl.supportedValuesOf('timeZone')
    },
    phoneNumbers() {
      return this.$store.getters.getPhones.map((phone) => {
        return phone.phone_number
      })
    },
  },
  watch: {
    showQrCodeDialog(newVal) {
      if (newVal && this.apiKey) {
        this.$nextTick(() => {
          this.generateQrCode(this.apiKey)
        })
      }
    },
  },
  async mounted() {
    await Promise.all([
      this.$store.dispatch('clearAxiosError'),
      this.$store.dispatch('loadUser'),
      this.$store.dispatch('loadPhones'),
    ])
    this.loadWebhooks()
    this.loadDiscordIntegrations()
    this.updateEmailNotifications()
    if (this.$route.hash) {
      await this.$vuetify.goTo(this.$route.hash)
    }
  },

  methods: {
    generateQrCode(text) {
      const canvas = this.$refs.qrCodeCanvas
      if (canvas) {
        QRCode.toCanvas(canvas, text, { errorCorrectionLevel: 'H' }, (err) => {
          if (err) {
            this.$store.dispatch('addNotification', {
              message: 'Failed to generate API key QR code',
              type: 'error',
            })
          }
        })
      }
    },
    updateEmailNotifications() {
      this.notificationSettings = {
        webhook_enabled:
          this.$store.getters.getUser.notification_webhook_enabled,
        message_status_enabled:
          this.$store.getters.getUser.notification_message_status_enabled,
        heartbeat_enabled:
          this.$store.getters.getUser.notification_heartbeat_enabled,
        newsletter_enabled:
          this.$store.getters.getUser.notification_newsletter_enabled,
      }
    },
    showEditPhone(phoneId) {
      const phone = this.$store.getters.getPhones.find((x) => x.id === phoneId)
      if (!phone) {
        return
      }
      this.activePhone = { ...phone }
      this.showPhoneEdit = true
      this.resetErrors()
    },

    onWebhookEdit(webhookId) {
      const webhook = this.webhooks.find((x) => x.id === webhookId)
      if (!webhook) {
        return
      }
      this.activeWebhook = {
        id: webhook.id,
        url: webhook.url,
        phone_numbers: webhook.phone_numbers.filter(
          (x) => this.phoneNumbers.find((y) => y === x) !== undefined,
        ),
        signing_key: webhook.signing_key,
        events: webhook.events,
      }
      this.showWebhookEdit = true
      this.resetErrors()
    },

    onDiscordEdit(discordId) {
      const discord = this.discords.find((x) => x.id === discordId)
      if (!discord) {
        return
      }
      this.activeDiscord = {
        id: discord.id,
        name: discord.name,
        server_id: discord.server_id,
        incoming_channel_id: discord.incoming_channel_id,
      }
      this.showDiscordEdit = true
      this.resetErrors()
    },

    onWebhookCreate() {
      this.activeWebhook = {
        id: null,
        url: '',
        signing_key: '',
        phone_numbers: this.$store.getters.getPhones.map(
          (phone) => phone.phone_number,
        ),
        events: [
          'message.phone.received',
          'message.phone.sent',
          'message.phone.delivered',
          'message.send.failed',
          'message.send.expired',
        ],
      }
      this.showWebhookEdit = true
      this.resetErrors()
    },

    onDiscordCreate() {
      this.activeDiscord = {
        id: null,
        name: '',
        server_id: '',
        incoming_channel_id: '',
        missed_call_auto_reply: '',
      }
      this.showDiscordEdit = true
      this.resetErrors()
    },

    async updatePhone() {
      this.updatingPhone = true
      await this.$store.dispatch('clearAxiosError')
      this.$store.dispatch('updatePhone', this.activePhone).finally(() => {
        if (!this.$store.getters.getAxiosError) {
          this.updatingPhone = false
          this.showPhoneEdit = false
          this.activePhone = null
        }
      })
    },

    resetErrors() {
      this.errorMessages = new ErrorMessages()
    },

    createDiscord() {
      this.resetErrors()
      this.updatingDiscord = true
      this.$store
        .dispatch('createDiscord', this.activeDiscord)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Discord integration created successfully',
            type: 'success',
          })
          this.showDiscordEdit = false
          this.loadDiscordIntegrations()
        })
        .catch((errors) => {
          this.errorMessages = errors
        })
        .finally(() => {
          this.updatingDiscord = false
        })
    },

    saveEmailNotifications() {
      this.updatingEmailNotifications = true
      this.$store
        .dispatch('saveEmailNotifications', this.notificationSettings)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Email notifications saved successfully',
            type: 'success',
          })
          this.updateEmailNotifications()
        })
        .finally(() => {
          this.updatingEmailNotifications = false
        })
    },

    updateDiscord() {
      this.resetErrors()
      this.updatingDiscord = true
      this.$store
        .dispatch('updateDiscordIntegration', this.activeDiscord)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Discord integration updated successfully',
            type: 'success',
          })
          this.showDiscordEdit = false
          this.loadDiscordIntegrations()
        })
        .catch((errors) => {
          this.errorMessages = errors
        })
        .finally(() => {
          this.updatingDiscord = false
        })
    },

    deleteDiscord(discordId) {
      this.updatingDiscord = true
      this.$store
        .dispatch('deleteDiscordIntegration', discordId)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Discord integration deleted successfully',
            type: 'success',
          })
          this.showDiscordEdit = false
          this.loadDiscordIntegrations()
        })
        .finally(() => {
          this.updatingDiscord = false
        })
    },

    createWebhook() {
      this.resetErrors()
      this.updatingWebhook = true
      this.$store
        .dispatch('createWebhook', this.activeWebhook)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Webhook created successfully',
            type: 'success',
          })
          this.showWebhookEdit = false
          this.loadWebhooks()
        })
        .catch((errors) => {
          this.errorMessages = errors
        })
        .finally(() => {
          this.updatingWebhook = false
        })
    },

    updateTimezone(timezone) {
      this.resetErrors()
      this.$store
        .dispatch('updateTimezone', timezone)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Timezone updated successfully',
            type: 'success',
          })
        })
        .catch(() => {
          this.$store.dispatch('addNotification', {
            message: 'Failed to update timezone',
            type: 'error',
          })
        })
    },

    updateWebhook() {
      this.resetErrors()
      this.updatingWebhook = true
      this.$store
        .dispatch('updateWebhook', this.activeWebhook)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Webhook updated successfully',
            type: 'success',
          })
          this.showWebhookEdit = false
          this.loadWebhooks()
        })
        .catch((errors) => {
          this.errorMessages = errors
        })
        .finally(() => {
          this.updatingWebhook = false
        })
    },

    rotateApiKey() {
      this.rotatingApiKey = true
      this.$store
        .dispatch('rotateApiKey', this.$store.getters.getUser.id)
        .finally(() => {
          this.rotatingApiKey = false
          this.showRotateApiKey = false
        })
    },

    deleteWebhook(webhookId) {
      this.updatingWebhook = true
      this.$store
        .dispatch('deleteWebhook', webhookId)
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Webhook deleted successfully',
            type: 'success',
          })
          this.showWebhookEdit = false
          this.loadWebhooks()
        })
        .finally(() => {
          this.updatingWebhook = false
        })
    },

    loadWebhooks() {
      this.loadingWebhooks = true
      this.$store
        .dispatch('getWebhooks')
        .then((webhooks) => {
          this.webhooks = webhooks
        })
        .finally(() => {
          this.loadingWebhooks = false
        })
    },

    loadDiscordIntegrations() {
      this.loadingDiscordIntegrations = true
      this.$store
        .dispatch('getDiscordIntegrations')
        .then((discords) => {
          this.discords = discords
        })
        .finally(() => {
          this.loadingDiscordIntegrations = false
        })
    },

    deleteUserAccount() {
      this.deletingAccount = true
      this.$store
        .dispatch('deleteUserAccount')
        .then((message) => {
          this.$store.dispatch('addNotification', {
            message: message ?? 'Your account has been deleted successfully',
            type: 'success',
          })
          this.$fire.auth.signOut().then(() => {
            this.$store.dispatch('setAuthUser', null)
            this.$store.dispatch('resetState')
            this.$store.dispatch('addNotification', {
              type: 'info',
              message: 'You have successfully logged out',
            })
            this.$router.push({ name: 'index' })
          })
        })
        .finally(() => {
          this.deletingAccount = false
          this.showDeleteAccountDialog = false
        })
    },

    deletePhone(phoneId) {
      this.updatingPhone = true
      this.$store.dispatch('deletePhone', phoneId).finally(() => {
        this.updatingPhone = false
        this.showPhoneEdit = false
        this.activePhone = null
      })
    },
  },
})
</script>
