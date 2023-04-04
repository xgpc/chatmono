import hand from "./index"




let reducer = (state = {  ...hand.state}, action: { type: string}) => {
    let newState = JSON.parse(JSON.stringify(state))

    const hands  = hand.actions

    for(const key in hands){
        console.log(key);
        if (key === action.type) {

            const fc = hands[key];
            fc(newState, action);
        }
        
    }

    return newState
}

export default reducer