import axios from 'axios';

//GET axios config...
export const getInstance = axios.create({
  baseURL: 'http://localhost:50004',
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  method: "get",
  timeout: 1000,
});

//POST axios config...
export const postInstance = axios.create({
  baseURL: 'http://localhost:50004',
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  data: {
    "boardId": "9",
    "boardName": null,
    "createdAt": null
  },
  method: "post",
  timeout: 1000,
});

//DELETE axios config...
export const deleteInstance = axios.create({
  baseURL: 'http://localhost:50004',
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  method: "delete",
  timeout: 1000,
});

//PATCH axios config...
export const patchInstance = axios.create({
  baseURL: 'http://localhost:50004',
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  method: "patch",
  timeout: 1000,
});