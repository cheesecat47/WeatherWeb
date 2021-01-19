import * as config from './Config';

export async function getBoardList(args) {
  try {
    const res = await config.getInstance(args);
    // console.log(res.data);
    //status:200 -> Ok
    return res.status === 200 ? res : "error";
  }
  catch (error) {
    console.log(error);
  }
}

export async function postBoard(args, info) {
  console.log(info);
  try {
    const res = await config.postInstance(args, {
      data: {
        "id": info,
        "boardId": info,
        "boardName": info,
        "createdAt": info
      }
    });
    //status:201 -> created
    return res.status === 201 ? res : "error";
  }
  catch (error) {
    console.log(error);
  }
}

export async function deleteBoard(args) {
  try {
    const res = await config.deleteInstance(args);
    return res.status === 200 ? res : "error";
  }
  catch (error) {
    console.log(error);
  }
}

export async function patchBoard(args) {
  try {
    const res = await config.patchInstance(args);
    return res.status === 200 ? res : "error";
  }
  catch (error) {
    console.log(error);
  }
}




