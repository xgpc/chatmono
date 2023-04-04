// import store from "@/store";
// type RootState =ReturnType<typeof store.getState>


// 不要使用以上那种方式
type RootState = ReturnType<typeof import("@/store").getState>

interface Window{
    __REDUX_DEVTOOLS_EXTENSION__:function;
}