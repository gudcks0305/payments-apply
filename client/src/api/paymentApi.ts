import ApiResponse from '@/types/base';
import axios from 'axios';
import { ref } from 'vue';
import { PaymentData, PortOnePaymentResponse } from '../types/payment';

interface PaymentInitResponse {
  id: string;
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
      const response = await axios.post<ApiResponse<PaymentInitResponse>>(
        `${import.meta.env.VITE_BASE_URL}/api/v1/payments`, 
        paymentData
      );
      console.log(response.data);
      // 서버에서 받은 결제 ID와 상점 거래 ID를 저장
      merchantUid.value = response.data.data?.id || '';
      
      console.log('결제 초기화 성공:', response.data.data?.id);
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

export const completePayment = async (paymentData: PortOnePaymentResponse): Promise<PaymentData | Error> => {
  try {
    const response = await axios.put<ApiResponse<PaymentData>>(
      `${import.meta.env.VITE_BASE_URL}/api/v1/payments/${paymentData.merchant_uid}/complete`,
      paymentData
    );
    if (response.data.code >= 400) {
      const error = new Error(response.data.message);
      alert(`결제 처리 중 오류가 발생했습니다: ${response.data.message}`);
      return error;
    }

    return response.data.data! ;
  } catch (err: any) {
    console.error('결제 확인 실패:', err);

    let errorMessage = '서버 에러가 발생했습니다. 결제 확인에 실패했습니다.'; // 기본 에러 메시지

    if (err.response) {
      const responseData = err.response.data;

      if (typeof responseData === 'string') {
        errorMessage = `서버 에러: ${responseData}`;
      } else if (responseData && responseData.message) {
        errorMessage = `서버 에러: ${responseData.message}`;
      } else if (responseData && responseData.error) {
        errorMessage = `서버 에러: ${responseData.error}`;
      } else {
        errorMessage = `서버 에러: ${err.response.status} - ${err.response.statusText}`;
      }
    }
    
    // 에러 발생 시 즉시 alert 표시
    alert(errorMessage);
    return new Error(errorMessage);
  }
};