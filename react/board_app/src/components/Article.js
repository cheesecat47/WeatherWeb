import React, { useState } from 'react';
import { connect } from "react-redux";

function Article({boardId}) {
  console.log(boardId);

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

export default Article;
