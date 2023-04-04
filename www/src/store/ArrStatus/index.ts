
interface IArrStatus {
    [key:string]:number[]
}

interface NObject {
    state:{}
    actions: {[key:string]:Function}
  }




const store:NObject = {
    state:{
        sarr:[1,2,3,4]
    } as IArrStatus,

    actions:{
        sarrpush(newState:IArrStatus, action:{val:number}){
            newState.sarr.push(action.val) 
        }
    }
}

export default store 


