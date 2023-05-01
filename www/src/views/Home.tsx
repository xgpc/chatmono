import React, { lazy,useState, useEffect } from 'react';

import IsMobile from '@/components/Ismobile';


const MobileView  = lazy(()=>import('@/views/MobileView')) 
const PcView  = lazy(()=>import('@/views/PcView')) 



const View: React.FC = () => {

  const [mobile, Setobile] = useState(false)


  useEffect(()=>{
    Setobile(IsMobile())
    return (()=>{
      console.log('刷新');
      
    })
  })


  return (

    <div style={{ height: '100vh' }}>

      {mobile == true && (<> <MobileView></MobileView></>)}
      {mobile == false && (<><PcView/></>)}

    </div>

  )


};

export default View;