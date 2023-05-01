import { useRoutes, Link } from "react-router-dom"
import router from "./router"
import Home from './views/Home';
import Notice from './components/Notice';

function App() {
  const outlet = useRoutes(router)



  return (
    <div className="App" >
     {/* {outlet}  */}
     <Notice></Notice>
     <Home/>
    </div>
  )
}

export default App
