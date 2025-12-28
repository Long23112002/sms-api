<template>
  <v-container
    fluid
    class="px-0 pt-0"
    :fill-height="$vuetify.breakpoint.lgAndUp"
  >
    <div class="w-full h-full">
      <v-app-bar height="60" :dense="$vuetify.breakpoint.mdAndDown">
        <v-btn icon to="/threads">
          <v-icon>{{ mdiArrowLeft }}</v-icon>
        </v-btn>
        <v-toolbar-title>
          <div class="py-16">Sử Dụng Tài Khoản</div>
        </v-toolbar-title>
        <v-progress-linear
          :active="loading"
          :indeterminate="loading"
          absolute
          bottom
        ></v-progress-linear>
      </v-app-bar>
      <v-container>
        <v-row>
          <v-col cols="12" md="9" offset-md="1" xl="8" offset-xl="2">
            <h5 class="text-h4 mb-3 mt-3">Gói Hiện Tại</h5>
            <v-row v-if="$store.getters.getUser">
              <v-col md="6" xl="4">
                <v-alert dense text prominent color="info">
                  <v-row align="center">
                    <v-col cols="12">
                      <h1
                        class="subtitle-1 font-weight-bold text-uppercase mt-3"
                      >
                        <span v-if="isOnFreePlan">{{ plan.name }}</span>
                        <span v-else-if="subscriptionIsCancelled"
                          ><span class="warning--text">{{ plan.name }}</span> →
                          Miễn Phí</span
                        >
                        <span v-else>{{ plan.name }}</span>
                      </h1>
                      <p
                        v-if="
                          !isOnFreePlan &&
                          !isOnLifetimePlan &&
                          !subscriptionIsCancelled
                        "
                        class="text--secondary"
                      >
                        Hóa đơn tiếp theo của bạn là <b>{{ formatVND(plan.price) }}</b> vào ngày
                        <b>{{
                          new Date(
                            $store.getters.getUser.subscription_renews_at,
                          ).toLocaleDateString()
                        }}</b>
                      </p>
                      <p v-if="isOnLifetimePlan" class="text--secondary">
                        Bạn đang ở gói trọn đời với giá
                        <b>{{ formatVND(plan.price) }}</b>
                      </p>
                      <p
                        v-else-if="subscriptionIsCancelled"
                        class="text--secondary"
                      >
                        Bạn sẽ được hạ xuống gói <b>MIỄN PHÍ</b> vào ngày
                        <b>{{
                          new Date(
                            $store.getters.getUser.subscription_ends_at,
                          ).toLocaleDateString()
                        }}</b>
                      </p>
                      <p v-else class="text--secondary">
                        {{ totalMessages }}/{{ plan.messagesPerMonth }} tin nhắn
                      </p>
                    </v-col>
                    <v-col cols="12" class="d-flex mb-2 mt-n6">
                      <loading-button
                        v-if="
                          !subscriptionIsCancelled &&
                          !isOnFreePlan &&
                          !isOnLifetimePlan
                        "
                        color="primary"
                        :loading="loading"
                        @click="updateDetails"
                      >
                        Cập Nhật Gói
                      </loading-button>
                      <v-btn
                        v-else-if="!isOnLifetimePlan"
                        color="primary"
                        :href="checkoutURL"
                        >Nâng Cấp Gói</v-btn
                      >
                      <v-spacer></v-spacer>
                      <v-dialog
                        v-if="
                          !subscriptionIsCancelled &&
                          !isOnFreePlan &&
                          !isOnLifetimePlan
                        "
                        v-model="dialog"
                        max-width="590"
                      >
                        <template #activator="{ on, attrs }">
                          <v-btn v-bind="attrs" color="error" text v-on="on">
                            Hủy Gói
                          </v-btn>
                        </template>
                        <v-card>
                          <v-card-text class="pt-4 mb-n6">
                            <h2 class="text--primary text-h5 mb-2">
                              Bạn có chắc chắn muốn hủy gói đăng ký của mình?
                            </h2>
                            <p>
                              Bạn sẽ được hạ xuống gói miễn phí vào cuối
                              chu kỳ thanh toán hiện tại vào ngày
                              <b>{{
                                new Date(
                                  $store.getters.getUser.subscription_renews_at,
                                ).toLocaleDateString()
                              }}</b>
                            </p>
                          </v-card-text>
                          <v-card-actions>
                            <v-btn color="primary" @click="dialog = false">
                              Giữ Gói Đăng Ký
                            </v-btn>
                            <v-spacer></v-spacer>
                            <loading-button
                              v-if="!isOnFreePlan"
                              :text="true"
                              :loading="loading"
                              color="error"
                              @click="cancelPlan"
                            >
                              Hủy Gói
                            </loading-button>
                          </v-card-actions>
                        </v-card>
                      </v-dialog>
                    </v-col>
                  </v-row>
                </v-alert>
              </v-col>
            </v-row>
            <h2 v-if="isOnFreePlan" class="text-h4 mt-4 mb-2">Nâng Cấp Gói</h2>
            <v-row v-if="isOnFreePlan">
              <v-col cols="12" md="6" xl="4">
                <v-hover v-slot="{ hover }">
                  <v-card
                    :color="hover ? 'black' : 'default'"
                    :href="checkoutURL"
                    outlined
                  >
                    <v-card-text>
                      <v-row align="center">
                        <v-col class="grow">
                          <h1
                            class="subtitle-1 font-weight-bold text-uppercase mt-3"
                          >
                            Pro - Hàng Tháng
                          </h1>
                          <p class="text--secondary">5,000 tin nhắn mỗi tháng</p>
                        </v-col>
                        <v-col class="shrink">
                          <span class="text-h5 text--primary">{{ formatVND(10) }}</span>/tháng
                        </v-col>
                      </v-row>
                    </v-card-text>
                  </v-card>
                </v-hover>
              </v-col>
              <v-col cols="12" md="6" xl="4">
                <v-hover v-slot="{ hover }">
                  <v-card
                    :color="hover ? 'black' : 'default'"
                    :href="checkoutURL"
                    outlined
                  >
                    <v-card-text>
                      <v-row align="center">
                        <v-col class="grow">
                          <h1
                            class="subtitle-1 font-weight-bold text-uppercase mt-3"
                          >
                            Pro - Hàng Năm
                            <v-chip small color="primary" class="mt-n1"
                              >2 tháng miễn phí</v-chip
                            >
                          </h1>
                          <p class="text--secondary">5,000 tin nhắn mỗi tháng</p>
                        </v-col>
                        <v-col class="shrink">
                          <span class="text-h5 text--primary">{{ formatVND(100) }}</span>/năm
                        </v-col>
                      </v-row>
                    </v-card-text>
                  </v-card>
                </v-hover>
              </v-col>
              <v-col cols="12" md="6" xl="4">
                <v-hover v-slot="{ hover }">
                  <v-card
                    :color="hover ? 'black' : 'default'"
                    :href="enterpriseCheckoutURL"
                    outlined
                  >
                    <v-card-text>
                      <v-row align="center">
                        <v-col class="grow">
                          <h1
                            class="subtitle-1 font-weight-bold text-uppercase mt-3"
                          >
                            100k - Hàng Tháng
                          </h1>
                          <p class="text--secondary">
                            100,000 tin nhắn mỗi tháng
                          </p>
                        </v-col>
                        <v-col class="shrink">
                          <span class="text-h5 text--primary">{{ formatVND(175) }}</span>/tháng
                        </v-col>
                      </v-row>
                    </v-card-text>
                  </v-card>
                </v-hover>
              </v-col>
            </v-row>
            <h5 class="text-h4 mb-3 mt-4">Tổng Quan</h5>
            <p class="text--secondary">
              Đây là tóm tắt các tin nhắn đã gửi và đã nhận trong
              <code
                v-if="$store.getters.getBillingUsage"
                class="font-weight-bold"
                >{{
                  $store.getters.getBillingUsage.start_timestamp | billingPeriod
                }}</code
              >.
            </p>
            <v-row v-if="$store.getters.getBillingUsage">
              <v-col cols="12" md="4">
                <v-alert
                  dark
                  dense
                  :icon="mdiCallMade"
                  prominent
                  type="info"
                  text
                >
                  <h2 class="text-h4 font-weight-bold mt-4">
                    {{ $store.getters.getBillingUsage.sent_messages | decimal }}
                  </h2>
                  <p class="text--secondary mt-n1">Tin Nhắn Đã Gửi</p>
                </v-alert>
              </v-col>
              <v-col cols="12" md="4">
                <v-alert
                  dark
                  dense
                  :icon="mdiCallReceived"
                  prominent
                  type="warning"
                  text
                >
                  <div class="d-flex">
                    <h2 class="text-h4 font-weight-bold mt-4">
                      {{
                        $store.getters.getBillingUsage.received_messages
                          | decimal
                      }}
                    </h2>
                  </div>
                  <p class="text--secondary mt-n1">Tin Nhắn Đã Nhận</p>
                </v-alert>
              </v-col>
              <v-col cols="12" md="4">
                <v-alert
                  dense
                  :icon="mdiCreditCard"
                  prominent
                  type="success"
                  text
                >
                  <h2 class="text-h4 font-weight-bold mt-4">
                    {{ $store.getters.getBillingUsage.total_cost | money }}
                  </h2>
                  <p class="text--secondary mt-n1">Tổng Chi Phí</p>
                </v-alert>
              </v-col>
            </v-row>
            <h5 class="text-h4 mb-3 mt-12">Lịch Sử Sử Dụng</h5>
            <p class="text--secondary">
              Tóm tắt tất cả các tin nhắn đã gửi và nhận trong 12 tháng qua
            </p>
            <v-simple-table>
              <template #default>
                <thead>
                  <tr class="text-uppercase">
                    <th class="text-left">Kỳ</th>
                    <th class="text-left">
                      Đã Gửi
                      <span v-if="$vuetify.breakpoint.lgAndUp">Tin Nhắn</span>
                    </th>
                    <th class="text-left">
                      Đã Nhận
                      <span v-if="$vuetify.breakpoint.lgAndUp">Tin Nhắn</span>
                    </th>
                    <th class="text-right">
                      <span v-if="$vuetify.breakpoint.lgAndUp">Tổng</span> Chi Phí
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="billingUsage in $store.getters
                      .getBillingUsageHistory"
                    :key="billingUsage.id"
                  >
                    <td>
                      {{ billingUsage.start_timestamp | billingPeriod }}
                    </td>
                    <td>
                      {{ billingUsage.sent_messages | decimal }}
                    </td>
                    <td>
                      {{ billingUsage.received_messages }}
                    </td>
                    <td class="text-right font-weight-bold">
                      {{ billingUsage.total_cost | money }}
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
          </v-col>
        </v-row>
      </v-container>
    </div>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue'
