import React from 'react';
// import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';

class Board extends React.Component {
  componentDidMount() {
    const { location, history } = this.props;
    if (location.state === undefined) {
      history.push("/");
    }
  }

  render() {
    const { location } = this.props;
    if (location.state) {
      return (
        <div className="board">
          <h3>{location.state.title}</h3>
          <p>article의 내용을 입력</p>
        </div>
      );
    }
    else {
      return null;
    }
  }
}

export default Board;