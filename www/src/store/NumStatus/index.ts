

interface NObject {
    state: {}
    actions: {[key:string]: Function}
  }


const store:NObject =  {
    state:{
        num:20
    },

    actions:{
        add(newState:{num:number}, action:{}){
            newState.num ++
        },
        add2(newState:{num:number}, action:{val:number}){
            newState.num += action.val
        }
    }
}

export default store