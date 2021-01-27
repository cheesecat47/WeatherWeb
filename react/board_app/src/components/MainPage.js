import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import * as service from '../services/BoardApi';
import { actionCreaters } from '../store';

//mainpage => 게시판 목록 보여주기

function MainPage({ boardName, boardId }) {

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
          <h3>Name:{boardName}
          </h3>
        </Link>
        <h4>Id:{boardId}</h4>
      </div>
    </div>
  );
}

// function mapStateToProps(state) {
//   console.log(state);
//   return { boardList: state };
// }
// function mapDispatchToProps(dispatch) {
//   return {
//     createBoard: (data) => {
//       dispatch(actionCreaters.postContent(data))
//     }
//   };
// }
export default MainPage;