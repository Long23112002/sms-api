<template>
  <div>
    <v-tabs v-model="tab" grow>
      <v-tab>Đăng Nhập</v-tab>
      <v-tab>Đăng Ký</v-tab>
    </v-tabs>

    <v-tabs-items v-model="tab">
      <!-- Sign In Tab -->
      <v-tab-item>
        <v-form ref="signInForm" @submit.prevent="signIn">
          <v-text-field
            v-model="signInEmail"
            label="Email"
            hint="Nhập địa chỉ email của bạn"
            type="email"
            :rules="emailRules"
            required
            outlined
            dense
            class="mt-4"
          ></v-text-field>
          <v-text-field
            v-model="signInPassword"
            label="Mật khẩu"
            hint="Nhập mật khẩu của bạn"
            :type="showSignInPassword ? 'text' : 'password'"
            :rules="passwordRules"
            required
            outlined
            dense
            :append-icon="showSignInPassword ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append="showSignInPassword = !showSignInPassword"
          ></v-text-field>
          <v-btn
            type="submit"
            color="primary"
            block
            :loading="signInLoading"
            :disabled="signInLoading"
            class="mt-4"
          >
            Đăng Nhập
          </v-btn>
        </v-form>
      </v-tab-item>

      <!-- Sign Up Tab -->
      <v-tab-item>
        <v-form ref="signUpForm" @submit.prevent="signUp">
          <v-text-field
            v-model="signUpEmail"
            label="Email"
            hint="Nhập địa chỉ email của bạn"
            type="email"
            :rules="emailRules"
            required
            outlined
            dense
            class="mt-4"
          ></v-text-field>
          <v-text-field
            v-model="signUpPassword"
            label="Mật khẩu"
            hint="Nhập mật khẩu của bạn"
            :type="showSignUpPassword ? 'text' : 'password'"
            :rules="passwordRules"
            required
            outlined
            dense
            :append-icon="showSignUpPassword ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append="showSignUpPassword = !showSignUpPassword"
          ></v-text-field>
          <v-text-field
            v-model="signUpConfirmPassword"
            label="Xác nhận mật khẩu"
            hint="Nhập lại mật khẩu để xác nhận"
            :type="showSignUpConfirmPassword ? 'text' : 'password'"
            :rules="confirmPasswordRules"
            required
            outlined
            dense
            :append-icon="showSignUpConfirmPassword ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append="showSignUpConfirmPassword = !showSignUpConfirmPassword"
          ></v-text-field>
          <v-btn
            type="submit"
            color="primary"
            block
            :loading="signUpLoading"
            :disabled="signUpLoading"
            class="mt-4"
          >
            Đăng Ký
          </v-btn>
        </v-form>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import axios from '@/plugins/axios'
import { NotificationRequest } from '~/store'

@Component
export default class CustomAuth extends Vue {
  tab = 0
  signInEmail = ''
  signInPassword = ''
  signUpEmail = ''
  signUpPassword = ''
  signUpConfirmPassword = ''
  signInLoading = false
  signUpLoading = false
  showSignInPassword = false
  showSignUpPassword = false
  showSignUpConfirmPassword = false

  @Prop({ required: false, type: String, default: '/' }) to!: string

  emailRules = [
    (v: string) => !!v || 'Email là bắt buộc',
    (v: string) => /.+@.+\..+/.test(v) || 'Email không hợp lệ',
  ]

  passwordRules = [
    (v: string) => !!v || 'Mật khẩu là bắt buộc',
    (v: string) => (v && v.length >= 6) || 'Mật khẩu phải có ít nhất 6 ký tự',
  ]

  get confirmPasswordRules() {
    return [
      (v: string) => !!v || 'Xác nhận mật khẩu là bắt buộc',
      (v: string) => {
        if (!v) return true
        return v === this.signUpPassword || 'Mật khẩu xác nhận không khớp'
      },
    ]
  }

  async signIn() {
    const form = this.$refs.signInForm as any
    if (!form.validate()) {
      return
    }

    this.signInLoading = true
    try {
      const response = await axios.post('/v1/auth/signin', {
        email: this.signInEmail,
        password: this.signInPassword,
      })

      const authData = response.data.data
      
      // Set API key for future requests and save to localStorage
      const { setApiKey, saveAuthUser } = await import('~/plugins/axios')
      setApiKey(authData.api_key)

      // Create auth user object
      const authUser = {
        email: authData.email,
        displayName: null,
        id: authData.user_id,
      }

      // Save to localStorage
      saveAuthUser(authUser)

      // Update store - set auth user directly
      this.$store.commit('setAuthUser', authUser)
      
      // Load user data from API
      try {
        await this.$store.dispatch('loadUser')
      } catch (error) {
        // User data will be loaded on next page
        console.log('User data will be loaded on next page')
      }

      this.$store.dispatch('addNotification', {
        message: 'Đăng nhập thành công!',
        type: 'success',
      } as NotificationRequest)

      this.$router.push({ path: this.to })
    } catch (error: any) {
      const errorMessage =
        error.response?.data?.message || 'Lỗi khi đăng nhập'
      this.$store.dispatch('addNotification', {
        message: errorMessage,
        type: 'error',
      } as NotificationRequest)
    } finally {
      this.signInLoading = false
    }
  }

  async signUp() {
    const form = this.$refs.signUpForm as any
    if (!form.validate()) {
      return
    }

    this.signUpLoading = true
    try {
      const response = await axios.post('/v1/auth/signup', {
        email: this.signUpEmail,
        password: this.signUpPassword,
        confirm_password: this.signUpConfirmPassword,
      })

      const authData = response.data.data
      
      // Set API key for future requests and save to localStorage
      const { setApiKey, saveAuthUser } = await import('~/plugins/axios')
      setApiKey(authData.api_key)

      // Create auth user object
      const authUser = {
        email: authData.email,
        displayName: null,
        id: authData.user_id,
      }

      // Save to localStorage
      saveAuthUser(authUser)

      // Update store - set auth user directly
      this.$store.commit('setAuthUser', authUser)
      
      // Load user data from API
      try {
        await this.$store.dispatch('loadUser')
      } catch (error) {
        // User data will be loaded on next page
        console.log('User data will be loaded on next page')
      }

      this.$store.dispatch('addNotification', {
        message: 'Tạo tài khoản thành công!',
        type: 'success',
      } as NotificationRequest)

      this.$router.push({ path: this.to })
    } catch (error: any) {
      const errorMessage =
        error.response?.data?.message || 'Lỗi khi đăng ký'
      this.$store.dispatch('addNotification', {
        message: errorMessage,
        type: 'error',
      } as NotificationRequest)
    } finally {
      this.signUpLoading = false
    }
  }
}
</script>

