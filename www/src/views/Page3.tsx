import React, { useState, useEffect } from 'react';


const View = () => {
    const [count, setCount] = useState(0);
  
    // 相当于 componentDidMount + componentDidUpdate
    useEffect(() => {
      document.title = `You clicked ${count} times`
    }, [count]);
  
    return (
      <div>
        <p>You clicked {count} times</p>
        <button onClick={() => setCount(count + 1)}>
          Click me
        </button>
      </div>
    );
  }


export default View