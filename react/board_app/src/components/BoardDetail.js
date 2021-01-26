import React, { useEffect } from 'react';
import { connect } from 'react-redux';
import * as services from '../services/ArticleApi';
import {Link} from 'react-router-dom';

function BoardDetail({ ownProps, articleList }) {
  const { boardId } = ownProps;
  useEffect(() => {
    services.getArticleList(`/article?boardId=${boardId}`);
  }, []);

  return (
    <div>
        {articleList.map(article =>(
          <Link to={{
            pathname: `/board/${boardId}/article/${article.articleId}`
          }}><p key={article.id}>
          {article.articleTitle}</p></Link>
        ))}
    </div>
  );
}

function mapStateToProps(state, ownProps) {
  return {
    articleList: state.articleReducer.data,
    ownProps
  }
}

export default connect(mapStateToProps)(BoardDetail);