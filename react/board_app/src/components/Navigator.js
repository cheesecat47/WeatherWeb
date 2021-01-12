import React from 'react';
import { Link } from 'react-router-dom';
// import PropTypes from 'prop-types';
import '../css/Navigator.css';

//게시판 id외에 넘어올 props 필요...?

function Navigator({ id,title }) {
  return (
    <nav>
      <ul className="nav-container">
        <li className="nav-item">
          <Link to={{
            pathname: "/",
            state: {
              fromNavigation: true
            }
          }}>Home</Link>
        </li>

        <li className="nav-item">
          <Link to={{
            pathname: `/board:${id}`,
            state: {
              fromNavigation: true
            }
          }}>{title}</Link>
        </li>

        <li className="nav-item">
          <Link to={{
            pathname: `/board:${id}`,
            state: {
              fromNavigation: true
            }
          }}>{title}</Link>
        </li>
      </ul>
    </nav>

  );
}

export default Navigator;
