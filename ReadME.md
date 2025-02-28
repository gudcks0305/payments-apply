# 프레임워크 및 라이브러리 설정 
- 언어 : GoLang
- 프레임워크 : Gin
- 라이브러리 
  - viper : 설정파일
  - swaggo : API 문서화
  - testify : 테스트
  - uber_fx : DI
## API 고려 사항 
- 결제 상태 확인을 위해 client가 외부 API 호출 전 요청을 서버에 저장
- 해당 결제 서버 ID를 화면에 넘겨주고 client는 해당 ID로 외부 결제 API에 요청
- 클라이언트는 redirect를 하고 나서 결제 상태를 확인하기 위해 서버에 해당 ID로 요청 ( complete API 이용)
- 동일 결제 ID에 대해서 이미 완료가 된 결제라면 해당 요청에 대해 서버에서 외부 api 에 대해 해당 결제 취소 요청
```mermaid
sequenceDiagram
  participant client
  participant server
  participant externalAPI

  client->>server: 결제 요청
  server->>server: 결제 정보 저장 및 결제 서버 ID 생성
  server->>client: 결제 서버 ID 전달
  Note right of server: 클라이언트, 결제 서버 ID 저장 및 화면에 전달
  client->>externalAPI: 결제 요청 (결제 서버 ID 포함)
  Note left of client: 결제 요청을 externalAPI에 전달
  externalAPI->>client: 응답 (redirect URL 포함)
  client->>client: redirect to 결제 완료 페이지
  client->>server: 결제 완료 상태 complete API 요청 (결제 서버 ID 포함)
  Note right of client: 결제 서버 ID를 서버에 전달하여 complete API 호출
  server->>server: 결제 상태 확인 및 최종 처리
  server->>client: 결제 완료 상태 응답

  alt 동일 결제 ID로 이미 완료된 결제 존재
    server->>externalAPI: 결제 취소 요청 (결제 서버 ID 포함)
    externalAPI->>server: 결제 취소 결과 응답
    server->>server: 결제 취소 처리
    server->>client: 결제 취소 상태 응답
  end
  
```


# 프론트엔드 설정
- 언어 : TypeScript
- 프레임워크 : Vue
