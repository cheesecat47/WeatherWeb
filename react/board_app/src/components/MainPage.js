import React from 'react';
import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';

function MainPage({ id, user_id, name, article }) {
  return (
    <div className="card" >
      <div className="board">
        <Link to={{
          pathname: `/board/${id}`,
          state: {
            id,
            name,
            user_id,
            article
          }
        }}>
          <h3>{name}</h3></Link>
        <div className="board_article">{ }
          {article.map(article => (
            <Link to={{
              pathname: `/article/${article.article_id}`,
              state: {
                id,
                name,
                user_id,
                article
              }
            }}> <p>{article.title}</p></Link>
          ))}
        </div>
      </div>
    </div>
  );
}

export default MainPage;