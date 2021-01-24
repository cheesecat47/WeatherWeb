import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import '../css/Board.css';

function Board({ location, history, match }) {
  const [title, setTitle] = useState("");
  const [ context, setContext ] = useState("");

  const onSubmit = (e) => {
    e.preventDefault();
    // console.log(e.target);
    //여기에서 action 보내서 post 요청하는 logic 필요... 
    setContext("");
    setTitle("");

  }

  const onTitleChange = (e) =>{
    setTitle(e.target.value);
  }
  const onContextChange = (e) => {
    setContext(e.target.value);
  }
  useEffect(() => {
    if (location.state === undefined) {
      //match.params로 id값 받아와서 데이터 요청해야할듯...
      // console.log(match.params);
      history.push('/');
    }
  }, [])

  if (location.state) {
    const { boardName, boardId } = location.state;
    const articleId = 1;

    return (
      <div className="container">
        <div className="board_name">
          {boardName}
        </div>
        <form className="write" onSubmit={onSubmit}>
          <details >
            <summary>새 글을 작성해주세요 </summary>
            <p><input type="text" name="article_title" value={title} className="write_title" onChange={onTitleChange} placeholder="글 제목" ></input></p>
            <textarea name="btn_write" value={context} cols="55" rows="20"onChange={onContextChange} placeholder="자유롭게 작성해주세요"></textarea>            <ul>
            <button> 작성완료</button>
            </ul>
          </details>
        </form>

        <div className="articles">
          <div className="article">
            {/* article axios.get으로 데이터 받아오기... */}
            <Link to={{
              pathname: `/board/${boardId}/article/${articleId}`
            }}>게시글 1</Link>
          </div>
        </div>
      </div>
    );
  }
  //여기 처리를 생각해야겠음...
  else {
    return null;
  }
}
export default Board;