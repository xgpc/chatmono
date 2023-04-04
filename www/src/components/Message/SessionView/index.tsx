import { log } from "console";
import { marked } from "marked"
import { ReactElement, JSXElementConstructor, ReactFragment, useEffect, useState } from "react"


interface IsessionData {
  role?: string,
  content?: string
}

import hljs from "highlight.js";



const App = ({ sessionList }: { sessionList: IsessionData[] }) => {

  useEffect(() => {

    for (let index = 0; index < sessionList.length; index++) {
      const element = sessionList[index];
      console.log(element);

    }

  }, []);

  // 配置marked
  let renderer = new marked.Renderer();
  // 配置链接，点击链接时，打开新的浏览器窗口
  renderer.link = function (href, title, text) {
    return `<a target="_blank" href="${href}" title="${title}">${text}</a>`
  }

  marked.setOptions({
    renderer: renderer,
    highlight: function (code, lang) {
      const language = hljs.getLanguage(lang) ? lang : 'plaintext';
      return hljs.highlight(code, { language }).value;
    },
    langPrefix: 'hljs language-', // highlight.js css expects a top-level 'hljs' class.
    pedantic: false,
    gfm: true,
    breaks: false,
    sanitize: false,
    smartLists: true,
    smartypants: false,
    xhtml: false
  });

  // 根据传入的数据渲染
  const addMessage = (message: IsessionData) => {

    switch (message.role) {
      case 'assistant':

        const content1 = marked.parse(message.content != undefined ? message.content : '')
        return <div className="site-html-preview-content" dangerouslySetInnerHTML={{ __html: content1 }}></div>
        break;

      case 'user':
        const markedData2 = marked.parse(message.content != undefined ? message.content : '')
        return <div className="site-html-preview-content" dangerouslySetInnerHTML={{ __html: markedData2 }}></div>
        break;
      default:
        break;
    }

  }


  return (
    <>
      <div id="chat-content" style={{height:'80%', overflow:'auto' ,background:'#EEEEEE' }}>

        {
          sessionList.map((item: IsessionData) => {
            return <div> {addMessage(item)} </div>
          })
        }
      </div>
    </>
  )

}

export default App;