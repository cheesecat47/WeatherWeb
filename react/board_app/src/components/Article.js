import React, { useState } from 'react';
import { actionCreaters } from '../store';
import { connect } from "react-redux";

function Article() {

  // const btnWrite = (e) =>{
  //   e.preventDefault();
  //   return <Write/>
  // }
 return (
    <div>
      <h1>Article</h1>
      <h3>article 내용~~~~</h3>
      {/* <button onClick={btnWrite}>글쓰기</button> */}
    </div>);
}

function mapStateToprops(state) {
  console.log(state);
  return { articleContent: state };
  // console.log(state);
}

function mapDispatchToProps(dispatch) {
  return {
    writeArticle: (input) => dispatch(actionCreaters.addContent(input))
  };
}
export default connect(mapStateToprops, mapDispatchToProps)(Article);
// connect (mapStateToprops,mapDispatchToProps)(Write);
