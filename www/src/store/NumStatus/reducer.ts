import handleNum from "./index";

// const defaultState = {

//     // num:handleNum.state.num

//     // 解构的写法
//     ...handleNum.state
// }

let reducer = (state = {...handleNum.state}, action: { type: string, val:number }) => {
    let newState = JSON.parse(JSON.stringify(state))


    for(const key in handleNum.actions){
        console.log(key);
        if (key === action.type) {
            const fc  = handleNum.actions[key];
            fc(newState, action)
        }
        
    }


    return newState
}

export default reducer