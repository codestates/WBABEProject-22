# Online ordering system
> Minimal API for online ordering in Go

## Install
```
go get -d github.com/[woncoh1]/oos
go mod tidy
swag init
go run main.go
```

## Tech stack
- Programming language: [Go](https://go.dev/)
- Web framework: [Gin](https://github.com/gin-gonic/gin)
- Database: [MongoDB](https://www.mongodb.com/) ([MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/))
- Documentation: [Swagger](https://swagger.io/) ([swag](https://github.com/swaggo/swag))
- Configuration: [TOML](https://toml.io/en/) ([go-toml](https://github.com/pelletier/go-toml/tree/v2))
- Logging: [Zap](https://github.com/uber-go/zap)

## Project layout
- User generated
  - config: TOML configuration
  - logger: Zap log generator
  - db: MongoDB database and collections
  - dto: data transfer objects for requests and responses
  - model: data entities
  - controller: request handlers
  - service: business logic that interacts with the DB
  - router
    - connect HTTP method, endpoint path, and request handler
    - HTTP server
    - custom middleware
- Code generated
  - docs: OAS2 documentation
  - logs: Zap logs
```
./oos
├── config
|   ├── config.go
│   └── config.toml
├── controller
│   └── order.go
├── go.mod
├── go.sum
├── logger
│   └── logger.go
├── logs
│   ├── oos_2022-12-09.log
│   └── oos_2022-12-11.log
├── main.go
├── model
│   └── model.go
├── router
│   └── router.go
├── dto
│   ├── request.go
│   └── response.go
├── run
└── util
```

## Features
- Customer: 주문자
  - Query available products and my orders
  - Submit orders to providers
  - Give feedbacks on products and orders by writing reviews
- Provider: 피주문자
  - Query all products and orders
  - Create products
  - Change order status

### Customer
- Product
  - `GET` 메뉴 전체 조회
    - 각 카테고리별 sorted 리스트 출력
    - Sort by: 추천, 평점, 재주문수, 최신
    - Return: 결과 5~10여개 임의 생성 출력 및 sorting 여부 표시
- Order
  - `GET` 현재 주문 내역 전체 조회
    - 현재 주문내역 리스트 및 상태 조회
    - 없으면 null 리턴
  - `GET` 과거 주문 내역 전체 조회
    - 과거 주문내역 리스트 최신순으로 출력
    - 없으면 null 리턴
  - `GET` 주문 상태 조회
    - 하기 provider의 **주문 상태 조회**에서도 재활용
    - e.g. 접수중, 조리중, 배달중, etc
  - `POST` 주문
    - UI의 메뉴 리스트에서 해당 메뉴 선택, 주문 요청 및 초기상태 저장
    - 주문 정보를 입력받아 주문 저장 (ex. 선택 메뉴 정보, 전화번호, 주소 등 정보를 입력받아 DB 저장)
    - 주문 내역 초기상태 저장
    - Return: 금일 주문 받은 일련번호 (주문번호)
  - `PUT` 메뉴 추가
    - 상태조회 후 `배달중`일 경우 실패 알림
    - Return: 성공 실패
    - 실패시 신규주문으로 전환
  - `PUT` 메뉴 변경
    - 상태조회 후 `조리중` 또는 `배달중`일 경우 실패 알림
    - Return: 성공 실패
- Review
  - `GET` 평점 및 리뷰 조회
    - UI의 메뉴 리스트에서 상기 리스트 출력에 따라 개별 메뉴를 선택했다고 가정
    - Return: 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터
  - `POST` 평점 및 리뷰 작성
    - 과거 주문 내역 중, 평점 및 리뷰 작성
    - 해당 주문내역을 기준으로 평점 정보와 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장
    - Return: 성공 여부

### Provider
- Product
  - `POST` 신규 메뉴 등록
    - 사업장에서 신규 메뉴 관련 정보를 등록하는 과정
    - e.g. 메뉴 이름, 주문가능여부, 한정수량, 원산지, 가격, 맵기정도, etc
    - Return: 성공 여부
  - `PUT` 기존 메뉴 수정
    - 사업장에서 기존의 메뉴 정보 변경기능(ex. 가격변경, 원산지 변경, soldout)
    - 금일 추천 메뉴 설정 변경, 리스트 출력
    - Return: 성공 여부
  - `DELETE` 기존 메뉴 삭제
    - 실제 데이터 백업이나 뷰플래그를 이용한 안보임 처리
    - 금일 추천 메뉴 설정 변경, 리스트 출력
    - Return: 성공 여부
- Order
  - `GET` 주문 내역 전체 조회
  - `PUT` 주문 상태 변경
    - e.g. 상태: 접수중/접수취소/추가접수/접수-조리중/배달중/배달완료 등을 이용 상태 저장
    - 각 단계별 사업장에서 상태 업데이트
    - e.g. **접수중 → 접수** or **접수취소 → 조리중** or **추가주문 → 배달중**
    - Return: 성공여부
