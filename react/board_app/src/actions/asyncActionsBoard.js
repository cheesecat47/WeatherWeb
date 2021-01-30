import {
    BOARD_FETCH_SUCCESS,
    BOARD_FETCH_REQUEST,
    BOARD_FETCH_ERROR,

} from './actionType';

//${type} 인자로 넘겨와서 board,article 구분할까?..
const fetchRequest = () => {
    return {
        type: BOARD_FETCH_REQUEST
    }
};

const fetchError = (error) => {
    return {
        type: BOARD_FETCH_ERROR,
        payload: error
    }
};

const fetchSuccess = (data) => {
    return {
        type: BOARD_FETCH_SUCCESS,
        payload: data
    }
};


export const actionCreators = {
    fetchRequest,
    fetchError,
    fetchSuccess,
}
