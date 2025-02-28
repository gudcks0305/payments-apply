<template>
  <main class="container">
    <header>
      <h1>PortOne 결제 시스템</h1>
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
</style>