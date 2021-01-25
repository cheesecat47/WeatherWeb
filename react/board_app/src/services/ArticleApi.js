import * as sendRequest from './sendRequest';
import { actionCreators } from '../actions/asyncActionsArticle';
import store from '../store.js';

export async function getArticleList(args) {
  const { dispatch } = store;
  //fetch request 시작 알림 
  dispatch(actionCreators.fetchRequest());

  try {
    const res = await sendRequest.sendGet(args);
    dispatch(actionCreators.fetchSuccess(res.data));
    // console.log(res.data);
    //status:200 -> Ok
    // return res.status === 200 ? res : "error";
  }
  catch (error) {
    dispatch(actionCreators.fetchError(error));
    console.log(error);
  }
}

export async function postArticle(args, info) {
  const { dispatch } = store;
  //fetch request 시작 알림
  dispatch(actionCreators.fetchRequest());

  try {
    const res = await sendRequest.sendPost(args, {
      data: {
        "id": info,
        "boardId": info,
        "boardName": info,
        "createdAt": info
      }
    });
    
    dispatch(actionCreators.fetchSuccess());
    //status:201 -> created
    // return res.status === 201 ? res : "error";
  }
  catch (error) {
    dispatch(actionCreators.fetchError(error));    
    console.log(error);
  }
}

// export async function deleteArticle(args) {
//   try {
//     const res = await sendRequest.sendDelete(args);
//     return res.status === 200 ? res : "error";
//   }
//   catch (error) {
//     console.log(error);
//   }
// }

// export async function sendPatch(args) {
//   try {
//     const res = await sendRequest.sendPatch(args);
//     return res.status === 200 ? res : "error";
//   }
//   catch (error) {
//     console.log(error);
//   }
// }




