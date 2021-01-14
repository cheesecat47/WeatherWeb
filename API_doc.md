# Domain Modeling

> [!NOTE]  
> Information the user should notice even if skimming

## board

---

```json
{
  "boardId": integer,
  "boardName": string,
  "createdAt": long
}
```

## article

---

```json
{
  "articleId": integer,
  "articleSummary": string,
  "articleTitle": string,
  "createdAt": long,
  "writerName": string
}
```

## articleDetail

---

```json
{
  "articleContent": string,
  "articleId": integer,
  "articleSummary": string,
  "articleTitle": string,
  "createdAt": long,
  "writerName": string
}
```

## user

---

```json
{
  "userId": string,
  "userName": string
}
```

# API endpoint

## 회원 관련

---

### 회원가입

---

```js
// @Route {POST} /signup

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Request Params
{
  "user_id": string,
  "user_pw": string,
  "user_name": string
}

// Response body
{
  "result": bool
}
```

### 회원탈퇴

```js
// @Route {POST} /signout

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Response body
{
  "boards": [board]
}
```

### 로그인

```js
// @Route {POST} /login

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Response body
{
  "boards": [board]
}
```

### 로그아웃

```js
// @Route {POST} /logout

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Response body
{
  "boards": [board]
}
```

### 아이디 중복 확인

---

```js
// @Route {GET} /valid/user/

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Request Params
{
  "userId": string
}

// Response body
{
  "result": bool
}
```

### 아이디 찾기

---

```js
// @Route {POST} /find/id

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Request Params
{
  "userName": string
}

// Response body
{
  "userId": string
}
```

### 인증 이메일(문자) 보내기

### 인증번호 유효성 검증 하기

### 비밀번호 재설정

---

```js
// @Route {POST} /find/pw

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Request Params
{
  "userId": string,
  "userName": string
}

// Response body
{
  "userPw": string
}
```

# 게시판

### 게시판 목록

---

```js
// @Route {GET} /board?page={page}&limit={limit}

// Reqeuest Header
{
  "Authorization": token, // Bearer Token for authority
  "Content-Type": "application-json"
}

// Response body
{
  "boards": [board]
}
```

board - crud
article - crud
