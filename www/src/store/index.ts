import { legacy_createStore, combineReducers} from 'redux'

import NumreStatus from "./NumStatus/reducer"
import ArrStatus from "./ArrStatus/reducer"

const reducers = combineReducers({
     NumreStatus,
     ArrStatus,
})

const store = legacy_createStore(
     reducers,
     window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__());

export default store