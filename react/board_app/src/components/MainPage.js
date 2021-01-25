import React from 'react';
import { Link } from 'react-router-dom';
import {connect} from 'react-redux';
import BoardDetail from './BoardDetail';

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
        <h4>{boardId}번 게시글</h4>
        {/* <BoardDetail boardId={boardId}/> */}
      </div>
    </div>
  );
}

function mapStateToProps(state) {
  return { articleList: state.articleReducer.data };
}
// function mapDispatchToProps(dispatch) {
//   return {
//     createBoard: (data) => {
//       dispatch(actionCreaters.postContent(data))
//     }
//   };
// }
export default connect(mapStateToProps) (MainPage);