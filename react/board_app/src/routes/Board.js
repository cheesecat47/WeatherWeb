import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import WriteArticle from '../components/WriteArticle';
import '../css/Board.css';

function Board({ location, history, match }) {
 
  useEffect(() => {
    if (location.state === undefined) {
      //match.params로 id값 받아와서 데이터 요청해야할듯...
      // console.log(match.params);
      //url로 접근하는 사용자 케이스 고려
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
        <WriteArticle />
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
  //여기 else 부분 match 사용??처리를 생각해야겠음...
  else {
    return null;
  }
}
export default Board;