import {
  mdiArrowLeft,
  mdiAccountCircle,
  mdiShieldCheck,
  mdiDelete,
  mdiCog,
  mdiContentSave,
  mdiEye,
  mdiEyeOff,
  mdiCallReceived,
  mdiCallMade,
  mdiCreditCard,
  mdiSquareEditOutline,
} from '@mdi/js'

type PaymentPlan = {
  name: string
  id: string
  price: number
  messagesPerMonth: number
}

export default Vue.extend({
  name: 'BillingIndex',
  middleware: ['auth'],
  data() {
    return {
      mdiEye,
      mdiEyeOff,
      mdiArrowLeft,
      mdiAccountCircle,
      mdiShieldCheck,
      mdiDelete,
      mdiCog,
      mdiContentSave,
      mdiCallReceived,
      mdiCallMade,
      mdiCreditCard,
      mdiSquareEditOutline,
      loading: true,
      dialog: false,
      plans: [
        {
          name: 'Miễn Phí',
          id: 'free',
          messagesPerMonth: 200,
          price: 0,
        },
        {
          name: 'PRO - Hàng Tháng',
          id: 'pro-monthly',
          messagesPerMonth: 5000,
          price: 10,
        },
        {
          name: 'PRO - Hàng Năm',
          id: 'pro-yearly',
          messagesPerMonth: 5000,
          price: 100,
        },
        {
          name: 'Ultra - Hàng Tháng',
          id: 'ultra-monthly',
          messagesPerMonth: 10000,
          price: 20,
        },
        {
          name: 'Ultra - Hàng Năm',
          id: 'ultra-yearly',
          messagesPerMonth: 10000,
          price: 200,
        },
        {
          name: '20k - Hàng Tháng',
          id: '20k-monthly',
          messagesPerMonth: 20000,
          price: 35,
        },
        {
          name: '20k - Hàng Năm',
          id: '20k-yearly',
          messagesPerMonth: 20000,
          price: 350,
        },
        {
          name: '50k - Hàng Tháng',
          id: '50k-monthly',
          messagesPerMonth: 50000,
          price: 89,
        },
        {
          name: '100k - Hàng Tháng',
          id: '100k-monthly',
          messagesPerMonth: 100000,
          price: 175,
        },
        {
          name: '200k - Hàng Tháng',
          id: '200k-monthly',
          messagesPerMonth: 200000,
          price: 350,
        },
        {
          name: 'PRO - Trọn Đời',
          id: 'pro-lifetime',
          messagesPerMonth: 10000,
          price: 1000,
        },
      ],
    }
  },
  head() {
    return {
      title: 'Sử Dụng & Thanh Toán - httpSMS',
    }
  },
  computed: {
    checkoutURL() {
      const checkoutUrl = this.$config.checkoutURL
      if (!checkoutUrl || checkoutUrl === '') {
        return '#'
      }
      try {
        const url = new URL(checkoutUrl)
      const user = this.$store.getters.getAuthUser
        if (user?.id) {
          url.searchParams.append('checkout[custom][user_id]', user.id)
        }
        if (user?.email) {
          url.searchParams.append('checkout[email]', user.email)
        }
        if (user?.displayName) {
          url.searchParams.append('checkout[name]', user.displayName)
        }
      return url.toString()
      } catch (e) {
        console.error('Invalid checkout URL:', checkoutUrl, e)
        return '#'
      }
    },
    enterpriseCheckoutURL() {
      const enterpriseUrl = this.$config.enterpriseCheckoutURL
      if (!enterpriseUrl || enterpriseUrl === '') {
        return '#'
      }
      try {
        const url = new URL(enterpriseUrl)
      const user = this.$store.getters.getAuthUser
        if (user?.id) {
          url.searchParams.append('checkout[custom][user_id]', user.id)
        }
        if (user?.email) {
          url.searchParams.append('checkout[email]', user.email)
        }
        if (user?.displayName) {
          url.searchParams.append('checkout[name]', user.displayName)
        }
      return url.toString()
      } catch (e) {
        console.error('Invalid enterprise checkout URL:', enterpriseUrl, e)
        return '#'
      }
    },

    plan(): PaymentPlan {
      return this.plans.find(
        (x) =>
          x.id === (this.$store.getters.getUser?.subscription_name || 'free'),
      )!
    },
    isOnFreePlan(): boolean {
      return this.plan.id === 'free'
    },
    isOnLifetimePlan(): boolean {
      return this.plan.id === 'pro-lifetime'
    },
    subscriptionIsCancelled(): boolean {
      return this.$store.getters.getUser?.subscription_status === 'cancelled'
    },
    totalMessages(): number {
      if (this.$store.getters.getBillingUsage == null) {
        return 0
      }
      return (
        this.$store.getters.getBillingUsage.sent_messages +
        this.$store.getters.getBillingUsage.received_messages
      )
    },
  },
  async mounted() {
    await this.loadData()
  },

  methods: {
    formatVND(usd: number): string {
      // Tỷ giá: 1 USD = 24,000 VND
      const vnd = Math.round(usd * 24000)
      return new Intl.NumberFormat('vi-VN', {
        style: 'currency',
        currency: 'VND',
      }).format(vnd)
    },
    async loadData() {
      await Promise.all([
        this.$store.dispatch('loadUser'),
        this.$store.dispatch('loadBillingUsage'),
        this.$store.dispatch('loadBillingUsageHistory'),
      ])
      this.loading = false
    },
    updateDetails() {
      this.loading = true
      this.$store
        .dispatch('getSubscriptionUpdateLink')
        .then((link: string) => {
          window.location.href = link
        })
        .catch(() => {
          this.loading = false
        })
    },
    cancelPlan() {
      this.loading = true
      this.$store
        .dispatch('cancelSubscription')
        .then(() => {
          this.$store.dispatch('addNotification', {
            message: 'Hủy gói đăng ký thành công',
            type: 'success',
          })
          this.$router.push('/')
        })
        .catch(() => {
          this.loading = false
        })
    },
  },
})
</script>
