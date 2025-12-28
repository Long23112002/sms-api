<template>
  <v-container fluid :fill-height="$vuetify.breakpoint.lgAndUp">
    <v-row v-if="$vuetify.breakpoint.lgAndUp" align="center" justify="center">
      <div>
        <v-img
          class="mx-auto mb-4"
          max-height="400"
          max-width="90%"
          contain
          :src="require('assets/img/person-texting.svg')"
        ></v-img>
        <div class="text-center">
          <h3 class="text-h5 mt-4">Chọn Một Tin Nhắn</h3>
          <p class="text--secondary">
            Đừng ngần ngại
            <a
              href=""
              target="_blank"
              class="text-decoration-none"
              >nhắn tin cho chúng tôi trên Zalo 0888880243</a
            >
            nếu bạn có bất kỳ câu hỏi nào
          </p>
        </div>
      </div>
    </v-row>
    <v-row v-else justify="end">
      <v-col class="px-0 py-0">
        <message-thread-header />
        <message-thread />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'ThreadsIndex',
  middleware: ['auth'],
  head() {
    return {
      title: 'Cuộc Trò Chuyện - httpSMS',
    }
  },
  async mounted() {
    await this.loadData()
  },

  methods: {
    async loadData() {
      await this.$store.dispatch('loadUser')
      await this.$store.dispatch('loadPhones')
      await this.$store.dispatch('loadThreads')
    },
  },
}
</script>
