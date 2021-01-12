import React from 'react';
import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';

function MainPage({ id, user_id, name, article }) {
console.log(article)

  return (
    <Link to={{
      pathname: `/board/${id}`,
      state: {
        id,
        user_id,
        name,
        article
      }
    }}>
      <div className="card">
        <div className="board">
          <h3>{name}</h3>
          {article.map(article => (
            <p>{article.title}</p>
        

          ))}
         
          
        </div>
      </div>
    </Link>
  );
}



export default MainPage;