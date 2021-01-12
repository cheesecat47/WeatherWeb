import React from 'react';
import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';

function MainPage({ id, user_id, title }) {
  return (
    <Link to={{
      pathname: `/board/${id}`,
      state: {
        id,
        user_id,
        title
      }
    }}>
      <div className="card">
        <div className="board">
          <h3>{title}</h3>
          <p>article title ~~~</p>
          <p>article title~~~</p>
          <p>article title~~~</p>
        </div>
      </div>
    </Link>
  );
}



export default MainPage;