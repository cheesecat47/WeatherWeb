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

// function Write({articleContent,writeArticle}) {
//   const [input, setInput] = useState("");

//   function onSubmit(e) {
//     e.preventDefault();
//     setInput("");
//     writeArticle(input);
//     // dispatch(addToDo(text));
//   }
//   const onChange = (e) => { setInput(e.target.value); }

//   return (
//     <div>
//       <h2>글 작성....</h2>
//       <form onSubmit={onSubmit}>
//         <input type="text" value={input} onChange={onChange}/>
//       <button>작성</button>
//       </form>
//       <button>취소</button>
//     </div>
//   )

// }

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
