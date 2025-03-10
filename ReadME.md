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
  externalAPI->>client: 응답
  client->>client: redirect to 결제 완료 페이지
  
  client->>server: 결제 완료 상태 complete API 요청 (결제 서버 ID 포함)
  Note right of client: 결제 검증 및 취소를 위해 비동기 Non blocking 요청
  Note right of client: 결제 Imp ID를 서버에 전달하여 complete API 호출
  externalAPI->>server: 결제 완료 상태 응답
  server->>server: 결제 상태 확인 및 최종 처리
  server->>client: 결제 상태 응답

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


## 실행 방법 
```shell
# make .env file @see .env.example
cat << EOF > .env
PAYMENTS_PORTONE_IMP_KEY=YOUR_IMP_KEY
PAYMENTS_PORTONE_IMP_SECRET=YOUR_IMP_SECRET
PAYMENTS_PORTONE_SERVER_PORT=8080
EOF

# run docker 
docker-compose up -d
```

http://localhost:3000 접속 후 결제 요청

## API 문서
@see
http://localhost:8080/swagger-ui/index.html

### API 목록
- POST api/v1/payments : 결제 기본 정보 생성 ( id 채번 및 기록 용 )
- PUT api/v1/payments/complete : 요구사항을 가장 basic 하게 구현한 api 
- GET api/v1/payments/imp/:impId : impId로 결제 정보 조회 ( 외부 API 호출 테스트 용 )
- POST api/v1/payments/imp/:impId/cancel : impId로 결제 취소 ( 외부 API 호출 테스트 용 )
- POST api/v1/payments/{merchantId}/cancel : merchantId로 결제 취소 ( 공식 문서 validate 참고 및 데이터 베이스 데이터 상태 변경 )
