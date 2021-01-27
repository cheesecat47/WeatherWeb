import React from 'react';
import { createStore, combineReducers } from 'redux';
import * as service from './services/BoardApi';
//useEffect({},[]) 두번째 인자에서 state ,data 등의 변화만을 감지해서 api에 데이터 요청하는
//형식으로 짜면 되겠다....

/*actionCreater function*/
const POST_BOARD = "POST_BOARD";
const DELETE = "DELTE";
const PATCh = "PATCH";
const GET_BOARD = "GET_BOARD";

//user 정보 갱신하는 , board 정보 갱신, article 정보 갱신 func...

const getContent = (data) => {
    return {
        type: GET_BOARD,
        data
    }
};

const postContent = (data) => {
    return {
        type: POST_BOARD,
        data
    };
};

//according to 'board' reducer function
const boardReducer = (state = [], action) => {
    switch (action.type) {
        case GET_BOARD:
            // console.log(action.boardList);
            // const stateObj = [action.boardList, ...state];
            // console.log(a);
            // console.log(action);
            return action.data;

        case POST_BOARD:
            console.log("action post에 왔다.", state, action.input)
            // const a = [{ input: action.input }, ...state];
            // console.log(a);
            return [{ input: action.input }, ...state];

        default: return state;
    }
}

//according to 'article' reducer function
const articleReducer = (state = [], action) => {
    switch (action.type) {
        case "GET1":
            // console.log(action.boardList);
            // const stateObj = [action.boardList, ...state];
            // console.log(a);
            return [2];

        case "POST1":
            // console.log("action post에 왔다.", state, action.input)
            // const a = [{ input: action.input }, ...state];
            // console.log(a);
            return [{ input: action.input }, ...state];

        default: return state;
    }
}

const rootReducer = combineReducers({
    boardReducer,
    articleReducer
})
const store = createStore(rootReducer);

//store.subscribe??
export const actionCreaters = {
    postContent,
    getContent
    // deleteContent,
    // updateContent
}

export default store;
