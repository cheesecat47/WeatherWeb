import React from 'react';
import { createStore } from 'redux';

//useEffect({},[]) 두번째 인자에서 state ,data 등의 변화만을 감지해서 api에 데이터 요청하는
//형식으로 짜면 되겠다....

// const apiData={user: [],board: [],article: [],articleDetail: []};
/*actionCreater function*/

const POST = "POST";
const DELETE = "DELTE";
const PATCh = "PATCH";
const GET = "GET";
//user 정보 갱신하는 , board 정보 갱신, article 정보 갱신 func...

const getContent = () => {
    return { type: GET };
};

const addContent = (input) => {
    return {
        type: POST,
        input
    };
};
// console.log(apiData.user);
// const setArticle = (data) => ({ articleID: data.articleID, articleSummary: data.articleSummary, articleTitle: data.articleTitle });
const reducer = (state = [], action) => {
    switch (action.type) {
        case GET:
            return state;
        case POST:
            // console.log("action post에 왔다.",state,action.input)
            // const a=[{input:action.input}, ...state];
            // console.log(a);
            return [{input:action.input}, ...state];
    }
}

const store = createStore(reducer);
export const actionCreaters = {
    addContent,
    getContent
    // deleteContent,
    // updateContent
}

export default store;
