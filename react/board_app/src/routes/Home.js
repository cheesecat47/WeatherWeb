import React from 'react';
// import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';
// import Board from '../components/MainPage';
import MainPage from '../components/MainPage';
import '../css/MainPage.css';
// import * as service from '../services/Api';
// import Board from './Board';

class Home extends React.Component {
  state = {
    isLoading: true,
    user: [{}],
    board: [{
      id: 1,
      name: "대구캠게시판",
      user_id: "안태건",
      article: [
        {
          article_id: 1,
          title:
            "대학생 정기권 신청 가능한가asdasdadsd",
          context: "아니 죽어도 안해주더라"
        },
        {
          article_id: 2,
          title: "대구캠 컴학 만세",
          context: "만만세입니다"
        }
      ]
    },
    {
      id: 2,
      name: "정보게시판",
      user_id: "한진규",
      article: [
        {
          article_id: 1,
          title:
            "정보 좀 주세요",
          context: "정보는 없단다..."
        },
        {
          article_id: 2,
          title: "제발 좀 주시죠",
          context: "나도 몰랑"
        }
      ]
    },
    {
      id: 3,
      name: "취업게시판",
      user_id: "신주용",
      article: [
        {
          article_id: 1,
          title:
            "취업 시켜 주시오",
          context: "넌 대학원 가거라"
        },
        {
          article_id: 2,
          title: "대학원 가기 싫어요",
          context: "싫으면 몰라"
        }
      ]
    }],

    // article: [
    // 	{
    // 		id: 1,
    // 		board_id: 1,
    // 		title: "대학원생 정기권 신청 가능한가?",
    // 		context: "밀리의 서재 구독 1년치 3명이거 끊어서하려합니다 한명남았어요 쪽지주세요~"
    // 	}]
  };


  /* api server 어떤 식으로 받아올지 정하고 state update
    getUserInfo = asnyc () => {
      const user= await axios.get("url");
      this.setState({user:user})
    }
    getBoardInfo = asnyc () => {
      const board= await axios.get("url");
      this.setState({board})
    }
    getArticleInfo = asnyc () => {
      const article= await axios.get("url");
      this.setState({article})
    }
  */
  componentDidMount() {
    // console.log(this.state)
    this.setState({ isLoading: false })
  }
  render() {
    const { isLoading, user, board } = this.state;

    return (
      <section className="container">
        {isLoading ? (
          <div className="loader">
            <span className="loader_text">"Loading"</span>
          </div>
        ) : (
            <div className="boards">
              {board.map(board => (
                <MainPage
                  key={board.id}
                  id={board.id}
                  name={board.name}
                  user_id={board.user_id}
                  article={board.article}
                />
              ))}
            </div>
          )}
      </section>
    );
  }
};

export default Home;