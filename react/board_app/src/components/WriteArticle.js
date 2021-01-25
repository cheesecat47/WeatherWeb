import React, { useState} from 'react';

function WriteArticle() {
  const [title, setTitle] = useState("");
  const [context, setContext] = useState("");
  const detailEle = document.querySelector("details");
  // console.log(detailEle);

  const onSubmit = (e) => {
    e.preventDefault();
    //여기에서 action 보내서 post 요청하는 logic 필요... 
    console.log(title,context);
    //글쓰기 창 접기...
    if (title || context) {
      detailEle.removeAttribute("open")
    }
    setContext("");
    setTitle("");  
  }

  const onTitleChange = (e) => {
    setTitle(e.target.value);
  }

  const onContextChange = (e) => {
    setContext(e.target.value);
  }

  return (
    <div>
      <form className="write" onSubmit={onSubmit}>
        <details >
          <summary>새 글을 작성해주세요 </summary>
          <p><input type="text" name="article_title" value={title} className="write_title" onChange={onTitleChange} placeholder="글 제목" ></input></p>
          <textarea name="btn_write" value={context} cols="55" rows="20" onChange={onContextChange} placeholder="자유롭게 작성해주세요"></textarea>
          <ul>
            <button> 작성완료</button>
          </ul>
        </details>
      </form>
    </div>
  )
}

export default WriteArticle;