import Vue from 'vue'
import { intervalToDuration, formatDuration } from 'date-fns'
import { parsePhoneNumber, isValidPhoneNumber } from 'libphonenumber-js'

export const formatPhoneNumber = (value: string) => {
  if (!isValidPhoneNumber(value)) {
    return value
  }
  const phoneNumber = parsePhoneNumber(value)
  if (phoneNumber) {
    return phoneNumber.formatInternational()
  }
  return value
}

Vue.filter('phoneNumber', (value: string): string => {
  return formatPhoneNumber(value)
})

Vue.filter('phoneCountry', (value: string): string => {
  const phoneNumber = parsePhoneNumber(value)
  if (phoneNumber && phoneNumber.country) {
    // @ts-ignore
    const regionNames = new Intl.DisplayNames(undefined, { type: 'region' })
    return regionNames.of(phoneNumber.country) ?? 'earth'
  }
  return 'Earth'
})

Vue.filter('timestamp', (value: string): string => {
  return new Date(value).toLocaleString()
})

Vue.filter('money', (value: string): string => {
  // Chuyển đổi từ USD sang VND (1 USD = 24,000 VND)
  const usd = parseFloat(value) || 0
  const vnd = Math.round(usd * 24000)
  return new Intl.NumberFormat('vi-VN', {
    style: 'currency',
    currency: 'VND',
  }).format(vnd)
})

Vue.filter('decimal', (value: string): string => {
  return new Intl.NumberFormat('en-US', {
    style: 'decimal',
  }).format(parseInt(value))
})

Vue.filter('billingPeriod', (value: string): string => {
  const options = {
    year: 'numeric',
    month: 'long',
  }
  // @ts-ignore
  return new Date(value).toLocaleDateString('vi-VN', options)
})

Vue.filter('humanizeTime', (value: string): string => {
  const durations = intervalToDuration({
    start: new Date(),
    end: new Date(value),
  })
  return formatDuration(durations)
})

Vue.filter('capitalize', (value: string): string => {
  return value.charAt(0).toUpperCase() + value.slice(1)
})
