import axios from 'axios';
import config from '../config';

//GET axios config...
export const sendGet = axios.create({
  baseURL: config.API_SERVER_URL,
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  method: "get",
  timeout: 1000,
});

//POST axios config...
export const sendPost = axios.create({
  baseURL: config.API_SERVER_URL,
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  // data: {
  //   "boardId": "9",
  //   "boardName": null,
  //   "createdAt": null
  // },
  method: "post",
  timeout: 1000,
});

//DELETE axios config...
export const sendDelete = axios.create({
  baseURL: config.API_SERVER_URL,
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  method: "delete",
  timeout: 1000,
});

//PATCH axios config...
export const sendPatch = axios.create({
  baseURL: config.API_SERVER_URL,
  //     header:{
  //   "Authorization": token, // Bearer Token for authority
  //   "Content-Type": "application-json"
  // }
  method: "patch",
  timeout: 1000,
});
