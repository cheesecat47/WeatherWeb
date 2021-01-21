import React, { useState, useEffect } from 'react';
import MainPage from '../components/MainPage';
import '../css/MainPage.css';
import * as service from '../services/BoardApi';
import { connect } from 'react-redux';
import { actionCreaters } from '../store';

// import Board from './Board';

function Home({ boardList, getBoardInfo, createBoard }) {
  //어떤 일 발생하는 것 등록하고 useEffect로 => loading false 변환하기, reload값으로 data post 된거 감지..
  const [isLoading, setLoading] = useState("true");
  const [reload, setReload] = useState(false);
  //submit 할 때 data...
  const [text, setText] = useState("");

  //게시판 생성 sumbit...
  const onSubmit = (e) => {
    e.preventDefault();
    setText("");
    console.log(text);
    const fetchData = async () => {
      const res = await service.postBoard('/boards', text);
      console.log(res);
    }
    //api server 동작하면 fetchData()주석 해제!!!
    // fetchData();
    setReload(true);
    // createBoard(text)
  }
  const onChange = (e) => setText(e.target.value);

  // const onClick = (e) => {
  //   console.log("생각좀합시다...");
  // }

  const renderBoard = (list) => {
    if (list)
      return list;
    else {
      console.log("mount 하기 전..")
      return [];
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      const res = await service.getBoardList('/boards');
      console.log(res.data);
      // setBoardList(res.data);
      getBoardInfo(res.data);
    }
    //api server 동작하면 fetchData()주석 해제!!!
    // fetchData();
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
              <input type="text" value={text} onChange={onChange} />
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

function mapStateToProps(state) {
  console.log(state);
  return { boardList: state.boardReducer };
}

function mapDispatchToProps(dispatch) {
  return {
    getBoardInfo: (data) => {
      dispatch(actionCreaters.getContent(data))
    }, createBoard: (data) => {
      dispatch(actionCreaters.postContent(data))
    }
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(Home);

