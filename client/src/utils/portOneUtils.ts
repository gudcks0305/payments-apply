import { PortonePaymentRequestParams, PortOnePaymentResponse } from '../types/payment';

const PORTONE_SHOP_ID = import.meta.env.VITE_PORTONE_SHOP_ID;
const CHANNEL_KEY = import.meta.env.VITE_CHANNEL_KEY;
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


export const requestPayment = (
  params: Omit<PortonePaymentRequestParams, 'channelKey'>
): Promise<PortOnePaymentResponse> => {
  if (!window.IMP) {
    return Promise.reject(new Error('PortOne SDK가 로드되지 않았습니다.'));
  }

  return new Promise((resolve) => {
    const requestParams: PortonePaymentRequestParams = {
      ...params,
      channelKey: CHANNEL_KEY,
    };
    
    window.IMP.request_pay(requestParams, (response: PortOnePaymentResponse) => {
      resolve(response);
    });
  });
}; 