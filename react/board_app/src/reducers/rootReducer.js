import {combineReducers} from 'redux';
import {boardReducer, articleReducer} from './reducers';

const rootReducer = combineReducers({
    boardReducer,
    articleReducer
})

export default rootReducer;