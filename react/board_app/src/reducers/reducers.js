const boardInitialState = {
    data: [],
    isFetching: false,
    error: null,
}

const articleInitialState = {
    data: [],
    isFetching: false,
    error: null
}

export const boardReducer = (state = boardInitialState, action) => {
    switch (action.type) {
        case "BOARD_FETCH_REQUEST":
            return { ...state, isFetching: true };

        case "BOARD_FETCH_SUCCESS":
            // console.log(action.payload);
            //post,delete,patch했을 때 state의 data가 전부 undefined 되는 것을 handling해야 할듯
            //post 했을때 받은 res.data를 state에 추가하고 rendering인지/ post하고 난 뒤 다시 get으로 받아오는건지...\
            return { data: action.payload, isFetching: false, error: null };

        case "BOARD_FETCH_ERROR":
            return { isFetching: false, error: action.payload, ...state }

        default: return state;
    }
}

export const articleReducer = (state = articleInitialState, action) => {
    switch (action.type) {
        case "ARTICLE_FETCH_REQUEST":
            return { ...state, isFetching: true };

        case "ARTICLE_FETCH_SUCCESS":
            //post,delete,patch했을 때 state의 data가 전부 undefined 되는 것을 handling해야 할듯
            //post 했을때 받은 res.data를 state에 추가하고 rendering인지/ post하고 난 뒤 다시 get으로 받아오는건지...\
            return { data: action.payload, isFetching: false, error: null};

        case "ARTICLE_FETCH_ERROR":
            return { isFetching: false, error: action.payload, ...state }

        default: return state;
    }
}