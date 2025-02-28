/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_PORTONE_SHOP_ID: string;
  readonly VITE_CHANNEL_KEY: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
