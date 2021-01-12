import React from 'react';
// import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';
// import Board from '../components/MainPage';
import MainPage from '../components/MainPage';
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
						"대학원생 정기권 신청 가능한가?",
					context: "context입니다"
				},
				{
					article_id: 2,
					title: "hello",
					context: "context2입니다."
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
						"대학원생 정기권 신청 가능한가?",
					context: "대학원생 정기권 신청 가능한가?"
				},
				{
					article_id: 2,
					title: "hello",
					context: "hi"
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
						"대학원생 정기권 신청 가능한가?",
					context: "대학원생 정기권 신청 가능한가?"
				},
				{
					article_id: 2,
					title: "hello",
					context: "hi"
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
						<div className="Boards">
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