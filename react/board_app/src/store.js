import { createStore} from 'redux';
import rootReducer from './reducers/rootReducer';
const store = createStore(rootReducer);

//store.subscribe??
export default store;
