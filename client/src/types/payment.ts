export interface PaymentRequestParams {
  channelKey: string;
  pay_method: string;
  merchant_uid: string;
  name: string;
  amount: number;
  buyer_email?: string;
  buyer_name?: string;
  buyer_tel?: string;
  buyer_addr?: string;
  buyer_postcode?: string;
  m_redirect_url?: string;
}

export interface PaymentFormData {
  merchantUid: string;
  productName: string;
  amount: number | null;
}

export interface PortOneResponse {
  success: boolean;
  error_code?: string;
  error_msg?: string;
  imp_uid?: string;
  merchant_uid?: string;
  pay_method?: string;
  paid_amount?: number;
  status?: string;
  name?: string;
  pg_provider?: string;
  pg_tid?: string;
  paid_at?: number;
}

export interface PaymentVerificationResponse {
  success: boolean;
  message: string;
  data?: any;
}