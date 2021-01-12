import React from 'react';
import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';

class Board extends React.Component {
  

  componentDidMount() {
    const { location, history } = this.props;
    if (location.state === undefined) {
      history.push("/");
    }
  }

  render() {
    console.log(this.props);
    const { location } = this.props;
    if (location.state) {
      return (
        <div className="board">
          <h3>{location.state.name}</h3>
          {location.state.article.map(article => (
            <h3><Link to = "/article" style={{ textDecoration: 'none'}}>제목 : {article.title}</Link>
            <h6>내용:{article.context}</h6>
            </h3>
            // <p>{article.context}</p>
          ))}
          {/* <p>{location.state.article.map(article => (
            article.context
          ))}</p> */}

        </div>
      );
    }
    else {
      return null;
    }
  }
}

export default Board;