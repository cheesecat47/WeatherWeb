import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import * as service from '../services/BoardApi';
import { actionCreaters } from '../store';

//게시판 목록 조회

function MainPage({ boardName, boardId}) {
  
  return (
    <div className="card" >
      <div className="board">
        <h3>Name:{boardName}
          {/* {renderBoard(boardList)} */}
        </h3>
        <h4>Id:{boardId}</h4>
        
        {/* {list[0].map(list => {
          <h3>{list.name}</h3>
        })} */}
        {/* {boardList.map(board => board.name)} */}
        {/* <Link to={{
          pathname: `/board/${id}`,
          state: {
            id: 
          }
          
        }}>
          <h3>{name}</h3></Link>
        <div className="board_article">{ }
            <Link to={{
              pathname: `/article/${article.article_id}`,
            }}> <p>{article.title}</p></Link> */}

      </div>
    </div>
  );
}

// function mapStateToProps(state) {
//   console.log(state);
//   return { boardList: state };
// }
function mapDispatchToProps(dispatch) {
  return {
    createBoard: (data) => {
      dispatch(actionCreaters.postContent(data))
    }
  };
}
export default connect(null, mapDispatchToProps)(MainPage);