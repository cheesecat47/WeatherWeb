import React, { useEffect } from 'react';
import { connect } from 'react-redux';
import * as services from '../services/ArticleApi';

function BoardDetail({ ownProps, articleList }) {
  const { boardId } = ownProps;
  useEffect(() => {
    // services.getArticleList(`/board/${boardId}/article`);
    services.getArticleList(`/board/${boardId}`)
    // console.log(articleList);
  }, []);

  return (
    <div>
      
        <p>{articleList.createdAt}</p>
    </div>
  );

}

function mapStateToProps(state, ownProps) {
  // console.log(state, ownProps);
  return {
    articleList: state.articleReducer.data,
    ownProps
  }
}

export default connect(mapStateToProps)(BoardDetail);