import { PaymentRequestParams, PortOneResponse } from '../types/payment';

const PORTONE_SHOP_ID = 'iamport00m';
const CHANNEL_KEY = 'channel-key-d879cf38-5530-4f0e-83e2-094e8f75e5ee';

declare global {
  interface Window {
    IMP: any;
  }
}

export const initPortOne = (): void => {
  if (!window.IMP) {
    console.error('PortOne SDK가 로드되지 않았습니다.');
    return;
  }
  
  window.IMP.init(PORTONE_SHOP_ID);
};

export const generateMerchantUid = (): string => {
  return `order-${new Date().getTime()}-${Math.floor(Math.random() * 1000)}`;
};

export const requestPayment = (
  params: Omit<PaymentRequestParams, 'channelKey'>
): Promise<PortOneResponse> => {
  if (!window.IMP) {
    return Promise.reject(new Error('PortOne SDK가 로드되지 않았습니다.'));
  }

  return new Promise((resolve, reject) => {
    
    const requestParams: PaymentRequestParams = {
      ...params,
      channelKey: CHANNEL_KEY,
    };
    
    window.IMP.request_pay(requestParams, (response: PortOneResponse) => {
      if (response.success) {
        resolve(response);
      } else {
        reject(new Error(response.error_msg || '결제에 실패했습니다.'));
      }
    });
  });
}; 