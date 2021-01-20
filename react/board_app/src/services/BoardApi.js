import sendRequest from './sendRequest';

export async function getBoardList(args) {
  try {
    const res = await sendRequest.getInstance(args);
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
    const res = await sendRequest.postInstance(args, {
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
    const res = await sendRequest.deleteInstance(args);
    return res.status === 200 ? res : "error";
  }
  catch (error) {
    console.log(error);
  }
}

export async function patchBoard(args) {
  try {
    const res = await sendRequest.patchInstance(args);
    return res.status === 200 ? res : "error";
  }
  catch (error) {
    console.log(error);
  }
}




