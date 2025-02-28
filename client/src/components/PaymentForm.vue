<template>
    <div class="container">
      <article>
        <h2>결제 정보 입력</h2>
        <form @submit.prevent="submitPayment">
          <div class="grid">
            <label for="merchantUid">
              가맹점 주문번호
              <input
                type="text"
                id="merchantUid"
                v-model="payment.merchantUid"
                placeholder="예: ORD20240501-001"
                required
              />
            </label>
          </div>
          
          <div class="grid">
            <label for="productName">
              결제대상 제품명
              <input
                type="text"
                id="productName"
                v-model="payment.productName"
                placeholder="예: 프리미엄 멤버십"
                required
              />
            </label>
          </div>
          
          <div class="grid">
            <label for="amount">
              결제금액
              <input
                type="number"
                id="amount"
                v-model="payment.amount"
                placeholder="예: 10000"
                min="100"
                required
              />
              <small>* 최소 결제 금액은 100원입니다.</small>
            </label>
          </div>
          
          <button type="submit" class="primary">결제하기</button>
        </form>
      </article>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, reactive } from 'vue'
  
  interface Payment {
    merchantUid: string;
    productName: string;
    amount: number | null;
  }
  
  export default defineComponent({
    name: 'PaymentForm',
    setup() {
      const payment = reactive<Payment>({
        merchantUid: '',
        productName: '',
        amount: null
      })
  
      const submitPayment = () => {
        if (!payment.amount || payment.amount < 100) {
          alert('결제 금액은 최소 100원 이상이어야 합니다.')
          return
        }
  
        console.log('결제 정보:', payment)
        // 여기에 실제 결제 요청 로직 구현
        alert(`${payment.productName} 상품에 대한 ${payment.amount}원 결제가 요청되었습니다.`)
      }
  
      return {
        payment,
        submitPayment
      }
    }
  })
  </script>
  
  <style scoped>
  .container {
    max-width: 600px;
    margin: 0 auto;
    padding-top: 2rem;
  }
  
  h2 {
    text-align: center;
    margin-bottom: 2rem;
  }
  
  button {
    margin-top: 1rem;
  }
  </style>