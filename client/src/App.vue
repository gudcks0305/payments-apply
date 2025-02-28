<template>
  <main class="container">
    <header>
      <h1>PortOne 결제 시스템</h1>
      <p>가맹점 주문번호, 결제대상 제품명, 결제금액을 입력하여 결제를 진행하세요.</p>
    </header>
    <PaymentForm />
  </main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import PaymentForm from './components/PaymentForm.vue'

export default defineComponent({
  name: 'App',
  components: {
    PaymentForm
  }
})
</script>

<style>
header {
  text-align: center;
  margin-bottom: 2rem;
  padding-top: 2rem;
}

body {
  font-family: 'Noto Sans KR', sans-serif;
}
</style>