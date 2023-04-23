function  IsMobile (){
 
    let plat =navigator.userAgent.match( // 判断不同端
    /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
    );
    
    return (        plat? true:false    )
}
 
 
export default IsMobile;