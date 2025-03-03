<template>
  <div class="container">
    <article>
      <h2>결제 정보 입력</h2>
      
      <form @submit.prevent="handleInitPayment">
        <div class="grid">
          <label for="merchantUid">
            가맹점 주문번호
            <input
              type="text"
              id="merchantUid"
              v-model="payment.merchantUid"
              placeholder="자동 생성됩니다"
              readonly
            />
            <small>* 주문번호는 결제 초기화 후 자동으로 생성됩니다.</small>
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
        
        <button type="submit" class="primary" :disabled="merchantUid !== ''">
          결제
        </button>
      </form>
      
      <div v-if="merchantUid" class="payment-info">
        <p>결제 ID: {{ merchantUid }}</p>
        <button @click="proceedToPayment" class="secondary">
          결제 진행하기
        </button>
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { completePayment, usePaymentApi } from '../api/paymentApi';
import { PortOnePaymentResponse } from '../types/payment';
import { requestPayment } from '../utils/portOneUtils';
const router = useRouter();

const { merchantUid, initializePayment } = usePaymentApi();
const isLoading = ref(false);

const payment = reactive({
  productName: '',
  amount: 10000,
  merchantUid: ''
});

const handleInitPayment = async () => {
  if (!payment.amount || payment.amount < 100) {
      alert('결제 금액은 최소 100원 이상이어야 합니다.');
    return;
  }
  
  isLoading.value = true;
  
  const success = await initializePayment({
    amount: payment.amount,
    productName: payment.productName,
    payMethod: 'card'
  });
  
  if (success) {
    payment.merchantUid = merchantUid.value;
  }
  
  isLoading.value = false;
};

const proceedToPayment = async () => {
  console.log('결제 진행: 결제 ID', merchantUid.value, '주문번호', merchantUid.value);
  
  try {
    const response = await requestPayment({
      merchant_uid: payment.merchantUid,
      name: payment.productName,
      amount: payment.amount,
      pay_method: 'card',
      buyer_email: 'test@test.com',
      buyer_name: '홍길동',
      buyer_tel: '01012345678',
      buyer_addr: '서울특별시 강남구 테헤란로 14길 6 남도빌딩 2층',
      buyer_postcode: '12345'
    });
    
    console.log(response);

    if (response.success) {
      alert('결제가 성공적으로 처리되었습니다.');
      
      handlePaymentComplete(response)
        .then(confirmedPayment => {
          console.log('비동기 결제 확인 완료:', confirmedPayment);
        })
        .catch(error => {
          console.error('비동기 결제 확인 실패:', error);
        });
    } else {
      alert(`결제에 실패했습니다: ${response.error_msg}`);
    }
  } catch (error) {
    console.error('결제 처리 중 오류 발생:', error);
    alert('결제 진행 중 오류가 발생했습니다: ' + (error instanceof Error ? error.message : String(error)));
  }
};

const handlePaymentComplete = async (paymentData: PortOnePaymentResponse) => {
  try {
    const result = await completePayment(paymentData);
    
    if (result instanceof Error) {
      // 백그라운드 처리이므로 사용자에게 알림을 표시하지 않음
      console.error('결제 확인 실패:', result.message);
      // 필요한 경우 로그 기록 또는 시스템 알림
      return null;
    }
    
    // 결제 성공 처리 (백그라운드)
    console.log('결제 확인 완료:', result);
    // 추가 성공 로직 (사용자에게 표시하지 않음)
    return result;
  } catch (error) {
    // 예상치 못한 오류 처리 (백그라운드)
    console.error('예상치 못한 오류:', error);
    // 필요한 경우 로그 기록 또는 시스템 알림
    return null;
  }
};
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

.loading {
  text-align: center;
  margin: 1rem 0;
  font-style: italic;
}

.payment-info {
  margin-top: 2rem;
  padding: 1rem;
  border-radius: 4px;
}

input[readonly] {
  background-color: #f2f2f2;
  color: #666;
}
</style>