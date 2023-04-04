import {useSelector, useDispatch} from 'react-redux'

import store from "@/store"

// type RootState = ReturnType<typeof store.getState>

const View = () => {
    const dispatch = useDispatch()
    // 对数字操作

    const {num} = useSelector((state:RootState)=>({
        num: state.NumreStatus.num
    }))


    // 修改参数数据
    const onClickNum = () =>{
        dispatch({type:"add"})
    }

    const onClickNum2 = () =>{
        dispatch({type:"add2", val:10})
    }



    const{ sarr } = useSelector((state:RootState)=>({
        sarr:state.ArrStatus.sarr
    }));


    const onClickArr = () => {
        dispatch({type:"sarrpush", val:1})
    }

    return (
        <div className='about'>
            <p>
                this is page1
            </p>

            <p>{num}</p>
            <button onClick={onClickNum}> 按钮</button>
            <button onClick={onClickNum2}> 按钮2</button>

            <p>
                -----------------------------------------------------
            </p>

            <p>{sarr}</p>
            <button onClick={onClickArr}> 添加元素</button>
            <p></p>
        </div>
    )
}


export default View