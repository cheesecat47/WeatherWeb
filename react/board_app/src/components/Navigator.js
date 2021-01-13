import React from 'react';
import { Link,withRouter } from 'react-router-dom';
// import PropTypes from 'prop-types';
import '../css/Navigator.css';

//게시판 id외에 넘어올 props 필요...?

function Navigator({ location }) {
  console.log(location);
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
{/* 
        <li className="nav-item">
          <Link to={{
            pathname: `/board/${location.state.id}`,
            state: {
              fromNavigation: true
            }
          }}>{location.state.name}</Link>
        </li> */}

        {/* <li className="nav-item">
          <Link to={{
            pathname: `/board/${location.state.id}`,
            state: {
              fromNavigation: true
            }
          }}>{name}</Link>
        </li> */}
      </ul>
    </nav>

  );
}

export default withRouter(Navigator);
