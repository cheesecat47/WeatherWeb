import React, { useState, useEffect } from 'react';
import MainPage from '../components/MainPage';
import '../css/MainPage.css';
import * as service from '../services/BoardApi';
import { connect } from 'react-redux';

function Home({ boardList }) {
  //어떤 일 발생하는 것 등록하고 useEffect로 => loading false 변환하기, reload값으로 data post 된거 감지..
  const [isLoading, setLoading] = useState("true");
  const [reload, setReload] = useState(false);
  //submit 할 때 data...
  const [newBoard, setNewBoard] = useState("");

  //게시판 생성 submit...
  const onSubmit = (e) => {
    e.preventDefault();
    //api server 동작하면 주석 해제!!!
    // service.postBoard('/board', newBoard);
    setNewBoard("");
    setReload(true);
  }
  const onChange = (e) => setNewBoard(e.target.value);

  useEffect(() => {
    //api server 동작하면 주석 해제!!!    
    // service.getBoardList('/board');
    // console.log(boardList);
    setLoading(false);
    setReload(false);
  }, [reload]);

  return (
    <section className="container">
      {isLoading ? (
        <div className="loader">
          <span className="loader_text">"Loading"</span>
        </div>
      ) : (
          <div className="boards">
            <h2>게시판 리스트</h2>
            <form onSubmit={onSubmit}>
              <input type="text" value={newBoard} onChange={onChange} />
              <button>게시판 생성</button>
            </form>
            {/* <button onClick={onClick}>게시판 삭제</button> */}
            {renderBoard(boardList).map(board => (
              <MainPage
                key={board.boardId}
                boardName={board.boardName}
                boardId={board.boardId} />
            ))}
          </div>)}
    </section>
  );
}

const renderBoard = (list) => {
  if (list) {
    // console.log(list);
    return list;
  }
  else {
    console.log("data 받기 전...")
    return [];
  }
};

function mapStateToProps(state) {
  // console.log(state.boardReducer);
  return { boardList: state.boardReducer.data };
}

// function mapDispatchToProps(dispatch) {
//   return {
//     getBoardInfo: (data) => {
//       dispatch(actionCreaters.getContent(data))
//     }, createBoard: (data) => {
//       dispatch(actionCreaters.postContent(data))
//     }
//   };
// }

export default connect(mapStateToProps)(Home);

