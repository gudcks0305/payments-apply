import axios from 'axios';
import { ref } from 'vue';

interface PaymentInitResponse {
  merchantUid: string;
  amount: number;
  productName: string;
}

interface PaymentInitRequest {
  amount: number;
  productName: string;
  payMethod: string;
}

export function usePaymentApi() {
  const merchantUid = ref<string>('');
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  const initializePayment = async (paymentData: PaymentInitRequest): Promise<boolean> => {
    isLoading.value = true;
    error.value = null;
    
    try {
      const response = await axios.post<PaymentInitResponse>(
        `${import.meta.env.VITE_BASE_URL}/api/v1/payments`, 
        paymentData
      );
      
      // 서버에서 받은 결제 ID와 상점 거래 ID를 저장
      merchantUid.value = response.data.data.id;
      
      console.log('결제 초기화 성공:', response.data.data.id);
      return true;
    } catch (err) {
      console.error('결제 초기화 실패:', err);
      error.value = err instanceof Error ? err.message : '결제 요청 중 오류가 발생했습니다';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  return {
    merchantUid,
    isLoading,
    error,
    initializePayment
  };
} 