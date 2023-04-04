import { ChangeEvent, useEffect, useState } from 'react'

import {marked} from 'marked' // 导入markdown转换器
import hljs from 'highlight.js' // 导入markdown代码块高亮包
import "highlight.js/styles/atom-one-light.css" // 代码块具体的高亮样式



const View = () => {


    var defaultData:string = String(`
        - 123
        - 32
        - 4243
    
    `)


    const [markdown, setMarkdown] =useState("# markdown\n```typescript\n// markdown样式\nimport {marked} from 'marked' // 导入markdown转换器\nimport hljs from 'highlight.js' // 导入markdown代码块高亮包\n```")
    // `# This is a header
    // and some **bold** text`
  
    const handleChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
      setMarkdown(e.target.value)

      let markdown_element: any = document.getElementById("Note-Container")
      markdown_element.innerHTML = Get_Note(e.target.value)
      hljs.highlightAll() // 高亮代码块
    }

    function Get_Note(s:string){
        return marked(markdown)
    }
    
    useEffect(() => {
        let markdown_element: any = document.getElementById("Note-Container")
        markdown_element.innerHTML = Get_Note(markdown)
        hljs.highlightAll() // 高亮代码块
    }, [])

    return (
        <div className="App">
        {/* <textarea placeholder='Enter your markdown text' className="textarea" input={markdown} onChange={(e)=>handleChange(e)} /> */}
            
        <div id="Note-Container"></div>

         <div className="htmlbox">How can I export HTML here?</div>
      </div>
    )
}


export default View