import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import * as sendRequest from '../services/sendRequest';

//mainpage => 게시판 목록 보여주기

function MainPage({ boardName, boardId, articleList }) {

  const [articleList2, setArticleTitle] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
      const res = await sendRequest.sendGet(`/article?boardId=${boardId}`);
      console.log(res.data);
      setArticleTitle(res.data);
    }
    //api server 완성되면 주석 해제...
    // fetchData();
  }, [])

  return (
    <div className="card" >
      <div className="board">
        <Link to={{
          pathname: `/board/${boardId}`,
          state: {
            boardName,
            boardId
          }
        }}>
          <h3>Name:{boardName}</h3>
        </Link>
        <h4>{boardId}번 게시글</h4>
        <div>
        </div>
        {articleList2.map(article => (
          <Link to={{
            pathname: `/board/${boardId}/article/${article.articleId}`
          }}>
            <p>{article.articleTitle}</p></Link>
        ))}
      </div>
    </div>
  )
}

function mapStateToProps(state) {
  // console.log(state.articleReducer.data);
  return { articleList: state.articleReducer.data };
}

export default connect(mapStateToProps)(MainPage);