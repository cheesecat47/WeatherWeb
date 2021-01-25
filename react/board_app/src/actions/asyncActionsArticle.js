import {
    ARTICLE_FETCH_SUCCESS,
    ARTICLE_FETCH_REQUEST,
    ARTICLE_FETCH_ERROR,

} from './actionType';

const fetchRequest = () => {
    return {
        type: ARTICLE_FETCH_REQUEST
    }
};

const fetchError = (error) => {
    return {
        type: ARTICLE_FETCH_ERROR,
        payload: error
    }
};

const fetchSuccess = (data) => {
    return {
        type: ARTICLE_FETCH_SUCCESS,
        payload: data
    }
};


export const actionCreators = {
    fetchRequest,
    fetchError,
    fetchSuccess,
}
