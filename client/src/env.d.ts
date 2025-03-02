/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_PORTONE_SHOP_ID: string
  readonly VITE_CHANNEL_KEY: string
  // 다른 환경 변수들...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
} 